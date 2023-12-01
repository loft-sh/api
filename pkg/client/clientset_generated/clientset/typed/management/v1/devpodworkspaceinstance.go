// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v3/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v3/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DevPodWorkspaceInstancesGetter has a method to return a DevPodWorkspaceInstanceInterface.
// A group's client should implement this interface.
type DevPodWorkspaceInstancesGetter interface {
	DevPodWorkspaceInstances(namespace string) DevPodWorkspaceInstanceInterface
}

// DevPodWorkspaceInstanceInterface has methods to work with DevPodWorkspaceInstance resources.
type DevPodWorkspaceInstanceInterface interface {
	Create(ctx context.Context, devPodWorkspaceInstance *v1.DevPodWorkspaceInstance, opts metav1.CreateOptions) (*v1.DevPodWorkspaceInstance, error)
	Update(ctx context.Context, devPodWorkspaceInstance *v1.DevPodWorkspaceInstance, opts metav1.UpdateOptions) (*v1.DevPodWorkspaceInstance, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DevPodWorkspaceInstance, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DevPodWorkspaceInstanceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DevPodWorkspaceInstance, err error)
	GetState(ctx context.Context, devPodWorkspaceInstanceName string, options metav1.GetOptions) (*v1.DevPodWorkspaceInstanceState, error)
	SetState(ctx context.Context, devPodWorkspaceInstanceName string, devPodWorkspaceInstanceState *v1.DevPodWorkspaceInstanceState, opts metav1.CreateOptions) (*v1.DevPodWorkspaceInstanceState, error)

	DevPodWorkspaceInstanceExpansion
}

// devPodWorkspaceInstances implements DevPodWorkspaceInstanceInterface
type devPodWorkspaceInstances struct {
	client rest.Interface
	ns     string
}

// newDevPodWorkspaceInstances returns a DevPodWorkspaceInstances
func newDevPodWorkspaceInstances(c *ManagementV1Client, namespace string) *devPodWorkspaceInstances {
	return &devPodWorkspaceInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the devPodWorkspaceInstance, and returns the corresponding devPodWorkspaceInstance object, and an error if there is any.
func (c *devPodWorkspaceInstances) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DevPodWorkspaceInstance, err error) {
	result = &v1.DevPodWorkspaceInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devpodworkspaceinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DevPodWorkspaceInstances that match those selectors.
func (c *devPodWorkspaceInstances) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DevPodWorkspaceInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DevPodWorkspaceInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devpodworkspaceinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested devPodWorkspaceInstances.
func (c *devPodWorkspaceInstances) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("devpodworkspaceinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a devPodWorkspaceInstance and creates it.  Returns the server's representation of the devPodWorkspaceInstance, and an error, if there is any.
func (c *devPodWorkspaceInstances) Create(ctx context.Context, devPodWorkspaceInstance *v1.DevPodWorkspaceInstance, opts metav1.CreateOptions) (result *v1.DevPodWorkspaceInstance, err error) {
	result = &v1.DevPodWorkspaceInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devpodworkspaceinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(devPodWorkspaceInstance).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a devPodWorkspaceInstance and updates it. Returns the server's representation of the devPodWorkspaceInstance, and an error, if there is any.
func (c *devPodWorkspaceInstances) Update(ctx context.Context, devPodWorkspaceInstance *v1.DevPodWorkspaceInstance, opts metav1.UpdateOptions) (result *v1.DevPodWorkspaceInstance, err error) {
	result = &v1.DevPodWorkspaceInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devpodworkspaceinstances").
		Name(devPodWorkspaceInstance.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(devPodWorkspaceInstance).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the devPodWorkspaceInstance and deletes it. Returns an error if one occurs.
func (c *devPodWorkspaceInstances) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devpodworkspaceinstances").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *devPodWorkspaceInstances) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devpodworkspaceinstances").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched devPodWorkspaceInstance.
func (c *devPodWorkspaceInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DevPodWorkspaceInstance, err error) {
	result = &v1.DevPodWorkspaceInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("devpodworkspaceinstances").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// GetState takes name of the devPodWorkspaceInstance, and returns the corresponding v1.DevPodWorkspaceInstanceState object, and an error if there is any.
func (c *devPodWorkspaceInstances) GetState(ctx context.Context, devPodWorkspaceInstanceName string, options metav1.GetOptions) (result *v1.DevPodWorkspaceInstanceState, err error) {
	result = &v1.DevPodWorkspaceInstanceState{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devpodworkspaceinstances").
		Name(devPodWorkspaceInstanceName).
		SubResource("state").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// SetState takes the representation of a devPodWorkspaceInstanceState and creates it.  Returns the server's representation of the devPodWorkspaceInstanceState, and an error, if there is any.
func (c *devPodWorkspaceInstances) SetState(ctx context.Context, devPodWorkspaceInstanceName string, devPodWorkspaceInstanceState *v1.DevPodWorkspaceInstanceState, opts metav1.CreateOptions) (result *v1.DevPodWorkspaceInstanceState, err error) {
	result = &v1.DevPodWorkspaceInstanceState{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devpodworkspaceinstances").
		Name(devPodWorkspaceInstanceName).
		SubResource("state").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(devPodWorkspaceInstanceState).
		Do(ctx).
		Into(result)
	return
}
