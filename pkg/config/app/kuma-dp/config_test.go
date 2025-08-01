package kumadp_test

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kumahq/kuma/pkg/config"
	kuma_dp "github.com/kumahq/kuma/pkg/config/app/kuma-dp"
	"github.com/kumahq/kuma/pkg/test/matchers"
)

var _ = Describe("Config", func() {
	It("should be loadable from configuration file", func() {
		// given
		cfg := kuma_dp.Config{}

		// when
		err := config.Load(filepath.Join("testdata", "valid-config.input.yaml"), &cfg)

		// then
		Expect(err).ToNot(HaveOccurred())

		// and
		Expect(cfg.ControlPlane.URL).To(Equal("https://kuma-control-plane.internal:5682"))
		Expect(cfg.Dataplane.DrainTime.Duration).To(Equal(60 * time.Second))
	})

	Context("with modified environment variables", func() {
		var backupEnvVars []string

		BeforeEach(func() {
			backupEnvVars = os.Environ()
		})

		AfterEach(func() {
			os.Clearenv()
			for _, envVar := range backupEnvVars {
				parts := strings.SplitN(envVar, "=", 2)
				os.Setenv(parts[0], parts[1])
			}
		})

		It("should be loadable from environment variables", func() {
			// setup
			env := map[string]string{
				"KUMA_CONTROL_PLANE_URL":                                        "https://kuma-control-plane.internal:5682",
				"KUMA_CONTROL_PLANE_RETRY_BACKOFF":                              "1s",
				"KUMA_CONTROL_PLANE_RETRY_MAX_DURATION":                         "10s",
				"KUMA_CONTROL_PLANE_BOOTSTRAP_SERVER_RETRY_BACKOFF":             "2s",
				"KUMA_CONTROL_PLANE_BOOTSTRAP_SERVER_RETRY_MAX_DURATION":        "11s",
				"KUMA_DATAPLANE_MESH":                                           "demo",
				"KUMA_DATAPLANE_NAME":                                           "example",
				"KUMA_DATAPLANE_DRAIN_TIME":                                     "60s",
				"KUMA_DATAPLANE_PROXY_TYPE":                                     "ingress",
				"KUMA_READINESS_PORT":                                           "9902",
				"KUMA_DATAPLANE_RUNTIME_BINARY_PATH":                            "envoy.sh",
				"KUMA_DATAPLANE_RUNTIME_CONFIG_DIR":                             "/var/run/envoy",
				"KUMA_DATAPLANE_RUNTIME_SOCKET_DIR":                             "/var/run/envoy",
				"KUMA_DATAPLANE_RUNTIME_WORK_DIR":                               "/var/run/envoy",
				"KUMA_DATAPLANE_RUNTIME_TOKEN_PATH":                             "/tmp/token",
				"KUMA_DATAPLANE_RUNTIME_ENVOY_LOG_LEVEL":                        "trace",
				"KUMA_DATAPLANE_RUNTIME_DYNAMIC_CONFIGURATION_REFRESH_INTERVAL": "5s",
				"KUMA_DATAPLANE_RUNTIME_UNIFIED_RESOURCE_NAMING_ENABLED":        "true",
				"KUMA_DNS_ENABLED":                                              "true",
				"KUMA_DNS_CORE_DNS_PORT":                                        "5300",
				"KUMA_DNS_ENVOY_DNS_PORT":                                       "5302",
				"KUMA_DNS_CORE_DNS_BINARY_PATH":                                 "/tmp/coredns",
				"KUMA_DNS_CORE_DNS_CONFIG_TEMPLATE_PATH":                        "/tmp/Corefile",
				"KUMA_DNS_CONFIG_DIR":                                           "/var/run/dnsserver",
				"KUMA_DNS_PROMETHEUS_PORT":                                      "6001",
				"KUMA_DNS_ENABLE_LOGGING":                                       "true",
			}
			for key, value := range env {
				os.Setenv(key, value)
			}

			// given
			cfg := kuma_dp.Config{}

			// when
			err := config.Load("", &cfg)

			// then
			Expect(err).ToNot(HaveOccurred())

			// and
			Expect(cfg.ControlPlane.URL).To(Equal("https://kuma-control-plane.internal:5682"))
			Expect(cfg.ControlPlane.Retry.Backoff.Duration).To(Equal(1 * time.Second))
			Expect(cfg.ControlPlane.Retry.MaxDuration.Duration).To(Equal(10 * time.Second))
			Expect(cfg.Dataplane.Mesh).To(Equal("demo"))
			Expect(cfg.Dataplane.Name).To(Equal("example"))
			Expect(cfg.Dataplane.DrainTime.Duration).To(Equal(60 * time.Second))
			Expect(cfg.DataplaneRuntime.BinaryPath).To(Equal("envoy.sh"))
			Expect(cfg.DataplaneRuntime.ConfigDir).To(Equal("/var/run/envoy"))
			Expect(cfg.DataplaneRuntime.SocketDir).To(Equal("/var/run/envoy"))
			Expect(cfg.DataplaneRuntime.WorkDir).To(Equal("/var/run/envoy"))
			Expect(cfg.DataplaneRuntime.TokenPath).To(Equal("/tmp/token"))
			Expect(cfg.DataplaneRuntime.EnvoyLogLevel).To(Equal("trace"))
			Expect(cfg.DataplaneRuntime.DynamicConfiguration.RefreshInterval.Duration).To(Equal(5 * time.Second))
			Expect(cfg.DataplaneRuntime.UnifiedResourceNamingEnabled).To(BeTrue())
			Expect(cfg.DNS.Enabled).To(BeTrue())
			Expect(cfg.DNS.CoreDNSPort).To(Equal(uint32(5300)))
			Expect(cfg.DNS.EnvoyDNSPort).To(Equal(uint32(5302)))
			Expect(cfg.DNS.CoreDNSBinaryPath).To(Equal("/tmp/coredns"))
			Expect(cfg.DNS.CoreDNSConfigTemplatePath).To(Equal("/tmp/Corefile"))
			Expect(cfg.DNS.ConfigDir).To(Equal("/var/run/dnsserver"))
			Expect(cfg.DNS.PrometheusPort).To(Equal(uint32(6001)))
			Expect(cfg.DNS.CoreDNSLogging).To(BeTrue())
		})
	})

	It("should have consistent defaults", func() {
		// given
		cfg := kuma_dp.DefaultConfig()

		// when
		actual, err := config.ToYAML(&cfg)

		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(matchers.MatchGoldenYAML("testdata", "default-config.golden.yaml"))
	})

	It("should have validators", func() {
		// given
		cfg := kuma_dp.Config{}

		// when
		err := config.Load(filepath.Join("testdata", "invalid-config.input.yaml"), &cfg)

		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error() + "\n").To(matchers.MatchGoldenEqual("testdata", "invalid-config.golden.txt"))
	})

	It("should have validators (valid DP proxyType, invalid DP name)", func() {
		// given
		cfg := kuma_dp.Config{}

		// when
		err := config.Load(filepath.Join("testdata", "invalid-dp-name-config.input.yaml"), &cfg)

		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error() + "\n").To(matchers.MatchGoldenEqual("testdata", "invalid-dp-name-config.golden.txt"))
	})

	DescribeTable("bad validation",
		func(fn func(*kuma_dp.Config)) {
			// given
			cfg := kuma_dp.Config{}

			Expect(config.Load(filepath.Join("testdata", "valid-config.input.yaml"), &cfg)).Should(Succeed())

			// when (apply a modification)
			fn(&cfg)

			// then
			Expect(cfg.Validate()).ShouldNot(Succeed())
		},
		Entry("unsupported proxy type", func(cfg *kuma_dp.Config) {
			cfg.Dataplane.ProxyType = "gateway"
		}),
		Entry("invalid cp url", func(cfg *kuma_dp.Config) {
			cfg.ControlPlane.URL = ":333"
		}),
	)
})
