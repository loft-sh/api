// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterAccessLister helps list ClusterAccesses.
// All objects returned here must be treated as read-only.
type ClusterAccessLister interface {
	// List lists all ClusterAccesses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*storagev1.ClusterAccess, err error)
	// Get retrieves the ClusterAccess from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*storagev1.ClusterAccess, error)
	ClusterAccessListerExpansion
}

// clusterAccessLister implements the ClusterAccessLister interface.
type clusterAccessLister struct {
	listers.ResourceIndexer[*storagev1.ClusterAccess]
}

// NewClusterAccessLister returns a new ClusterAccessLister.
func NewClusterAccessLister(indexer cache.Indexer) ClusterAccessLister {
	return &clusterAccessLister{listers.New[*storagev1.ClusterAccess](indexer, storagev1.Resource("clusteraccess"))}
}
