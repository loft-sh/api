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

// FakeClusters implements ClusterInterface
type FakeClusters struct {
	Fake *FakeStorageV1
}

var clustersResource = schema.GroupVersionResource{Group: "storage.loft.sh", Version: "v1", Resource: "clusters"}

var clustersKind = schema.GroupVersionKind{Group: "storage.loft.sh", Version: "v1", Kind: "Cluster"}

// Get takes name of the cluster, and returns the corresponding cluster object, and an error if there is any.
func (c *FakeClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *storagev1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustersResource, name), &storagev1.Cluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.Cluster), err
}

// List takes label and field selectors, and returns the list of Clusters that match those selectors.
func (c *FakeClusters) List(ctx context.Context, opts v1.ListOptions) (result *storagev1.ClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustersResource, clustersKind, opts), &storagev1.ClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1.ClusterList{ListMeta: obj.(*storagev1.ClusterList).ListMeta}
	for _, item := range obj.(*storagev1.ClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusters.
func (c *FakeClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustersResource, opts))
}

// Create takes the representation of a cluster and creates it.  Returns the server's representation of the cluster, and an error, if there is any.
func (c *FakeClusters) Create(ctx context.Context, cluster *storagev1.Cluster, opts v1.CreateOptions) (result *storagev1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustersResource, cluster), &storagev1.Cluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.Cluster), err
}

// Update takes the representation of a cluster and updates it. Returns the server's representation of the cluster, and an error, if there is any.
func (c *FakeClusters) Update(ctx context.Context, cluster *storagev1.Cluster, opts v1.UpdateOptions) (result *storagev1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustersResource, cluster), &storagev1.Cluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.Cluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusters) UpdateStatus(ctx context.Context, cluster *storagev1.Cluster, opts v1.UpdateOptions) (*storagev1.Cluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustersResource, "status", cluster), &storagev1.Cluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.Cluster), err
}

// Delete takes name of the cluster and deletes it. Returns an error if one occurs.
func (c *FakeClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clustersResource, name), &storagev1.Cluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustersResource, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1.ClusterList{})
	return err
}

// Patch applies the patch and returns the patched cluster.
func (c *FakeClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *storagev1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustersResource, name, pt, data, subresources...), &storagev1.Cluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.Cluster), err
}
