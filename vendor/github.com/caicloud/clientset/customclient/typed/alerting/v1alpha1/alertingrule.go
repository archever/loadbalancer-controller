/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	scheme "github.com/caicloud/clientset/customclient/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/alerting/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AlertingRulesGetter has a method to return a AlertingRuleInterface.
// A group's client should implement this interface.
type AlertingRulesGetter interface {
	AlertingRules() AlertingRuleInterface
}

// AlertingRuleInterface has methods to work with AlertingRule resources.
type AlertingRuleInterface interface {
	Create(*v1alpha1.AlertingRule) (*v1alpha1.AlertingRule, error)
	Update(*v1alpha1.AlertingRule) (*v1alpha1.AlertingRule, error)
	UpdateStatus(*v1alpha1.AlertingRule) (*v1alpha1.AlertingRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AlertingRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AlertingRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AlertingRule, err error)
	AlertingRuleExpansion
}

// alertingRules implements AlertingRuleInterface
type alertingRules struct {
	client rest.Interface
}

// newAlertingRules returns a AlertingRules
func newAlertingRules(c *AlertingV1alpha1Client) *alertingRules {
	return &alertingRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the alertingRule, and returns the corresponding alertingRule object, and an error if there is any.
func (c *alertingRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AlertingRule, err error) {
	result = &v1alpha1.AlertingRule{}
	err = c.client.Get().
		Resource("alertingrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AlertingRules that match those selectors.
func (c *alertingRules) List(opts v1.ListOptions) (result *v1alpha1.AlertingRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AlertingRuleList{}
	err = c.client.Get().
		Resource("alertingrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested alertingRules.
func (c *alertingRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("alertingrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a alertingRule and creates it.  Returns the server's representation of the alertingRule, and an error, if there is any.
func (c *alertingRules) Create(alertingRule *v1alpha1.AlertingRule) (result *v1alpha1.AlertingRule, err error) {
	result = &v1alpha1.AlertingRule{}
	err = c.client.Post().
		Resource("alertingrules").
		Body(alertingRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a alertingRule and updates it. Returns the server's representation of the alertingRule, and an error, if there is any.
func (c *alertingRules) Update(alertingRule *v1alpha1.AlertingRule) (result *v1alpha1.AlertingRule, err error) {
	result = &v1alpha1.AlertingRule{}
	err = c.client.Put().
		Resource("alertingrules").
		Name(alertingRule.Name).
		Body(alertingRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *alertingRules) UpdateStatus(alertingRule *v1alpha1.AlertingRule) (result *v1alpha1.AlertingRule, err error) {
	result = &v1alpha1.AlertingRule{}
	err = c.client.Put().
		Resource("alertingrules").
		Name(alertingRule.Name).
		SubResource("status").
		Body(alertingRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the alertingRule and deletes it. Returns an error if one occurs.
func (c *alertingRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("alertingrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *alertingRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("alertingrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched alertingRule.
func (c *alertingRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AlertingRule, err error) {
	result = &v1alpha1.AlertingRule{}
	err = c.client.Patch(pt).
		Resource("alertingrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}