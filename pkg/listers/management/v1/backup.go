// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackupLister helps list Backups.
// All objects returned here must be treated as read-only.
type BackupLister interface {
	// List lists all Backups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Backup, err error)
	// Get retrieves the Backup from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Backup, error)
	BackupListerExpansion
}

// backupLister implements the BackupLister interface.
type backupLister struct {
	indexer cache.Indexer
}

// NewBackupLister returns a new BackupLister.
func NewBackupLister(indexer cache.Indexer) BackupLister {
	return &backupLister{indexer: indexer}
}

// List lists all Backups in the indexer.
func (s *backupLister) List(selector labels.Selector) (ret []*v1.Backup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Backup))
	})
	return ret, err
}

// Get retrieves the Backup from the index for a given name.
func (s *backupLister) Get(name string) (*v1.Backup, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("backup"), name)
	}
	return obj.(*v1.Backup), nil
}
