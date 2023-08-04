// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/loft-sh/api/v3/pkg/apis/management/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRunners implements RunnerInterface
type FakeRunners struct {
	Fake *FakeManagementV1
}

var runnersResource = v1.SchemeGroupVersion.WithResource("runners")

var runnersKind = v1.SchemeGroupVersion.WithKind("Runner")

// Get takes name of the runner, and returns the corresponding runner object, and an error if there is any.
func (c *FakeRunners) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Runner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(runnersResource, name), &v1.Runner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Runner), err
}

// List takes label and field selectors, and returns the list of Runners that match those selectors.
func (c *FakeRunners) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RunnerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(runnersResource, runnersKind, opts), &v1.RunnerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RunnerList{ListMeta: obj.(*v1.RunnerList).ListMeta}
	for _, item := range obj.(*v1.RunnerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested runners.
func (c *FakeRunners) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(runnersResource, opts))
}

// Create takes the representation of a runner and creates it.  Returns the server's representation of the runner, and an error, if there is any.
func (c *FakeRunners) Create(ctx context.Context, runner *v1.Runner, opts metav1.CreateOptions) (result *v1.Runner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(runnersResource, runner), &v1.Runner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Runner), err
}

// Update takes the representation of a runner and updates it. Returns the server's representation of the runner, and an error, if there is any.
func (c *FakeRunners) Update(ctx context.Context, runner *v1.Runner, opts metav1.UpdateOptions) (result *v1.Runner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(runnersResource, runner), &v1.Runner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Runner), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRunners) UpdateStatus(ctx context.Context, runner *v1.Runner, opts metav1.UpdateOptions) (*v1.Runner, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(runnersResource, "status", runner), &v1.Runner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Runner), err
}

// Delete takes name of the runner and deletes it. Returns an error if one occurs.
func (c *FakeRunners) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(runnersResource, name, opts), &v1.Runner{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRunners) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(runnersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RunnerList{})
	return err
}

// Patch applies the patch and returns the patched runner.
func (c *FakeRunners) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Runner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(runnersResource, name, pt, data, subresources...), &v1.Runner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Runner), err
}

// GetConfig takes name of the runner, and returns the corresponding runnerConfig object, and an error if there is any.
func (c *FakeRunners) GetConfig(ctx context.Context, runnerName string, options metav1.GetOptions) (result *v1.RunnerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(runnersResource, "config", runnerName), &v1.RunnerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RunnerConfig), err
}
