// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ClusterLister helps list Clusters.
// All objects returned here must be treated as read-only.
type ClusterLister interface {
	// List lists all Clusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Cluster, err error)
	// Get retrieves the Cluster from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Cluster, error)
	ClusterListerExpansion
}

// clusterLister implements the ClusterLister interface.
type clusterLister struct {
	listers.ResourceIndexer[*v1.Cluster]
}

// NewClusterLister returns a new ClusterLister.
func NewClusterLister(indexer cache.Indexer) ClusterLister {
	return &clusterLister{listers.New[*v1.Cluster](indexer, v1.Resource("cluster"))}
}
