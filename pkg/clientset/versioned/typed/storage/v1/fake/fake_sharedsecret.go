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

// FakeSharedSecrets implements SharedSecretInterface
type FakeSharedSecrets struct {
	Fake *FakeStorageV1
	ns   string
}

var sharedsecretsResource = v1.SchemeGroupVersion.WithResource("sharedsecrets")

var sharedsecretsKind = v1.SchemeGroupVersion.WithKind("SharedSecret")

// Get takes name of the sharedSecret, and returns the corresponding sharedSecret object, and an error if there is any.
func (c *FakeSharedSecrets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SharedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharedsecretsResource, c.ns, name), &v1.SharedSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SharedSecret), err
}

// List takes label and field selectors, and returns the list of SharedSecrets that match those selectors.
func (c *FakeSharedSecrets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SharedSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharedsecretsResource, sharedsecretsKind, c.ns, opts), &v1.SharedSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.SharedSecretList{ListMeta: obj.(*v1.SharedSecretList).ListMeta}
	for _, item := range obj.(*v1.SharedSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sharedSecrets.
func (c *FakeSharedSecrets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharedsecretsResource, c.ns, opts))

}

// Create takes the representation of a sharedSecret and creates it.  Returns the server's representation of the sharedSecret, and an error, if there is any.
func (c *FakeSharedSecrets) Create(ctx context.Context, sharedSecret *v1.SharedSecret, opts metav1.CreateOptions) (result *v1.SharedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharedsecretsResource, c.ns, sharedSecret), &v1.SharedSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SharedSecret), err
}

// Update takes the representation of a sharedSecret and updates it. Returns the server's representation of the sharedSecret, and an error, if there is any.
func (c *FakeSharedSecrets) Update(ctx context.Context, sharedSecret *v1.SharedSecret, opts metav1.UpdateOptions) (result *v1.SharedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharedsecretsResource, c.ns, sharedSecret), &v1.SharedSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SharedSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSharedSecrets) UpdateStatus(ctx context.Context, sharedSecret *v1.SharedSecret, opts metav1.UpdateOptions) (*v1.SharedSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sharedsecretsResource, "status", c.ns, sharedSecret), &v1.SharedSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SharedSecret), err
}

// Delete takes name of the sharedSecret and deletes it. Returns an error if one occurs.
func (c *FakeSharedSecrets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(sharedsecretsResource, c.ns, name, opts), &v1.SharedSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSharedSecrets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sharedsecretsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.SharedSecretList{})
	return err
}

// Patch applies the patch and returns the patched sharedSecret.
func (c *FakeSharedSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SharedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharedsecretsResource, c.ns, name, pt, data, subresources...), &v1.SharedSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SharedSecret), err
}
