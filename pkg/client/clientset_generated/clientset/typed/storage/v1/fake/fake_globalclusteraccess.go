// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGlobalClusterAccesses implements GlobalClusterAccessInterface
type FakeGlobalClusterAccesses struct {
	Fake *FakeStorageV1
}

var globalclusteraccessesResource = schema.GroupVersionResource{Group: "storage.loft.sh", Version: "v1", Resource: "globalclusteraccesses"}

var globalclusteraccessesKind = schema.GroupVersionKind{Group: "storage.loft.sh", Version: "v1", Kind: "GlobalClusterAccess"}

// Get takes name of the globalClusterAccess, and returns the corresponding globalClusterAccess object, and an error if there is any.
func (c *FakeGlobalClusterAccesses) Get(ctx context.Context, name string, options v1.GetOptions) (result *storagev1.GlobalClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(globalclusteraccessesResource, name), &storagev1.GlobalClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.GlobalClusterAccess), err
}

// List takes label and field selectors, and returns the list of GlobalClusterAccesses that match those selectors.
func (c *FakeGlobalClusterAccesses) List(ctx context.Context, opts v1.ListOptions) (result *storagev1.GlobalClusterAccessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(globalclusteraccessesResource, globalclusteraccessesKind, opts), &storagev1.GlobalClusterAccessList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1.GlobalClusterAccessList{ListMeta: obj.(*storagev1.GlobalClusterAccessList).ListMeta}
	for _, item := range obj.(*storagev1.GlobalClusterAccessList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalClusterAccesses.
func (c *FakeGlobalClusterAccesses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(globalclusteraccessesResource, opts))
}

// Create takes the representation of a globalClusterAccess and creates it.  Returns the server's representation of the globalClusterAccess, and an error, if there is any.
func (c *FakeGlobalClusterAccesses) Create(ctx context.Context, globalClusterAccess *storagev1.GlobalClusterAccess, opts v1.CreateOptions) (result *storagev1.GlobalClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(globalclusteraccessesResource, globalClusterAccess), &storagev1.GlobalClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.GlobalClusterAccess), err
}

// Update takes the representation of a globalClusterAccess and updates it. Returns the server's representation of the globalClusterAccess, and an error, if there is any.
func (c *FakeGlobalClusterAccesses) Update(ctx context.Context, globalClusterAccess *storagev1.GlobalClusterAccess, opts v1.UpdateOptions) (result *storagev1.GlobalClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(globalclusteraccessesResource, globalClusterAccess), &storagev1.GlobalClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.GlobalClusterAccess), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlobalClusterAccesses) UpdateStatus(ctx context.Context, globalClusterAccess *storagev1.GlobalClusterAccess, opts v1.UpdateOptions) (*storagev1.GlobalClusterAccess, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(globalclusteraccessesResource, "status", globalClusterAccess), &storagev1.GlobalClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.GlobalClusterAccess), err
}

// Delete takes name of the globalClusterAccess and deletes it. Returns an error if one occurs.
func (c *FakeGlobalClusterAccesses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(globalclusteraccessesResource, name), &storagev1.GlobalClusterAccess{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalClusterAccesses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(globalclusteraccessesResource, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1.GlobalClusterAccessList{})
	return err
}

// Patch applies the patch and returns the patched globalClusterAccess.
func (c *FakeGlobalClusterAccesses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *storagev1.GlobalClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(globalclusteraccessesResource, name, pt, data, subresources...), &storagev1.GlobalClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.GlobalClusterAccess), err
}
