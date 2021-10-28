// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GlobalClusterRoleTemplatesGetter has a method to return a GlobalClusterRoleTemplateInterface.
// A group's client should implement this interface.
type GlobalClusterRoleTemplatesGetter interface {
	GlobalClusterRoleTemplates() GlobalClusterRoleTemplateInterface
}

// GlobalClusterRoleTemplateInterface has methods to work with GlobalClusterRoleTemplate resources.
type GlobalClusterRoleTemplateInterface interface {
	Create(ctx context.Context, globalClusterRoleTemplate *v1.GlobalClusterRoleTemplate, opts metav1.CreateOptions) (*v1.GlobalClusterRoleTemplate, error)
	Update(ctx context.Context, globalClusterRoleTemplate *v1.GlobalClusterRoleTemplate, opts metav1.UpdateOptions) (*v1.GlobalClusterRoleTemplate, error)
	UpdateStatus(ctx context.Context, globalClusterRoleTemplate *v1.GlobalClusterRoleTemplate, opts metav1.UpdateOptions) (*v1.GlobalClusterRoleTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.GlobalClusterRoleTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GlobalClusterRoleTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GlobalClusterRoleTemplate, err error)
	GlobalClusterRoleTemplateExpansion
}

// globalClusterRoleTemplates implements GlobalClusterRoleTemplateInterface
type globalClusterRoleTemplates struct {
	client rest.Interface
}

// newGlobalClusterRoleTemplates returns a GlobalClusterRoleTemplates
func newGlobalClusterRoleTemplates(c *ManagementV1Client) *globalClusterRoleTemplates {
	return &globalClusterRoleTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the globalClusterRoleTemplate, and returns the corresponding globalClusterRoleTemplate object, and an error if there is any.
func (c *globalClusterRoleTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GlobalClusterRoleTemplate, err error) {
	result = &v1.GlobalClusterRoleTemplate{}
	err = c.client.Get().
		Resource("globalclusterroletemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalClusterRoleTemplates that match those selectors.
func (c *globalClusterRoleTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GlobalClusterRoleTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GlobalClusterRoleTemplateList{}
	err = c.client.Get().
		Resource("globalclusterroletemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalClusterRoleTemplates.
func (c *globalClusterRoleTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("globalclusterroletemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a globalClusterRoleTemplate and creates it.  Returns the server's representation of the globalClusterRoleTemplate, and an error, if there is any.
func (c *globalClusterRoleTemplates) Create(ctx context.Context, globalClusterRoleTemplate *v1.GlobalClusterRoleTemplate, opts metav1.CreateOptions) (result *v1.GlobalClusterRoleTemplate, err error) {
	result = &v1.GlobalClusterRoleTemplate{}
	err = c.client.Post().
		Resource("globalclusterroletemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalClusterRoleTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a globalClusterRoleTemplate and updates it. Returns the server's representation of the globalClusterRoleTemplate, and an error, if there is any.
func (c *globalClusterRoleTemplates) Update(ctx context.Context, globalClusterRoleTemplate *v1.GlobalClusterRoleTemplate, opts metav1.UpdateOptions) (result *v1.GlobalClusterRoleTemplate, err error) {
	result = &v1.GlobalClusterRoleTemplate{}
	err = c.client.Put().
		Resource("globalclusterroletemplates").
		Name(globalClusterRoleTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalClusterRoleTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *globalClusterRoleTemplates) UpdateStatus(ctx context.Context, globalClusterRoleTemplate *v1.GlobalClusterRoleTemplate, opts metav1.UpdateOptions) (result *v1.GlobalClusterRoleTemplate, err error) {
	result = &v1.GlobalClusterRoleTemplate{}
	err = c.client.Put().
		Resource("globalclusterroletemplates").
		Name(globalClusterRoleTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalClusterRoleTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the globalClusterRoleTemplate and deletes it. Returns an error if one occurs.
func (c *globalClusterRoleTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("globalclusterroletemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalClusterRoleTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("globalclusterroletemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched globalClusterRoleTemplate.
func (c *globalClusterRoleTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GlobalClusterRoleTemplate, err error) {
	result = &v1.GlobalClusterRoleTemplate{}
	err = c.client.Patch(pt).
		Resource("globalclusterroletemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
