// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RegisterVirtualClusterLister helps list RegisterVirtualClusters.
// All objects returned here must be treated as read-only.
type RegisterVirtualClusterLister interface {
	// List lists all RegisterVirtualClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.RegisterVirtualCluster, err error)
	// Get retrieves the RegisterVirtualCluster from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.RegisterVirtualCluster, error)
	RegisterVirtualClusterListerExpansion
}

// registerVirtualClusterLister implements the RegisterVirtualClusterLister interface.
type registerVirtualClusterLister struct {
	indexer cache.Indexer
}

// NewRegisterVirtualClusterLister returns a new RegisterVirtualClusterLister.
func NewRegisterVirtualClusterLister(indexer cache.Indexer) RegisterVirtualClusterLister {
	return &registerVirtualClusterLister{indexer: indexer}
}

// List lists all RegisterVirtualClusters in the indexer.
func (s *registerVirtualClusterLister) List(selector labels.Selector) (ret []*v1.RegisterVirtualCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RegisterVirtualCluster))
	})
	return ret, err
}

// Get retrieves the RegisterVirtualCluster from the index for a given name.
func (s *registerVirtualClusterLister) Get(name string) (*v1.RegisterVirtualCluster, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("registervirtualcluster"), name)
	}
	return obj.(*v1.RegisterVirtualCluster), nil
}
