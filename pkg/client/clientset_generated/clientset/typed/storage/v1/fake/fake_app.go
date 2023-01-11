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

// FakeApps implements AppInterface
type FakeApps struct {
	Fake *FakeStorageV1
}

var appsResource = schema.GroupVersionResource{Group: "storage.loft.sh", Version: "v1", Resource: "apps"}

var appsKind = schema.GroupVersionKind{Group: "storage.loft.sh", Version: "v1", Kind: "App"}

// Get takes name of the app, and returns the corresponding app object, and an error if there is any.
func (c *FakeApps) Get(ctx context.Context, name string, options v1.GetOptions) (result *storagev1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(appsResource, name), &storagev1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.App), err
}

// List takes label and field selectors, and returns the list of Apps that match those selectors.
func (c *FakeApps) List(ctx context.Context, opts v1.ListOptions) (result *storagev1.AppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(appsResource, appsKind, opts), &storagev1.AppList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1.AppList{ListMeta: obj.(*storagev1.AppList).ListMeta}
	for _, item := range obj.(*storagev1.AppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apps.
func (c *FakeApps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(appsResource, opts))
}

// Create takes the representation of a app and creates it.  Returns the server's representation of the app, and an error, if there is any.
func (c *FakeApps) Create(ctx context.Context, app *storagev1.App, opts v1.CreateOptions) (result *storagev1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(appsResource, app), &storagev1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.App), err
}

// Update takes the representation of a app and updates it. Returns the server's representation of the app, and an error, if there is any.
func (c *FakeApps) Update(ctx context.Context, app *storagev1.App, opts v1.UpdateOptions) (result *storagev1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(appsResource, app), &storagev1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.App), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApps) UpdateStatus(ctx context.Context, app *storagev1.App, opts v1.UpdateOptions) (*storagev1.App, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(appsResource, "status", app), &storagev1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.App), err
}

// Delete takes name of the app and deletes it. Returns an error if one occurs.
func (c *FakeApps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(appsResource, name, opts), &storagev1.App{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(appsResource, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1.AppList{})
	return err
}

// Patch applies the patch and returns the patched app.
func (c *FakeApps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *storagev1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(appsResource, name, pt, data, subresources...), &storagev1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.App), err
}
