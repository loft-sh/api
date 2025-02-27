// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// TeamLister helps list Teams.
// All objects returned here must be treated as read-only.
type TeamLister interface {
	// List lists all Teams in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*storagev1.Team, err error)
	// Get retrieves the Team from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*storagev1.Team, error)
	TeamListerExpansion
}

// teamLister implements the TeamLister interface.
type teamLister struct {
	listers.ResourceIndexer[*storagev1.Team]
}

// NewTeamLister returns a new TeamLister.
func NewTeamLister(indexer cache.Indexer) TeamLister {
	return &teamLister{listers.New[*storagev1.Team](indexer, storagev1.Resource("team"))}
}
