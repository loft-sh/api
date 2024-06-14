// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApps implements AppInterface
type FakeApps struct {
	Fake *FakeManagementV1
}

var appsResource = v1.SchemeGroupVersion.WithResource("apps")

var appsKind = v1.SchemeGroupVersion.WithKind("App")

// Get takes name of the app, and returns the corresponding app object, and an error if there is any.
func (c *FakeApps) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(appsResource, name), &v1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.App), err
}

// List takes label and field selectors, and returns the list of Apps that match those selectors.
func (c *FakeApps) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(appsResource, appsKind, opts), &v1.AppList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.AppList{ListMeta: obj.(*v1.AppList).ListMeta}
	for _, item := range obj.(*v1.AppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apps.
func (c *FakeApps) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(appsResource, opts))
}

// Create takes the representation of a app and creates it.  Returns the server's representation of the app, and an error, if there is any.
func (c *FakeApps) Create(ctx context.Context, app *v1.App, opts metav1.CreateOptions) (result *v1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(appsResource, app), &v1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.App), err
}

// Update takes the representation of a app and updates it. Returns the server's representation of the app, and an error, if there is any.
func (c *FakeApps) Update(ctx context.Context, app *v1.App, opts metav1.UpdateOptions) (result *v1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(appsResource, app), &v1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.App), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApps) UpdateStatus(ctx context.Context, app *v1.App, opts metav1.UpdateOptions) (*v1.App, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(appsResource, "status", app), &v1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.App), err
}

// Delete takes name of the app and deletes it. Returns an error if one occurs.
func (c *FakeApps) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(appsResource, name, opts), &v1.App{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApps) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(appsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.AppList{})
	return err
}

// Patch applies the patch and returns the patched app.
func (c *FakeApps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(appsResource, name, pt, data, subresources...), &v1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.App), err
}

// GetCredentials takes name of the app, and returns the corresponding appCredentials object, and an error if there is any.
func (c *FakeApps) GetCredentials(ctx context.Context, appName string, options metav1.GetOptions) (result *v1.AppCredentials, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(appsResource, "credentials", appName), &v1.AppCredentials{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.AppCredentials), err
}