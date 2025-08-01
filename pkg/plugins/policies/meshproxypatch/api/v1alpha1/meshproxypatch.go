// +kubebuilder:object:generate=true
package v1alpha1

import (
	common_api "github.com/kumahq/kuma/api/common/v1alpha1"
)

// MeshProxyPatch is a policy that lets you patch standard Envoy configuration
// generated by Kuma Control Plane.
type MeshProxyPatch struct {
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined inplace.
	TargetRef *common_api.TargetRef `json:"targetRef,omitempty"`

	// Default is a configuration specific to the group of destinations
	// referenced in 'targetRef'.
	Default Conf `json:"default"`
}

type Conf struct {
	// AppendModifications is a list of modifications applied on the selected proxy.
	AppendModifications *[]Modification `json:"appendModifications,omitempty"`
}

type Modification struct {
	// Cluster is a modification of Envoy's Cluster resource.
	Cluster *ClusterMod `json:"cluster,omitempty"`
	// Listener is a modification of Envoy's Listener resource.
	Listener *ListenerMod `json:"listener,omitempty"`
	// NetworkFilter is a modification of Envoy Listener's filter.
	NetworkFilter *NetworkFilterMod `json:"networkFilter,omitempty"`
	// HTTPFilter is a modification of Envoy HTTP Filter
	// available in HTTP Connection Manager in a Listener resource.
	HTTPFilter *HTTPFilterMod `json:"httpFilter,omitempty"`
	// VirtualHost is a modification of Envoy's VirtualHost
	// referenced in HTTP Connection Manager in a Listener resource.
	VirtualHost *VirtualHostMod `json:"virtualHost,omitempty"`
}

// ModOperation is modification operation on Envoy's resource.
type ModOperation string

const (
	// ModOpAdd adds a new resource.
	ModOpAdd ModOperation = "Add"
	// ModOpRemove removes an existing resource.
	ModOpRemove ModOperation = "Remove"
	// ModOpPatch modifies an existing resource.
	ModOpPatch ModOperation = "Patch"
	// ModOpAddFirst adds a resources as a first resource in the list.
	// Only applicable to NetworkFilter and HTTPFilter.
	ModOpAddFirst ModOperation = "AddFirst"
	// ModOpAddLast adds a resources as a last resource in the list.
	// Only applicable to NetworkFilter and HTTPFilter.
	ModOpAddLast ModOperation = "AddLast"
	// ModOpAddBefore adds a resources before existing resource in the list.
	// If existing resource does not exist, the new resource won't be added.
	// Only applicable to NetworkFilter and HTTPFilter.
	ModOpAddBefore ModOperation = "AddBefore"
	// ModOpAddAfter adds a resources after existing resource in the list.
	// If existing resource does not exist, the new resource won't be added.
	// Only applicable to NetworkFilter and HTTPFilter.
	ModOpAddAfter ModOperation = "AddAfter"
)

// ClusterMod is a modification of Envoy's Cluster resource.
type ClusterMod struct {
	// Match is a set of conditions that have to be matched for modification operation to happen.
	Match *ClusterMatch `json:"match,omitempty"`
	// Operation to execute on matched cluster.
	// +kubebuilder:validation:Enum=Add;Remove;Patch
	Operation ModOperation `json:"operation"`
	// Value of xDS resource in YAML format to add or patch.
	Value *string `json:"value,omitempty"`
	// JsonPatches specifies list of jsonpatches to apply to on Envoy's Cluster
	// resource
	JsonPatches *[]common_api.JsonPatchBlock `json:"jsonPatches,omitempty"`
}

// ClusterMatch is a set of conditions on cluster resource.
type ClusterMatch struct {
	// Origin is the name of the component or plugin that generated the resource.
	//
	// Here is the list of well-known origins:
	// inbound - resources generated for handling incoming traffic.
	// outbound - resources generated for handling outgoing traffic.
	// transparent - resources generated for transparent proxy functionality.
	// prometheus - resources generated when Prometheus metrics are enabled.
	// direct-access - resources generated for Direct Access functionality.
	// ingress - resources generated for Zone Ingress.
	// egress - resources generated for Zone Egress.
	// gateway - resources generated for MeshGateway.
	//
	// The list is not complete, because policy plugins can introduce new resources.
	// For example MeshTrace plugin can create Cluster with "mesh-trace" origin.
	Origin *string `json:"origin,omitempty"`
	// Name of the cluster to match.
	Name *string `json:"name,omitempty"`
}

// ListenerMod is a modification of Envoy's Listener resource.
type ListenerMod struct {
	// Match is a set of conditions that have to be matched for modification operation to happen.
	Match *ListenerMatch `json:"match,omitempty"`
	// Operation to execute on matched listener.
	// +kubebuilder:validation:Enum=Add;Remove;Patch
	Operation ModOperation `json:"operation"`
	// Value of xDS resource in YAML format to add or patch.
	Value *string `json:"value,omitempty"`
	// JsonPatches specifies list of jsonpatches to apply to on Envoy's Listener
	// resource
	JsonPatches *[]common_api.JsonPatchBlock `json:"jsonPatches,omitempty"`
}

// ListenerMatch is a set of conditions that have to be matched for modification operation to happen.
type ListenerMatch struct {
	// Origin is the name of the component or plugin that generated the resource.
	//
	// Here is the list of well-known origins:
	// inbound - resources generated for handling incoming traffic.
	// outbound - resources generated for handling outgoing traffic.
	// transparent - resources generated for transparent proxy functionality.
	// prometheus - resources generated when Prometheus metrics are enabled.
	// direct-access - resources generated for Direct Access functionality.
	// ingress - resources generated for Zone Ingress.
	// egress - resources generated for Zone Egress.
	// gateway - resources generated for MeshGateway.
	//
	// The list is not complete, because policy plugins can introduce new resources.
	// For example MeshTrace plugin can create Cluster with "mesh-trace" origin.
	Origin *string `json:"origin,omitempty"`
	// Name of the listener to match.
	Name *string `json:"name,omitempty"`
	// Tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]
	Tags *map[string]string `json:"tags,omitempty"`
}

// NetworkFilterMod is a modification of Envoy Listener's filter.
type NetworkFilterMod struct {
	// Match is a set of conditions that have to be matched for modification operation to happen.
	Match *NetworkFilterMatch `json:"match,omitempty"`
	// Operation to execute on matched listener.
	// +kubebuilder:validation:Enum=Remove;Patch;AddFirst;AddBefore;AddAfter;AddLast
	Operation ModOperation `json:"operation"`
	// Value of xDS resource in YAML format to add or patch.
	Value *string `json:"value,omitempty"`
	// JsonPatches specifies list of jsonpatches to apply to on Envoy Listener's
	// filter.
	JsonPatches *[]common_api.JsonPatchBlock `json:"jsonPatches,omitempty"`
}

// NetworkFilterMatch is a set of conditions that have to be matched for modification operation to happen.
type NetworkFilterMatch struct {
	// Origin is the name of the component or plugin that generated the resource.
	//
	// Here is the list of well-known origins:
	// inbound - resources generated for handling incoming traffic.
	// outbound - resources generated for handling outgoing traffic.
	// transparent - resources generated for transparent proxy functionality.
	// prometheus - resources generated when Prometheus metrics are enabled.
	// direct-access - resources generated for Direct Access functionality.
	// ingress - resources generated for Zone Ingress.
	// egress - resources generated for Zone Egress.
	// gateway - resources generated for MeshGateway.
	//
	// The list is not complete, because policy plugins can introduce new resources.
	// For example MeshTrace plugin can create Cluster with "mesh-trace" origin.
	Origin *string `json:"origin,omitempty"`
	// Name of the network filter. For example "envoy.filters.network.ratelimit"
	Name *string `json:"name,omitempty"`
	// Name of the listener to match.
	ListenerName *string `json:"listenerName,omitempty"`
	// Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]
	ListenerTags *map[string]string `json:"listenerTags,omitempty"`
}

// HTTPFilterMod is a modification of Envoy HTTP Filter
// available in HTTP Connection Manager in a Listener resource.
type HTTPFilterMod struct {
	// Match is a set of conditions that have to be matched for modification operation to happen.
	Match *HTTPFilterMatch `json:"match,omitempty"`
	// Operation to execute on matched listener.
	// +kubebuilder:validation:Enum=Remove;Patch;AddFirst;AddBefore;AddAfter;AddLast
	Operation ModOperation `json:"operation"`
	// Value of xDS resource in YAML format to add or patch.
	Value *string `json:"value,omitempty"`
	// JsonPatches specifies list of jsonpatches to apply to on Envoy's
	// HTTP Filter available in HTTP Connection Manager in a Listener resource.
	JsonPatches *[]common_api.JsonPatchBlock `json:"jsonPatches,omitempty"`
}

// HTTPFilterMatch is a set of conditions that have to be matched for modification operation to happen.
type HTTPFilterMatch struct {
	// Origin is the name of the component or plugin that generated the resource.
	//
	// Here is the list of well-known origins:
	// inbound - resources generated for handling incoming traffic.
	// outbound - resources generated for handling outgoing traffic.
	// transparent - resources generated for transparent proxy functionality.
	// prometheus - resources generated when Prometheus metrics are enabled.
	// direct-access - resources generated for Direct Access functionality.
	// ingress - resources generated for Zone Ingress.
	// egress - resources generated for Zone Egress.
	// gateway - resources generated for MeshGateway.
	//
	// The list is not complete, because policy plugins can introduce new resources.
	// For example MeshTrace plugin can create Cluster with "mesh-trace" origin.
	Origin *string `json:"origin,omitempty"`
	// Name of the HTTP filter. For example "envoy.filters.http.local_ratelimit"
	Name *string `json:"name,omitempty"`
	// Name of the listener to match.
	ListenerName *string `json:"listenerName,omitempty"`
	// Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]
	ListenerTags *map[string]string `json:"listenerTags,omitempty"`
}

// VirtualHostMod is a modification of Envoy's VirtualHost
// referenced in HTTP Connection Manager in a Listener resource.
type VirtualHostMod struct {
	// Match is a set of conditions that have to be matched for modification operation to happen.
	// +kuma:nolint // https://github.com/kumahq/kuma/issues/14107
	Match *VirtualHostMatch `json:"match"`
	// Operation to execute on matched listener.
	// +kubebuilder:validation:Enum=Add;Remove;Patch
	Operation ModOperation `json:"operation"`
	// Value of xDS resource in YAML format to add or patch.
	Value *string `json:"value,omitempty"`
	// JsonPatches specifies list of jsonpatches to apply to on Envoy's
	// VirtualHost resource
	JsonPatches *[]common_api.JsonPatchBlock `json:"jsonPatches,omitempty"`
}

// VirtualHostMatch is a set of conditions that have to be matched for modification operation to happen.
type VirtualHostMatch struct {
	// Origin is the name of the component or plugin that generated the resource.
	//
	// Here is the list of well-known origins:
	// inbound - resources generated for handling incoming traffic.
	// outbound - resources generated for handling outgoing traffic.
	// transparent - resources generated for transparent proxy functionality.
	// prometheus - resources generated when Prometheus metrics are enabled.
	// direct-access - resources generated for Direct Access functionality.
	// ingress - resources generated for Zone Ingress.
	// egress - resources generated for Zone Egress.
	// gateway - resources generated for MeshGateway.
	//
	// The list is not complete, because policy plugins can introduce new resources.
	// For example MeshTrace plugin can create Cluster with "mesh-trace" origin.
	Origin *string `json:"origin,omitempty"`
	// Name of the VirtualHost to match.
	Name *string `json:"name,omitempty"`
	// Name of the RouteConfiguration resource to match.
	RouteConfigurationName *string `json:"routeConfigurationName,omitempty"`
}
