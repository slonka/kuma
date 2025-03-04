// +kubebuilder:object:generate=true
package v1alpha1

import (
	"fmt"

	hostnamegenerator_api "github.com/kumahq/kuma/pkg/core/resources/apis/hostnamegenerator/api/v1alpha1"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	meshservice_api "github.com/kumahq/kuma/pkg/core/resources/apis/meshservice/api/v1alpha1"
)

// MeshMultiZoneService allows users to create a service that spawns across multiple zones
// It aggregates existing MeshServices by labels.
// +kuma:policy:is_policy=false
// +kuma:policy:has_status=true
// +kuma:policy:is_referenceable_in_to=true
// +kuma:policy:short_name=mzsvc
// MeshMultizoneServices are only created on global
// +kuma:policy:kds_flags=model.GlobalToZonesFlag
// +kubebuilder:printcolumn:JSONPath=".status.addresses[0].hostname",name=Hostname,type=string
type MeshMultiZoneService struct {
	// Selector is a way to select multiple MeshServices
	Selector Selector `json:"selector"`
	// Ports is a list of ports from selected MeshServices
	// +kubebuilder:validation:MinItems=1
	Ports []Port `json:"ports"`
}

type Port struct {
	Name *string `json:"name,omitempty"`
	Port uint32  `json:"port"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=tcp
	AppProtocol core_mesh.Protocol `json:"appProtocol"`
}

type Selector struct {
	// MeshService selects MeshServices
	MeshService MeshServiceSelector `json:"meshService"`
}

type MeshServiceSelector struct {
	// MatchLabels matches multiple MeshServices by labels
	MatchLabels map[string]string `json:"matchLabels"`
}

type MeshMultiZoneServiceStatus struct {
	// Addresses is a list of addresses generated by HostnameGenerator
	Addresses []hostnamegenerator_api.Address `json:"addresses,omitempty"`
	// VIPs is a list of assigned Kuma VIPs.
	VIPs []meshservice_api.VIP `json:"vips,omitempty"`
	// MeshServices is a list of matched MeshServices
	MeshServices []MatchedMeshService `json:"meshServices,omitempty"`
	// Status of hostnames generator applied on this resource
	HostnameGenerators []hostnamegenerator_api.HostnameGeneratorStatus `json:"hostnameGenerators,omitempty"`
}

type MatchedMeshService struct {
	// Name is a core name of MeshService
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Zone      string `json:"zone"`
	Mesh      string `json:"mesh"`
}

func (m *MatchedMeshService) FullyQualifiedName() string {
	return fmt.Sprintf("%s/%s/%s/%s", m.Name, m.Namespace, m.Zone, m.Mesh)
}
