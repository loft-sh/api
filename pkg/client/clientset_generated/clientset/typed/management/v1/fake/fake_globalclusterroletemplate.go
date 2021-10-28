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

// FakeGlobalClusterRoleTemplates implements GlobalClusterRoleTemplateInterface
type FakeGlobalClusterRoleTemplates struct {
	Fake *FakeManagementV1
}

var globalclusterroletemplatesResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "globalclusterroletemplates"}

var globalclusterroletemplatesKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "GlobalClusterRoleTemplate"}

// Get takes name of the globalClusterRoleTemplate, and returns the corresponding globalClusterRoleTemplate object, and an error if there is any.
func (c *FakeGlobalClusterRoleTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.GlobalClusterRoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(globalclusterroletemplatesResource, name), &managementv1.GlobalClusterRoleTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.GlobalClusterRoleTemplate), err
}

// List takes label and field selectors, and returns the list of GlobalClusterRoleTemplates that match those selectors.
func (c *FakeGlobalClusterRoleTemplates) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.GlobalClusterRoleTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(globalclusterroletemplatesResource, globalclusterroletemplatesKind, opts), &managementv1.GlobalClusterRoleTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.GlobalClusterRoleTemplateList{ListMeta: obj.(*managementv1.GlobalClusterRoleTemplateList).ListMeta}
	for _, item := range obj.(*managementv1.GlobalClusterRoleTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalClusterRoleTemplates.
func (c *FakeGlobalClusterRoleTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(globalclusterroletemplatesResource, opts))
}

// Create takes the representation of a globalClusterRoleTemplate and creates it.  Returns the server's representation of the globalClusterRoleTemplate, and an error, if there is any.
func (c *FakeGlobalClusterRoleTemplates) Create(ctx context.Context, globalClusterRoleTemplate *managementv1.GlobalClusterRoleTemplate, opts v1.CreateOptions) (result *managementv1.GlobalClusterRoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(globalclusterroletemplatesResource, globalClusterRoleTemplate), &managementv1.GlobalClusterRoleTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.GlobalClusterRoleTemplate), err
}

// Update takes the representation of a globalClusterRoleTemplate and updates it. Returns the server's representation of the globalClusterRoleTemplate, and an error, if there is any.
func (c *FakeGlobalClusterRoleTemplates) Update(ctx context.Context, globalClusterRoleTemplate *managementv1.GlobalClusterRoleTemplate, opts v1.UpdateOptions) (result *managementv1.GlobalClusterRoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(globalclusterroletemplatesResource, globalClusterRoleTemplate), &managementv1.GlobalClusterRoleTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.GlobalClusterRoleTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlobalClusterRoleTemplates) UpdateStatus(ctx context.Context, globalClusterRoleTemplate *managementv1.GlobalClusterRoleTemplate, opts v1.UpdateOptions) (*managementv1.GlobalClusterRoleTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(globalclusterroletemplatesResource, "status", globalClusterRoleTemplate), &managementv1.GlobalClusterRoleTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.GlobalClusterRoleTemplate), err
}

// Delete takes name of the globalClusterRoleTemplate and deletes it. Returns an error if one occurs.
func (c *FakeGlobalClusterRoleTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(globalclusterroletemplatesResource, name), &managementv1.GlobalClusterRoleTemplate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalClusterRoleTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(globalclusterroletemplatesResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.GlobalClusterRoleTemplateList{})
	return err
}

// Patch applies the patch and returns the patched globalClusterRoleTemplate.
func (c *FakeGlobalClusterRoleTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.GlobalClusterRoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(globalclusterroletemplatesResource, name, pt, data, subresources...), &managementv1.GlobalClusterRoleTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.GlobalClusterRoleTemplate), err
}
