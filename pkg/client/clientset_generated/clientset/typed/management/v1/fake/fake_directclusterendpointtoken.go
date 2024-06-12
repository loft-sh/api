// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/loft-sh/api/v3/pkg/apis/management/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDirectClusterEndpointTokens implements DirectClusterEndpointTokenInterface
type FakeDirectClusterEndpointTokens struct {
	Fake *FakeManagementV1
}

var directclusterendpointtokensResource = v1.SchemeGroupVersion.WithResource("directclusterendpointtokens")

var directclusterendpointtokensKind = v1.SchemeGroupVersion.WithKind("DirectClusterEndpointToken")

// Get takes name of the directClusterEndpointToken, and returns the corresponding directClusterEndpointToken object, and an error if there is any.
func (c *FakeDirectClusterEndpointTokens) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DirectClusterEndpointToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(directclusterendpointtokensResource, name), &v1.DirectClusterEndpointToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DirectClusterEndpointToken), err
}

// List takes label and field selectors, and returns the list of DirectClusterEndpointTokens that match those selectors.
func (c *FakeDirectClusterEndpointTokens) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DirectClusterEndpointTokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(directclusterendpointtokensResource, directclusterendpointtokensKind, opts), &v1.DirectClusterEndpointTokenList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.DirectClusterEndpointTokenList{ListMeta: obj.(*v1.DirectClusterEndpointTokenList).ListMeta}
	for _, item := range obj.(*v1.DirectClusterEndpointTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested directClusterEndpointTokens.
func (c *FakeDirectClusterEndpointTokens) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(directclusterendpointtokensResource, opts))
}

// Create takes the representation of a directClusterEndpointToken and creates it.  Returns the server's representation of the directClusterEndpointToken, and an error, if there is any.
func (c *FakeDirectClusterEndpointTokens) Create(ctx context.Context, directClusterEndpointToken *v1.DirectClusterEndpointToken, opts metav1.CreateOptions) (result *v1.DirectClusterEndpointToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(directclusterendpointtokensResource, directClusterEndpointToken), &v1.DirectClusterEndpointToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DirectClusterEndpointToken), err
}

// Update takes the representation of a directClusterEndpointToken and updates it. Returns the server's representation of the directClusterEndpointToken, and an error, if there is any.
func (c *FakeDirectClusterEndpointTokens) Update(ctx context.Context, directClusterEndpointToken *v1.DirectClusterEndpointToken, opts metav1.UpdateOptions) (result *v1.DirectClusterEndpointToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(directclusterendpointtokensResource, directClusterEndpointToken), &v1.DirectClusterEndpointToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DirectClusterEndpointToken), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDirectClusterEndpointTokens) UpdateStatus(ctx context.Context, directClusterEndpointToken *v1.DirectClusterEndpointToken, opts metav1.UpdateOptions) (*v1.DirectClusterEndpointToken, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(directclusterendpointtokensResource, "status", directClusterEndpointToken), &v1.DirectClusterEndpointToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DirectClusterEndpointToken), err
}

// Delete takes name of the directClusterEndpointToken and deletes it. Returns an error if one occurs.
func (c *FakeDirectClusterEndpointTokens) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(directclusterendpointtokensResource, name, opts), &v1.DirectClusterEndpointToken{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDirectClusterEndpointTokens) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(directclusterendpointtokensResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.DirectClusterEndpointTokenList{})
	return err
}

// Patch applies the patch and returns the patched directClusterEndpointToken.
func (c *FakeDirectClusterEndpointTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DirectClusterEndpointToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(directclusterendpointtokensResource, name, pt, data, subresources...), &v1.DirectClusterEndpointToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DirectClusterEndpointToken), err
}
