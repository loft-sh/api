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

// LoftUpgradesGetter has a method to return a LoftUpgradeInterface.
// A group's client should implement this interface.
type LoftUpgradesGetter interface {
	LoftUpgrades() LoftUpgradeInterface
}

// LoftUpgradeInterface has methods to work with LoftUpgrade resources.
type LoftUpgradeInterface interface {
	Create(ctx context.Context, loftUpgrade *v1.LoftUpgrade, opts metav1.CreateOptions) (*v1.LoftUpgrade, error)
	Update(ctx context.Context, loftUpgrade *v1.LoftUpgrade, opts metav1.UpdateOptions) (*v1.LoftUpgrade, error)
	UpdateStatus(ctx context.Context, loftUpgrade *v1.LoftUpgrade, opts metav1.UpdateOptions) (*v1.LoftUpgrade, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.LoftUpgrade, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LoftUpgradeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LoftUpgrade, err error)
	LoftUpgradeExpansion
}

// loftUpgrades implements LoftUpgradeInterface
type loftUpgrades struct {
	client rest.Interface
}

// newLoftUpgrades returns a LoftUpgrades
func newLoftUpgrades(c *ManagementV1Client) *loftUpgrades {
	return &loftUpgrades{
		client: c.RESTClient(),
	}
}

// Get takes name of the loftUpgrade, and returns the corresponding loftUpgrade object, and an error if there is any.
func (c *loftUpgrades) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LoftUpgrade, err error) {
	result = &v1.LoftUpgrade{}
	err = c.client.Get().
		Resource("loftupgrades").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LoftUpgrades that match those selectors.
func (c *loftUpgrades) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LoftUpgradeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LoftUpgradeList{}
	err = c.client.Get().
		Resource("loftupgrades").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loftUpgrades.
func (c *loftUpgrades) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("loftupgrades").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a loftUpgrade and creates it.  Returns the server's representation of the loftUpgrade, and an error, if there is any.
func (c *loftUpgrades) Create(ctx context.Context, loftUpgrade *v1.LoftUpgrade, opts metav1.CreateOptions) (result *v1.LoftUpgrade, err error) {
	result = &v1.LoftUpgrade{}
	err = c.client.Post().
		Resource("loftupgrades").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loftUpgrade).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a loftUpgrade and updates it. Returns the server's representation of the loftUpgrade, and an error, if there is any.
func (c *loftUpgrades) Update(ctx context.Context, loftUpgrade *v1.LoftUpgrade, opts metav1.UpdateOptions) (result *v1.LoftUpgrade, err error) {
	result = &v1.LoftUpgrade{}
	err = c.client.Put().
		Resource("loftupgrades").
		Name(loftUpgrade.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loftUpgrade).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *loftUpgrades) UpdateStatus(ctx context.Context, loftUpgrade *v1.LoftUpgrade, opts metav1.UpdateOptions) (result *v1.LoftUpgrade, err error) {
	result = &v1.LoftUpgrade{}
	err = c.client.Put().
		Resource("loftupgrades").
		Name(loftUpgrade.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loftUpgrade).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the loftUpgrade and deletes it. Returns an error if one occurs.
func (c *loftUpgrades) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("loftupgrades").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loftUpgrades) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("loftupgrades").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched loftUpgrade.
func (c *loftUpgrades) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LoftUpgrade, err error) {
	result = &v1.LoftUpgrade{}
	err = c.client.Patch(pt).
		Resource("loftupgrades").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
