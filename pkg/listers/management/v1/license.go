// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LicenseLister helps list Licenses.
// All objects returned here must be treated as read-only.
type LicenseLister interface {
	// List lists all Licenses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.License, err error)
	// Get retrieves the License from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.License, error)
	LicenseListerExpansion
}

// licenseLister implements the LicenseLister interface.
type licenseLister struct {
	indexer cache.Indexer
}

// NewLicenseLister returns a new LicenseLister.
func NewLicenseLister(indexer cache.Indexer) LicenseLister {
	return &licenseLister{indexer: indexer}
}

// List lists all Licenses in the indexer.
func (s *licenseLister) List(selector labels.Selector) (ret []*v1.License, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.License))
	})
	return ret, err
}

// Get retrieves the License from the index for a given name.
func (s *licenseLister) Get(name string) (*v1.License, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("license"), name)
	}
	return obj.(*v1.License), nil
}
