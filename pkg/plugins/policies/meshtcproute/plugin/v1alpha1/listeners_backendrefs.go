package v1alpha1

import (
	common_api "github.com/kumahq/kuma/api/common/v1alpha1"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	core_xds "github.com/kumahq/kuma/pkg/plugins/policies/core/rules"
	meshhttproute_api "github.com/kumahq/kuma/pkg/plugins/policies/meshhttproute/api/v1alpha1"
	api "github.com/kumahq/kuma/pkg/plugins/policies/meshtcproute/api/v1alpha1"
	"strings"
)

func getBackendRefs(toRulesTCP core_xds.Rules, toRulesHTTP core_xds.Rules, serviceName string, protocol core_mesh.Protocol, backendRef common_api.BackendRef) []common_api.BackendRef {
	var service core_xds.Subset
	if backendRef.Kind == common_api.MeshExternalService {
		service = core_xds.MeshExternalService(backendRef)
	} else {
		service = core_xds.MeshService(serviceName)
	}

	tcpConf := core_xds.ComputeConf[api.Rule](toRulesTCP, service)

	// If the outbounds protocol is http-like and there exists MeshHTTPRoute
	// with rule targeting the same MeshService as MeshTCPRoute, it should take
	// precedence over the latter
	switch protocol {
	case core_mesh.ProtocolHTTP, core_mesh.ProtocolHTTP2, core_mesh.ProtocolGRPC:
		// If we have an >= HTTP service, don't manage routing with
		// MeshTCPRoutes if we either don't have any MeshTCPRoutes or we have
		// MeshHTTPRoutes
		httpConf := core_xds.ComputeConf[meshhttproute_api.PolicyDefault](
			toRulesHTTP,
			service,
		)
		if tcpConf == nil || httpConf != nil {
			return nil
		}
	default:
	}

	if tcpConf != nil {
		if backendRef.Kind == common_api.MeshExternalService {
			parts := strings.Split(serviceName, ".")
			if len(parts) == 2 {
				return mapName(tcpConf.Default.BackendRefs, parts[1])
			}
		} else {
			return tcpConf.Default.BackendRefs
		}
	}
	return []common_api.BackendRef{
		backendRef,
	}
}

func mapName(refs []common_api.BackendRef, namespace string) []common_api.BackendRef {
	var newRefs []common_api.BackendRef
	for _, ref := range refs {
		newRef := ref
		newRef.Name = ref.Name + "." + namespace
		newRefs = append(newRefs, newRef)
	}

	return newRefs
}
