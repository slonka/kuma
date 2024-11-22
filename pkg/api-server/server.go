package api_server

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/bakito/go-log-logr-adapter/adapter"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-logr/logr"
	"github.com/mrichman/hargo"
	"github.com/pkg/errors"
	http_prometheus "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful"

	"github.com/kumahq/kuma/pkg/api-server/authn"
	"github.com/kumahq/kuma/pkg/api-server/filters"
	api_server "github.com/kumahq/kuma/pkg/config/api-server"
	kuma_cp "github.com/kumahq/kuma/pkg/config/app/kuma-cp"
	config_core "github.com/kumahq/kuma/pkg/config/core"
	config_store "github.com/kumahq/kuma/pkg/config/core/resources/store"
	config_types "github.com/kumahq/kuma/pkg/config/types"
	"github.com/kumahq/kuma/pkg/core"
	resources_access "github.com/kumahq/kuma/pkg/core/resources/access"
	"github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/apis/system"
	"github.com/kumahq/kuma/pkg/core/resources/manager"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/registry"
	"github.com/kumahq/kuma/pkg/core/runtime"
	"github.com/kumahq/kuma/pkg/dns/vips"
	"github.com/kumahq/kuma/pkg/insights/globalinsight"
	kuma_log "github.com/kumahq/kuma/pkg/log"
	"github.com/kumahq/kuma/pkg/plugins/authn/api-server/certs"
	"github.com/kumahq/kuma/pkg/plugins/resources/k8s"
	secrets_k8s "github.com/kumahq/kuma/pkg/plugins/secrets/k8s"
	tokens_server "github.com/kumahq/kuma/pkg/tokens/builtin/server"
	kuma_srv "github.com/kumahq/kuma/pkg/util/http/server"
	util_prometheus "github.com/kumahq/kuma/pkg/util/prometheus"
	"github.com/kumahq/kuma/pkg/version"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	"github.com/kumahq/kuma/pkg/xds/hooks"
	"github.com/kumahq/kuma/pkg/xds/server"
)

var log = core.Log.WithName("api-server")

var requestsUUID = uuid.NewString()

type ApiServer struct {
	mux        *http.ServeMux
	config     api_server.ApiServerConfig
	httpReady  atomic.Bool
	httpsReady atomic.Bool
}

func (a *ApiServer) NeedLeaderElection() bool {
	return false
}

func (a *ApiServer) Address() string {
	return net.JoinHostPort(a.config.HTTP.Interface, strconv.FormatUint(uint64(a.config.HTTP.Port), 10))
}

func (a *ApiServer) Config() api_server.ApiServerConfig {
	return a.config
}

func init() {
	// turn off escape & character so the link in "next" fields for resources is user friendly
	restful.NewEncoder = func(w io.Writer) *json.Encoder {
		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)
		return encoder
	}
	restful.MarshalIndent = func(v interface{}, prefix, indent string) ([]byte, error) {
		var buf bytes.Buffer
		encoder := restful.NewEncoder(&buf)
		encoder.SetIndent(prefix, indent)
		if err := encoder.Encode(v); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
}

func NewApiServer(
	rt runtime.Runtime,
	meshContextBuilder xds_context.MeshContextBuilder,
	defs []model.ResourceTypeDescriptor,
	cfg *kuma_cp.Config,
	xdsHooks []hooks.ResourceSetHook,
) (*ApiServer, error) {
	serverConfig := cfg.ApiServer
	container := restful.NewContainer()

	promMiddleware := middleware.New(middleware.Config{
		Recorder: http_prometheus.NewRecorder(http_prometheus.Config{
			Registry: rt.Metrics(),
			Prefix:   "api_server",
		}),
	})
	container.Filter(util_prometheus.MetricsHandler("", promMiddleware))

	// NOTE: This must come before any filters that make HTTP calls
	container.Filter(otelrestful.OTelFilter("api-server"))

	if cfg.ApiServer.Authn.LocalhostIsAdmin {
		container.Filter(authn.LocalhostAuthenticator)
	}
	container.Filter(rt.APIServerAuthenticator())

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{restful.HEADER_AccessControlAllowOrigin},
		AllowedDomains: serverConfig.CorsAllowedDomains,
		Container:      container,
	}
	container.Filter(cors.Filter)

	// We create a WebService and set up resources endpoints and index endpoint instead of creating WebService
	// for every resource like /meshes/{mesh}/traffic-permissions, /meshes/{mesh}/traffic-log etc.
	// because go-restful detects it as a clash (you cannot register 2 WebServices with path /meshes/)
	ws := new(restful.WebService)
	ws.
		Path(cfg.ApiServer.BasePath).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	addResourcesEndpoints(
		ws,
		defs,
		rt.ResourceManager(),
		cfg,
		rt.Access().ResourceAccess,
		rt.GlobalInsightService(),
		meshContextBuilder,
		xdsHooks,
	)
	addPoliciesWsEndpoints(ws, cfg.IsFederatedZoneCP(), cfg.ApiServer.ReadOnly, defs)
	addInspectEndpoints(ws, cfg, meshContextBuilder, rt.ResourceManager())
	addInspectEnvoyAdminEndpoints(ws, cfg, rt.ResourceManager(), rt.Access().EnvoyAdminAccess, rt.EnvoyAdminClient())
	addInspectMeshServiceEndpoints(ws, rt.ResourceManager(), rt.Access().ResourceAccess)
	addZoneEndpoints(ws, rt.ResourceManager())
	guiUrl := ""
	if cfg.ApiServer.GUI.Enabled && !cfg.IsFederatedZoneCP() {
		guiUrl = cfg.ApiServer.GUI.BasePath
		if cfg.ApiServer.GUI.RootUrl != "" {
			guiUrl = cfg.ApiServer.GUI.RootUrl
		}
	}
	apiUrl := cfg.ApiServer.BasePath
	if cfg.ApiServer.RootUrl != "" {
		apiUrl = cfg.ApiServer.RootUrl
	}

	err := addConfigEndpoints(ws, rt.Access().ControlPlaneMetadataAccess, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "could not create configuration webservice")
	}
	if err := addIndexWsEndpoints(ws, rt.GetInstanceId, rt.GetClusterId, guiUrl); err != nil {
		return nil, errors.Wrap(err, "could not create index webservice")
	}
	addWhoamiEndpoints(ws)

	ws.SetDynamicRoutes(true)
	if err := rt.APIWebServiceCustomize()(ws); err != nil {
		return nil, errors.Wrap(err, "couldn't customize webservice")
	}

	container.Add(ws)

	path := cfg.ApiServer.BasePath
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	container.Add(tokens_server.NewWebservice(
		path+"tokens",
		rt.TokenIssuers().DataplaneToken,
		rt.TokenIssuers().ZoneToken,
		rt.Access().DataplaneTokenAccess,
		rt.Access().ZoneTokenAccess,
	))
	guiPath := cfg.ApiServer.GUI.BasePath
	if !strings.HasSuffix(guiPath, "/") {
		guiPath += "/"
	}
	basePath := guiPath
	if cfg.ApiServer.GUI.RootUrl != "" {
		u, err := url.Parse(cfg.ApiServer.GUI.RootUrl)
		if err != nil {
			return nil, errors.New("Gui.RootUrl is not a valid url")
		}
		basePath = u.Path
	}
	basePath = strings.TrimSuffix(basePath, "/")
	guiHandler, err := NewGuiHandler(guiPath, !cfg.ApiServer.GUI.Enabled || cfg.IsFederatedZoneCP(), GuiConfig{
		BaseGuiPath: basePath,
		ApiUrl:      apiUrl,
		Version:     version.Build.Version,
		Product:     version.Product,
		BasedOnKuma: version.Build.BasedOnKuma,
		Mode:        cfg.Mode,
		Environment: cfg.Environment,
		StoreType:   cfg.Store.Type,
		ReadOnly:    cfg.ApiServer.ReadOnly,
	})
	if err != nil {
		return nil, err
	}
	container.Handle(guiPath, guiHandler)

	newApiServer := &ApiServer{
		mux:    container.ServeMux,
		config: *serverConfig,
	}

	container.Filter(func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		request.Request = request.Request.WithContext(logr.NewContext(
			request.Request.Context(),
			kuma_log.AddFieldsFromCtx(
				core.Log.WithName("rest"),
				request.Request.Context(),
				rt.Extensions(),
			),
		))

		chain.ProcessFilter(request, response)
	})

	rt.APIInstaller().Install(container)

	return newApiServer, nil
}

func addResourcesEndpoints(
	ws *restful.WebService,
	defs []model.ResourceTypeDescriptor,
	resManager manager.ResourceManager,
	cfg *kuma_cp.Config,
	resourceAccess resources_access.ResourceAccess,
	globalInsightService globalinsight.GlobalInsightService,
	meshContextBuilder xds_context.MeshContextBuilder,
	xdsHooks []hooks.ResourceSetHook,
) {
	ws.Filter(HARFilter)
	globalInsightsEndpoints := globalInsightsEndpoints{
		resManager:     resManager,
		resourceAccess: resourceAccess,
	}
	globalInsightsEndpoints.addEndpoint(ws)

	globalInsightEndpoint := globalInsightEndpoint{
		globalInsightService: globalInsightService,
	}
	globalInsightEndpoint.addEndpoint(ws)

	var k8sMapper k8s.ResourceMapperFunc
	var k8sSecretMapper k8s.ResourceMapperFunc
	switch cfg.Store.Type {
	case config_store.KubernetesStore:
		k8sMapper = k8s.NewKubernetesMapper(k8s.NewSimpleKubeFactory())
		k8sSecretMapper = secrets_k8s.NewKubernetesMapper()
	default:
		k8sMapper = k8s.NewInferenceMapper(cfg.Store.Kubernetes.SystemNamespace, k8s.NewSimpleKubeFactory())
		k8sSecretMapper = secrets_k8s.NewInferenceMapper(cfg.Store.Kubernetes.SystemNamespace)
	}
	for _, definition := range defs {
		defType := definition.Name
		if ShouldBeReadOnly(definition.KDSFlags, cfg) {
			definition.ReadOnly = true
		}
		endpoints := resourceEndpoints{
			k8sMapper:                    k8sMapper,
			mode:                         cfg.Mode,
			federatedZone:                cfg.IsFederatedZoneCP(),
			resManager:                   resManager,
			descriptor:                   definition,
			resourceAccess:               resourceAccess,
			filter:                       filters.Resource(definition),
			meshContextBuilder:           meshContextBuilder,
			disableOriginLabelValidation: cfg.Multizone.Zone.DisableOriginLabelValidation,
			xdsHooks:                     xdsHooks,
			systemNamespace:              cfg.Store.Kubernetes.SystemNamespace,
			isK8s:                        cfg.Environment == config_core.KubernetesEnvironment,
		}
		if cfg.Mode == config_core.Zone && cfg.Multizone != nil && cfg.Multizone.Zone != nil {
			endpoints.zoneName = cfg.Multizone.Zone.Name
		}
		if defType == system.SecretType || defType == system.GlobalSecretType {
			endpoints.k8sMapper = k8sSecretMapper
		}
		switch defType {
		case mesh.ServiceInsightType:
			// ServiceInsight is a bit different
			ep := serviceInsightEndpoints{
				resourceEndpoints: endpoints,
				addressPortGenerator: func(svc string) string {
					return fmt.Sprintf("%s.%s:%d", svc, cfg.DNSServer.Domain, cfg.DNSServer.ServiceVipPort)
				},
			}
			ep.addCreateOrUpdateEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
			ep.addDeleteEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
			ep.addFindEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
			ep.addListEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
			ep.addListEndpoint(ws, "/"+definition.WsPath) // listing all resources in all meshes
		default:
			switch definition.Scope {
			case model.ScopeMesh:
				endpoints.addCreateOrUpdateEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
				endpoints.addDeleteEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
				endpoints.addFindEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
				endpoints.addListEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
				endpoints.addListEndpoint(ws, "/"+definition.WsPath) // listing all resources in all meshes
			case model.ScopeGlobal:
				endpoints.addCreateOrUpdateEndpoint(ws, "/"+definition.WsPath)
				endpoints.addDeleteEndpoint(ws, "/"+definition.WsPath)
				endpoints.addFindEndpoint(ws, "/"+definition.WsPath)
				endpoints.addListEndpoint(ws, "/"+definition.WsPath)
			}
		}
	}
}

func ShouldBeReadOnly(kdsFlag model.KDSFlagType, cfg *kuma_cp.Config) bool {
	if cfg.ApiServer.ReadOnly {
		return true
	}
	if kdsFlag == model.KDSDisabledFlag {
		return false
	}
	if cfg.Mode == config_core.Global && !kdsFlag.Has(model.GlobalToAllZonesFlag) {
		return true
	}
	if cfg.IsFederatedZoneCP() && !kdsFlag.Has(model.ZoneToGlobalFlag) {
		return true
	}
	return false
}

func (a *ApiServer) Ready() bool {
	return a.httpReady.Load() && a.httpsReady.Load()
}

func (a *ApiServer) Start(stop <-chan struct{}) error {
	errChan := make(chan error)

	var httpServer, httpsServer *http.Server
	if a.config.HTTP.Enabled {
		httpServer = &http.Server{
			ReadHeaderTimeout: time.Second,
			Addr:              net.JoinHostPort(a.config.HTTP.Interface, strconv.FormatUint(uint64(a.config.HTTP.Port), 10)),
			Handler:           a.mux,
			ErrorLog:          adapter.ToStd(log),
		}
		if err := kuma_srv.StartServer(log, httpServer, &a.httpReady, errChan); err != nil {
			return err
		}
	} else {
		a.httpReady.Store(true)
	}
	if a.config.HTTPS.Enabled {
		tlsConfig, err := configureTLS(a.config)
		if err != nil {
			return err
		}
		httpsServer = &http.Server{
			ReadHeaderTimeout: time.Second,
			Addr:              net.JoinHostPort(a.config.HTTPS.Interface, strconv.FormatUint(uint64(a.config.HTTPS.Port), 10)),
			Handler:           a.mux,
			TLSConfig:         tlsConfig,
			ErrorLog:          adapter.ToStd(log),
		}
		if err := kuma_srv.StartServer(log, httpsServer, &a.httpsReady, errChan); err != nil {
			return err
		}
	} else {
		a.httpsReady.Store(true)
	}
	select {
	case <-stop:
		log.Info("stopping down API Server")
		a.httpReady.Store(false)
		a.httpsReady.Store(false)
		if httpServer != nil {
			return httpServer.Shutdown(context.Background())
		}
		if httpsServer != nil {
			return httpsServer.Shutdown(context.Background())
		}
	case err := <-errChan:
		return err
	}
	return nil
}

func configureTLS(cfg api_server.ApiServerConfig) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(cfg.HTTPS.TlsCertFile, cfg.HTTPS.TlsKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load TLS certificate")
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12, // to pass gosec (in practice it's always set after.
	}
	tlsConfig.MinVersion, err = config_types.TLSVersion(cfg.HTTPS.TlsMinVersion)
	if err != nil {
		return nil, err
	}
	tlsConfig.CipherSuites, err = config_types.TLSCiphers(cfg.HTTPS.TlsCipherSuites)
	if err != nil {
		return nil, err
	}
	clientCertPool := x509.NewCertPool()
	if cfg.Auth.ClientCertsDir != "" {
		log.Info("loading client certificates")
		files, err := os.ReadDir(cfg.Auth.ClientCertsDir)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if !strings.HasSuffix(file.Name(), ".pem") && !strings.HasSuffix(file.Name(), ".crt") {
				log.Info("skipping file without .pem or .crt extension", "file", file.Name())
				continue
			}
			log.Info("adding client certificate", "file", file.Name())
			path := filepath.Join(cfg.Auth.ClientCertsDir, file.Name())
			caCert, err := os.ReadFile(path)
			if err != nil {
				return nil, errors.Wrapf(err, "could not read certificate %q", path)
			}
			if !clientCertPool.AppendCertsFromPEM(caCert) {
				return nil, errors.Errorf("failed to load PEM client certificate from %q", path)
			}
		}
	}
	if cfg.HTTPS.TlsCaFile != "" {
		file, err := os.ReadFile(cfg.HTTPS.TlsCaFile)
		if err != nil {
			return nil, err
		}
		if !clientCertPool.AppendCertsFromPEM(file) {
			return nil, errors.Errorf("failed to load PEM client certificate from %q", cfg.HTTPS.TlsCaFile)
		}
	}

	tlsConfig.ClientCAs = clientCertPool
	if cfg.HTTPS.RequireClientCert {
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	} else if cfg.Authn.Type == certs.PluginName {
		tlsConfig.ClientAuth = tls.VerifyClientCertIfGiven // client certs are required only for some endpoints when using admin client cert
	}
	return tlsConfig, nil
}

func SetupServer(rt runtime.Runtime) error {
	cfg := rt.Config()
	apiServer, err := NewApiServer(
		rt,
		xds_context.NewMeshContextBuilder(
			rt.ResourceManager(),
			server.MeshResourceTypes(),
			net.LookupIP,
			cfg.Multizone.Zone.Name,
			vips.NewPersistence(rt.ResourceManager(), rt.ConfigManager(), cfg.Experimental.UseTagFirstVirtualOutboundModel),
			cfg.DNSServer.Domain,
			cfg.DNSServer.ServiceVipPort,
			xds_context.AnyToAnyReachableServicesGraphBuilder,
		),
		registry.Global().ObjectDescriptors(model.HasWsEnabled()),
		&cfg,
		rt.XDS().Hooks.ResourceSetHooks(),
	)
	if err != nil {
		return err
	}
	return rt.Add(apiServer)
}

func HARFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	start := time.Now()

	// Capture the request body
	var requestBody bytes.Buffer
	var bodyContents []byte
	if request.Request.Body != nil {
		log.Info("reading post data")
		tee := io.TeeReader(request.Request.Body, &requestBody)
		request.Request.Body = io.NopCloser(&requestBody)
		var err error
		bodyContents, err = io.ReadAll(tee) // Read the request body into the buffer
		if err != nil {
			log.Error(err, "error reading buffer")
		}
	}

	// Capture response
	recorder := &ResponseRecorder{
		ResponseWriter: response.ResponseWriter,
		Body:           &bytes.Buffer{},
	}
	response.ResponseWriter = recorder

	// Proceed with the next filter/handler
	chain.ProcessFilter(request, response)

	// Construct HAR entry
	harEntry := hargo.Entry{
		StartedDateTime: start.Format(time.RFC3339),
		Time:            float32(time.Since(start).Milliseconds()),
		Request: hargo.Request{
			Method:      request.Request.Method,
			URL:         "http://" + request.Request.Host + request.Request.URL.String(),
			HTTPVersion: request.Request.Proto,
			Headers:     headersToHAR(request.Request.Header),
			QueryString: queryToHAR(request.Request.URL.Query()),
			PostData:    postDataToHAR(bodyContents, request.Request.Header.Get("Content-Type")),
			HeaderSize: -1,
			BodySize:    int(requestBody.Len()),
		},
		Response: hargo.Response{
			Status:      recorder.StatusCode,
			StatusText:  http.StatusText(recorder.StatusCode),
			HTTPVersion: request.Request.Proto,
			Headers:     headersToHAR(recorder.Header()),
			Content: hargo.Content{
				MimeType: recorder.Header().Get("Content-Type"),
				Text:     recorder.Body.String(),
			},
			HeadersSize: -1,
			BodySize:    int(recorder.Body.Len()),
		},
	}

	saveHAR(harEntry)
}

// ResponseRecorder is used to capture the response data.
type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       *bytes.Buffer
}

func (r *ResponseRecorder) WriteHeader(code int) {
	r.StatusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *ResponseRecorder) Write(data []byte) (int, error) {
	r.Body.Write(data)
	return r.ResponseWriter.Write(data)
}

// Save HAR entry to file
func saveHAR(entry hargo.Entry) {
	harFile := "requests-" + requestsUUID + ".har"

	// If the HAR file doesn't exist, create a new HAR log
	var log hargo.Har
	if _, err := os.Stat(harFile); os.IsNotExist(err) {
		log = hargo.Har{
			Log: hargo.Log{
				Version: "1.2",
				Creator: hargo.Creator{Name: "HARFilter", Version: "1.0"},
				Entries: []hargo.Entry{},
			},
		}
	} else {
		// Load existing HAR file
		file, err := os.Open(harFile)
		if err != nil {
			fmt.Println("Error opening HAR file:", err)
			return
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&log); err != nil {
			fmt.Println("Error decoding HAR file:", err)
			return
		}
	}

	// Append the new entry
	log.Log.Entries = append(log.Log.Entries, entry)

	// Write back to the HAR file
	file, err := os.Create(harFile)
	if err != nil {
		fmt.Println("Error creating HAR file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&log)
	if err != nil {
		fmt.Println("Error encoding HAR file:", err)
	}
}

// Convert headers to HAR format
func headersToHAR(headers http.Header) []hargo.NVP {
	var harHeaders []hargo.NVP
	for name, values := range headers {
		for _, value := range values {
			harHeaders = append(harHeaders, hargo.NVP{Name: name, Value: value})
		}
	}
	return harHeaders
}

// Convert query parameters to HAR format
func queryToHAR(query map[string][]string) []hargo.NVP {
	var harQuery = []hargo.NVP{}
	for name, values := range query {
		for _, value := range values {
			harQuery = append(harQuery, hargo.NVP{Name: name, Value: value})
		}
	}
	return harQuery
}

// Convert POST data to HAR format
func postDataToHAR(body []byte, contentType string) hargo.PostData {
	var postData hargo.PostData

	if len(body) == 0 {
		log.Info("Post data empty")
		return postData // Return empty PostData if body is empty
	}
	log.Info("Post data received")

	// Try to unmarshal the body as JSON first
	var jsonData interface{}
	err := json.Unmarshal(body, &jsonData)
	if err == nil {
		// If unmarshalling succeeds, store it as JSON
		postData = hargo.PostData{
			MimeType: "application/json",
			Text:     string(body),
		}
		return postData
	} else {
		log.Error(err, "error unmarshaling json", "value", string(body))
	}

	// If JSON unmarshalling fails, check if it's application/x-www-form-urlencoded
	if strings.Contains(contentType, "application/x-www-form-urlencoded") {
		values, err := url.ParseQuery(string(body))
		if err == nil {
			for key, vals := range values {
				for _, val := range vals {
					postData.Params = append(postData.Params, hargo.PostParam{Name: key, Value: val})
				}
			}
		}
		postData.MimeType = contentType
		return postData
	}

	// Fallback to raw text for other content types
	postData = hargo.PostData{
		MimeType: contentType,
		Text:     string(body),
	}

	return postData
}
