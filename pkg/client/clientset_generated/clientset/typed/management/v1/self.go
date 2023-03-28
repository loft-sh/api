// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v2/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v2/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SelvesGetter has a method to return a SelfInterface.
// A group's client should implement this interface.
type SelvesGetter interface {
	Selves() SelfInterface
}

// SelfInterface has methods to work with Self resources.
type SelfInterface interface {
	Create(ctx context.Context, self *v1.Self, opts metav1.CreateOptions) (*v1.Self, error)
	Update(ctx context.Context, self *v1.Self, opts metav1.UpdateOptions) (*v1.Self, error)
	UpdateStatus(ctx context.Context, self *v1.Self, opts metav1.UpdateOptions) (*v1.Self, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Self, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SelfList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Self, err error)
	SelfExpansion
}

// selves implements SelfInterface
type selves struct {
	client rest.Interface
}

// newSelves returns a Selves
func newSelves(c *ManagementV1Client) *selves {
	return &selves{
		client: c.RESTClient(),
	}
}

// Get takes name of the self, and returns the corresponding self object, and an error if there is any.
func (c *selves) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Self, err error) {
	result = &v1.Self{}
	err = c.client.Get().
		Resource("selves").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Selves that match those selectors.
func (c *selves) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SelfList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SelfList{}
	err = c.client.Get().
		Resource("selves").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested selves.
func (c *selves) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("selves").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a self and creates it.  Returns the server's representation of the self, and an error, if there is any.
func (c *selves) Create(ctx context.Context, self *v1.Self, opts metav1.CreateOptions) (result *v1.Self, err error) {
	result = &v1.Self{}
	err = c.client.Post().
		Resource("selves").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(self).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a self and updates it. Returns the server's representation of the self, and an error, if there is any.
func (c *selves) Update(ctx context.Context, self *v1.Self, opts metav1.UpdateOptions) (result *v1.Self, err error) {
	result = &v1.Self{}
	err = c.client.Put().
		Resource("selves").
		Name(self.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(self).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *selves) UpdateStatus(ctx context.Context, self *v1.Self, opts metav1.UpdateOptions) (result *v1.Self, err error) {
	result = &v1.Self{}
	err = c.client.Put().
		Resource("selves").
		Name(self.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(self).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the self and deletes it. Returns an error if one occurs.
func (c *selves) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("selves").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *selves) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("selves").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched self.
func (c *selves) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Self, err error) {
	result = &v1.Self{}
	err = c.client.Patch(pt).
		Resource("selves").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
