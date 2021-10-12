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

// FakeUsers implements UserInterface
type FakeUsers struct {
	Fake *FakeManagementV1
}

var usersResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "users"}

var usersKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "User"}

// Get takes name of the user, and returns the corresponding user object, and an error if there is any.
func (c *FakeUsers) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(usersResource, name), &managementv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.User), err
}

// List takes label and field selectors, and returns the list of Users that match those selectors.
func (c *FakeUsers) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.UserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(usersResource, usersKind, opts), &managementv1.UserList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.UserList{ListMeta: obj.(*managementv1.UserList).ListMeta}
	for _, item := range obj.(*managementv1.UserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested users.
func (c *FakeUsers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(usersResource, opts))
}

// Create takes the representation of a user and creates it.  Returns the server's representation of the user, and an error, if there is any.
func (c *FakeUsers) Create(ctx context.Context, user *managementv1.User, opts v1.CreateOptions) (result *managementv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(usersResource, user), &managementv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.User), err
}

// Update takes the representation of a user and updates it. Returns the server's representation of the user, and an error, if there is any.
func (c *FakeUsers) Update(ctx context.Context, user *managementv1.User, opts v1.UpdateOptions) (result *managementv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(usersResource, user), &managementv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.User), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUsers) UpdateStatus(ctx context.Context, user *managementv1.User, opts v1.UpdateOptions) (*managementv1.User, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(usersResource, "status", user), &managementv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.User), err
}

// Delete takes name of the user and deletes it. Returns an error if one occurs.
func (c *FakeUsers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(usersResource, name), &managementv1.User{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUsers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(usersResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.UserList{})
	return err
}

// Patch applies the patch and returns the patched user.
func (c *FakeUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(usersResource, name, pt, data, subresources...), &managementv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.User), err
}

// GetProfile takes name of the user, and returns the corresponding userProfile object, and an error if there is any.
func (c *FakeUsers) GetProfile(ctx context.Context, userName string, options v1.GetOptions) (result *managementv1.UserProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(usersResource, "profile", userName), &managementv1.UserProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.UserProfile), err
}

// UpdateProfile takes the representation of a userProfile and creates it.  Returns the server's representation of the userProfile, and an error, if there is any.
func (c *FakeUsers) UpdateProfile(ctx context.Context, userName string, userProfile *managementv1.UserProfile, opts v1.CreateOptions) (result *managementv1.UserProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateSubresourceAction(usersResource, userName, "profile", userProfile), &managementv1.UserProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.UserProfile), err
}

// ListSpaces takes name of the user, and returns the corresponding userSpaces object, and an error if there is any.
func (c *FakeUsers) ListSpaces(ctx context.Context, userName string, options v1.GetOptions) (result *managementv1.UserSpaces, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(usersResource, "spaces", userName), &managementv1.UserSpaces{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.UserSpaces), err
}

// ListQuotas takes name of the user, and returns the corresponding userQuotas object, and an error if there is any.
func (c *FakeUsers) ListQuotas(ctx context.Context, userName string, options v1.GetOptions) (result *managementv1.UserQuotas, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(usersResource, "quotas", userName), &managementv1.UserQuotas{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.UserQuotas), err
}

// ListTeams takes name of the user, and returns the corresponding userTeams object, and an error if there is any.
func (c *FakeUsers) ListTeams(ctx context.Context, userName string, options v1.GetOptions) (result *managementv1.UserTeams, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(usersResource, "teams", userName), &managementv1.UserTeams{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.UserTeams), err
}

// ListClusters takes name of the user, and returns the corresponding userClusters object, and an error if there is any.
func (c *FakeUsers) ListClusters(ctx context.Context, userName string, options v1.GetOptions) (result *managementv1.UserClusters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(usersResource, "clusters", userName), &managementv1.UserClusters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.UserClusters), err
}

// ListVirtualClusters takes name of the user, and returns the corresponding userVirtualClusters object, and an error if there is any.
func (c *FakeUsers) ListVirtualClusters(ctx context.Context, userName string, options v1.GetOptions) (result *managementv1.UserVirtualClusters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(usersResource, "virtualclusters", userName), &managementv1.UserVirtualClusters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.UserVirtualClusters), err
}

// ListAccessKeys takes name of the user, and returns the corresponding userAccessKeys object, and an error if there is any.
func (c *FakeUsers) ListAccessKeys(ctx context.Context, userName string, options v1.GetOptions) (result *managementv1.UserAccessKeys, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(usersResource, "accesskeys", userName), &managementv1.UserAccessKeys{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.UserAccessKeys), err
}
