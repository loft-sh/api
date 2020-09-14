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

// FakeClusterAccountTemplates implements ClusterAccountTemplateInterface
type FakeClusterAccountTemplates struct {
	Fake *FakeManagementV1
}

var clusteraccounttemplatesResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "clusteraccounttemplates"}

var clusteraccounttemplatesKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "ClusterAccountTemplate"}

// Get takes name of the clusterAccountTemplate, and returns the corresponding clusterAccountTemplate object, and an error if there is any.
func (c *FakeClusterAccountTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.ClusterAccountTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusteraccounttemplatesResource, name), &managementv1.ClusterAccountTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterAccountTemplate), err
}

// List takes label and field selectors, and returns the list of ClusterAccountTemplates that match those selectors.
func (c *FakeClusterAccountTemplates) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.ClusterAccountTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusteraccounttemplatesResource, clusteraccounttemplatesKind, opts), &managementv1.ClusterAccountTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.ClusterAccountTemplateList{ListMeta: obj.(*managementv1.ClusterAccountTemplateList).ListMeta}
	for _, item := range obj.(*managementv1.ClusterAccountTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterAccountTemplates.
func (c *FakeClusterAccountTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusteraccounttemplatesResource, opts))
}

// Create takes the representation of a clusterAccountTemplate and creates it.  Returns the server's representation of the clusterAccountTemplate, and an error, if there is any.
func (c *FakeClusterAccountTemplates) Create(ctx context.Context, clusterAccountTemplate *managementv1.ClusterAccountTemplate, opts v1.CreateOptions) (result *managementv1.ClusterAccountTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusteraccounttemplatesResource, clusterAccountTemplate), &managementv1.ClusterAccountTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterAccountTemplate), err
}

// Update takes the representation of a clusterAccountTemplate and updates it. Returns the server's representation of the clusterAccountTemplate, and an error, if there is any.
func (c *FakeClusterAccountTemplates) Update(ctx context.Context, clusterAccountTemplate *managementv1.ClusterAccountTemplate, opts v1.UpdateOptions) (result *managementv1.ClusterAccountTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusteraccounttemplatesResource, clusterAccountTemplate), &managementv1.ClusterAccountTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterAccountTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterAccountTemplates) UpdateStatus(ctx context.Context, clusterAccountTemplate *managementv1.ClusterAccountTemplate, opts v1.UpdateOptions) (*managementv1.ClusterAccountTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusteraccounttemplatesResource, "status", clusterAccountTemplate), &managementv1.ClusterAccountTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterAccountTemplate), err
}

// Delete takes name of the clusterAccountTemplate and deletes it. Returns an error if one occurs.
func (c *FakeClusterAccountTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusteraccounttemplatesResource, name), &managementv1.ClusterAccountTemplate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterAccountTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusteraccounttemplatesResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.ClusterAccountTemplateList{})
	return err
}

// Patch applies the patch and returns the patched clusterAccountTemplate.
func (c *FakeClusterAccountTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.ClusterAccountTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusteraccounttemplatesResource, name, pt, data, subresources...), &managementv1.ClusterAccountTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.ClusterAccountTemplate), err
}
