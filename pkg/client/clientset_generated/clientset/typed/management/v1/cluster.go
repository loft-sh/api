// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v3/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v3/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClustersGetter has a method to return a ClusterInterface.
// A group's client should implement this interface.
type ClustersGetter interface {
	Clusters() ClusterInterface
}

// ClusterInterface has methods to work with Cluster resources.
type ClusterInterface interface {
	Create(ctx context.Context, cluster *v1.Cluster, opts metav1.CreateOptions) (*v1.Cluster, error)
	Update(ctx context.Context, cluster *v1.Cluster, opts metav1.UpdateOptions) (*v1.Cluster, error)
	UpdateStatus(ctx context.Context, cluster *v1.Cluster, opts metav1.UpdateOptions) (*v1.Cluster, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Cluster, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Cluster, err error)
	ListAccess(ctx context.Context, clusterName string, options metav1.GetOptions) (*v1.ClusterMemberAccess, error)
	ListMembers(ctx context.Context, clusterName string, options metav1.GetOptions) (*v1.ClusterMembers, error)
	ListVirtualClusterDefaults(ctx context.Context, clusterName string, options metav1.GetOptions) (*v1.ClusterVirtualClusterDefaults, error)
	GetAgentConfig(ctx context.Context, clusterName string, options metav1.GetOptions) (*v1.ClusterAgentConfig, error)
	GetAccessKey(ctx context.Context, clusterName string, options metav1.GetOptions) (*v1.ClusterAccessKey, error)

	ClusterExpansion
}

// clusters implements ClusterInterface
type clusters struct {
	client rest.Interface
}

// newClusters returns a Clusters
func newClusters(c *ManagementV1Client) *clusters {
	return &clusters{
		client: c.RESTClient(),
	}
}

// Get takes name of the cluster, and returns the corresponding cluster object, and an error if there is any.
func (c *clusters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Cluster, err error) {
	result = &v1.Cluster{}
	err = c.client.Get().
		Resource("clusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Clusters that match those selectors.
func (c *clusters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterList{}
	err = c.client.Get().
		Resource("clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusters.
func (c *clusters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cluster and creates it.  Returns the server's representation of the cluster, and an error, if there is any.
func (c *clusters) Create(ctx context.Context, cluster *v1.Cluster, opts metav1.CreateOptions) (result *v1.Cluster, err error) {
	result = &v1.Cluster{}
	err = c.client.Post().
		Resource("clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cluster and updates it. Returns the server's representation of the cluster, and an error, if there is any.
func (c *clusters) Update(ctx context.Context, cluster *v1.Cluster, opts metav1.UpdateOptions) (result *v1.Cluster, err error) {
	result = &v1.Cluster{}
	err = c.client.Put().
		Resource("clusters").
		Name(cluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusters) UpdateStatus(ctx context.Context, cluster *v1.Cluster, opts metav1.UpdateOptions) (result *v1.Cluster, err error) {
	result = &v1.Cluster{}
	err = c.client.Put().
		Resource("clusters").
		Name(cluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cluster and deletes it. Returns an error if one occurs.
func (c *clusters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cluster.
func (c *clusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Cluster, err error) {
	result = &v1.Cluster{}
	err = c.client.Patch(pt).
		Resource("clusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ListAccess takes name of the cluster, and returns the corresponding v1.ClusterMemberAccess object, and an error if there is any.
func (c *clusters) ListAccess(ctx context.Context, clusterName string, options metav1.GetOptions) (result *v1.ClusterMemberAccess, err error) {
	result = &v1.ClusterMemberAccess{}
	err = c.client.Get().
		Resource("clusters").
		Name(clusterName).
		SubResource("memberaccess").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// ListMembers takes name of the cluster, and returns the corresponding v1.ClusterMembers object, and an error if there is any.
func (c *clusters) ListMembers(ctx context.Context, clusterName string, options metav1.GetOptions) (result *v1.ClusterMembers, err error) {
	result = &v1.ClusterMembers{}
	err = c.client.Get().
		Resource("clusters").
		Name(clusterName).
		SubResource("members").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// ListVirtualClusterDefaults takes name of the cluster, and returns the corresponding v1.ClusterVirtualClusterDefaults object, and an error if there is any.
func (c *clusters) ListVirtualClusterDefaults(ctx context.Context, clusterName string, options metav1.GetOptions) (result *v1.ClusterVirtualClusterDefaults, err error) {
	result = &v1.ClusterVirtualClusterDefaults{}
	err = c.client.Get().
		Resource("clusters").
		Name(clusterName).
		SubResource("virtualclusterdefaults").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// GetAgentConfig takes name of the cluster, and returns the corresponding v1.ClusterAgentConfig object, and an error if there is any.
func (c *clusters) GetAgentConfig(ctx context.Context, clusterName string, options metav1.GetOptions) (result *v1.ClusterAgentConfig, err error) {
	result = &v1.ClusterAgentConfig{}
	err = c.client.Get().
		Resource("clusters").
		Name(clusterName).
		SubResource("agentconfig").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// GetAccessKey takes name of the cluster, and returns the corresponding v1.ClusterAccessKey object, and an error if there is any.
func (c *clusters) GetAccessKey(ctx context.Context, clusterName string, options metav1.GetOptions) (result *v1.ClusterAccessKey, err error) {
	result = &v1.ClusterAccessKey{}
	err = c.client.Get().
		Resource("clusters").
		Name(clusterName).
		SubResource("accesskey").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}
