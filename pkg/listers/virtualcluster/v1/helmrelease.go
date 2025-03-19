// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	virtualclusterv1 "github.com/loft-sh/api/v4/pkg/apis/virtualcluster/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// HelmReleaseLister helps list HelmReleases.
// All objects returned here must be treated as read-only.
type HelmReleaseLister interface {
	// List lists all HelmReleases in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*virtualclusterv1.HelmRelease, err error)
	// HelmReleases returns an object that can list and get HelmReleases.
	HelmReleases(namespace string) HelmReleaseNamespaceLister
	HelmReleaseListerExpansion
}

// helmReleaseLister implements the HelmReleaseLister interface.
type helmReleaseLister struct {
	listers.ResourceIndexer[*virtualclusterv1.HelmRelease]
}

// NewHelmReleaseLister returns a new HelmReleaseLister.
func NewHelmReleaseLister(indexer cache.Indexer) HelmReleaseLister {
	return &helmReleaseLister{listers.New[*virtualclusterv1.HelmRelease](indexer, virtualclusterv1.Resource("helmrelease"))}
}

// HelmReleases returns an object that can list and get HelmReleases.
func (s *helmReleaseLister) HelmReleases(namespace string) HelmReleaseNamespaceLister {
	return helmReleaseNamespaceLister{listers.NewNamespaced[*virtualclusterv1.HelmRelease](s.ResourceIndexer, namespace)}
}

// HelmReleaseNamespaceLister helps list and get HelmReleases.
// All objects returned here must be treated as read-only.
type HelmReleaseNamespaceLister interface {
	// List lists all HelmReleases in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*virtualclusterv1.HelmRelease, err error)
	// Get retrieves the HelmRelease from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*virtualclusterv1.HelmRelease, error)
	HelmReleaseNamespaceListerExpansion
}

// helmReleaseNamespaceLister implements the HelmReleaseNamespaceLister
// interface.
type helmReleaseNamespaceLister struct {
	listers.ResourceIndexer[*virtualclusterv1.HelmRelease]
}
