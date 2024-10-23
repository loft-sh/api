// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
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
	listers.ResourceIndexer[*v1.ClusterRoleTemplate]
}

// NewClusterRoleTemplateLister returns a new ClusterRoleTemplateLister.
func NewClusterRoleTemplateLister(indexer cache.Indexer) ClusterRoleTemplateLister {
	return &clusterRoleTemplateLister{listers.New[*v1.ClusterRoleTemplate](indexer, v1.Resource("clusterroletemplate"))}
}
