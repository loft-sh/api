// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	managementv1 "github.com/loft-sh/api/pkg/apis/management/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDirectClusterEndpointTokens implements DirectClusterEndpointTokenInterface
type FakeDirectClusterEndpointTokens struct {
	Fake *FakeManagementV1
}

var directclusterendpointtokensResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "directclusterendpointtokens"}

var directclusterendpointtokensKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "DirectClusterEndpointToken"}

// Get takes name of the directClusterEndpointToken, and returns the corresponding directClusterEndpointToken object, and an error if there is any.
func (c *FakeDirectClusterEndpointTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.DirectClusterEndpointToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(directclusterendpointtokensResource, name), &managementv1.DirectClusterEndpointToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.DirectClusterEndpointToken), err
}

// List takes label and field selectors, and returns the list of DirectClusterEndpointTokens that match those selectors.
func (c *FakeDirectClusterEndpointTokens) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.DirectClusterEndpointTokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(directclusterendpointtokensResource, directclusterendpointtokensKind, opts), &managementv1.DirectClusterEndpointTokenList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.DirectClusterEndpointTokenList{ListMeta: obj.(*managementv1.DirectClusterEndpointTokenList).ListMeta}
	for _, item := range obj.(*managementv1.DirectClusterEndpointTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested directClusterEndpointTokens.
func (c *FakeDirectClusterEndpointTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(directclusterendpointtokensResource, opts))
}

// Create takes the representation of a directClusterEndpointToken and creates it.  Returns the server's representation of the directClusterEndpointToken, and an error, if there is any.
func (c *FakeDirectClusterEndpointTokens) Create(ctx context.Context, directClusterEndpointToken *managementv1.DirectClusterEndpointToken, opts v1.CreateOptions) (result *managementv1.DirectClusterEndpointToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(directclusterendpointtokensResource, directClusterEndpointToken), &managementv1.DirectClusterEndpointToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.DirectClusterEndpointToken), err
}

// Update takes the representation of a directClusterEndpointToken and updates it. Returns the server's representation of the directClusterEndpointToken, and an error, if there is any.
func (c *FakeDirectClusterEndpointTokens) Update(ctx context.Context, directClusterEndpointToken *managementv1.DirectClusterEndpointToken, opts v1.UpdateOptions) (result *managementv1.DirectClusterEndpointToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(directclusterendpointtokensResource, directClusterEndpointToken), &managementv1.DirectClusterEndpointToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.DirectClusterEndpointToken), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDirectClusterEndpointTokens) UpdateStatus(ctx context.Context, directClusterEndpointToken *managementv1.DirectClusterEndpointToken, opts v1.UpdateOptions) (*managementv1.DirectClusterEndpointToken, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(directclusterendpointtokensResource, "status", directClusterEndpointToken), &managementv1.DirectClusterEndpointToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.DirectClusterEndpointToken), err
}

// Delete takes name of the directClusterEndpointToken and deletes it. Returns an error if one occurs.
func (c *FakeDirectClusterEndpointTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(directclusterendpointtokensResource, name), &managementv1.DirectClusterEndpointToken{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDirectClusterEndpointTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(directclusterendpointtokensResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.DirectClusterEndpointTokenList{})
	return err
}

// Patch applies the patch and returns the patched directClusterEndpointToken.
func (c *FakeDirectClusterEndpointTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.DirectClusterEndpointToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(directclusterendpointtokensResource, name, pt, data, subresources...), &managementv1.DirectClusterEndpointToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.DirectClusterEndpointToken), err
}