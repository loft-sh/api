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

// FakeTeams implements TeamInterface
type FakeTeams struct {
	Fake *FakeManagementV1
}

var teamsResource = v1.SchemeGroupVersion.WithResource("teams")

var teamsKind = v1.SchemeGroupVersion.WithKind("Team")

// Get takes name of the team, and returns the corresponding team object, and an error if there is any.
func (c *FakeTeams) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Team, err error) {
	emptyResult := &v1.Team{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(teamsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Team), err
}

// List takes label and field selectors, and returns the list of Teams that match those selectors.
func (c *FakeTeams) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TeamList, err error) {
	emptyResult := &v1.TeamList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(teamsResource, teamsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.TeamList{ListMeta: obj.(*v1.TeamList).ListMeta}
	for _, item := range obj.(*v1.TeamList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested teams.
func (c *FakeTeams) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(teamsResource, opts))
}

// Create takes the representation of a team and creates it.  Returns the server's representation of the team, and an error, if there is any.
func (c *FakeTeams) Create(ctx context.Context, team *v1.Team, opts metav1.CreateOptions) (result *v1.Team, err error) {
	emptyResult := &v1.Team{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(teamsResource, team, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Team), err
}

// Update takes the representation of a team and updates it. Returns the server's representation of the team, and an error, if there is any.
func (c *FakeTeams) Update(ctx context.Context, team *v1.Team, opts metav1.UpdateOptions) (result *v1.Team, err error) {
	emptyResult := &v1.Team{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(teamsResource, team, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Team), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTeams) UpdateStatus(ctx context.Context, team *v1.Team, opts metav1.UpdateOptions) (result *v1.Team, err error) {
	emptyResult := &v1.Team{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(teamsResource, "status", team, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Team), err
}

// Delete takes name of the team and deletes it. Returns an error if one occurs.
func (c *FakeTeams) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(teamsResource, name, opts), &v1.Team{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTeams) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(teamsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.TeamList{})
	return err
}

// Patch applies the patch and returns the patched team.
func (c *FakeTeams) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Team, err error) {
	emptyResult := &v1.Team{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(teamsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Team), err
}

// ListClusters takes name of the team, and returns the corresponding teamClusters object, and an error if there is any.
func (c *FakeTeams) ListClusters(ctx context.Context, teamName string, options metav1.GetOptions) (result *v1.TeamClusters, err error) {
	emptyResult := &v1.TeamClusters{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceActionWithOptions(teamsResource, "clusters", teamName, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.TeamClusters), err
}

// ListAccessKeys takes name of the team, and returns the corresponding teamAccessKeys object, and an error if there is any.
func (c *FakeTeams) ListAccessKeys(ctx context.Context, teamName string, options metav1.GetOptions) (result *v1.TeamAccessKeys, err error) {
	emptyResult := &v1.TeamAccessKeys{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceActionWithOptions(teamsResource, "accesskeys", teamName, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.TeamAccessKeys), err
}
