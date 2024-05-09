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

// FakeClusterAccesses implements ClusterAccessInterface
type FakeClusterAccesses struct {
	Fake *FakeManagementV1
}

var clusteraccessesResource = v1.SchemeGroupVersion.WithResource("clusteraccesses")

var clusteraccessesKind = v1.SchemeGroupVersion.WithKind("ClusterAccess")

// Get takes name of the clusterAccess, and returns the corresponding clusterAccess object, and an error if there is any.
func (c *FakeClusterAccesses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusteraccessesResource, name), &v1.ClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterAccess), err
}

// List takes label and field selectors, and returns the list of ClusterAccesses that match those selectors.
func (c *FakeClusterAccesses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterAccessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusteraccessesResource, clusteraccessesKind, opts), &v1.ClusterAccessList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ClusterAccessList{ListMeta: obj.(*v1.ClusterAccessList).ListMeta}
	for _, item := range obj.(*v1.ClusterAccessList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterAccesses.
func (c *FakeClusterAccesses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusteraccessesResource, opts))
}

// Create takes the representation of a clusterAccess and creates it.  Returns the server's representation of the clusterAccess, and an error, if there is any.
func (c *FakeClusterAccesses) Create(ctx context.Context, clusterAccess *v1.ClusterAccess, opts metav1.CreateOptions) (result *v1.ClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusteraccessesResource, clusterAccess), &v1.ClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterAccess), err
}

// Update takes the representation of a clusterAccess and updates it. Returns the server's representation of the clusterAccess, and an error, if there is any.
func (c *FakeClusterAccesses) Update(ctx context.Context, clusterAccess *v1.ClusterAccess, opts metav1.UpdateOptions) (result *v1.ClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusteraccessesResource, clusterAccess), &v1.ClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterAccess), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterAccesses) UpdateStatus(ctx context.Context, clusterAccess *v1.ClusterAccess, opts metav1.UpdateOptions) (*v1.ClusterAccess, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusteraccessesResource, "status", clusterAccess), &v1.ClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterAccess), err
}

// Delete takes name of the clusterAccess and deletes it. Returns an error if one occurs.
func (c *FakeClusterAccesses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusteraccessesResource, name, opts), &v1.ClusterAccess{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterAccesses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusteraccessesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ClusterAccessList{})
	return err
}

// Patch applies the patch and returns the patched clusterAccess.
func (c *FakeClusterAccesses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusteraccessesResource, name, pt, data, subresources...), &v1.ClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterAccess), err
}
