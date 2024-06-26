// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AgentAuditEventLister helps list AgentAuditEvents.
// All objects returned here must be treated as read-only.
type AgentAuditEventLister interface {
	// List lists all AgentAuditEvents in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AgentAuditEvent, err error)
	// Get retrieves the AgentAuditEvent from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AgentAuditEvent, error)
	AgentAuditEventListerExpansion
}

// agentAuditEventLister implements the AgentAuditEventLister interface.
type agentAuditEventLister struct {
	indexer cache.Indexer
}

// NewAgentAuditEventLister returns a new AgentAuditEventLister.
func NewAgentAuditEventLister(indexer cache.Indexer) AgentAuditEventLister {
	return &agentAuditEventLister{indexer: indexer}
}

// List lists all AgentAuditEvents in the indexer.
func (s *agentAuditEventLister) List(selector labels.Selector) (ret []*v1.AgentAuditEvent, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AgentAuditEvent))
	})
	return ret, err
}

// Get retrieves the AgentAuditEvent from the index for a given name.
func (s *agentAuditEventLister) Get(name string) (*v1.AgentAuditEvent, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("agentauditevent"), name)
	}
	return obj.(*v1.AgentAuditEvent), nil
}
