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

// FakeDevPodEnvironmentTemplates implements DevPodEnvironmentTemplateInterface
type FakeDevPodEnvironmentTemplates struct {
	Fake *FakeStorageV1
}

var devpodenvironmenttemplatesResource = v1.SchemeGroupVersion.WithResource("devpodenvironmenttemplates")

var devpodenvironmenttemplatesKind = v1.SchemeGroupVersion.WithKind("DevPodEnvironmentTemplate")

// Get takes name of the devPodEnvironmentTemplate, and returns the corresponding devPodEnvironmentTemplate object, and an error if there is any.
func (c *FakeDevPodEnvironmentTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DevPodEnvironmentTemplate, err error) {
	emptyResult := &v1.DevPodEnvironmentTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(devpodenvironmenttemplatesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.DevPodEnvironmentTemplate), err
}

// List takes label and field selectors, and returns the list of DevPodEnvironmentTemplates that match those selectors.
func (c *FakeDevPodEnvironmentTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DevPodEnvironmentTemplateList, err error) {
	emptyResult := &v1.DevPodEnvironmentTemplateList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(devpodenvironmenttemplatesResource, devpodenvironmenttemplatesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.DevPodEnvironmentTemplateList{ListMeta: obj.(*v1.DevPodEnvironmentTemplateList).ListMeta}
	for _, item := range obj.(*v1.DevPodEnvironmentTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devPodEnvironmentTemplates.
func (c *FakeDevPodEnvironmentTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(devpodenvironmenttemplatesResource, opts))
}

// Create takes the representation of a devPodEnvironmentTemplate and creates it.  Returns the server's representation of the devPodEnvironmentTemplate, and an error, if there is any.
func (c *FakeDevPodEnvironmentTemplates) Create(ctx context.Context, devPodEnvironmentTemplate *v1.DevPodEnvironmentTemplate, opts metav1.CreateOptions) (result *v1.DevPodEnvironmentTemplate, err error) {
	emptyResult := &v1.DevPodEnvironmentTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(devpodenvironmenttemplatesResource, devPodEnvironmentTemplate, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.DevPodEnvironmentTemplate), err
}

// Update takes the representation of a devPodEnvironmentTemplate and updates it. Returns the server's representation of the devPodEnvironmentTemplate, and an error, if there is any.
func (c *FakeDevPodEnvironmentTemplates) Update(ctx context.Context, devPodEnvironmentTemplate *v1.DevPodEnvironmentTemplate, opts metav1.UpdateOptions) (result *v1.DevPodEnvironmentTemplate, err error) {
	emptyResult := &v1.DevPodEnvironmentTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(devpodenvironmenttemplatesResource, devPodEnvironmentTemplate, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.DevPodEnvironmentTemplate), err
}

// Delete takes name of the devPodEnvironmentTemplate and deletes it. Returns an error if one occurs.
func (c *FakeDevPodEnvironmentTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(devpodenvironmenttemplatesResource, name, opts), &v1.DevPodEnvironmentTemplate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevPodEnvironmentTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(devpodenvironmenttemplatesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.DevPodEnvironmentTemplateList{})
	return err
}

// Patch applies the patch and returns the patched devPodEnvironmentTemplate.
func (c *FakeDevPodEnvironmentTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DevPodEnvironmentTemplate, err error) {
	emptyResult := &v1.DevPodEnvironmentTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(devpodenvironmenttemplatesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.DevPodEnvironmentTemplate), err
}
