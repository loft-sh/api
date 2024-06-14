// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualClusterInstances implements VirtualClusterInstanceInterface
type FakeVirtualClusterInstances struct {
	Fake *FakeStorageV1
	ns   string
}

var virtualclusterinstancesResource = v1.SchemeGroupVersion.WithResource("virtualclusterinstances")

var virtualclusterinstancesKind = v1.SchemeGroupVersion.WithKind("VirtualClusterInstance")

// Get takes name of the virtualClusterInstance, and returns the corresponding virtualClusterInstance object, and an error if there is any.
func (c *FakeVirtualClusterInstances) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualclusterinstancesResource, c.ns, name), &v1.VirtualClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualClusterInstance), err
}

// List takes label and field selectors, and returns the list of VirtualClusterInstances that match those selectors.
func (c *FakeVirtualClusterInstances) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualClusterInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualclusterinstancesResource, virtualclusterinstancesKind, c.ns, opts), &v1.VirtualClusterInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.VirtualClusterInstanceList{ListMeta: obj.(*v1.VirtualClusterInstanceList).ListMeta}
	for _, item := range obj.(*v1.VirtualClusterInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualClusterInstances.
func (c *FakeVirtualClusterInstances) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualclusterinstancesResource, c.ns, opts))

}

// Create takes the representation of a virtualClusterInstance and creates it.  Returns the server's representation of the virtualClusterInstance, and an error, if there is any.
func (c *FakeVirtualClusterInstances) Create(ctx context.Context, virtualClusterInstance *v1.VirtualClusterInstance, opts metav1.CreateOptions) (result *v1.VirtualClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualclusterinstancesResource, c.ns, virtualClusterInstance), &v1.VirtualClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualClusterInstance), err
}

// Update takes the representation of a virtualClusterInstance and updates it. Returns the server's representation of the virtualClusterInstance, and an error, if there is any.
func (c *FakeVirtualClusterInstances) Update(ctx context.Context, virtualClusterInstance *v1.VirtualClusterInstance, opts metav1.UpdateOptions) (result *v1.VirtualClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualclusterinstancesResource, c.ns, virtualClusterInstance), &v1.VirtualClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualClusterInstance), err
}

// Delete takes name of the virtualClusterInstance and deletes it. Returns an error if one occurs.
func (c *FakeVirtualClusterInstances) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualclusterinstancesResource, c.ns, name, opts), &v1.VirtualClusterInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualClusterInstances) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualclusterinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.VirtualClusterInstanceList{})
	return err
}

// Patch applies the patch and returns the patched virtualClusterInstance.
func (c *FakeVirtualClusterInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualclusterinstancesResource, c.ns, name, pt, data, subresources...), &v1.VirtualClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualClusterInstance), err
}