// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package imagepolicy

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageCondition) DeepCopyInto(out *ImageCondition) {
	*out = *in
	if in.OnResources != nil {
		in, out := &in.OnResources, &out.OnResources
		*out = make([]schema.GroupResource, len(*in))
		copy(*out, *in)
	}
	if in.MatchRegistries != nil {
		in, out := &in.MatchRegistries, &out.MatchRegistries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchDockerImageLabels != nil {
		in, out := &in.MatchDockerImageLabels, &out.MatchDockerImageLabels
		*out = make([]ValueCondition, len(*in))
		copy(*out, *in)
	}
	if in.MatchImageLabels != nil {
		in, out := &in.MatchImageLabels, &out.MatchImageLabels
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchImageLabelSelectors != nil {
		in, out := &in.MatchImageLabelSelectors, &out.MatchImageLabelSelectors
		*out = make([]labels.Selector, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = (*in)[i].DeepCopySelector()
			}
		}
	}
	if in.MatchImageAnnotations != nil {
		in, out := &in.MatchImageAnnotations, &out.MatchImageAnnotations
		*out = make([]ValueCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageCondition.
func (in *ImageCondition) DeepCopy() *ImageCondition {
	if in == nil {
		return nil
	}
	out := new(ImageCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageExecutionPolicyRule) DeepCopyInto(out *ImageExecutionPolicyRule) {
	*out = *in
	in.ImageCondition.DeepCopyInto(&out.ImageCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageExecutionPolicyRule.
func (in *ImageExecutionPolicyRule) DeepCopy() *ImageExecutionPolicyRule {
	if in == nil {
		return nil
	}
	out := new(ImageExecutionPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicyConfig) DeepCopyInto(out *ImagePolicyConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ResolutionRules != nil {
		in, out := &in.ResolutionRules, &out.ResolutionRules
		*out = make([]ImageResolutionPolicyRule, len(*in))
		copy(*out, *in)
	}
	if in.ExecutionRules != nil {
		in, out := &in.ExecutionRules, &out.ExecutionRules
		*out = make([]ImageExecutionPolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicyConfig.
func (in *ImagePolicyConfig) DeepCopy() *ImagePolicyConfig {
	if in == nil {
		return nil
	}
	out := new(ImagePolicyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagePolicyConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageResolutionPolicyRule) DeepCopyInto(out *ImageResolutionPolicyRule) {
	*out = *in
	out.TargetResource = in.TargetResource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageResolutionPolicyRule.
func (in *ImageResolutionPolicyRule) DeepCopy() *ImageResolutionPolicyRule {
	if in == nil {
		return nil
	}
	out := new(ImageResolutionPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueCondition) DeepCopyInto(out *ValueCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueCondition.
func (in *ValueCondition) DeepCopy() *ValueCondition {
	if in == nil {
		return nil
	}
	out := new(ValueCondition)
	in.DeepCopyInto(out)
	return out
}
