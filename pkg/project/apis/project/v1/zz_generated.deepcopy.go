// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Project, InType: reflect.TypeOf(&Project{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ProjectList, InType: reflect.TypeOf(&ProjectList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ProjectRequest, InType: reflect.TypeOf(&ProjectRequest{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ProjectSpec, InType: reflect.TypeOf(&ProjectSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ProjectStatus, InType: reflect.TypeOf(&ProjectStatus{})},
	)
}

// DeepCopy_v1_Project is an autogenerated deepcopy function.
func DeepCopy_v1_Project(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Project)
		out := out.(*Project)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if err := DeepCopy_v1_ProjectSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_ProjectList is an autogenerated deepcopy function.
func DeepCopy_v1_ProjectList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ProjectList)
		out := out.(*ProjectList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Project, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_Project(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_ProjectRequest is an autogenerated deepcopy function.
func DeepCopy_v1_ProjectRequest(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ProjectRequest)
		out := out.(*ProjectRequest)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		return nil
	}
}

// DeepCopy_v1_ProjectSpec is an autogenerated deepcopy function.
func DeepCopy_v1_ProjectSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ProjectSpec)
		out := out.(*ProjectSpec)
		*out = *in
		if in.Finalizers != nil {
			in, out := &in.Finalizers, &out.Finalizers
			*out = make([]api_v1.FinalizerName, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1_ProjectStatus is an autogenerated deepcopy function.
func DeepCopy_v1_ProjectStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ProjectStatus)
		out := out.(*ProjectStatus)
		*out = *in
		return nil
	}
}
