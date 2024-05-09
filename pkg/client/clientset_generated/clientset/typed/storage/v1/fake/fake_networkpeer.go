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

// FakeNetworkPeers implements NetworkPeerInterface
type FakeNetworkPeers struct {
	Fake *FakeStorageV1
}

var networkpeersResource = v1.SchemeGroupVersion.WithResource("networkpeers")

var networkpeersKind = v1.SchemeGroupVersion.WithKind("NetworkPeer")

// Get takes name of the networkPeer, and returns the corresponding networkPeer object, and an error if there is any.
func (c *FakeNetworkPeers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NetworkPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(networkpeersResource, name), &v1.NetworkPeer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.NetworkPeer), err
}

// List takes label and field selectors, and returns the list of NetworkPeers that match those selectors.
func (c *FakeNetworkPeers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NetworkPeerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(networkpeersResource, networkpeersKind, opts), &v1.NetworkPeerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.NetworkPeerList{ListMeta: obj.(*v1.NetworkPeerList).ListMeta}
	for _, item := range obj.(*v1.NetworkPeerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPeers.
func (c *FakeNetworkPeers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(networkpeersResource, opts))
}

// Create takes the representation of a networkPeer and creates it.  Returns the server's representation of the networkPeer, and an error, if there is any.
func (c *FakeNetworkPeers) Create(ctx context.Context, networkPeer *v1.NetworkPeer, opts metav1.CreateOptions) (result *v1.NetworkPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(networkpeersResource, networkPeer), &v1.NetworkPeer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.NetworkPeer), err
}

// Update takes the representation of a networkPeer and updates it. Returns the server's representation of the networkPeer, and an error, if there is any.
func (c *FakeNetworkPeers) Update(ctx context.Context, networkPeer *v1.NetworkPeer, opts metav1.UpdateOptions) (result *v1.NetworkPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(networkpeersResource, networkPeer), &v1.NetworkPeer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.NetworkPeer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkPeers) UpdateStatus(ctx context.Context, networkPeer *v1.NetworkPeer, opts metav1.UpdateOptions) (*v1.NetworkPeer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(networkpeersResource, "status", networkPeer), &v1.NetworkPeer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.NetworkPeer), err
}

// Delete takes name of the networkPeer and deletes it. Returns an error if one occurs.
func (c *FakeNetworkPeers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(networkpeersResource, name, opts), &v1.NetworkPeer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkPeers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(networkpeersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.NetworkPeerList{})
	return err
}

// Patch applies the patch and returns the patched networkPeer.
func (c *FakeNetworkPeers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(networkpeersResource, name, pt, data, subresources...), &v1.NetworkPeer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.NetworkPeer), err
}
