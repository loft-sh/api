package managerfactory

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RemoteRunner can be registered in the factory and will be called if a new manager is created
type RemoteRunner interface {
	Register(cluster string, clusterManager ctrl.Manager, managementManager ctrl.Manager) error
}

// SharedManagerFactory is the interface for retrieving managers
type SharedManagerFactory interface {
	CreateManager(cluster string, localClient client.Client) error
	RegisterRunner(RemoteRunner)
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
