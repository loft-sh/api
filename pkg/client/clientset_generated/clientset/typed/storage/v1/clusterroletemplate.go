// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	scheme "github.com/loft-sh/api/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterRoleTemplatesGetter has a method to return a ClusterRoleTemplateInterface.
// A group's client should implement this interface.
type ClusterRoleTemplatesGetter interface {
	ClusterRoleTemplates() ClusterRoleTemplateInterface
}

// ClusterRoleTemplateInterface has methods to work with ClusterRoleTemplate resources.
type ClusterRoleTemplateInterface interface {
	Create(ctx context.Context, clusterRoleTemplate *v1.ClusterRoleTemplate, opts metav1.CreateOptions) (*v1.ClusterRoleTemplate, error)
	Update(ctx context.Context, clusterRoleTemplate *v1.ClusterRoleTemplate, opts metav1.UpdateOptions) (*v1.ClusterRoleTemplate, error)
	UpdateStatus(ctx context.Context, clusterRoleTemplate *v1.ClusterRoleTemplate, opts metav1.UpdateOptions) (*v1.ClusterRoleTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterRoleTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterRoleTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterRoleTemplate, err error)
	ClusterRoleTemplateExpansion
}

// clusterRoleTemplates implements ClusterRoleTemplateInterface
type clusterRoleTemplates struct {
	client rest.Interface
}

// newClusterRoleTemplates returns a ClusterRoleTemplates
func newClusterRoleTemplates(c *StorageV1Client) *clusterRoleTemplates {
	return &clusterRoleTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterRoleTemplate, and returns the corresponding clusterRoleTemplate object, and an error if there is any.
func (c *clusterRoleTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterRoleTemplate, err error) {
	result = &v1.ClusterRoleTemplate{}
	err = c.client.Get().
		Resource("clusterroletemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterRoleTemplates that match those selectors.
func (c *clusterRoleTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterRoleTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterRoleTemplateList{}
	err = c.client.Get().
		Resource("clusterroletemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterRoleTemplates.
func (c *clusterRoleTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterroletemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterRoleTemplate and creates it.  Returns the server's representation of the clusterRoleTemplate, and an error, if there is any.
func (c *clusterRoleTemplates) Create(ctx context.Context, clusterRoleTemplate *v1.ClusterRoleTemplate, opts metav1.CreateOptions) (result *v1.ClusterRoleTemplate, err error) {
	result = &v1.ClusterRoleTemplate{}
	err = c.client.Post().
		Resource("clusterroletemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterRoleTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterRoleTemplate and updates it. Returns the server's representation of the clusterRoleTemplate, and an error, if there is any.
func (c *clusterRoleTemplates) Update(ctx context.Context, clusterRoleTemplate *v1.ClusterRoleTemplate, opts metav1.UpdateOptions) (result *v1.ClusterRoleTemplate, err error) {
	result = &v1.ClusterRoleTemplate{}
	err = c.client.Put().
		Resource("clusterroletemplates").
		Name(clusterRoleTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterRoleTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterRoleTemplates) UpdateStatus(ctx context.Context, clusterRoleTemplate *v1.ClusterRoleTemplate, opts metav1.UpdateOptions) (result *v1.ClusterRoleTemplate, err error) {
	result = &v1.ClusterRoleTemplate{}
	err = c.client.Put().
		Resource("clusterroletemplates").
		Name(clusterRoleTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterRoleTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterRoleTemplate and deletes it. Returns an error if one occurs.
func (c *clusterRoleTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterroletemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterRoleTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterroletemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterRoleTemplate.
func (c *clusterRoleTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterRoleTemplate, err error) {
	result = &v1.ClusterRoleTemplate{}
	err = c.client.Patch(pt).
		Resource("clusterroletemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
