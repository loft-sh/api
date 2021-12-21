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

// FakeEvents implements EventInterface
type FakeEvents struct {
	Fake *FakeManagementV1
}

var eventsResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "events"}

var eventsKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "Event"}

// Get takes name of the event, and returns the corresponding event object, and an error if there is any.
func (c *FakeEvents) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(eventsResource, name), &managementv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Event), err
}

// List takes label and field selectors, and returns the list of Events that match those selectors.
func (c *FakeEvents) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.EventList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(eventsResource, eventsKind, opts), &managementv1.EventList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.EventList{ListMeta: obj.(*managementv1.EventList).ListMeta}
	for _, item := range obj.(*managementv1.EventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested events.
func (c *FakeEvents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(eventsResource, opts))
}

// Create takes the representation of a event and creates it.  Returns the server's representation of the event, and an error, if there is any.
func (c *FakeEvents) Create(ctx context.Context, event *managementv1.Event, opts v1.CreateOptions) (result *managementv1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(eventsResource, event), &managementv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Event), err
}

// Update takes the representation of a event and updates it. Returns the server's representation of the event, and an error, if there is any.
func (c *FakeEvents) Update(ctx context.Context, event *managementv1.Event, opts v1.UpdateOptions) (result *managementv1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(eventsResource, event), &managementv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Event), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEvents) UpdateStatus(ctx context.Context, event *managementv1.Event, opts v1.UpdateOptions) (*managementv1.Event, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(eventsResource, "status", event), &managementv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Event), err
}

// Delete takes name of the event and deletes it. Returns an error if one occurs.
func (c *FakeEvents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(eventsResource, name, opts), &managementv1.Event{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEvents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(eventsResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.EventList{})
	return err
}

// Patch applies the patch and returns the patched event.
func (c *FakeEvents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(eventsResource, name, pt, data, subresources...), &managementv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Event), err
}
