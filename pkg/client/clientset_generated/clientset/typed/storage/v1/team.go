// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	scheme "github.com/loft-sh/api/v4/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TeamsGetter has a method to return a TeamInterface.
// A group's client should implement this interface.
type TeamsGetter interface {
	Teams() TeamInterface
}

// TeamInterface has methods to work with Team resources.
type TeamInterface interface {
	Create(ctx context.Context, team *v1.Team, opts metav1.CreateOptions) (*v1.Team, error)
	Update(ctx context.Context, team *v1.Team, opts metav1.UpdateOptions) (*v1.Team, error)
	UpdateStatus(ctx context.Context, team *v1.Team, opts metav1.UpdateOptions) (*v1.Team, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Team, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TeamList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Team, err error)
	TeamExpansion
}

// teams implements TeamInterface
type teams struct {
	client rest.Interface
}

// newTeams returns a Teams
func newTeams(c *StorageV1Client) *teams {
	return &teams{
		client: c.RESTClient(),
	}
}

// Get takes name of the team, and returns the corresponding team object, and an error if there is any.
func (c *teams) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Team, err error) {
	result = &v1.Team{}
	err = c.client.Get().
		Resource("teams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Teams that match those selectors.
func (c *teams) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TeamList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TeamList{}
	err = c.client.Get().
		Resource("teams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested teams.
func (c *teams) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("teams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a team and creates it.  Returns the server's representation of the team, and an error, if there is any.
func (c *teams) Create(ctx context.Context, team *v1.Team, opts metav1.CreateOptions) (result *v1.Team, err error) {
	result = &v1.Team{}
	err = c.client.Post().
		Resource("teams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(team).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a team and updates it. Returns the server's representation of the team, and an error, if there is any.
func (c *teams) Update(ctx context.Context, team *v1.Team, opts metav1.UpdateOptions) (result *v1.Team, err error) {
	result = &v1.Team{}
	err = c.client.Put().
		Resource("teams").
		Name(team.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(team).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *teams) UpdateStatus(ctx context.Context, team *v1.Team, opts metav1.UpdateOptions) (result *v1.Team, err error) {
	result = &v1.Team{}
	err = c.client.Put().
		Resource("teams").
		Name(team.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(team).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the team and deletes it. Returns an error if one occurs.
func (c *teams) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("teams").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *teams) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("teams").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched team.
func (c *teams) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Team, err error) {
	result = &v1.Team{}
	err = c.client.Patch(pt).
		Resource("teams").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
