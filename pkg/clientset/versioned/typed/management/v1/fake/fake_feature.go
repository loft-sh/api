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

// FakeFeatures implements FeatureInterface
type FakeFeatures struct {
	Fake *FakeManagementV1
}

var featuresResource = v1.SchemeGroupVersion.WithResource("features")

var featuresKind = v1.SchemeGroupVersion.WithKind("Feature")

// Get takes name of the feature, and returns the corresponding feature object, and an error if there is any.
func (c *FakeFeatures) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Feature, err error) {
	emptyResult := &v1.Feature{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(featuresResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Feature), err
}

// List takes label and field selectors, and returns the list of Features that match those selectors.
func (c *FakeFeatures) List(ctx context.Context, opts metav1.ListOptions) (result *v1.FeatureList, err error) {
	emptyResult := &v1.FeatureList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(featuresResource, featuresKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.FeatureList{ListMeta: obj.(*v1.FeatureList).ListMeta}
	for _, item := range obj.(*v1.FeatureList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested features.
func (c *FakeFeatures) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(featuresResource, opts))
}

// Create takes the representation of a feature and creates it.  Returns the server's representation of the feature, and an error, if there is any.
func (c *FakeFeatures) Create(ctx context.Context, feature *v1.Feature, opts metav1.CreateOptions) (result *v1.Feature, err error) {
	emptyResult := &v1.Feature{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(featuresResource, feature, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Feature), err
}

// Update takes the representation of a feature and updates it. Returns the server's representation of the feature, and an error, if there is any.
func (c *FakeFeatures) Update(ctx context.Context, feature *v1.Feature, opts metav1.UpdateOptions) (result *v1.Feature, err error) {
	emptyResult := &v1.Feature{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(featuresResource, feature, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Feature), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFeatures) UpdateStatus(ctx context.Context, feature *v1.Feature, opts metav1.UpdateOptions) (result *v1.Feature, err error) {
	emptyResult := &v1.Feature{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(featuresResource, "status", feature, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Feature), err
}

// Delete takes name of the feature and deletes it. Returns an error if one occurs.
func (c *FakeFeatures) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(featuresResource, name, opts), &v1.Feature{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFeatures) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(featuresResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.FeatureList{})
	return err
}

// Patch applies the patch and returns the patched feature.
func (c *FakeFeatures) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Feature, err error) {
	emptyResult := &v1.Feature{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(featuresResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Feature), err
}
