// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLoftUpgrades implements LoftUpgradeInterface
type FakeLoftUpgrades struct {
	Fake *FakeManagementV1
}

var loftupgradesResource = v1.SchemeGroupVersion.WithResource("loftupgrades")

var loftupgradesKind = v1.SchemeGroupVersion.WithKind("LoftUpgrade")

// Get takes name of the loftUpgrade, and returns the corresponding loftUpgrade object, and an error if there is any.
func (c *FakeLoftUpgrades) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LoftUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(loftupgradesResource, name), &v1.LoftUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LoftUpgrade), err
}

// List takes label and field selectors, and returns the list of LoftUpgrades that match those selectors.
func (c *FakeLoftUpgrades) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LoftUpgradeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(loftupgradesResource, loftupgradesKind, opts), &v1.LoftUpgradeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.LoftUpgradeList{ListMeta: obj.(*v1.LoftUpgradeList).ListMeta}
	for _, item := range obj.(*v1.LoftUpgradeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loftUpgrades.
func (c *FakeLoftUpgrades) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(loftupgradesResource, opts))
}

// Create takes the representation of a loftUpgrade and creates it.  Returns the server's representation of the loftUpgrade, and an error, if there is any.
func (c *FakeLoftUpgrades) Create(ctx context.Context, loftUpgrade *v1.LoftUpgrade, opts metav1.CreateOptions) (result *v1.LoftUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(loftupgradesResource, loftUpgrade), &v1.LoftUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LoftUpgrade), err
}

// Update takes the representation of a loftUpgrade and updates it. Returns the server's representation of the loftUpgrade, and an error, if there is any.
func (c *FakeLoftUpgrades) Update(ctx context.Context, loftUpgrade *v1.LoftUpgrade, opts metav1.UpdateOptions) (result *v1.LoftUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(loftupgradesResource, loftUpgrade), &v1.LoftUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LoftUpgrade), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoftUpgrades) UpdateStatus(ctx context.Context, loftUpgrade *v1.LoftUpgrade, opts metav1.UpdateOptions) (*v1.LoftUpgrade, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(loftupgradesResource, "status", loftUpgrade), &v1.LoftUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LoftUpgrade), err
}

// Delete takes name of the loftUpgrade and deletes it. Returns an error if one occurs.
func (c *FakeLoftUpgrades) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(loftupgradesResource, name, opts), &v1.LoftUpgrade{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoftUpgrades) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(loftupgradesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.LoftUpgradeList{})
	return err
}

// Patch applies the patch and returns the patched loftUpgrade.
func (c *FakeLoftUpgrades) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LoftUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(loftupgradesResource, name, pt, data, subresources...), &v1.LoftUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LoftUpgrade), err
}
