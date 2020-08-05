// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	"github.com/loft-sh/api/pkg/client/clientset_generated/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type StorageV1Interface interface {
	RESTClient() rest.Interface
	ClustersGetter
	TeamsGetter
	UsersGetter
	VirtualClustersGetter
}

// StorageV1Client is used to interact with features provided by the storage.loft.sh group.
type StorageV1Client struct {
	restClient rest.Interface
}

func (c *StorageV1Client) Clusters() ClusterInterface {
	return newClusters(c)
}

func (c *StorageV1Client) Teams() TeamInterface {
	return newTeams(c)
}

func (c *StorageV1Client) Users() UserInterface {
	return newUsers(c)
}

func (c *StorageV1Client) VirtualClusters(namespace string) VirtualClusterInterface {
	return newVirtualClusters(c, namespace)
}

// NewForConfig creates a new StorageV1Client for the given config.
func NewForConfig(c *rest.Config) (*StorageV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &StorageV1Client{client}, nil
}

// NewForConfigOrDie creates a new StorageV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *StorageV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new StorageV1Client for the given RESTClient.
func New(c rest.Interface) *StorageV1Client {
	return &StorageV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *StorageV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
