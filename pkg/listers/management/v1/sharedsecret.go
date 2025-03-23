// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// SharedSecretLister helps list SharedSecrets.
// All objects returned here must be treated as read-only.
type SharedSecretLister interface {
	// List lists all SharedSecrets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*managementv1.SharedSecret, err error)
	// SharedSecrets returns an object that can list and get SharedSecrets.
	SharedSecrets(namespace string) SharedSecretNamespaceLister
	SharedSecretListerExpansion
}

// sharedSecretLister implements the SharedSecretLister interface.
type sharedSecretLister struct {
	listers.ResourceIndexer[*managementv1.SharedSecret]
}

// NewSharedSecretLister returns a new SharedSecretLister.
func NewSharedSecretLister(indexer cache.Indexer) SharedSecretLister {
	return &sharedSecretLister{listers.New[*managementv1.SharedSecret](indexer, managementv1.Resource("sharedsecret"))}
}

// SharedSecrets returns an object that can list and get SharedSecrets.
func (s *sharedSecretLister) SharedSecrets(namespace string) SharedSecretNamespaceLister {
	return sharedSecretNamespaceLister{listers.NewNamespaced[*managementv1.SharedSecret](s.ResourceIndexer, namespace)}
}

// SharedSecretNamespaceLister helps list and get SharedSecrets.
// All objects returned here must be treated as read-only.
type SharedSecretNamespaceLister interface {
	// List lists all SharedSecrets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*managementv1.SharedSecret, err error)
	// Get retrieves the SharedSecret from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*managementv1.SharedSecret, error)
	SharedSecretNamespaceListerExpansion
}

// sharedSecretNamespaceLister implements the SharedSecretNamespaceLister
// interface.
type sharedSecretNamespaceLister struct {
	listers.ResourceIndexer[*managementv1.SharedSecret]
}
