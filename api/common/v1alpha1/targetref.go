// +kubebuilder:object:generate=true
package v1alpha1

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	util_maps "github.com/kumahq/kuma/pkg/util/maps"
	"github.com/kumahq/kuma/pkg/util/pointer"
)

type TargetRefKind string

var (
	Mesh                 TargetRefKind = "Mesh"
	Dataplane            TargetRefKind = "Dataplane"
	MeshSubset           TargetRefKind = "MeshSubset"
	MeshGateway          TargetRefKind = "MeshGateway"
	MeshService          TargetRefKind = "MeshService"
	MeshExternalService  TargetRefKind = "MeshExternalService"
	MeshMultiZoneService TargetRefKind = "MeshMultiZoneService"
	MeshServiceSubset    TargetRefKind = "MeshServiceSubset"
	MeshHTTPRoute        TargetRefKind = "MeshHTTPRoute"
)

var order = map[TargetRefKind]int{
	Mesh:                 1,
	Dataplane:            2,
	MeshSubset:           3,
	MeshGateway:          4,
	MeshService:          5,
	MeshExternalService:  6,
	MeshMultiZoneService: 7,
	MeshServiceSubset:    8,
	MeshHTTPRoute:        9,
}

// +kubebuilder:validation:Enum=Sidecar;Gateway
type TargetRefProxyType string

var (
	Sidecar TargetRefProxyType = "Sidecar"
	Gateway TargetRefProxyType = "Gateway"
)

func (k TargetRefKind) Compare(o TargetRefKind) int {
	return order[k] - order[o]
}

func (k TargetRefKind) IsRealResource() bool {
	switch k {
	case MeshSubset, MeshServiceSubset:
		return false
	default:
		return true
	}
}

// These are the kinds that can be used in Kuma policies before support for
// actual resources (e.g., MeshExternalService, MeshMultiZoneService, and MeshService) was introduced.
func (k TargetRefKind) IsOldKind() bool {
	switch k {
	case Mesh, MeshSubset, MeshServiceSubset, MeshService, MeshGateway, MeshHTTPRoute:
		return true
	default:
		return false
	}
}

func AllTargetRefKinds() []TargetRefKind {
	keys := util_maps.AllKeys(order)
	sort.Sort(TargetRefKindSlice(keys))
	return keys
}

type TargetRefKindSlice []TargetRefKind

func (x TargetRefKindSlice) Len() int           { return len(x) }
func (x TargetRefKindSlice) Less(i, j int) bool { return string(x[i]) < string(x[j]) }
func (x TargetRefKindSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// TargetRef defines structure that allows attaching policy to various objects
type TargetRef struct {
	// This is needed to not sync policies with empty topLevelTarget ref to old zones that does not support it
	// This can be removed in 2.11.x
	UsesSyntacticSugar bool `json:"-"`

	// Kind of the referenced resource
	// +kubebuilder:validation:Enum=Mesh;MeshSubset;MeshGateway;MeshService;MeshExternalService;MeshMultiZoneService;MeshServiceSubset;MeshHTTPRoute;Dataplane
	Kind TargetRefKind `json:"kind"`
	// Name of the referenced resource. Can only be used with kinds: `MeshService`,
	// `MeshServiceSubset` and `MeshGatewayRoute`
	Name *string `json:"name,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags *map[string]string `json:"tags,omitempty"`
	// Mesh is reserved for future use to identify cross mesh resources.
	Mesh *string `json:"mesh,omitempty"`
	// ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
	// all data plane types are targeted by the policy.
	ProxyTypes *[]TargetRefProxyType `json:"proxyTypes,omitempty"`
	// Namespace specifies the namespace of target resource. If empty only resources in policy namespace
	// will be targeted.
	Namespace *string `json:"namespace,omitempty"`
	// Labels are used to select group of MeshServices that match labels. Either Labels or
	// Name and Namespace can be used.
	Labels *map[string]string `json:"labels,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
}

func (t TargetRef) CompareDataplaneKind(other TargetRef) int {
	if t.Kind != Dataplane || other.Kind != Dataplane {
		return 0
	}
	if selectsNameAndNamespace(t) && selectsLabels(other) {
		return 1
	}
	if selectsLabels(t) && selectsNameAndNamespace(other) {
		return -1
	}
	if pointer.Deref(t.SectionName) != "" && pointer.Deref(other.SectionName) == "" {
		return 1
	}
	if pointer.Deref(t.SectionName) == "" && pointer.Deref(other.SectionName) != "" {
		return -1
	}
	return 0
}

func selectsNameAndNamespace(tr TargetRef) bool {
	return pointer.Deref(tr.Name) != ""
}

func selectsLabels(tr TargetRef) bool {
	return tr.Labels != nil
}

func IncludesGateways(ref TargetRef) bool {
	isGateway := ref.Kind == MeshGateway
	isMeshKind := ref.Kind == Mesh || ref.Kind == MeshSubset
	isGatewayInProxyTypes := len(pointer.Deref(ref.ProxyTypes)) == 0 || slices.Contains(pointer.Deref(ref.ProxyTypes), Gateway)
	isGatewayCompatible := isMeshKind && isGatewayInProxyTypes
	isMeshHTTPRoute := ref.Kind == MeshHTTPRoute

	return isGateway || isGatewayCompatible || isMeshHTTPRoute
}

// BackendRef defines where to forward traffic.
type BackendRef struct {
	// +kuma:nolint // https://github.com/kumahq/kuma/issues/14107
	TargetRef `json:","`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=1
	// +kuma:nolint // https://github.com/kumahq/kuma/issues/14107
	Weight *uint `json:"weight,omitempty"`
	// Port is only supported when this ref refers to a real MeshService object
	Port *uint32 `json:"port,omitempty"`
}

func (b BackendRef) ReferencesRealObject() bool {
	switch b.Kind {
	case MeshService:
		return b.Port != nil
	case MeshServiceSubset:
		return false
	// empty targetRef should not be treated as real object
	case "":
		return false
	default:
		return true
	}
}

// MatchesHash is used to hash route matches to determine the origin resource
// for a ref
type MatchesHash string

type BackendRefHash string

// Hash returns a hash of the BackendRef
func (in BackendRef) Hash() BackendRefHash {
	keys := util_maps.SortedKeys(pointer.Deref(in.Tags))
	orderedTags := make([]string, 0, len(keys))
	for _, k := range keys {
		orderedTags = append(orderedTags, fmt.Sprintf("%s=%s", k, pointer.Deref(in.Tags)[k]))
	}

	keys = util_maps.SortedKeys(pointer.Deref(in.Labels))
	orderedLabels := make([]string, 0, len(pointer.Deref(in.Labels)))
	for _, k := range keys {
		orderedLabels = append(orderedLabels, fmt.Sprintf("%s=%s", k, pointer.Deref(in.Labels)[k]))
	}

	name := in.Name
	if in.Port != nil {
		name = pointer.To(fmt.Sprintf("%s_svc_%d", pointer.Deref(in.Name), *in.Port))
	}
	return BackendRefHash(fmt.Sprintf("%s/%s/%s/%s/%s", in.Kind, pointer.Deref(name), strings.Join(orderedTags, "/"), strings.Join(orderedLabels, "/"), pointer.Deref(in.Mesh)))
}
