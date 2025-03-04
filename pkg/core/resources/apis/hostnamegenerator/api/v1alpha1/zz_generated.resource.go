// Generated by tools/policy-gen.
// Run "make generate" to update this file.

// nolint:whitespace
package v1alpha1

import (
	_ "embed"
	"errors"
	"fmt"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	"k8s.io/kube-openapi/pkg/validation/strfmt"
	"k8s.io/kube-openapi/pkg/validation/validate"
	"sigs.k8s.io/yaml"

	"github.com/kumahq/kuma/pkg/core/resources/model"
)

//go:embed schema.yaml
var rawSchema []byte

func init() {
	var structuralSchema *schema.Structural
	var v1JsonSchemaProps *apiextensionsv1.JSONSchemaProps
	var validator *validate.SchemaValidator
	if rawSchema != nil {
		if err := yaml.Unmarshal(rawSchema, &v1JsonSchemaProps); err != nil {
			panic(err)
		}
		var jsonSchemaProps apiextensions.JSONSchemaProps
		err := apiextensionsv1.Convert_v1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(v1JsonSchemaProps, &jsonSchemaProps, nil)
		if err != nil {
			panic(err)
		}
		structuralSchema, err = schema.NewStructural(&jsonSchemaProps)
		if err != nil {
			panic(err)
		}
		schemaObject := structuralSchema.ToKubeOpenAPI()
		validator = validate.NewSchemaValidator(schemaObject, nil, "", strfmt.Default)
	}
	rawSchema = nil
	HostnameGeneratorResourceTypeDescriptor.Validator = validator
	HostnameGeneratorResourceTypeDescriptor.StructuralSchema = structuralSchema
}

const (
	HostnameGeneratorType model.ResourceType = "HostnameGenerator"
)

var _ model.Resource = &HostnameGeneratorResource{}

type HostnameGeneratorResource struct {
	Meta model.ResourceMeta
	Spec *HostnameGenerator
}

func NewHostnameGeneratorResource() *HostnameGeneratorResource {
	return &HostnameGeneratorResource{
		Spec: &HostnameGenerator{},
	}
}

func (t *HostnameGeneratorResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *HostnameGeneratorResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *HostnameGeneratorResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *HostnameGeneratorResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*HostnameGenerator)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &HostnameGenerator{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *HostnameGeneratorResource) GetStatus() model.ResourceStatus {
	return nil
}

func (t *HostnameGeneratorResource) SetStatus(_ model.ResourceStatus) error {
	return errors.New("status not supported")
}

func (t *HostnameGeneratorResource) Descriptor() model.ResourceTypeDescriptor {
	return HostnameGeneratorResourceTypeDescriptor
}

func (t *HostnameGeneratorResource) Validate() error {
	if v, ok := interface{}(t).(interface{ validate() error }); !ok {
		return nil
	} else {
		return v.validate()
	}
}

var _ model.ResourceList = &HostnameGeneratorResourceList{}

type HostnameGeneratorResourceList struct {
	Items      []*HostnameGeneratorResource
	Pagination model.Pagination
}

func (l *HostnameGeneratorResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *HostnameGeneratorResourceList) GetItemType() model.ResourceType {
	return HostnameGeneratorType
}

func (l *HostnameGeneratorResourceList) NewItem() model.Resource {
	return NewHostnameGeneratorResource()
}

func (l *HostnameGeneratorResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*HostnameGeneratorResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*HostnameGeneratorResource)(nil), r)
	}
}

func (l *HostnameGeneratorResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *HostnameGeneratorResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var HostnameGeneratorResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                         HostnameGeneratorType,
	Resource:                     NewHostnameGeneratorResource(),
	ResourceList:                 &HostnameGeneratorResourceList{},
	Scope:                        model.ScopeGlobal,
	KDSFlags:                     model.GlobalToZonesFlag | model.ZoneToGlobalFlag,
	WsPath:                       "hostnamegenerators",
	KumactlArg:                   "hostnamegenerator",
	KumactlListArg:               "hostnamegenerators",
	AllowToInspect:               false,
	IsPolicy:                     false,
	IsExperimental:               false,
	SingularDisplayName:          "Hostname Generator",
	PluralDisplayName:            "Hostname Generators",
	IsPluginOriginated:           true,
	IsTargetRefBased:             false,
	HasToTargetRef:               false,
	HasFromTargetRef:             false,
	HasRulesTargetRef:            false,
	HasStatus:                    false,
	AllowedOnSystemNamespaceOnly: true,
	IsReferenceableInTo:          false,
	ShortName:                    "hg",
	IsFromAsRules:                false,
}
