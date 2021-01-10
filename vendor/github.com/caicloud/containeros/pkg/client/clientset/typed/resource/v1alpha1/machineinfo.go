/*
Copyright 2021 bytedance authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/caicloud/containeros/pkg/apis/resource/v1alpha1"
	scheme "github.com/caicloud/containeros/pkg/client/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineInfosGetter has a method to return a MachineInfoInterface.
// A group's client should implement this interface.
type MachineInfosGetter interface {
	MachineInfos() MachineInfoInterface
}

// MachineInfoInterface has methods to work with MachineInfo resources.
type MachineInfoInterface interface {
	Create(ctx context.Context, machineInfo *v1alpha1.MachineInfo, opts v1.CreateOptions) (*v1alpha1.MachineInfo, error)
	Update(ctx context.Context, machineInfo *v1alpha1.MachineInfo, opts v1.UpdateOptions) (*v1alpha1.MachineInfo, error)
	UpdateStatus(ctx context.Context, machineInfo *v1alpha1.MachineInfo, opts v1.UpdateOptions) (*v1alpha1.MachineInfo, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MachineInfo, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MachineInfoList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MachineInfo, err error)
	MachineInfoExpansion
}

// machineInfos implements MachineInfoInterface
type machineInfos struct {
	client rest.Interface
}

// newMachineInfos returns a MachineInfos
func newMachineInfos(c *ResourceV1alpha1Client) *machineInfos {
	return &machineInfos{
		client: c.RESTClient(),
	}
}

// Get takes name of the machineInfo, and returns the corresponding machineInfo object, and an error if there is any.
func (c *machineInfos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MachineInfo, err error) {
	result = &v1alpha1.MachineInfo{}
	err = c.client.Get().
		Resource("machineinfos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineInfos that match those selectors.
func (c *machineInfos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MachineInfoList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MachineInfoList{}
	err = c.client.Get().
		Resource("machineinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineInfos.
func (c *machineInfos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("machineinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a machineInfo and creates it.  Returns the server's representation of the machineInfo, and an error, if there is any.
func (c *machineInfos) Create(ctx context.Context, machineInfo *v1alpha1.MachineInfo, opts v1.CreateOptions) (result *v1alpha1.MachineInfo, err error) {
	result = &v1alpha1.MachineInfo{}
	err = c.client.Post().
		Resource("machineinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineInfo).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a machineInfo and updates it. Returns the server's representation of the machineInfo, and an error, if there is any.
func (c *machineInfos) Update(ctx context.Context, machineInfo *v1alpha1.MachineInfo, opts v1.UpdateOptions) (result *v1alpha1.MachineInfo, err error) {
	result = &v1alpha1.MachineInfo{}
	err = c.client.Put().
		Resource("machineinfos").
		Name(machineInfo.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineInfo).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *machineInfos) UpdateStatus(ctx context.Context, machineInfo *v1alpha1.MachineInfo, opts v1.UpdateOptions) (result *v1alpha1.MachineInfo, err error) {
	result = &v1alpha1.MachineInfo{}
	err = c.client.Put().
		Resource("machineinfos").
		Name(machineInfo.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineInfo).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the machineInfo and deletes it. Returns an error if one occurs.
func (c *machineInfos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("machineinfos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineInfos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("machineinfos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched machineInfo.
func (c *machineInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MachineInfo, err error) {
	result = &v1alpha1.MachineInfo{}
	err = c.client.Patch(pt).
		Resource("machineinfos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
