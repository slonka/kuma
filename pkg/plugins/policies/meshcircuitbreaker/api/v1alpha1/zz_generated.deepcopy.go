//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	commonv1alpha1 "github.com/kumahq/kuma/api/common/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Conf) DeepCopyInto(out *Conf) {
	*out = *in
	if in.ConnectionLimits != nil {
		in, out := &in.ConnectionLimits, &out.ConnectionLimits
		*out = new(ConnectionLimits)
		(*in).DeepCopyInto(*out)
	}
	if in.OutlierDetection != nil {
		in, out := &in.OutlierDetection, &out.OutlierDetection
		*out = new(OutlierDetection)
		(*in).DeepCopyInto(*out)
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
func (in *ConnectionLimits) DeepCopyInto(out *ConnectionLimits) {
	*out = *in
	if in.MaxConnections != nil {
		in, out := &in.MaxConnections, &out.MaxConnections
		*out = new(uint32)
		**out = **in
	}
	if in.MaxConnectionPools != nil {
		in, out := &in.MaxConnectionPools, &out.MaxConnectionPools
		*out = new(uint32)
		**out = **in
	}
	if in.MaxPendingRequests != nil {
		in, out := &in.MaxPendingRequests, &out.MaxPendingRequests
		*out = new(uint32)
		**out = **in
	}
	if in.MaxRetries != nil {
		in, out := &in.MaxRetries, &out.MaxRetries
		*out = new(uint32)
		**out = **in
	}
	if in.MaxRequests != nil {
		in, out := &in.MaxRequests, &out.MaxRequests
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionLimits.
func (in *ConnectionLimits) DeepCopy() *ConnectionLimits {
	if in == nil {
		return nil
	}
	out := new(ConnectionLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetectorFailurePercentageFailures) DeepCopyInto(out *DetectorFailurePercentageFailures) {
	*out = *in
	if in.MinimumHosts != nil {
		in, out := &in.MinimumHosts, &out.MinimumHosts
		*out = new(uint32)
		**out = **in
	}
	if in.RequestVolume != nil {
		in, out := &in.RequestVolume, &out.RequestVolume
		*out = new(uint32)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetectorFailurePercentageFailures.
func (in *DetectorFailurePercentageFailures) DeepCopy() *DetectorFailurePercentageFailures {
	if in == nil {
		return nil
	}
	out := new(DetectorFailurePercentageFailures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetectorGatewayFailures) DeepCopyInto(out *DetectorGatewayFailures) {
	*out = *in
	if in.Consecutive != nil {
		in, out := &in.Consecutive, &out.Consecutive
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetectorGatewayFailures.
func (in *DetectorGatewayFailures) DeepCopy() *DetectorGatewayFailures {
	if in == nil {
		return nil
	}
	out := new(DetectorGatewayFailures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetectorLocalOriginFailures) DeepCopyInto(out *DetectorLocalOriginFailures) {
	*out = *in
	if in.Consecutive != nil {
		in, out := &in.Consecutive, &out.Consecutive
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetectorLocalOriginFailures.
func (in *DetectorLocalOriginFailures) DeepCopy() *DetectorLocalOriginFailures {
	if in == nil {
		return nil
	}
	out := new(DetectorLocalOriginFailures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetectorSuccessRateFailures) DeepCopyInto(out *DetectorSuccessRateFailures) {
	*out = *in
	if in.MinimumHosts != nil {
		in, out := &in.MinimumHosts, &out.MinimumHosts
		*out = new(uint32)
		**out = **in
	}
	if in.RequestVolume != nil {
		in, out := &in.RequestVolume, &out.RequestVolume
		*out = new(uint32)
		**out = **in
	}
	if in.StandardDeviationFactor != nil {
		in, out := &in.StandardDeviationFactor, &out.StandardDeviationFactor
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetectorSuccessRateFailures.
func (in *DetectorSuccessRateFailures) DeepCopy() *DetectorSuccessRateFailures {
	if in == nil {
		return nil
	}
	out := new(DetectorSuccessRateFailures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetectorTotalFailures) DeepCopyInto(out *DetectorTotalFailures) {
	*out = *in
	if in.Consecutive != nil {
		in, out := &in.Consecutive, &out.Consecutive
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetectorTotalFailures.
func (in *DetectorTotalFailures) DeepCopy() *DetectorTotalFailures {
	if in == nil {
		return nil
	}
	out := new(DetectorTotalFailures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Detectors) DeepCopyInto(out *Detectors) {
	*out = *in
	if in.TotalFailures != nil {
		in, out := &in.TotalFailures, &out.TotalFailures
		*out = new(DetectorTotalFailures)
		(*in).DeepCopyInto(*out)
	}
	if in.GatewayFailures != nil {
		in, out := &in.GatewayFailures, &out.GatewayFailures
		*out = new(DetectorGatewayFailures)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalOriginFailures != nil {
		in, out := &in.LocalOriginFailures, &out.LocalOriginFailures
		*out = new(DetectorLocalOriginFailures)
		(*in).DeepCopyInto(*out)
	}
	if in.SuccessRate != nil {
		in, out := &in.SuccessRate, &out.SuccessRate
		*out = new(DetectorSuccessRateFailures)
		(*in).DeepCopyInto(*out)
	}
	if in.FailurePercentage != nil {
		in, out := &in.FailurePercentage, &out.FailurePercentage
		*out = new(DetectorFailurePercentageFailures)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Detectors.
func (in *Detectors) DeepCopy() *Detectors {
	if in == nil {
		return nil
	}
	out := new(Detectors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *From) DeepCopyInto(out *From) {
	*out = *in
	in.TargetRef.DeepCopyInto(&out.TargetRef)
	in.Default.DeepCopyInto(&out.Default)
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
func (in *MeshCircuitBreaker) DeepCopyInto(out *MeshCircuitBreaker) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshCircuitBreaker.
func (in *MeshCircuitBreaker) DeepCopy() *MeshCircuitBreaker {
	if in == nil {
		return nil
	}
	out := new(MeshCircuitBreaker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutlierDetection) DeepCopyInto(out *OutlierDetection) {
	*out = *in
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(v1.Duration)
		**out = **in
	}
	if in.BaseEjectionTime != nil {
		in, out := &in.BaseEjectionTime, &out.BaseEjectionTime
		*out = new(v1.Duration)
		**out = **in
	}
	if in.MaxEjectionPercent != nil {
		in, out := &in.MaxEjectionPercent, &out.MaxEjectionPercent
		*out = new(uint32)
		**out = **in
	}
	if in.SplitExternalAndLocalErrors != nil {
		in, out := &in.SplitExternalAndLocalErrors, &out.SplitExternalAndLocalErrors
		*out = new(bool)
		**out = **in
	}
	if in.Detectors != nil {
		in, out := &in.Detectors, &out.Detectors
		*out = new(Detectors)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthyPanicThreshold != nil {
		in, out := &in.HealthyPanicThreshold, &out.HealthyPanicThreshold
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetection.
func (in *OutlierDetection) DeepCopy() *OutlierDetection {
	if in == nil {
		return nil
	}
	out := new(OutlierDetection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	in.Default.DeepCopyInto(&out.Default)
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
func (in *To) DeepCopyInto(out *To) {
	*out = *in
	in.TargetRef.DeepCopyInto(&out.TargetRef)
	in.Default.DeepCopyInto(&out.Default)
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
