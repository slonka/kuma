//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	commonv1alpha1 "github.com/kumahq/kuma/api/common/v1alpha1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	if in.Tcp != nil {
		in, out := &in.Tcp, &out.Tcp
		*out = new(TCPBackend)
		(*in).DeepCopyInto(*out)
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(FileBackend)
		(*in).DeepCopyInto(*out)
	}
	if in.OpenTelemetry != nil {
		in, out := &in.OpenTelemetry, &out.OpenTelemetry
		*out = new(OtelBackend)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Conf) DeepCopyInto(out *Conf) {
	*out = *in
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = new([]Backend)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Backend, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conf.
func (in *Conf) DeepCopy() *Conf {
	if in == nil {
		return nil
	}
	out := new(Conf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileBackend) DeepCopyInto(out *FileBackend) {
	*out = *in
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(Format)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileBackend.
func (in *FileBackend) DeepCopy() *FileBackend {
	if in == nil {
		return nil
	}
	out := new(FileBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Format) DeepCopyInto(out *Format) {
	*out = *in
	if in.Plain != nil {
		in, out := &in.Plain, &out.Plain
		*out = new(string)
		**out = **in
	}
	if in.Json != nil {
		in, out := &in.Json, &out.Json
		*out = new([]JsonValue)
		if **in != nil {
			in, out := *in, *out
			*out = make([]JsonValue, len(*in))
			copy(*out, *in)
		}
	}
	if in.OmitEmptyValues != nil {
		in, out := &in.OmitEmptyValues, &out.OmitEmptyValues
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Format.
func (in *Format) DeepCopy() *Format {
	if in == nil {
		return nil
	}
	out := new(Format)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *From) DeepCopyInto(out *From) {
	*out = *in
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(commonv1alpha1.TargetRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(Conf)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new From.
func (in *From) DeepCopy() *From {
	if in == nil {
		return nil
	}
	out := new(From)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsonValue) DeepCopyInto(out *JsonValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsonValue.
func (in *JsonValue) DeepCopy() *JsonValue {
	if in == nil {
		return nil
	}
	out := new(JsonValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshAccessLog) DeepCopyInto(out *MeshAccessLog) {
	*out = *in
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(commonv1alpha1.TargetRef)
		(*in).DeepCopyInto(*out)
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new([]To)
		if **in != nil {
			in, out := *in, *out
			*out = make([]To, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new([]From)
		if **in != nil {
			in, out := *in, *out
			*out = make([]From, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = new([]Rule)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Rule, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshAccessLog.
func (in *MeshAccessLog) DeepCopy() *MeshAccessLog {
	if in == nil {
		return nil
	}
	out := new(MeshAccessLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OtelBackend) DeepCopyInto(out *OtelBackend) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = new([]JsonValue)
		if **in != nil {
			in, out := *in, *out
			*out = make([]JsonValue, len(*in))
			copy(*out, *in)
		}
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OtelBackend.
func (in *OtelBackend) DeepCopy() *OtelBackend {
	if in == nil {
		return nil
	}
	out := new(OtelBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(Conf)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPBackend) DeepCopyInto(out *TCPBackend) {
	*out = *in
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(Format)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPBackend.
func (in *TCPBackend) DeepCopy() *TCPBackend {
	if in == nil {
		return nil
	}
	out := new(TCPBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *To) DeepCopyInto(out *To) {
	*out = *in
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(commonv1alpha1.TargetRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(Conf)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new To.
func (in *To) DeepCopy() *To {
	if in == nil {
		return nil
	}
	out := new(To)
	in.DeepCopyInto(out)
	return out
}
