// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterRoleTemplateLister helps list ClusterRoleTemplates.
// All objects returned here must be treated as read-only.
type ClusterRoleTemplateLister interface {
	// List lists all ClusterRoleTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterRoleTemplate, err error)
	// Get retrieves the ClusterRoleTemplate from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterRoleTemplate, error)
	ClusterRoleTemplateListerExpansion
}

// clusterRoleTemplateLister implements the ClusterRoleTemplateLister interface.
type clusterRoleTemplateLister struct {
	indexer cache.Indexer
}

// NewClusterRoleTemplateLister returns a new ClusterRoleTemplateLister.
func NewClusterRoleTemplateLister(indexer cache.Indexer) ClusterRoleTemplateLister {
	return &clusterRoleTemplateLister{indexer: indexer}
}

// List lists all ClusterRoleTemplates in the indexer.
func (s *clusterRoleTemplateLister) List(selector labels.Selector) (ret []*v1.ClusterRoleTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterRoleTemplate))
	})
	return ret, err
}

// Get retrieves the ClusterRoleTemplate from the index for a given name.
func (s *clusterRoleTemplateLister) Get(name string) (*v1.ClusterRoleTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterroletemplate"), name)
	}
	return obj.(*v1.ClusterRoleTemplate), nil
}
