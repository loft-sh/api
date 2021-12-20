// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	managementv1 "github.com/loft-sh/api/v2/pkg/apis/management/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterRoleTemplates implements ClusterRoleTemplateInterface
type FakeClusterRoleTemplates struct {
	Fake *FakeManagementV1
}

var clusterroletemplatesResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "clusterroletemplates"}

var clusterroletemplatesKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "ClusterRoleTemplate"}

// Get takes name of the clusterRoleTemplate, and returns the corresponding clusterRoleTemplate object, and an error if there is any.
func (c *FakeClusterRoleTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.ClusterRoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterroletemplatesResource, name), &managementv1.ClusterRoleTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterRoleTemplate), err
}

// List takes label and field selectors, and returns the list of ClusterRoleTemplates that match those selectors.
func (c *FakeClusterRoleTemplates) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.ClusterRoleTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterroletemplatesResource, clusterroletemplatesKind, opts), &managementv1.ClusterRoleTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.ClusterRoleTemplateList{ListMeta: obj.(*managementv1.ClusterRoleTemplateList).ListMeta}
	for _, item := range obj.(*managementv1.ClusterRoleTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRoleTemplates.
func (c *FakeClusterRoleTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterroletemplatesResource, opts))
}

// Create takes the representation of a clusterRoleTemplate and creates it.  Returns the server's representation of the clusterRoleTemplate, and an error, if there is any.
func (c *FakeClusterRoleTemplates) Create(ctx context.Context, clusterRoleTemplate *managementv1.ClusterRoleTemplate, opts v1.CreateOptions) (result *managementv1.ClusterRoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterroletemplatesResource, clusterRoleTemplate), &managementv1.ClusterRoleTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterRoleTemplate), err
}

// Update takes the representation of a clusterRoleTemplate and updates it. Returns the server's representation of the clusterRoleTemplate, and an error, if there is any.
func (c *FakeClusterRoleTemplates) Update(ctx context.Context, clusterRoleTemplate *managementv1.ClusterRoleTemplate, opts v1.UpdateOptions) (result *managementv1.ClusterRoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterroletemplatesResource, clusterRoleTemplate), &managementv1.ClusterRoleTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterRoleTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterRoleTemplates) UpdateStatus(ctx context.Context, clusterRoleTemplate *managementv1.ClusterRoleTemplate, opts v1.UpdateOptions) (*managementv1.ClusterRoleTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterroletemplatesResource, "status", clusterRoleTemplate), &managementv1.ClusterRoleTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterRoleTemplate), err
}

// Delete takes name of the clusterRoleTemplate and deletes it. Returns an error if one occurs.
func (c *FakeClusterRoleTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterroletemplatesResource, name, opts), &managementv1.ClusterRoleTemplate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRoleTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterroletemplatesResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.ClusterRoleTemplateList{})
	return err
}

// Patch applies the patch and returns the patched clusterRoleTemplate.
func (c *FakeClusterRoleTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.ClusterRoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterroletemplatesResource, name, pt, data, subresources...), &managementv1.ClusterRoleTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterRoleTemplate), err
}
