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

// FakeClusterConnects implements ClusterConnectInterface
type FakeClusterConnects struct {
	Fake *FakeManagementV1
}

var clusterconnectsResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "clusterconnects"}

var clusterconnectsKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "ClusterConnect"}

// Get takes name of the clusterConnect, and returns the corresponding clusterConnect object, and an error if there is any.
func (c *FakeClusterConnects) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.ClusterConnect, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterconnectsResource, name), &managementv1.ClusterConnect{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterConnect), err
}

// List takes label and field selectors, and returns the list of ClusterConnects that match those selectors.
func (c *FakeClusterConnects) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.ClusterConnectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterconnectsResource, clusterconnectsKind, opts), &managementv1.ClusterConnectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.ClusterConnectList{ListMeta: obj.(*managementv1.ClusterConnectList).ListMeta}
	for _, item := range obj.(*managementv1.ClusterConnectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterConnects.
func (c *FakeClusterConnects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterconnectsResource, opts))
}

// Create takes the representation of a clusterConnect and creates it.  Returns the server's representation of the clusterConnect, and an error, if there is any.
func (c *FakeClusterConnects) Create(ctx context.Context, clusterConnect *managementv1.ClusterConnect, opts v1.CreateOptions) (result *managementv1.ClusterConnect, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterconnectsResource, clusterConnect), &managementv1.ClusterConnect{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterConnect), err
}

// Update takes the representation of a clusterConnect and updates it. Returns the server's representation of the clusterConnect, and an error, if there is any.
func (c *FakeClusterConnects) Update(ctx context.Context, clusterConnect *managementv1.ClusterConnect, opts v1.UpdateOptions) (result *managementv1.ClusterConnect, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterconnectsResource, clusterConnect), &managementv1.ClusterConnect{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterConnect), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterConnects) UpdateStatus(ctx context.Context, clusterConnect *managementv1.ClusterConnect, opts v1.UpdateOptions) (*managementv1.ClusterConnect, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterconnectsResource, "status", clusterConnect), &managementv1.ClusterConnect{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterConnect), err
}

// Delete takes name of the clusterConnect and deletes it. Returns an error if one occurs.
func (c *FakeClusterConnects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterconnectsResource, name, opts), &managementv1.ClusterConnect{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterConnects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterconnectsResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.ClusterConnectList{})
	return err
}

// Patch applies the patch and returns the patched clusterConnect.
func (c *FakeClusterConnects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.ClusterConnect, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterconnectsResource, name, pt, data, subresources...), &managementv1.ClusterConnect{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterConnect), err
}
