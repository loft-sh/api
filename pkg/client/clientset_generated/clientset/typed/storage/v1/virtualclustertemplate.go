// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v2/pkg/apis/storage/v1"
	scheme "github.com/loft-sh/api/v2/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualClusterTemplatesGetter has a method to return a VirtualClusterTemplateInterface.
// A group's client should implement this interface.
type VirtualClusterTemplatesGetter interface {
	VirtualClusterTemplates() VirtualClusterTemplateInterface
}

// VirtualClusterTemplateInterface has methods to work with VirtualClusterTemplate resources.
type VirtualClusterTemplateInterface interface {
	Create(ctx context.Context, virtualClusterTemplate *v1.VirtualClusterTemplate, opts metav1.CreateOptions) (*v1.VirtualClusterTemplate, error)
	Update(ctx context.Context, virtualClusterTemplate *v1.VirtualClusterTemplate, opts metav1.UpdateOptions) (*v1.VirtualClusterTemplate, error)
	UpdateStatus(ctx context.Context, virtualClusterTemplate *v1.VirtualClusterTemplate, opts metav1.UpdateOptions) (*v1.VirtualClusterTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VirtualClusterTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VirtualClusterTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualClusterTemplate, err error)
	VirtualClusterTemplateExpansion
}

// virtualClusterTemplates implements VirtualClusterTemplateInterface
type virtualClusterTemplates struct {
	client rest.Interface
}

// newVirtualClusterTemplates returns a VirtualClusterTemplates
func newVirtualClusterTemplates(c *StorageV1Client) *virtualClusterTemplates {
	return &virtualClusterTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the virtualClusterTemplate, and returns the corresponding virtualClusterTemplate object, and an error if there is any.
func (c *virtualClusterTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualClusterTemplate, err error) {
	result = &v1.VirtualClusterTemplate{}
	err = c.client.Get().
		Resource("virtualclustertemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualClusterTemplates that match those selectors.
func (c *virtualClusterTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualClusterTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VirtualClusterTemplateList{}
	err = c.client.Get().
		Resource("virtualclustertemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualClusterTemplates.
func (c *virtualClusterTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("virtualclustertemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualClusterTemplate and creates it.  Returns the server's representation of the virtualClusterTemplate, and an error, if there is any.
func (c *virtualClusterTemplates) Create(ctx context.Context, virtualClusterTemplate *v1.VirtualClusterTemplate, opts metav1.CreateOptions) (result *v1.VirtualClusterTemplate, err error) {
	result = &v1.VirtualClusterTemplate{}
	err = c.client.Post().
		Resource("virtualclustertemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualClusterTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualClusterTemplate and updates it. Returns the server's representation of the virtualClusterTemplate, and an error, if there is any.
func (c *virtualClusterTemplates) Update(ctx context.Context, virtualClusterTemplate *v1.VirtualClusterTemplate, opts metav1.UpdateOptions) (result *v1.VirtualClusterTemplate, err error) {
	result = &v1.VirtualClusterTemplate{}
	err = c.client.Put().
		Resource("virtualclustertemplates").
		Name(virtualClusterTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualClusterTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *virtualClusterTemplates) UpdateStatus(ctx context.Context, virtualClusterTemplate *v1.VirtualClusterTemplate, opts metav1.UpdateOptions) (result *v1.VirtualClusterTemplate, err error) {
	result = &v1.VirtualClusterTemplate{}
	err = c.client.Put().
		Resource("virtualclustertemplates").
		Name(virtualClusterTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualClusterTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualClusterTemplate and deletes it. Returns an error if one occurs.
func (c *virtualClusterTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("virtualclustertemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualClusterTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("virtualclustertemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualClusterTemplate.
func (c *virtualClusterTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualClusterTemplate, err error) {
	result = &v1.VirtualClusterTemplate{}
	err = c.client.Patch(pt).
		Resource("virtualclustertemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
