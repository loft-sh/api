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

// ClusterAccountTemplatesGetter has a method to return a ClusterAccountTemplateInterface.
// A group's client should implement this interface.
type ClusterAccountTemplatesGetter interface {
	ClusterAccountTemplates() ClusterAccountTemplateInterface
}

// ClusterAccountTemplateInterface has methods to work with ClusterAccountTemplate resources.
type ClusterAccountTemplateInterface interface {
	Create(ctx context.Context, clusterAccountTemplate *v1.ClusterAccountTemplate, opts metav1.CreateOptions) (*v1.ClusterAccountTemplate, error)
	Update(ctx context.Context, clusterAccountTemplate *v1.ClusterAccountTemplate, opts metav1.UpdateOptions) (*v1.ClusterAccountTemplate, error)
	UpdateStatus(ctx context.Context, clusterAccountTemplate *v1.ClusterAccountTemplate, opts metav1.UpdateOptions) (*v1.ClusterAccountTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterAccountTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterAccountTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterAccountTemplate, err error)
	ClusterAccountTemplateExpansion
}

// clusterAccountTemplates implements ClusterAccountTemplateInterface
type clusterAccountTemplates struct {
	client rest.Interface
}

// newClusterAccountTemplates returns a ClusterAccountTemplates
func newClusterAccountTemplates(c *StorageV1Client) *clusterAccountTemplates {
	return &clusterAccountTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterAccountTemplate, and returns the corresponding clusterAccountTemplate object, and an error if there is any.
func (c *clusterAccountTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterAccountTemplate, err error) {
	result = &v1.ClusterAccountTemplate{}
	err = c.client.Get().
		Resource("clusteraccounttemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterAccountTemplates that match those selectors.
func (c *clusterAccountTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterAccountTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterAccountTemplateList{}
	err = c.client.Get().
		Resource("clusteraccounttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterAccountTemplates.
func (c *clusterAccountTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusteraccounttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterAccountTemplate and creates it.  Returns the server's representation of the clusterAccountTemplate, and an error, if there is any.
func (c *clusterAccountTemplates) Create(ctx context.Context, clusterAccountTemplate *v1.ClusterAccountTemplate, opts metav1.CreateOptions) (result *v1.ClusterAccountTemplate, err error) {
	result = &v1.ClusterAccountTemplate{}
	err = c.client.Post().
		Resource("clusteraccounttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAccountTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterAccountTemplate and updates it. Returns the server's representation of the clusterAccountTemplate, and an error, if there is any.
func (c *clusterAccountTemplates) Update(ctx context.Context, clusterAccountTemplate *v1.ClusterAccountTemplate, opts metav1.UpdateOptions) (result *v1.ClusterAccountTemplate, err error) {
	result = &v1.ClusterAccountTemplate{}
	err = c.client.Put().
		Resource("clusteraccounttemplates").
		Name(clusterAccountTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAccountTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterAccountTemplates) UpdateStatus(ctx context.Context, clusterAccountTemplate *v1.ClusterAccountTemplate, opts metav1.UpdateOptions) (result *v1.ClusterAccountTemplate, err error) {
	result = &v1.ClusterAccountTemplate{}
	err = c.client.Put().
		Resource("clusteraccounttemplates").
		Name(clusterAccountTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAccountTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterAccountTemplate and deletes it. Returns an error if one occurs.
func (c *clusterAccountTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusteraccounttemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterAccountTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusteraccounttemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterAccountTemplate.
func (c *clusterAccountTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterAccountTemplate, err error) {
	result = &v1.ClusterAccountTemplate{}
	err = c.client.Patch(pt).
		Resource("clusteraccounttemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
