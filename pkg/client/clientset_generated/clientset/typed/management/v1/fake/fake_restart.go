// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	managementv1 "github.com/loft-sh/api/pkg/apis/management/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRestarts implements RestartInterface
type FakeRestarts struct {
	Fake *FakeManagementV1
}

var restartsResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "restarts"}

var restartsKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "Restart"}

// Get takes name of the restart, and returns the corresponding restart object, and an error if there is any.
func (c *FakeRestarts) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.Restart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(restartsResource, name), &managementv1.Restart{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Restart), err
}

// List takes label and field selectors, and returns the list of Restarts that match those selectors.
func (c *FakeRestarts) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.RestartList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(restartsResource, restartsKind, opts), &managementv1.RestartList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.RestartList{ListMeta: obj.(*managementv1.RestartList).ListMeta}
	for _, item := range obj.(*managementv1.RestartList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested restarts.
func (c *FakeRestarts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(restartsResource, opts))
}

// Create takes the representation of a restart and creates it.  Returns the server's representation of the restart, and an error, if there is any.
func (c *FakeRestarts) Create(ctx context.Context, restart *managementv1.Restart, opts v1.CreateOptions) (result *managementv1.Restart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(restartsResource, restart), &managementv1.Restart{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Restart), err
}

// Update takes the representation of a restart and updates it. Returns the server's representation of the restart, and an error, if there is any.
func (c *FakeRestarts) Update(ctx context.Context, restart *managementv1.Restart, opts v1.UpdateOptions) (result *managementv1.Restart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(restartsResource, restart), &managementv1.Restart{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Restart), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRestarts) UpdateStatus(ctx context.Context, restart *managementv1.Restart, opts v1.UpdateOptions) (*managementv1.Restart, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(restartsResource, "status", restart), &managementv1.Restart{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Restart), err
}

// Delete takes name of the restart and deletes it. Returns an error if one occurs.
func (c *FakeRestarts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(restartsResource, name), &managementv1.Restart{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRestarts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(restartsResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.RestartList{})
	return err
}

// Patch applies the patch and returns the patched restart.
func (c *FakeRestarts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.Restart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(restartsResource, name, pt, data, subresources...), &managementv1.Restart{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Restart), err
}
