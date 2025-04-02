// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// BackupLister helps list Backups.
// All objects returned here must be treated as read-only.
type BackupLister interface {
	// List lists all Backups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*managementv1.Backup, err error)
	// Get retrieves the Backup from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*managementv1.Backup, error)
	BackupListerExpansion
}

// backupLister implements the BackupLister interface.
type backupLister struct {
	listers.ResourceIndexer[*managementv1.Backup]
}

// NewBackupLister returns a new BackupLister.
func NewBackupLister(indexer cache.Indexer) BackupLister {
	return &backupLister{listers.New[*managementv1.Backup](indexer, managementv1.Resource("backup"))}
}
