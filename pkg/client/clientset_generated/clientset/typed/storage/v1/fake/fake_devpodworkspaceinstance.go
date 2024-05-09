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

// FakeDevPodWorkspaceInstances implements DevPodWorkspaceInstanceInterface
type FakeDevPodWorkspaceInstances struct {
	Fake *FakeStorageV1
	ns   string
}

var devpodworkspaceinstancesResource = v1.SchemeGroupVersion.WithResource("devpodworkspaceinstances")

var devpodworkspaceinstancesKind = v1.SchemeGroupVersion.WithKind("DevPodWorkspaceInstance")

// Get takes name of the devPodWorkspaceInstance, and returns the corresponding devPodWorkspaceInstance object, and an error if there is any.
func (c *FakeDevPodWorkspaceInstances) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DevPodWorkspaceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devpodworkspaceinstancesResource, c.ns, name), &v1.DevPodWorkspaceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DevPodWorkspaceInstance), err
}

// List takes label and field selectors, and returns the list of DevPodWorkspaceInstances that match those selectors.
func (c *FakeDevPodWorkspaceInstances) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DevPodWorkspaceInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devpodworkspaceinstancesResource, devpodworkspaceinstancesKind, c.ns, opts), &v1.DevPodWorkspaceInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.DevPodWorkspaceInstanceList{ListMeta: obj.(*v1.DevPodWorkspaceInstanceList).ListMeta}
	for _, item := range obj.(*v1.DevPodWorkspaceInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devPodWorkspaceInstances.
func (c *FakeDevPodWorkspaceInstances) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devpodworkspaceinstancesResource, c.ns, opts))

}

// Create takes the representation of a devPodWorkspaceInstance and creates it.  Returns the server's representation of the devPodWorkspaceInstance, and an error, if there is any.
func (c *FakeDevPodWorkspaceInstances) Create(ctx context.Context, devPodWorkspaceInstance *v1.DevPodWorkspaceInstance, opts metav1.CreateOptions) (result *v1.DevPodWorkspaceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devpodworkspaceinstancesResource, c.ns, devPodWorkspaceInstance), &v1.DevPodWorkspaceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DevPodWorkspaceInstance), err
}

// Update takes the representation of a devPodWorkspaceInstance and updates it. Returns the server's representation of the devPodWorkspaceInstance, and an error, if there is any.
func (c *FakeDevPodWorkspaceInstances) Update(ctx context.Context, devPodWorkspaceInstance *v1.DevPodWorkspaceInstance, opts metav1.UpdateOptions) (result *v1.DevPodWorkspaceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devpodworkspaceinstancesResource, c.ns, devPodWorkspaceInstance), &v1.DevPodWorkspaceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DevPodWorkspaceInstance), err
}

// Delete takes name of the devPodWorkspaceInstance and deletes it. Returns an error if one occurs.
func (c *FakeDevPodWorkspaceInstances) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(devpodworkspaceinstancesResource, c.ns, name, opts), &v1.DevPodWorkspaceInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevPodWorkspaceInstances) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devpodworkspaceinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.DevPodWorkspaceInstanceList{})
	return err
}

// Patch applies the patch and returns the patched devPodWorkspaceInstance.
func (c *FakeDevPodWorkspaceInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DevPodWorkspaceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devpodworkspaceinstancesResource, c.ns, name, pt, data, subresources...), &v1.DevPodWorkspaceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DevPodWorkspaceInstance), err
}
