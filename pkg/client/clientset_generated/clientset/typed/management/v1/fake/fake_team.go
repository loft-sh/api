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

// FakeTeams implements TeamInterface
type FakeTeams struct {
	Fake *FakeManagementV1
}

var teamsResource = schema.GroupVersionResource{Group: "management.loft.sh", Version: "v1", Resource: "teams"}

var teamsKind = schema.GroupVersionKind{Group: "management.loft.sh", Version: "v1", Kind: "Team"}

// Get takes name of the team, and returns the corresponding team object, and an error if there is any.
func (c *FakeTeams) Get(ctx context.Context, name string, options v1.GetOptions) (result *managementv1.Team, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(teamsResource, name), &managementv1.Team{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Team), err
}

// List takes label and field selectors, and returns the list of Teams that match those selectors.
func (c *FakeTeams) List(ctx context.Context, opts v1.ListOptions) (result *managementv1.TeamList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(teamsResource, teamsKind, opts), &managementv1.TeamList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &managementv1.TeamList{ListMeta: obj.(*managementv1.TeamList).ListMeta}
	for _, item := range obj.(*managementv1.TeamList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested teams.
func (c *FakeTeams) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(teamsResource, opts))
}

// Create takes the representation of a team and creates it.  Returns the server's representation of the team, and an error, if there is any.
func (c *FakeTeams) Create(ctx context.Context, team *managementv1.Team, opts v1.CreateOptions) (result *managementv1.Team, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(teamsResource, team), &managementv1.Team{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Team), err
}

// Update takes the representation of a team and updates it. Returns the server's representation of the team, and an error, if there is any.
func (c *FakeTeams) Update(ctx context.Context, team *managementv1.Team, opts v1.UpdateOptions) (result *managementv1.Team, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(teamsResource, team), &managementv1.Team{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Team), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTeams) UpdateStatus(ctx context.Context, team *managementv1.Team, opts v1.UpdateOptions) (*managementv1.Team, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(teamsResource, "status", team), &managementv1.Team{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Team), err
}

// Delete takes name of the team and deletes it. Returns an error if one occurs.
func (c *FakeTeams) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(teamsResource, name), &managementv1.Team{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTeams) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(teamsResource, listOpts)

	_, err := c.Fake.Invokes(action, &managementv1.TeamList{})
	return err
}

// Patch applies the patch and returns the patched team.
func (c *FakeTeams) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *managementv1.Team, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(teamsResource, name, pt, data, subresources...), &managementv1.Team{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.Team), err
}

// ListSpaces takes name of the team, and returns the corresponding teamSpaces object, and an error if there is any.
func (c *FakeTeams) ListSpaces(ctx context.Context, teamName string, options v1.GetOptions) (result *managementv1.TeamSpaces, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(teamsResource, "spaces", teamName), &managementv1.TeamSpaces{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.TeamSpaces), err
}

// ListClusters takes name of the team, and returns the corresponding teamClusters object, and an error if there is any.
func (c *FakeTeams) ListClusters(ctx context.Context, teamName string, options v1.GetOptions) (result *managementv1.TeamClusters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(teamsResource, "clusters", teamName), &managementv1.TeamClusters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*managementv1.TeamClusters), err
}
