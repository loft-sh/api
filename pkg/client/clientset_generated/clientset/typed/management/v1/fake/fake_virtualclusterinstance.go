// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	managementv1 "github.com/loft-sh/api/v3/pkg/apis/management/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualClusterInstances implements VirtualClusterInstanceInterface
type FakeVirtualClusterInstances struct {
	Fake *FakeManagementV1
	ns   string
}

var virtualclusterinstancesResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "virtualclusterinstances"}

var virtualclusterinstancesKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "VirtualClusterInstance"}

// Get takes name of the virtualClusterInstance, and returns the corresponding virtualClusterInstance object, and an error if there is any.
func (c *FakeVirtualClusterInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.VirtualClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualclusterinstancesResource, c.ns, name), &managementv1.VirtualClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.VirtualClusterInstance), err
}

// List takes label and field selectors, and returns the list of VirtualClusterInstances that match those selectors.
func (c *FakeVirtualClusterInstances) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.VirtualClusterInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualclusterinstancesResource, virtualclusterinstancesKind, c.ns, opts), &managementv1.VirtualClusterInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.VirtualClusterInstanceList{ListMeta: obj.(*managementv1.VirtualClusterInstanceList).ListMeta}
	for _, item := range obj.(*managementv1.VirtualClusterInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualClusterInstances.
func (c *FakeVirtualClusterInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualclusterinstancesResource, c.ns, opts))

}

// Create takes the representation of a virtualClusterInstance and creates it.  Returns the server's representation of the virtualClusterInstance, and an error, if there is any.
func (c *FakeVirtualClusterInstances) Create(ctx context.Context, virtualClusterInstance *managementv1.VirtualClusterInstance, opts v1.CreateOptions) (result *managementv1.VirtualClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualclusterinstancesResource, c.ns, virtualClusterInstance), &managementv1.VirtualClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.VirtualClusterInstance), err
}

// Update takes the representation of a virtualClusterInstance and updates it. Returns the server's representation of the virtualClusterInstance, and an error, if there is any.
func (c *FakeVirtualClusterInstances) Update(ctx context.Context, virtualClusterInstance *managementv1.VirtualClusterInstance, opts v1.UpdateOptions) (result *managementv1.VirtualClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualclusterinstancesResource, c.ns, virtualClusterInstance), &managementv1.VirtualClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.VirtualClusterInstance), err
}

// Delete takes name of the virtualClusterInstance and deletes it. Returns an error if one occurs.
func (c *FakeVirtualClusterInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualclusterinstancesResource, c.ns, name, opts), &managementv1.VirtualClusterInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualClusterInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualclusterinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.VirtualClusterInstanceList{})
	return err
}

// Patch applies the patch and returns the patched virtualClusterInstance.
func (c *FakeVirtualClusterInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.VirtualClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualclusterinstancesResource, c.ns, name, pt, data, subresources...), &managementv1.VirtualClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.VirtualClusterInstance), err
}

// GetKubeConfig takes the representation of a virtualClusterInstanceKubeConfig and creates it.  Returns the server's representation of the virtualClusterInstanceKubeConfig, and an error, if there is any.
func (c *FakeVirtualClusterInstances) GetKubeConfig(ctx context.Context, virtualClusterInstanceName string, virtualClusterInstanceKubeConfig *managementv1.VirtualClusterInstanceKubeConfig, opts v1.CreateOptions) (result *managementv1.VirtualClusterInstanceKubeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateSubresourceAction(virtualclusterinstancesResource, virtualClusterInstanceName, "kubeconfig", c.ns, virtualClusterInstanceKubeConfig), &managementv1.VirtualClusterInstanceKubeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.VirtualClusterInstanceKubeConfig), err
}
