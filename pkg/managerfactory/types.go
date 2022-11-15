package managerfactory

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SharedManagerFactory is the interface for retrieving managers
type SharedManagerFactory interface {
	Cluster(cluster string) ClusterClientAccess
	Management() ManagementClientAccess
}

// ClusterClientAccess holds the functions for cluster access
type ClusterClientAccess interface {
	Config() (*rest.Config, error)
	UncachedClient() (client.Client, error)
}

// ManagementClientAccess holds the functions for management access
type ManagementClientAccess interface {
	Config() *rest.Config
	UncachedClient() client.Client
	CachedClient() client.Client
}
