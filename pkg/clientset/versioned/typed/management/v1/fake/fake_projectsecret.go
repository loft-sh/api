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

// FakeProjectSecrets implements ProjectSecretInterface
type FakeProjectSecrets struct {
	Fake *FakeManagementV1
	ns   string
}

var projectsecretsResource = v1.SchemeGroupVersion.WithResource("projectsecrets")

var projectsecretsKind = v1.SchemeGroupVersion.WithKind("ProjectSecret")

// Get takes name of the projectSecret, and returns the corresponding projectSecret object, and an error if there is any.
func (c *FakeProjectSecrets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ProjectSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(projectsecretsResource, c.ns, name), &v1.ProjectSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ProjectSecret), err
}

// List takes label and field selectors, and returns the list of ProjectSecrets that match those selectors.
func (c *FakeProjectSecrets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ProjectSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(projectsecretsResource, projectsecretsKind, c.ns, opts), &v1.ProjectSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ProjectSecretList{ListMeta: obj.(*v1.ProjectSecretList).ListMeta}
	for _, item := range obj.(*v1.ProjectSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projectSecrets.
func (c *FakeProjectSecrets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(projectsecretsResource, c.ns, opts))

}

// Create takes the representation of a projectSecret and creates it.  Returns the server's representation of the projectSecret, and an error, if there is any.
func (c *FakeProjectSecrets) Create(ctx context.Context, projectSecret *v1.ProjectSecret, opts metav1.CreateOptions) (result *v1.ProjectSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(projectsecretsResource, c.ns, projectSecret), &v1.ProjectSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ProjectSecret), err
}

// Update takes the representation of a projectSecret and updates it. Returns the server's representation of the projectSecret, and an error, if there is any.
func (c *FakeProjectSecrets) Update(ctx context.Context, projectSecret *v1.ProjectSecret, opts metav1.UpdateOptions) (result *v1.ProjectSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(projectsecretsResource, c.ns, projectSecret), &v1.ProjectSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ProjectSecret), err
}

// Delete takes name of the projectSecret and deletes it. Returns an error if one occurs.
func (c *FakeProjectSecrets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(projectsecretsResource, c.ns, name, opts), &v1.ProjectSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjectSecrets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(projectsecretsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ProjectSecretList{})
	return err
}

// Patch applies the patch and returns the patched projectSecret.
func (c *FakeProjectSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ProjectSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(projectsecretsResource, c.ns, name, pt, data, subresources...), &v1.ProjectSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ProjectSecret), err
}