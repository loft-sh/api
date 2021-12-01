// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	managementv1 "github.com/loft-sh/api/v2/pkg/apis/management/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLoftUpgrades implements LoftUpgradeInterface
type FakeLoftUpgrades struct {
	Fake *FakeManagementV1
}

var loftupgradesResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "loftupgrades"}

var loftupgradesKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "LoftUpgrade"}

// Get takes name of the loftUpgrade, and returns the corresponding loftUpgrade object, and an error if there is any.
func (c *FakeLoftUpgrades) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.LoftUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(loftupgradesResource, name), &managementv1.LoftUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.LoftUpgrade), err
}

// List takes label and field selectors, and returns the list of LoftUpgrades that match those selectors.
func (c *FakeLoftUpgrades) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.LoftUpgradeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(loftupgradesResource, loftupgradesKind, opts), &managementv1.LoftUpgradeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.LoftUpgradeList{ListMeta: obj.(*managementv1.LoftUpgradeList).ListMeta}
	for _, item := range obj.(*managementv1.LoftUpgradeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loftUpgrades.
func (c *FakeLoftUpgrades) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(loftupgradesResource, opts))
}

// Create takes the representation of a loftUpgrade and creates it.  Returns the server's representation of the loftUpgrade, and an error, if there is any.
func (c *FakeLoftUpgrades) Create(ctx context.Context, loftUpgrade *managementv1.LoftUpgrade, opts v1.CreateOptions) (result *managementv1.LoftUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(loftupgradesResource, loftUpgrade), &managementv1.LoftUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.LoftUpgrade), err
}

// Update takes the representation of a loftUpgrade and updates it. Returns the server's representation of the loftUpgrade, and an error, if there is any.
func (c *FakeLoftUpgrades) Update(ctx context.Context, loftUpgrade *managementv1.LoftUpgrade, opts v1.UpdateOptions) (result *managementv1.LoftUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(loftupgradesResource, loftUpgrade), &managementv1.LoftUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.LoftUpgrade), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoftUpgrades) UpdateStatus(ctx context.Context, loftUpgrade *managementv1.LoftUpgrade, opts v1.UpdateOptions) (*managementv1.LoftUpgrade, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(loftupgradesResource, "status", loftUpgrade), &managementv1.LoftUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.LoftUpgrade), err
}

// Delete takes name of the loftUpgrade and deletes it. Returns an error if one occurs.
func (c *FakeLoftUpgrades) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(loftupgradesResource, name), &managementv1.LoftUpgrade{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoftUpgrades) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(loftupgradesResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.LoftUpgradeList{})
	return err
}

// Patch applies the patch and returns the patched loftUpgrade.
func (c *FakeLoftUpgrades) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.LoftUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(loftupgradesResource, name, pt, data, subresources...), &managementv1.LoftUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.LoftUpgrade), err
}
