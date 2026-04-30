package meshproxy

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/sync/errgroup"

	mesh_proto "github.com/kumahq/kuma/v2/api/mesh/v1alpha1"
	meshidentity_api "github.com/kumahq/kuma/v2/pkg/core/resources/apis/meshidentity/api/v1alpha1"
	meshtrust_api "github.com/kumahq/kuma/v2/pkg/core/resources/apis/meshtrust/api/v1alpha1"
	"github.com/kumahq/kuma/v2/pkg/core/resources/model"
	"github.com/kumahq/kuma/v2/pkg/core/resources/model/rest"
	"github.com/kumahq/kuma/v2/pkg/kds/hash"
	"github.com/kumahq/kuma/v2/pkg/test/resources/builders"
	. "github.com/kumahq/kuma/v2/test/framework"
	"github.com/kumahq/kuma/v2/test/framework/client"
	"github.com/kumahq/kuma/v2/test/framework/deployments/democlient"
	"github.com/kumahq/kuma/v2/test/framework/deployments/testserver"
	"github.com/kumahq/kuma/v2/test/framework/deployments/zoneproxy"
	"github.com/kumahq/kuma/v2/test/framework/envs/multizone"
)

func Migration() {
	namespace := "meshproxy-migration"

	type meshCase struct {
		mesh          string
		zone1Client   string
		zone2Server   string
		zone2Instance string
	}

	meshes := []meshCase{
		{
			mesh:          "meshproxymig1",
			zone1Client:   "democlient1-zone1",
			zone2Server:   "testserver1-zone2",
			zone2Instance: "testserver1-zone2",
		},
		{
			mesh:          "meshproxymig2",
			zone1Client:   "democlient2-zone1",
			zone2Server:   "testserver2-zone2",
			zone2Instance: "testserver2-zone2",
		},
	}

	setupMeshIdentity := func(meshName string) {
		GinkgoHelper()

		// Use a per-mesh identity name so the bundled provider's
		// auto-generated CA secrets ("<identity>-root-ca",
		// "<identity>-private-key") don't collide across meshes in the same
		// kuma-system namespace.
		identityName := "identity-" + meshName

		meshIdentityYAML := fmt.Sprintf(`
type: MeshIdentity
name: %[2]s
mesh: %[1]s
spec:
  selector:
    dataplane:
      matchLabels: {}
  spiffeID:
    trustDomain: "{{ .Mesh }}.{{ .Zone }}.mesh.local"
  provider:
    type: Bundled
    bundled:
      meshTrustCreation: Enabled
      insecureAllowSelfSigned: true
      certificateParameters:
        expiry: 24h
      autogenerate:
        enabled: true
`, meshName, identityName)
		Expect(NewClusterSetup().
			Install(YamlUniversal(meshIdentityYAML)).
			Setup(multizone.Global)).To(Succeed())

		hashedIdentityName := hash.HashedName(meshName, identityName)
		Expect(WaitForResource(
			meshidentity_api.MeshIdentityResourceTypeDescriptor,
			model.ResourceKey{Mesh: meshName, Name: fmt.Sprintf("%s.%s", hashedIdentityName, Config.KumaNamespace)},
			multizone.KubeZone1, multizone.KubeZone2,
		)).To(Succeed())

		getMeshTrust := func(hashValues ...string) *meshtrust_api.MeshTrust {
			var trust *meshtrust_api.MeshTrust
			Eventually(func(g Gomega) {
				out, err := multizone.Global.GetKumactlOptions().RunKumactlAndGetOutput(
					"get", "meshtrust", "-m", meshName,
					hash.HashedName(meshName, hashedIdentityName, hashValues...),
					"-ojson",
				)
				g.Expect(err).ToNot(HaveOccurred())
				r, err := rest.JSON.Unmarshal([]byte(out), meshtrust_api.MeshTrustResourceTypeDescriptor)
				g.Expect(err).ToNot(HaveOccurred())
				trust = r.GetSpec().(*meshtrust_api.MeshTrust)
			}, "2m", "1s").Should(Succeed())
			return trust
		}

		installTrustToGlobal := func(trust *meshtrust_api.MeshTrust, sourceZoneName string) {
			yaml := builders.MeshTrust().
				WithName("meshtrust-of-zone-" + sourceZoneName).
				WithMesh(meshName).
				WithCA(trust.CABundles[0].PEM.Value).
				WithTrustDomain(trust.TrustDomain).
				UniYaml()
			Expect(NewClusterSetup().
				Install(YamlUniversal(yaml)).
				Setup(multizone.Global)).To(Succeed())
		}

		trustZone1 := getMeshTrust(multizone.KubeZone1.Name(), Config.KumaNamespace)
		installTrustToGlobal(trustZone1, multizone.KubeZone1.Name())

		trustZone2 := getMeshTrust(multizone.KubeZone2.Name(), Config.KumaNamespace)
		installTrustToGlobal(trustZone2, multizone.KubeZone2.Name())
	}

	BeforeAll(func() {
		setup := NewClusterSetup()
		for _, tc := range meshes {
			setup.
				Install(Yaml(
					builders.Mesh().
						WithName(tc.mesh).
						WithMeshServicesEnabled(mesh_proto.Mesh_MeshServices_Exclusive),
				)).
				Install(MeshTrafficPermissionAllowAllUniversal(tc.mesh))
		}
		Expect(setup.Setup(multizone.Global)).To(Succeed())

		for _, tc := range meshes {
			Expect(WaitForMesh(tc.mesh, multizone.Zones())).To(Succeed())
		}

		kubeZone1Install := []InstallFunc{}
		kubeZone2Install := []InstallFunc{}
		for _, tc := range meshes {
			kubeZone1Install = append(kubeZone1Install,
				democlient.Install(
					democlient.WithNamespace(namespace),
					democlient.WithMesh(tc.mesh),
					democlient.WithName(tc.zone1Client),
				),
			)
			kubeZone2Install = append(kubeZone2Install,
				testserver.Install(
					testserver.WithNamespace(namespace),
					testserver.WithMesh(tc.mesh),
					testserver.WithName(tc.zone2Server),
					testserver.WithEchoArgs("echo", "--instance", tc.zone2Instance),
				),
			)
		}

		group := errgroup.Group{}
		NewClusterSetup().
			Install(NamespaceWithSidecarInjection(namespace)).
			Install(Parallel(kubeZone1Install...)).
			SetupInGroup(multizone.KubeZone1, &group)
		NewClusterSetup().
			Install(NamespaceWithSidecarInjection(namespace)).
			Install(Parallel(kubeZone2Install...)).
			SetupInGroup(multizone.KubeZone2, &group)

		Expect(group.Wait()).To(Succeed())

		for _, tc := range meshes {
			setupMeshIdentity(tc.mesh)
		}
	})

	AfterEachFailure(func() {
		for _, tc := range meshes {
			DebugUniversal(multizone.Global, tc.mesh)
			DebugKube(multizone.KubeZone1, tc.mesh, namespace)
			DebugKube(multizone.KubeZone2, tc.mesh, namespace)
		}
	})

	E2EAfterAll(func() {
		Expect(multizone.KubeZone1.TriggerDeleteNamespace(namespace)).To(Succeed())
		Expect(multizone.KubeZone2.TriggerDeleteNamespace(namespace)).To(Succeed())
		for _, tc := range meshes {
			Expect(multizone.Global.DeleteMesh(tc.mesh)).To(Succeed())
		}
	})

	It("should migrate traffic to mesh-scoped zone proxies without request drops for two meshes", func() {
		expectRequest := func(g Gomega, source *K8sCluster, sourceClient, targetServer, targetZone, expectedInstance string) {
			GinkgoHelper()

			response, err := client.CollectEchoResponse(
				source,
				sourceClient,
				fmt.Sprintf("http://%s.%s.svc.%s.mesh.local:80", targetServer, namespace, targetZone),
				client.FromKubernetesPod(namespace, sourceClient),
			)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(response.Instance).To(Equal(expectedInstance))
		}

		expectCrossZoneTraffic := func(g Gomega, tc meshCase) {
			GinkgoHelper()

			expectRequest(g, multizone.KubeZone1, tc.zone1Client, tc.zone2Server, multizone.KubeZone2.ZoneName(), tc.zone2Instance)
		}

		assertAllMeshesReachable := func() {
			GinkgoHelper()

			for _, tc := range meshes {
				tc := tc
				Eventually(func(g Gomega) {
					expectCrossZoneTraffic(g, tc)
				}, "60s", "1s").MustPassRepeatedly(3).Should(Succeed())
			}
		}

		deployMeshScopedZoneProxies := func(mesh string) {
			GinkgoHelper()

			const ingressPort = uint32(11001)
			const egressPort = uint32(11002)

			proxyName := fmt.Sprintf("mesh-zone-proxy-%s", mesh)
			group := errgroup.Group{}
			NewClusterSetup().
				Install(zoneproxy.Install(
					zoneproxy.WithName(proxyName),
					zoneproxy.WithNamespace(namespace),
					zoneproxy.WithMesh(mesh),
					zoneproxy.WithIngressPort(ingressPort),
					zoneproxy.WithEgressPort(egressPort),
				)).
				SetupInGroup(multizone.KubeZone1, &group)
			NewClusterSetup().
				Install(zoneproxy.Install(
					zoneproxy.WithName(proxyName),
					zoneproxy.WithNamespace(namespace),
					zoneproxy.WithMesh(mesh),
					zoneproxy.WithIngressPort(ingressPort),
					zoneproxy.WithEgressPort(egressPort),
				)).
				SetupInGroup(multizone.KubeZone2, &group)
			Expect(group.Wait()).To(Succeed())
		}

		readUpstreamRequests := func(cluster *K8sCluster, app, mesh string) float64 {
			GinkgoHelper()

			stats, err := cluster.GetEnvoyAdminTunnel(app, namespace).
				GetStats(fmt.Sprintf(".*%s.*upstream_rq_total", mesh))
			Expect(err).ToNot(HaveOccurred())

			total := 0.0
			for _, item := range stats.Stats {
				if !strings.Contains(item.Name, "upstream_rq_total") {
					continue
				}
				value, ok := item.Value.(float64)
				if !ok {
					continue
				}
				total += value
			}
			return total
		}

		assertMeshScopedProxiesUsed := func(tc meshCase) {
			GinkgoHelper()

			// The mesh in this test does not enable zone egress, so the
			// source-side egress is bypassed - traffic goes sidecar -> remote
			// zone ingress directly. Only verify that the destination zone's
			// new mesh-scoped ingress receives at least some cross-zone
			// requests once it is deployed. Both legacy and mesh-scoped
			// ingresses can coexist; the success criterion for the migration
			// is that the new mesh-scoped ingress is reachable and is
			// observed routing traffic to the local backend.
			ingressApp := fmt.Sprintf("mesh-zone-proxy-%s-ingress", tc.mesh)

			Eventually(func(g Gomega) {
				_, err := client.CollectEchoResponse(
					multizone.KubeZone1,
					tc.zone1Client,
					fmt.Sprintf("http://%s.%s.svc.%s.mesh.local:80", tc.zone2Server, namespace, multizone.KubeZone2.ZoneName()),
					client.FromKubernetesPod(namespace, tc.zone1Client),
				)
				g.Expect(err).ToNot(HaveOccurred())

				g.Expect(readUpstreamRequests(multizone.KubeZone2, ingressApp, tc.mesh)).To(BeNumerically(">", 0))
			}, "60s", "1s").Should(Succeed())
		}

		assertAllMeshesReachable()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var (
			reqErrMu sync.Mutex
			reqErr   error
		)
		recordRequestError := func(err error) {
			GinkgoHelper()

			if err == nil {
				return
			}
			reqErrMu.Lock()
			defer reqErrMu.Unlock()
			if reqErr == nil {
				reqErr = err
			}
		}
		for _, tc := range meshes {
			tc := tc
			go func() {
				defer GinkgoRecover()

				for {
					select {
					case <-ctx.Done():
						return
					default:
					}

					_, err := client.CollectEchoResponse(
						multizone.KubeZone1,
						tc.zone1Client,
						fmt.Sprintf("http://%s.%s.svc.%s.mesh.local:80", tc.zone2Server, namespace, multizone.KubeZone2.ZoneName()),
						client.FromKubernetesPod(namespace, tc.zone1Client),
					)
					recordRequestError(err)

					time.Sleep(200 * time.Millisecond)
				}
			}()
		}

		deployMeshScopedZoneProxies(meshes[0].mesh)
		assertAllMeshesReachable()
		assertMeshScopedProxiesUsed(meshes[0])

		deployMeshScopedZoneProxies(meshes[1].mesh)
		assertAllMeshesReachable()
		assertMeshScopedProxiesUsed(meshes[1])

		Consistently(func(g Gomega) {
			reqErrMu.Lock()
			defer reqErrMu.Unlock()
			g.Expect(reqErr).ToNot(HaveOccurred())
		}, "10s", "500ms").Should(Succeed())
	})
}
