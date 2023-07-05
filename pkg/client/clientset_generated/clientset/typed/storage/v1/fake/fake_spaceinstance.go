// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSpaceInstances implements SpaceInstanceInterface
type FakeSpaceInstances struct {
	Fake *FakeStorageV1
	ns   string
}

var spaceinstancesResource = v1.SchemeGroupVersion.WithResource("spaceinstances")

var spaceinstancesKind = v1.SchemeGroupVersion.WithKind("SpaceInstance")

// Get takes name of the spaceInstance, and returns the corresponding spaceInstance object, and an error if there is any.
func (c *FakeSpaceInstances) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SpaceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(spaceinstancesResource, c.ns, name), &v1.SpaceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SpaceInstance), err
}

// List takes label and field selectors, and returns the list of SpaceInstances that match those selectors.
func (c *FakeSpaceInstances) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SpaceInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(spaceinstancesResource, spaceinstancesKind, c.ns, opts), &v1.SpaceInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.SpaceInstanceList{ListMeta: obj.(*v1.SpaceInstanceList).ListMeta}
	for _, item := range obj.(*v1.SpaceInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spaceInstances.
func (c *FakeSpaceInstances) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(spaceinstancesResource, c.ns, opts))

}

// Create takes the representation of a spaceInstance and creates it.  Returns the server's representation of the spaceInstance, and an error, if there is any.
func (c *FakeSpaceInstances) Create(ctx context.Context, spaceInstance *v1.SpaceInstance, opts metav1.CreateOptions) (result *v1.SpaceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(spaceinstancesResource, c.ns, spaceInstance), &v1.SpaceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SpaceInstance), err
}

// Update takes the representation of a spaceInstance and updates it. Returns the server's representation of the spaceInstance, and an error, if there is any.
func (c *FakeSpaceInstances) Update(ctx context.Context, spaceInstance *v1.SpaceInstance, opts metav1.UpdateOptions) (result *v1.SpaceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(spaceinstancesResource, c.ns, spaceInstance), &v1.SpaceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SpaceInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpaceInstances) UpdateStatus(ctx context.Context, spaceInstance *v1.SpaceInstance, opts metav1.UpdateOptions) (*v1.SpaceInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(spaceinstancesResource, "status", c.ns, spaceInstance), &v1.SpaceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SpaceInstance), err
}

// Delete takes name of the spaceInstance and deletes it. Returns an error if one occurs.
func (c *FakeSpaceInstances) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(spaceinstancesResource, c.ns, name, opts), &v1.SpaceInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpaceInstances) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(spaceinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.SpaceInstanceList{})
	return err
}

// Patch applies the patch and returns the patched spaceInstance.
func (c *FakeSpaceInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SpaceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(spaceinstancesResource, c.ns, name, pt, data, subresources...), &v1.SpaceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SpaceInstance), err
}
