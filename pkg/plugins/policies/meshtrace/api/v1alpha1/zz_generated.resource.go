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
	MeshTraceResourceTypeDescriptor.Validator = validator
	MeshTraceResourceTypeDescriptor.StructuralSchema = structuralSchema
}

const (
	MeshTraceType model.ResourceType = "MeshTrace"
)

var _ model.Resource = &MeshTraceResource{}

type MeshTraceResource struct {
	Meta model.ResourceMeta
	Spec *MeshTrace
}

func NewMeshTraceResource() *MeshTraceResource {
	return &MeshTraceResource{
		Spec: &MeshTrace{},
	}
}

func (t *MeshTraceResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *MeshTraceResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *MeshTraceResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *MeshTraceResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*MeshTrace)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &MeshTrace{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *MeshTraceResource) GetStatus() model.ResourceStatus {
	return nil
}

func (t *MeshTraceResource) SetStatus(_ model.ResourceStatus) error {
	return errors.New("status not supported")
}

func (t *MeshTraceResource) Descriptor() model.ResourceTypeDescriptor {
	return MeshTraceResourceTypeDescriptor
}

func (t *MeshTraceResource) Validate() error {
	if v, ok := interface{}(t).(interface{ validate() error }); !ok {
		return nil
	} else {
		return v.validate()
	}
}

var _ model.ResourceList = &MeshTraceResourceList{}

type MeshTraceResourceList struct {
	Items      []*MeshTraceResource
	Pagination model.Pagination
}

func (l *MeshTraceResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *MeshTraceResourceList) GetItemType() model.ResourceType {
	return MeshTraceType
}

func (l *MeshTraceResourceList) NewItem() model.Resource {
	return NewMeshTraceResource()
}

func (l *MeshTraceResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*MeshTraceResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*MeshTraceResource)(nil), r)
	}
}

func (l *MeshTraceResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *MeshTraceResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var MeshTraceResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                         MeshTraceType,
	Resource:                     NewMeshTraceResource(),
	ResourceList:                 &MeshTraceResourceList{},
	Scope:                        model.ScopeMesh,
	KDSFlags:                     model.GlobalToZonesFlag | model.ZoneToGlobalFlag | model.SyncedAcrossZonesFlag,
	WsPath:                       "meshtraces",
	KumactlArg:                   "meshtrace",
	KumactlListArg:               "meshtraces",
	AllowToInspect:               true,
	IsPolicy:                     true,
	IsDestination:                false,
	IsExperimental:               false,
	SingularDisplayName:          "Mesh Trace",
	PluralDisplayName:            "Mesh Traces",
	IsPluginOriginated:           true,
	IsTargetRefBased:             true,
	HasToTargetRef:               false,
	HasFromTargetRef:             false,
	HasRulesTargetRef:            false,
	HasStatus:                    false,
	AllowedOnSystemNamespaceOnly: false,
	IsReferenceableInTo:          false,
	ShortName:                    "mtr",
	IsFromAsRules:                false,
}
