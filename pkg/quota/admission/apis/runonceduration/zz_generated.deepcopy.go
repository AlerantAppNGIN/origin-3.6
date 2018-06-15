// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package runonceduration

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunOnceDurationConfig) DeepCopyInto(out *RunOnceDurationConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ActiveDeadlineSecondsLimit != nil {
		in, out := &in.ActiveDeadlineSecondsLimit, &out.ActiveDeadlineSecondsLimit
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunOnceDurationConfig.
func (in *RunOnceDurationConfig) DeepCopy() *RunOnceDurationConfig {
	if in == nil {
		return nil
	}
	out := new(RunOnceDurationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunOnceDurationConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
