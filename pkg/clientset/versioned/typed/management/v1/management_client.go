// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	http "net/http"

	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ManagementV1Interface interface {
	RESTClient() rest.Interface
	AgentAuditEventsGetter
	AnnouncementsGetter
	AppsGetter
	BackupsGetter
	ClustersGetter
	ClusterAccessesGetter
	ClusterRoleTemplatesGetter
	ConfigsGetter
	ConvertVirtualClusterConfigsGetter
	DevPodEnvironmentTemplatesGetter
	DevPodWorkspaceInstancesGetter
	DevPodWorkspacePresetsGetter
	DevPodWorkspaceTemplatesGetter
	DirectClusterEndpointTokensGetter
	EventsGetter
	FeaturesGetter
	IngressAuthTokensGetter
	LicensesGetter
	LicenseTokensGetter
	LoftUpgradesGetter
	OIDCClientsGetter
	OwnedAccessKeysGetter
	ProjectsGetter
	ProjectSecretsGetter
	RedirectTokensGetter
	RegisterVirtualClustersGetter
	ResetAccessKeysGetter
	SelvesGetter
	SelfSubjectAccessReviewsGetter
	SharedSecretsGetter
	SpaceInstancesGetter
	SpaceTemplatesGetter
	SubjectAccessReviewsGetter
	TasksGetter
	TeamsGetter
	TranslateVClusterResourceNamesGetter
	UsersGetter
	VirtualClusterInstancesGetter
	VirtualClusterSchemasGetter
	VirtualClusterTemplatesGetter
}

// ManagementV1Client is used to interact with features provided by the management.loft.sh group.
type ManagementV1Client struct {
	restClient rest.Interface
}

func (c *ManagementV1Client) AgentAuditEvents() AgentAuditEventInterface {
	return newAgentAuditEvents(c)
}

func (c *ManagementV1Client) Announcements() AnnouncementInterface {
	return newAnnouncements(c)
}

func (c *ManagementV1Client) Apps() AppInterface {
	return newApps(c)
}

func (c *ManagementV1Client) Backups() BackupInterface {
	return newBackups(c)
}

func (c *ManagementV1Client) Clusters() ClusterInterface {
	return newClusters(c)
}

func (c *ManagementV1Client) ClusterAccesses() ClusterAccessInterface {
	return newClusterAccesses(c)
}

func (c *ManagementV1Client) ClusterRoleTemplates() ClusterRoleTemplateInterface {
	return newClusterRoleTemplates(c)
}

func (c *ManagementV1Client) Configs() ConfigInterface {
	return newConfigs(c)
}

func (c *ManagementV1Client) ConvertVirtualClusterConfigs() ConvertVirtualClusterConfigInterface {
	return newConvertVirtualClusterConfigs(c)
}

func (c *ManagementV1Client) DevPodEnvironmentTemplates() DevPodEnvironmentTemplateInterface {
	return newDevPodEnvironmentTemplates(c)
}

func (c *ManagementV1Client) DevPodWorkspaceInstances(namespace string) DevPodWorkspaceInstanceInterface {
	return newDevPodWorkspaceInstances(c, namespace)
}

func (c *ManagementV1Client) DevPodWorkspacePresets() DevPodWorkspacePresetInterface {
	return newDevPodWorkspacePresets(c)
}

func (c *ManagementV1Client) DevPodWorkspaceTemplates() DevPodWorkspaceTemplateInterface {
	return newDevPodWorkspaceTemplates(c)
}

func (c *ManagementV1Client) DirectClusterEndpointTokens() DirectClusterEndpointTokenInterface {
	return newDirectClusterEndpointTokens(c)
}

func (c *ManagementV1Client) Events() EventInterface {
	return newEvents(c)
}

func (c *ManagementV1Client) Features() FeatureInterface {
	return newFeatures(c)
}

func (c *ManagementV1Client) IngressAuthTokens() IngressAuthTokenInterface {
	return newIngressAuthTokens(c)
}

func (c *ManagementV1Client) Licenses() LicenseInterface {
	return newLicenses(c)
}

func (c *ManagementV1Client) LicenseTokens() LicenseTokenInterface {
	return newLicenseTokens(c)
}

func (c *ManagementV1Client) LoftUpgrades() LoftUpgradeInterface {
	return newLoftUpgrades(c)
}

func (c *ManagementV1Client) OIDCClients() OIDCClientInterface {
	return newOIDCClients(c)
}

func (c *ManagementV1Client) OwnedAccessKeys() OwnedAccessKeyInterface {
	return newOwnedAccessKeys(c)
}

func (c *ManagementV1Client) Projects() ProjectInterface {
	return newProjects(c)
}

func (c *ManagementV1Client) ProjectSecrets(namespace string) ProjectSecretInterface {
	return newProjectSecrets(c, namespace)
}

func (c *ManagementV1Client) RedirectTokens() RedirectTokenInterface {
	return newRedirectTokens(c)
}

func (c *ManagementV1Client) RegisterVirtualClusters() RegisterVirtualClusterInterface {
	return newRegisterVirtualClusters(c)
}

func (c *ManagementV1Client) ResetAccessKeys() ResetAccessKeyInterface {
	return newResetAccessKeys(c)
}

func (c *ManagementV1Client) Selves() SelfInterface {
	return newSelves(c)
}

func (c *ManagementV1Client) SelfSubjectAccessReviews() SelfSubjectAccessReviewInterface {
	return newSelfSubjectAccessReviews(c)
}

func (c *ManagementV1Client) SharedSecrets(namespace string) SharedSecretInterface {
	return newSharedSecrets(c, namespace)
}

func (c *ManagementV1Client) SpaceInstances(namespace string) SpaceInstanceInterface {
	return newSpaceInstances(c, namespace)
}

func (c *ManagementV1Client) SpaceTemplates() SpaceTemplateInterface {
	return newSpaceTemplates(c)
}

func (c *ManagementV1Client) SubjectAccessReviews() SubjectAccessReviewInterface {
	return newSubjectAccessReviews(c)
}

func (c *ManagementV1Client) Tasks() TaskInterface {
	return newTasks(c)
}

func (c *ManagementV1Client) Teams() TeamInterface {
	return newTeams(c)
}

func (c *ManagementV1Client) TranslateVClusterResourceNames() TranslateVClusterResourceNameInterface {
	return newTranslateVClusterResourceNames(c)
}

func (c *ManagementV1Client) Users() UserInterface {
	return newUsers(c)
}

func (c *ManagementV1Client) VirtualClusterInstances(namespace string) VirtualClusterInstanceInterface {
	return newVirtualClusterInstances(c, namespace)
}

func (c *ManagementV1Client) VirtualClusterSchemas() VirtualClusterSchemaInterface {
	return newVirtualClusterSchemas(c)
}

func (c *ManagementV1Client) VirtualClusterTemplates() VirtualClusterTemplateInterface {
	return newVirtualClusterTemplates(c)
}

// NewForConfig creates a new ManagementV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ManagementV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ManagementV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ManagementV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
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
	gv := managementv1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

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
