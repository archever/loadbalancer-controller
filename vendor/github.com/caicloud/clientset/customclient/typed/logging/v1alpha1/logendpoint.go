/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	scheme "github.com/caicloud/clientset/customclient/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/logging/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LogEndpointsGetter has a method to return a LogEndpointInterface.
// A group's client should implement this interface.
type LogEndpointsGetter interface {
	LogEndpoints() LogEndpointInterface
}

// LogEndpointInterface has methods to work with LogEndpoint resources.
type LogEndpointInterface interface {
	Create(*v1alpha1.LogEndpoint) (*v1alpha1.LogEndpoint, error)
	Update(*v1alpha1.LogEndpoint) (*v1alpha1.LogEndpoint, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LogEndpoint, error)
	List(opts v1.ListOptions) (*v1alpha1.LogEndpointList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogEndpoint, err error)
	LogEndpointExpansion
}

// logEndpoints implements LogEndpointInterface
type logEndpoints struct {
	client rest.Interface
}

// newLogEndpoints returns a LogEndpoints
func newLogEndpoints(c *LoggingV1alpha1Client) *logEndpoints {
	return &logEndpoints{
		client: c.RESTClient(),
	}
}

// Get takes name of the logEndpoint, and returns the corresponding logEndpoint object, and an error if there is any.
func (c *logEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.LogEndpoint, err error) {
	result = &v1alpha1.LogEndpoint{}
	err = c.client.Get().
		Resource("logendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LogEndpoints that match those selectors.
func (c *logEndpoints) List(opts v1.ListOptions) (result *v1alpha1.LogEndpointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LogEndpointList{}
	err = c.client.Get().
		Resource("logendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested logEndpoints.
func (c *logEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("logendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a logEndpoint and creates it.  Returns the server's representation of the logEndpoint, and an error, if there is any.
func (c *logEndpoints) Create(logEndpoint *v1alpha1.LogEndpoint) (result *v1alpha1.LogEndpoint, err error) {
	result = &v1alpha1.LogEndpoint{}
	err = c.client.Post().
		Resource("logendpoints").
		Body(logEndpoint).
		Do().
		Into(result)
	return
}

// Update takes the representation of a logEndpoint and updates it. Returns the server's representation of the logEndpoint, and an error, if there is any.
func (c *logEndpoints) Update(logEndpoint *v1alpha1.LogEndpoint) (result *v1alpha1.LogEndpoint, err error) {
	result = &v1alpha1.LogEndpoint{}
	err = c.client.Put().
		Resource("logendpoints").
		Name(logEndpoint.Name).
		Body(logEndpoint).
		Do().
		Into(result)
	return
}

// Delete takes name of the logEndpoint and deletes it. Returns an error if one occurs.
func (c *logEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("logendpoints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *logEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("logendpoints").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched logEndpoint.
func (c *logEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogEndpoint, err error) {
	result = &v1alpha1.LogEndpoint{}
	err = c.client.Patch(pt).
		Resource("logendpoints").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
