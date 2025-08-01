package v1alpha1

import (
	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_resource "github.com/envoyproxy/go-control-plane/pkg/resource/v3"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	"github.com/kumahq/kuma/pkg/core/kri"
	core_plugins "github.com/kumahq/kuma/pkg/core/plugins"
	"github.com/kumahq/kuma/pkg/core/resources/apis/core/destinationname"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	meshexternalservice_api "github.com/kumahq/kuma/pkg/core/resources/apis/meshexternalservice/api/v1alpha1"
	core_xds "github.com/kumahq/kuma/pkg/core/xds"
	xds_types "github.com/kumahq/kuma/pkg/core/xds/types"
	"github.com/kumahq/kuma/pkg/plugins/policies/core/matchers"
	core_rules "github.com/kumahq/kuma/pkg/plugins/policies/core/rules"
	"github.com/kumahq/kuma/pkg/plugins/policies/core/rules/outbound"
	"github.com/kumahq/kuma/pkg/plugins/policies/core/rules/subsetutils"
	policies_xds "github.com/kumahq/kuma/pkg/plugins/policies/core/xds"
	api "github.com/kumahq/kuma/pkg/plugins/policies/meshhealthcheck/api/v1alpha1"
	plugin_xds "github.com/kumahq/kuma/pkg/plugins/policies/meshhealthcheck/plugin/xds"
	gateway_plugin "github.com/kumahq/kuma/pkg/plugins/runtime/gateway"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
)

var _ core_plugins.EgressPolicyPlugin = &plugin{}

type plugin struct{}

func NewPlugin() core_plugins.Plugin {
	return &plugin{}
}

func (p plugin) MatchedPolicies(dataplane *core_mesh.DataplaneResource, resources xds_context.Resources, opts ...core_plugins.MatchedPoliciesOption) (core_xds.TypedMatchingPolicies, error) {
	return matchers.MatchedPolicies(api.MeshHealthCheckType, dataplane, resources, opts...)
}

func (p plugin) EgressMatchedPolicies(tags map[string]string, resources xds_context.Resources, opts ...core_plugins.MatchedPoliciesOption) (core_xds.TypedMatchingPolicies, error) {
	return matchers.EgressMatchedPolicies(api.MeshHealthCheckType, tags, resources, opts...)
}

func (p plugin) Apply(rs *core_xds.ResourceSet, ctx xds_context.Context, proxy *core_xds.Proxy) error {
	if proxy.ZoneEgressProxy != nil {
		return applyToEgressRealResources(rs, proxy)
	}
	policies, ok := proxy.Policies.Dynamic[api.MeshHealthCheckType]
	if !ok {
		return nil
	}

	clusters := policies_xds.GatherClusters(rs)

	if err := applyToOutbounds(policies.ToRules, clusters.Outbound, clusters.OutboundSplit, proxy.Outbounds, proxy.Dataplane, ctx.Mesh); err != nil {
		return err
	}

	if err := applyToGateways(ctx.Mesh, proxy, rs, policies.GatewayRules, clusters.Gateway); err != nil {
		return err
	}

	if err := applyToRealResources(rs, policies.ToRules.ResourceRules, ctx.Mesh, proxy.Dataplane.Spec.TagSet()); err != nil {
		return err
	}

	return nil
}

func applyToOutbounds(
	rules core_rules.ToRules,
	outboundClusters map[string]*envoy_cluster.Cluster,
	outboundSplitClusters map[string][]*envoy_cluster.Cluster,
	outbounds xds_types.Outbounds,
	dataplane *core_mesh.DataplaneResource,
	meshCtx xds_context.MeshContext,
) error {
	targetedClusters := policies_xds.GatherTargetedClusters(
		outbounds,
		outboundSplitClusters,
		outboundClusters,
	)

	for cluster, serviceName := range targetedClusters {
		if err := configure(dataplane, rules.Rules, subsetutils.KumaServiceTagElement(serviceName), meshCtx.GetServiceProtocol(serviceName), cluster); err != nil {
			return err
		}
	}

	return nil
}

func applyToGateways(
	meshCtx xds_context.MeshContext,
	proxy *core_xds.Proxy,
	rs *core_xds.ResourceSet,
	gatewayRules core_rules.GatewayRules,
	gatewayClusters map[string]*envoy_cluster.Cluster,
) error {
	resourcesByOrigin := rs.IndexByOrigin(core_xds.NonMeshExternalService)

	for _, listenerInfo := range gateway_plugin.ExtractGatewayListeners(proxy) {
		for _, listenerHostname := range listenerInfo.ListenerHostnames {
			inboundListener := core_rules.NewInboundListenerHostname(
				proxy.Dataplane.Spec.GetNetworking().Address,
				listenerInfo.Listener.Port,
				listenerHostname.Hostname,
			)
			rules, ok := gatewayRules.ToRules.ByListenerAndHostname[inboundListener]
			if !ok {
				continue
			}
			for _, hostInfo := range listenerHostname.HostInfos {
				destinations := gateway_plugin.RouteDestinationsMutable(hostInfo.Entries())
				for _, dest := range destinations {
					clusterName, err := dest.Destination.DestinationClusterName(hostInfo.Host.Tags)
					if err != nil {
						continue
					}
					cluster, ok := gatewayClusters[clusterName]
					if !ok {
						continue
					}

					serviceName := dest.Destination[mesh_proto.ServiceTag]

					if err := configure(
						proxy.Dataplane,
						rules.Rules,
						subsetutils.KumaServiceTagElement(serviceName),
						toProtocol(listenerInfo.Listener.Protocol),
						cluster,
					); err != nil {
						return err
					}

					if dest.BackendRef == nil {
						continue
					}
					if realRef := dest.BackendRef.ResourceOrNil(); realRef != nil {
						resources := resourcesByOrigin[*realRef]
						if err := applyToRealResource(
							meshCtx,
							rules.ResourceRules,
							proxy.Dataplane.Spec.TagSet(),
							*realRef,
							resources,
						); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func toProtocol(p mesh_proto.MeshGateway_Listener_Protocol) core_mesh.Protocol {
	switch p {
	case mesh_proto.MeshGateway_Listener_HTTP, mesh_proto.MeshGateway_Listener_HTTPS:
		return core_mesh.ProtocolHTTP
	case mesh_proto.MeshGateway_Listener_TCP, mesh_proto.MeshGateway_Listener_TLS:
		return core_mesh.ProtocolTCP
	}
	return core_mesh.ProtocolTCP
}

func configure(
	dataplane *core_mesh.DataplaneResource,
	rules core_rules.Rules,
	element subsetutils.Element,
	protocol core_mesh.Protocol,
	cluster *envoy_cluster.Cluster,
) error {
	conf := core_rules.ComputeConf[api.Conf](rules, element)
	if conf == nil {
		return nil
	}

	configurer := plugin_xds.Configurer{
		Conf:     *conf,
		Protocol: protocol,
		Tags:     dataplane.Spec.TagSet(),
	}

	if err := configurer.Configure(cluster); err != nil {
		return err
	}
	return nil
}

func applyToEgressRealResources(rs *core_xds.ResourceSet, proxy *core_xds.Proxy) error {
	indexed := rs.IndexByOrigin()
	for _, meshResources := range proxy.ZoneEgressProxy.MeshResourcesList {
		meshExternalServices := meshResources.ListOrEmpty(meshexternalservice_api.MeshExternalServiceType)
		for _, mes := range meshExternalServices.GetItems() {
			meshExtSvc := mes.(*meshexternalservice_api.MeshExternalServiceResource)
			policies, ok := meshResources.Dynamic[destinationname.MustResolve(false, meshExtSvc, meshExtSvc.Spec.Match)]
			if !ok {
				continue
			}
			mhc, ok := policies[api.MeshHealthCheckType]
			if !ok {
				continue
			}
			for mesID, typedResources := range indexed {
				conf := mhc.ToRules.ResourceRules.Compute(mesID, meshResources)
				if conf == nil {
					continue
				}

				for typ, resources := range typedResources {
					switch typ {
					case envoy_resource.ClusterType:
						err := configureClusters(resources, conf.Conf[0].(api.Conf), mesh_proto.MultiValueTagSet{})
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

func applyToRealResource(meshCtx xds_context.MeshContext, rules outbound.ResourceRules, tagSet mesh_proto.MultiValueTagSet, uri kri.Identifier, resourcesByType core_xds.ResourcesByType) error {
	conf := rules.Compute(uri, meshCtx.Resources)
	if conf == nil {
		return nil
	}

	for typ, resources := range resourcesByType {
		switch typ {
		case envoy_resource.ClusterType:
			err := configureClusters(resources, conf.Conf[0].(api.Conf), tagSet)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func applyToRealResources(rs *core_xds.ResourceSet, rules outbound.ResourceRules, meshCtx xds_context.MeshContext, tagSet mesh_proto.MultiValueTagSet) error {
	for uri, resType := range rs.IndexByOrigin(core_xds.NonMeshExternalService) {
		if err := applyToRealResource(meshCtx, rules, tagSet, uri, resType); err != nil {
			return err
		}
	}
	return nil
}

func configureClusters(resources []*core_xds.Resource, conf api.Conf, tagSet mesh_proto.MultiValueTagSet) error {
	for _, resource := range resources {
		configurer := plugin_xds.Configurer{
			Conf:     conf,
			Protocol: resource.Protocol,
			Tags:     tagSet,
		}
		err := configurer.Configure(resource.Resource.(*envoy_cluster.Cluster))
		if err != nil {
			return err
		}
	}
	return nil
}
