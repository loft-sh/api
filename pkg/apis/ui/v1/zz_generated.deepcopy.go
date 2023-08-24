//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NavBarButton) DeepCopyInto(out *NavBarButton) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NavBarButton.
func (in *NavBarButton) DeepCopy() *NavBarButton {
	if in == nil {
		return nil
	}
	out := new(NavBarButton)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UISettings) DeepCopyInto(out *UISettings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UISettings.
func (in *UISettings) DeepCopy() *UISettings {
	if in == nil {
		return nil
	}
	out := new(UISettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UISettings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UISettingsConfig) DeepCopyInto(out *UISettingsConfig) {
	*out = *in
	if in.CustomCSS != nil {
		in, out := &in.CustomCSS, &out.CustomCSS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomJavaScript != nil {
		in, out := &in.CustomJavaScript, &out.CustomJavaScript
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NavBarButtons != nil {
		in, out := &in.NavBarButtons, &out.NavBarButtons
		*out = make([]NavBarButton, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UISettingsConfig.
func (in *UISettingsConfig) DeepCopy() *UISettingsConfig {
	if in == nil {
		return nil
	}
	out := new(UISettingsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UISettingsSpec) DeepCopyInto(out *UISettingsSpec) {
	*out = *in
	in.UISettingsConfig.DeepCopyInto(&out.UISettingsConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UISettingsSpec.
func (in *UISettingsSpec) DeepCopy() *UISettingsSpec {
	if in == nil {
		return nil
	}
	out := new(UISettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UISettingsStatus) DeepCopyInto(out *UISettingsStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UISettingsStatus.
func (in *UISettingsStatus) DeepCopy() *UISettingsStatus {
	if in == nil {
		return nil
	}
	out := new(UISettingsStatus)
	in.DeepCopyInto(out)
	return out
}
