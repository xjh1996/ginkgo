/*
Copyright 2020 bytedance authors. All rights reserved.
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

// ClusterInfosGetter has a method to return a ClusterInfoInterface.
// A group's client should implement this interface.
type ClusterInfosGetter interface {
	ClusterInfos() ClusterInfoInterface
}

// ClusterInfoInterface has methods to work with ClusterInfo resources.
type ClusterInfoInterface interface {
	Create(ctx context.Context, clusterInfo *v1alpha1.ClusterInfo, opts v1.CreateOptions) (*v1alpha1.ClusterInfo, error)
	Update(ctx context.Context, clusterInfo *v1alpha1.ClusterInfo, opts v1.UpdateOptions) (*v1alpha1.ClusterInfo, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterInfo, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterInfoList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterInfo, err error)
	ClusterInfoExpansion
}

// clusterInfos implements ClusterInfoInterface
type clusterInfos struct {
	client rest.Interface
}

// newClusterInfos returns a ClusterInfos
func newClusterInfos(c *ResourceV1alpha1Client) *clusterInfos {
	return &clusterInfos{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterInfo, and returns the corresponding clusterInfo object, and an error if there is any.
func (c *clusterInfos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterInfo, err error) {
	result = &v1alpha1.ClusterInfo{}
	err = c.client.Get().
		Resource("clusterinfos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterInfos that match those selectors.
func (c *clusterInfos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterInfoList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterInfoList{}
	err = c.client.Get().
		Resource("clusterinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterInfos.
func (c *clusterInfos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterInfo and creates it.  Returns the server's representation of the clusterInfo, and an error, if there is any.
func (c *clusterInfos) Create(ctx context.Context, clusterInfo *v1alpha1.ClusterInfo, opts v1.CreateOptions) (result *v1alpha1.ClusterInfo, err error) {
	result = &v1alpha1.ClusterInfo{}
	err = c.client.Post().
		Resource("clusterinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterInfo).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterInfo and updates it. Returns the server's representation of the clusterInfo, and an error, if there is any.
func (c *clusterInfos) Update(ctx context.Context, clusterInfo *v1alpha1.ClusterInfo, opts v1.UpdateOptions) (result *v1alpha1.ClusterInfo, err error) {
	result = &v1alpha1.ClusterInfo{}
	err = c.client.Put().
		Resource("clusterinfos").
		Name(clusterInfo.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterInfo).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterInfo and deletes it. Returns an error if one occurs.
func (c *clusterInfos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterinfos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterInfos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterinfos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterInfo.
func (c *clusterInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterInfo, err error) {
	result = &v1alpha1.ClusterInfo{}
	err = c.client.Patch(pt).
		Resource("clusterinfos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
