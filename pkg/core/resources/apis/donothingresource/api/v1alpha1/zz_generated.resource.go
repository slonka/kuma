// Generated by tools/policy-gen.
// Run "make generate" to update this file.

// nolint:whitespace
package v1alpha1

import (
	_ "embed"
	"fmt"

	"k8s.io/kube-openapi/pkg/validation/spec"
	"sigs.k8s.io/yaml"

	"github.com/kumahq/kuma/pkg/core/resources/model"
)

//go:embed schema.yaml
var rawSchema []byte

func init() {
	var schema spec.Schema
	if err := yaml.Unmarshal(rawSchema, &schema); err != nil {
		panic(err)
	}
	rawSchema = nil
	DoNothingResourceResourceTypeDescriptor.Schema = &schema
}

const (
	DoNothingResourceType model.ResourceType = "DoNothingResource"
)

var _ model.Resource = &DoNothingResourceResource{}

type DoNothingResourceResource struct {
	Meta model.ResourceMeta
	Spec *DoNothingResource
}

func NewDoNothingResourceResource() *DoNothingResourceResource {
	return &DoNothingResourceResource{
		Spec: &DoNothingResource{},
	}
}

func (t *DoNothingResourceResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *DoNothingResourceResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *DoNothingResourceResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *DoNothingResourceResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*DoNothingResource)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &DoNothingResource{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *DoNothingResourceResource) Descriptor() model.ResourceTypeDescriptor {
	return DoNothingResourceResourceTypeDescriptor
}

func (t *DoNothingResourceResource) Validate() error {
	if v, ok := interface{}(t).(interface{ validate() error }); !ok {
		return nil
	} else {
		return v.validate()
	}
}

var _ model.ResourceList = &DoNothingResourceResourceList{}

type DoNothingResourceResourceList struct {
	Items      []*DoNothingResourceResource
	Pagination model.Pagination
}

func (l *DoNothingResourceResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *DoNothingResourceResourceList) GetItemType() model.ResourceType {
	return DoNothingResourceType
}

func (l *DoNothingResourceResourceList) NewItem() model.Resource {
	return NewDoNothingResourceResource()
}

func (l *DoNothingResourceResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*DoNothingResourceResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*DoNothingResourceResource)(nil), r)
	}
}

func (l *DoNothingResourceResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *DoNothingResourceResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var DoNothingResourceResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                DoNothingResourceType,
	Resource:            NewDoNothingResourceResource(),
	ResourceList:        &DoNothingResourceResourceList{},
	Scope:               model.ScopeMesh,
	KDSFlags:            model.GlobalToAllZonesFlag | model.ZoneToGlobalFlag,
	WsPath:              "donothingresources",
	KumactlArg:          "donothingresource",
	KumactlListArg:      "donothingresources",
	AllowToInspect:      false,
	IsPolicy:            false,
	IsExperimental:      false,
	SingularDisplayName: "Do Nothing Resource",
	PluralDisplayName:   "Do Nothing Resources",
	IsPluginOriginated:  false,
	IsTargetRefBased:    false,
	HasToTargetRef:      false,
	HasFromTargetRef:    false,
}
