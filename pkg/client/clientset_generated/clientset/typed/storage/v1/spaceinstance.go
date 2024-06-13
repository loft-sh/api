// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	scheme "github.com/loft-sh/api/v3/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SpaceInstancesGetter has a method to return a SpaceInstanceInterface.
// A group's client should implement this interface.
type SpaceInstancesGetter interface {
	SpaceInstances(namespace string) SpaceInstanceInterface
}

// SpaceInstanceInterface has methods to work with SpaceInstance resources.
type SpaceInstanceInterface interface {
	Create(ctx context.Context, spaceInstance *v1.SpaceInstance, opts metav1.CreateOptions) (*v1.SpaceInstance, error)
	Update(ctx context.Context, spaceInstance *v1.SpaceInstance, opts metav1.UpdateOptions) (*v1.SpaceInstance, error)
	UpdateStatus(ctx context.Context, spaceInstance *v1.SpaceInstance, opts metav1.UpdateOptions) (*v1.SpaceInstance, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SpaceInstance, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SpaceInstanceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SpaceInstance, err error)
	SpaceInstanceExpansion
}

// spaceInstances implements SpaceInstanceInterface
type spaceInstances struct {
	client rest.Interface
	ns     string
}

// newSpaceInstances returns a SpaceInstances
func newSpaceInstances(c *StorageV1Client, namespace string) *spaceInstances {
	return &spaceInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the spaceInstance, and returns the corresponding spaceInstance object, and an error if there is any.
func (c *spaceInstances) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SpaceInstance, err error) {
	result = &v1.SpaceInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spaceinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SpaceInstances that match those selectors.
func (c *spaceInstances) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SpaceInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SpaceInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spaceinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spaceInstances.
func (c *spaceInstances) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("spaceinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a spaceInstance and creates it.  Returns the server's representation of the spaceInstance, and an error, if there is any.
func (c *spaceInstances) Create(ctx context.Context, spaceInstance *v1.SpaceInstance, opts metav1.CreateOptions) (result *v1.SpaceInstance, err error) {
	result = &v1.SpaceInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("spaceinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spaceInstance).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a spaceInstance and updates it. Returns the server's representation of the spaceInstance, and an error, if there is any.
func (c *spaceInstances) Update(ctx context.Context, spaceInstance *v1.SpaceInstance, opts metav1.UpdateOptions) (result *v1.SpaceInstance, err error) {
	result = &v1.SpaceInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spaceinstances").
		Name(spaceInstance.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spaceInstance).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *spaceInstances) UpdateStatus(ctx context.Context, spaceInstance *v1.SpaceInstance, opts metav1.UpdateOptions) (result *v1.SpaceInstance, err error) {
	result = &v1.SpaceInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spaceinstances").
		Name(spaceInstance.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spaceInstance).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the spaceInstance and deletes it. Returns an error if one occurs.
func (c *spaceInstances) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spaceinstances").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spaceInstances) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spaceinstances").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched spaceInstance.
func (c *spaceInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SpaceInstance, err error) {
	result = &v1.SpaceInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("spaceinstances").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
