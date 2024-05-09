// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	scheme "github.com/loft-sh/api/v4/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DevPodWorkspaceTemplatesGetter has a method to return a DevPodWorkspaceTemplateInterface.
// A group's client should implement this interface.
type DevPodWorkspaceTemplatesGetter interface {
	DevPodWorkspaceTemplates() DevPodWorkspaceTemplateInterface
}

// DevPodWorkspaceTemplateInterface has methods to work with DevPodWorkspaceTemplate resources.
type DevPodWorkspaceTemplateInterface interface {
	Create(ctx context.Context, devPodWorkspaceTemplate *v1.DevPodWorkspaceTemplate, opts metav1.CreateOptions) (*v1.DevPodWorkspaceTemplate, error)
	Update(ctx context.Context, devPodWorkspaceTemplate *v1.DevPodWorkspaceTemplate, opts metav1.UpdateOptions) (*v1.DevPodWorkspaceTemplate, error)
	UpdateStatus(ctx context.Context, devPodWorkspaceTemplate *v1.DevPodWorkspaceTemplate, opts metav1.UpdateOptions) (*v1.DevPodWorkspaceTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DevPodWorkspaceTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DevPodWorkspaceTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DevPodWorkspaceTemplate, err error)
	DevPodWorkspaceTemplateExpansion
}

// devPodWorkspaceTemplates implements DevPodWorkspaceTemplateInterface
type devPodWorkspaceTemplates struct {
	client rest.Interface
}

// newDevPodWorkspaceTemplates returns a DevPodWorkspaceTemplates
func newDevPodWorkspaceTemplates(c *StorageV1Client) *devPodWorkspaceTemplates {
	return &devPodWorkspaceTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the devPodWorkspaceTemplate, and returns the corresponding devPodWorkspaceTemplate object, and an error if there is any.
func (c *devPodWorkspaceTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DevPodWorkspaceTemplate, err error) {
	result = &v1.DevPodWorkspaceTemplate{}
	err = c.client.Get().
		Resource("devpodworkspacetemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DevPodWorkspaceTemplates that match those selectors.
func (c *devPodWorkspaceTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DevPodWorkspaceTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DevPodWorkspaceTemplateList{}
	err = c.client.Get().
		Resource("devpodworkspacetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested devPodWorkspaceTemplates.
func (c *devPodWorkspaceTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("devpodworkspacetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a devPodWorkspaceTemplate and creates it.  Returns the server's representation of the devPodWorkspaceTemplate, and an error, if there is any.
func (c *devPodWorkspaceTemplates) Create(ctx context.Context, devPodWorkspaceTemplate *v1.DevPodWorkspaceTemplate, opts metav1.CreateOptions) (result *v1.DevPodWorkspaceTemplate, err error) {
	result = &v1.DevPodWorkspaceTemplate{}
	err = c.client.Post().
		Resource("devpodworkspacetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(devPodWorkspaceTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a devPodWorkspaceTemplate and updates it. Returns the server's representation of the devPodWorkspaceTemplate, and an error, if there is any.
func (c *devPodWorkspaceTemplates) Update(ctx context.Context, devPodWorkspaceTemplate *v1.DevPodWorkspaceTemplate, opts metav1.UpdateOptions) (result *v1.DevPodWorkspaceTemplate, err error) {
	result = &v1.DevPodWorkspaceTemplate{}
	err = c.client.Put().
		Resource("devpodworkspacetemplates").
		Name(devPodWorkspaceTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(devPodWorkspaceTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *devPodWorkspaceTemplates) UpdateStatus(ctx context.Context, devPodWorkspaceTemplate *v1.DevPodWorkspaceTemplate, opts metav1.UpdateOptions) (result *v1.DevPodWorkspaceTemplate, err error) {
	result = &v1.DevPodWorkspaceTemplate{}
	err = c.client.Put().
		Resource("devpodworkspacetemplates").
		Name(devPodWorkspaceTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(devPodWorkspaceTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the devPodWorkspaceTemplate and deletes it. Returns an error if one occurs.
func (c *devPodWorkspaceTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("devpodworkspacetemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *devPodWorkspaceTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("devpodworkspacetemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched devPodWorkspaceTemplate.
func (c *devPodWorkspaceTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DevPodWorkspaceTemplate, err error) {
	result = &v1.DevPodWorkspaceTemplate{}
	err = c.client.Patch(pt).
		Resource("devpodworkspacetemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
