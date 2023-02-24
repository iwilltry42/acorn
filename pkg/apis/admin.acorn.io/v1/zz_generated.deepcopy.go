//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	internal_acorn_iov1 "github.com/acorn-io/acorn/pkg/apis/internal.acorn.io/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVolumeClass) DeepCopyInto(out *ClusterVolumeClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AllowedAccessModes != nil {
		in, out := &in.AllowedAccessModes, &out.AllowedAccessModes
		*out = make(internal_acorn_iov1.AccessModes, len(*in))
		copy(*out, *in)
	}
	out.Size = in.Size
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVolumeClass.
func (in *ClusterVolumeClass) DeepCopy() *ClusterVolumeClass {
	if in == nil {
		return nil
	}
	out := new(ClusterVolumeClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVolumeClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVolumeClassList) DeepCopyInto(out *ClusterVolumeClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterVolumeClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVolumeClassList.
func (in *ClusterVolumeClassList) DeepCopy() *ClusterVolumeClassList {
	if in == nil {
		return nil
	}
	out := new(ClusterVolumeClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVolumeClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectVolumeClass) DeepCopyInto(out *ProjectVolumeClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AllowedAccessModes != nil {
		in, out := &in.AllowedAccessModes, &out.AllowedAccessModes
		*out = make(internal_acorn_iov1.AccessModes, len(*in))
		copy(*out, *in)
	}
	out.Size = in.Size
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectVolumeClass.
func (in *ProjectVolumeClass) DeepCopy() *ProjectVolumeClass {
	if in == nil {
		return nil
	}
	out := new(ProjectVolumeClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectVolumeClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectVolumeClassList) DeepCopyInto(out *ProjectVolumeClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectVolumeClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectVolumeClassList.
func (in *ProjectVolumeClassList) DeepCopy() *ProjectVolumeClassList {
	if in == nil {
		return nil
	}
	out := new(ProjectVolumeClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectVolumeClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
