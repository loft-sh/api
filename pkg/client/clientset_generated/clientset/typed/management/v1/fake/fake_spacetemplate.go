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

// FakeSpaceTemplates implements SpaceTemplateInterface
type FakeSpaceTemplates struct {
	Fake *FakeManagementV1
}

var spacetemplatesResource = v1.SchemeGroupVersion.WithResource("spacetemplates")

var spacetemplatesKind = v1.SchemeGroupVersion.WithKind("SpaceTemplate")

// Get takes name of the spaceTemplate, and returns the corresponding spaceTemplate object, and an error if there is any.
func (c *FakeSpaceTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SpaceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(spacetemplatesResource, name), &v1.SpaceTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SpaceTemplate), err
}

// List takes label and field selectors, and returns the list of SpaceTemplates that match those selectors.
func (c *FakeSpaceTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SpaceTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(spacetemplatesResource, spacetemplatesKind, opts), &v1.SpaceTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.SpaceTemplateList{ListMeta: obj.(*v1.SpaceTemplateList).ListMeta}
	for _, item := range obj.(*v1.SpaceTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spaceTemplates.
func (c *FakeSpaceTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(spacetemplatesResource, opts))
}

// Create takes the representation of a spaceTemplate and creates it.  Returns the server's representation of the spaceTemplate, and an error, if there is any.
func (c *FakeSpaceTemplates) Create(ctx context.Context, spaceTemplate *v1.SpaceTemplate, opts metav1.CreateOptions) (result *v1.SpaceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(spacetemplatesResource, spaceTemplate), &v1.SpaceTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SpaceTemplate), err
}

// Update takes the representation of a spaceTemplate and updates it. Returns the server's representation of the spaceTemplate, and an error, if there is any.
func (c *FakeSpaceTemplates) Update(ctx context.Context, spaceTemplate *v1.SpaceTemplate, opts metav1.UpdateOptions) (result *v1.SpaceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(spacetemplatesResource, spaceTemplate), &v1.SpaceTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SpaceTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpaceTemplates) UpdateStatus(ctx context.Context, spaceTemplate *v1.SpaceTemplate, opts metav1.UpdateOptions) (*v1.SpaceTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(spacetemplatesResource, "status", spaceTemplate), &v1.SpaceTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SpaceTemplate), err
}

// Delete takes name of the spaceTemplate and deletes it. Returns an error if one occurs.
func (c *FakeSpaceTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(spacetemplatesResource, name, opts), &v1.SpaceTemplate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpaceTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(spacetemplatesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.SpaceTemplateList{})
	return err
}

// Patch applies the patch and returns the patched spaceTemplate.
func (c *FakeSpaceTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SpaceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(spacetemplatesResource, name, pt, data, subresources...), &v1.SpaceTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SpaceTemplate), err
}
