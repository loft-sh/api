// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/pkg/apis/management/v1"
	"github.com/loft-sh/api/pkg/client/clientset_generated/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type ManagementV1Interface interface {
	RESTClient() rest.Interface
	AnnouncementsGetter
	ClustersGetter
	ClusterAccountTemplatesGetter
	ClusterConnectsGetter
	ClusterRolesGetter
	ConfigsGetter
	FeaturesGetter
	LicensesGetter
	LicenseTokensGetter
	LoftRestartsGetter
	LoftUpgradesGetter
	SelfSubjectAccessReviewsGetter
	SharedSecretsGetter
	SubjectAccessReviewsGetter
	TeamsGetter
	UsersGetter
}

// ManagementV1Client is used to interact with features provided by the management.loft.sh group.
type ManagementV1Client struct {
	restClient rest.Interface
}

func (c *ManagementV1Client) Announcements() AnnouncementInterface {
	return newAnnouncements(c)
}

func (c *ManagementV1Client) Clusters() ClusterInterface {
	return newClusters(c)
}

func (c *ManagementV1Client) ClusterAccountTemplates() ClusterAccountTemplateInterface {
	return newClusterAccountTemplates(c)
}

func (c *ManagementV1Client) ClusterConnects() ClusterConnectInterface {
	return newClusterConnects(c)
}

func (c *ManagementV1Client) ClusterRoles() ClusterRoleInterface {
	return newClusterRoles(c)
}

func (c *ManagementV1Client) Configs() ConfigInterface {
	return newConfigs(c)
}

func (c *ManagementV1Client) Features() FeatureInterface {
	return newFeatures(c)
}

func (c *ManagementV1Client) Licenses() LicenseInterface {
	return newLicenses(c)
}

func (c *ManagementV1Client) LicenseTokens() LicenseTokenInterface {
	return newLicenseTokens(c)
}

func (c *ManagementV1Client) LoftRestarts() LoftRestartInterface {
	return newLoftRestarts(c)
}

func (c *ManagementV1Client) LoftUpgrades() LoftUpgradeInterface {
	return newLoftUpgrades(c)
}

func (c *ManagementV1Client) SelfSubjectAccessReviews() SelfSubjectAccessReviewInterface {
	return newSelfSubjectAccessReviews(c)
}

func (c *ManagementV1Client) SharedSecrets(namespace string) SharedSecretInterface {
	return newSharedSecrets(c, namespace)
}

func (c *ManagementV1Client) SubjectAccessReviews() SubjectAccessReviewInterface {
	return newSubjectAccessReviews(c)
}

func (c *ManagementV1Client) Teams() TeamInterface {
	return newTeams(c)
}

func (c *ManagementV1Client) Users() UserInterface {
	return newUsers(c)
}

// NewForConfig creates a new ManagementV1Client for the given config.
func NewForConfig(c *rest.Config) (*ManagementV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ManagementV1Client{client}, nil
}

// NewForConfigOrDie creates a new ManagementV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ManagementV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ManagementV1Client for the given RESTClient.
func New(c rest.Interface) *ManagementV1Client {
	return &ManagementV1Client{c}
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
func (c *ManagementV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
