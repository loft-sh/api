// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTasks implements TaskInterface
type FakeTasks struct {
	Fake *FakeStorageV1
}

var tasksResource = v1.SchemeGroupVersion.WithResource("tasks")

var tasksKind = v1.SchemeGroupVersion.WithKind("Task")

// Get takes name of the task, and returns the corresponding task object, and an error if there is any.
func (c *FakeTasks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Task, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(tasksResource, name), &v1.Task{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Task), err
}

// List takes label and field selectors, and returns the list of Tasks that match those selectors.
func (c *FakeTasks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TaskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(tasksResource, tasksKind, opts), &v1.TaskList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.TaskList{ListMeta: obj.(*v1.TaskList).ListMeta}
	for _, item := range obj.(*v1.TaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tasks.
func (c *FakeTasks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(tasksResource, opts))
}

// Create takes the representation of a task and creates it.  Returns the server's representation of the task, and an error, if there is any.
func (c *FakeTasks) Create(ctx context.Context, task *v1.Task, opts metav1.CreateOptions) (result *v1.Task, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(tasksResource, task), &v1.Task{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Task), err
}

// Update takes the representation of a task and updates it. Returns the server's representation of the task, and an error, if there is any.
func (c *FakeTasks) Update(ctx context.Context, task *v1.Task, opts metav1.UpdateOptions) (result *v1.Task, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(tasksResource, task), &v1.Task{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Task), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTasks) UpdateStatus(ctx context.Context, task *v1.Task, opts metav1.UpdateOptions) (*v1.Task, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(tasksResource, "status", task), &v1.Task{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Task), err
}

// Delete takes name of the task and deletes it. Returns an error if one occurs.
func (c *FakeTasks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(tasksResource, name, opts), &v1.Task{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTasks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(tasksResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.TaskList{})
	return err
}

// Patch applies the patch and returns the patched task.
func (c *FakeTasks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Task, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(tasksResource, name, pt, data, subresources...), &v1.Task{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Task), err
}
