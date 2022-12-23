// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	storagev1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProjects implements ProjectInterface
type FakeProjects struct {
	Fake *FakeStorageV1
}

var projectsResource = schema.GroupVersionResource{Group: "storage.loft.sh", Version: "v1", Resource: "projects"}

var projectsKind = schema.GroupVersionKind{Group: "storage.loft.sh", Version: "v1", Kind: "Project"}

// Get takes name of the project, and returns the corresponding project object, and an error if there is any.
func (c *FakeProjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *storagev1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(projectsResource, name), &storagev1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.Project), err
}

// List takes label and field selectors, and returns the list of Projects that match those selectors.
func (c *FakeProjects) List(ctx context.Context, opts v1.ListOptions) (result *storagev1.ProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(projectsResource, projectsKind, opts), &storagev1.ProjectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1.ProjectList{ListMeta: obj.(*storagev1.ProjectList).ListMeta}
	for _, item := range obj.(*storagev1.ProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projects.
func (c *FakeProjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(projectsResource, opts))
}

// Create takes the representation of a project and creates it.  Returns the server's representation of the project, and an error, if there is any.
func (c *FakeProjects) Create(ctx context.Context, project *storagev1.Project, opts v1.CreateOptions) (result *storagev1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(projectsResource, project), &storagev1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.Project), err
}

// Update takes the representation of a project and updates it. Returns the server's representation of the project, and an error, if there is any.
func (c *FakeProjects) Update(ctx context.Context, project *storagev1.Project, opts v1.UpdateOptions) (result *storagev1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(projectsResource, project), &storagev1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.Project), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjects) UpdateStatus(ctx context.Context, project *storagev1.Project, opts v1.UpdateOptions) (*storagev1.Project, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(projectsResource, "status", project), &storagev1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.Project), err
}

// Delete takes name of the project and deletes it. Returns an error if one occurs.
func (c *FakeProjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(projectsResource, name, opts), &storagev1.Project{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(projectsResource, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1.ProjectList{})
	return err
}

// Patch applies the patch and returns the patched project.
func (c *FakeProjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *storagev1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(projectsResource, name, pt, data, subresources...), &storagev1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.Project), err
}
