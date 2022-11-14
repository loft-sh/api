package managerfactory

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SharedManagerFactory is the interface for retrieving managers
type SharedManagerFactory interface {
	CreateManager(cluster string, localClient client.Client) error
	Stop(cluster string)

	Cluster(cluster string) ClusterClientAccess
	Management() ManagementClientAccess
}

// ClusterClientAccess holds the functions for cluster access
type ClusterClientAccess interface {
	Scheme() (*runtime.Scheme, error)
	Config() (*rest.Config, error)
	RESTMapper() (meta.RESTMapper, error)
	UncachedClient() (client.Client, error)
	CachedClient() (client.Client, error)
}

// ManagementClientAccess holds the functions for management access
type ManagementClientAccess interface {
	Scheme() *runtime.Scheme
	Config() *rest.Config
	RESTMapper() meta.RESTMapper
	UncachedClient() client.Client
	CachedClient() client.Client
}
