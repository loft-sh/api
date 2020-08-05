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
	Register(cluster string, mgr ctrl.Manager) error
}

// SharedManagerFactory is the interface for retrieving managers
type SharedManagerFactory interface {
	CreateManager(cluster string, localClient client.Client) error
	GetClient(cluster string) (client.Client, error)
	GetUncachedClient(cluster string) (client.Client, error)
	GetRESTMapper(cluster string) (meta.RESTMapper, error)
	GetClientConfig(cluster string) (*rest.Config, error)
	GetScheme() *runtime.Scheme
	GetManagementConfig() *rest.Config
	GetManagementClient() client.Client
	GetManagementRESTMapper() meta.RESTMapper
	RegisterRunner(RemoteRunner)
	Stop(cluster string)
}
