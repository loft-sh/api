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

// FakeAnnouncements implements AnnouncementInterface
type FakeAnnouncements struct {
	Fake *FakeManagementV1
}

var announcementsResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "announcements"}

var announcementsKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "Announcement"}

// Get takes name of the announcement, and returns the corresponding announcement object, and an error if there is any.
func (c *FakeAnnouncements) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.Announcement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(announcementsResource, name), &managementv1.Announcement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Announcement), err
}

// List takes label and field selectors, and returns the list of Announcements that match those selectors.
func (c *FakeAnnouncements) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.AnnouncementList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(announcementsResource, announcementsKind, opts), &managementv1.AnnouncementList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.AnnouncementList{ListMeta: obj.(*managementv1.AnnouncementList).ListMeta}
	for _, item := range obj.(*managementv1.AnnouncementList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested announcements.
func (c *FakeAnnouncements) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(announcementsResource, opts))
}

// Create takes the representation of a announcement and creates it.  Returns the server's representation of the announcement, and an error, if there is any.
func (c *FakeAnnouncements) Create(ctx context.Context, announcement *managementv1.Announcement, opts v1.CreateOptions) (result *managementv1.Announcement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(announcementsResource, announcement), &managementv1.Announcement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Announcement), err
}

// Update takes the representation of a announcement and updates it. Returns the server's representation of the announcement, and an error, if there is any.
func (c *FakeAnnouncements) Update(ctx context.Context, announcement *managementv1.Announcement, opts v1.UpdateOptions) (result *managementv1.Announcement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(announcementsResource, announcement), &managementv1.Announcement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Announcement), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAnnouncements) UpdateStatus(ctx context.Context, announcement *managementv1.Announcement, opts v1.UpdateOptions) (*managementv1.Announcement, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(announcementsResource, "status", announcement), &managementv1.Announcement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Announcement), err
}

// Delete takes name of the announcement and deletes it. Returns an error if one occurs.
func (c *FakeAnnouncements) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(announcementsResource, name), &managementv1.Announcement{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAnnouncements) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(announcementsResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.AnnouncementList{})
	return err
}

// Patch applies the patch and returns the patched announcement.
func (c *FakeAnnouncements) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.Announcement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(announcementsResource, name, pt, data, subresources...), &managementv1.Announcement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Announcement), err
}
