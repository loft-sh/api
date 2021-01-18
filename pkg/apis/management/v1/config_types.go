package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User holds the user information
// +k8s:openapi-gen=true
// +resource:path=configs,rest=ConfigREST
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSpec   `json:"spec,omitempty"`
	Status ConfigStatus `json:"status,omitempty"`
}

// ConfigSpec holds the specification
type ConfigSpec struct {
	// Authentication holds the information for authentication
	Authentication Authentication `json:"auth,omitempty"`

	// OIDC holds oidc provider relevant information
	OIDC *OIDC `json:"oidc,omitempty"`

	// Apps holds configuration around apps
	// +optional
	Apps *Apps `json:"apps,omitempty"`
}

// ConfigStatus holds the status
type ConfigStatus struct {
}

// Apps holds configuration for apps that should be shown
type Apps struct {
	// If this option is true, loft will not try to parse the default apps
	// +optional
	NoDefault bool `json:"noDefault,omitempty"`

	// These are additional repositories that are parsed by loft
	// +optional
	Repositories []storagev1.HelmChartRepository `json:"repositories,omitempty"`

	// Predefined apps that can be selected in the Spaces > Space menu
	// +optional
	PredefinedApps []PredefinedApp `json:"predefinedApps,omitempty"`
}

// PredefinedApp holds information about a predefined app
type PredefinedApp struct {
	// Chart holds the repo/chart name of the predefined app
	Chart string `json:"chart"`

	// InitialVersion holds the initial version of this app.
	// This version will be selected automatically.
	// +optional
	InitialVersion string `json:"initialVersion,omitempty"`

	// InitialValues holds the initial values for this app.
	// The values will be prefilled automatically. There are certain
	// placeholders that can be used within the values that are replaced
	// by the loft UI automatically.
	// +optional
	InitialValues string `json:"initialValues,omitempty"`

	// Holds the cluster names where to display this app
	// +optional
	Clusters []string `json:"clusters,omitempty"`

	// Title is the name that should be displayed for the predefined app.
	// If empty the chart name is used.
	// +optional
	Title string `json:"title,omitempty"`

	// IconURL specifies an url to the icon that should be displayed for this app.
	// If none is specified the icon from the chart metadata is used.
	// +optional
	IconURL string `json:"iconUrl,omitempty"`

	// ReadmeURL specifies an url to the readme page of this predefined app. If empty
	// an url will be constructed to artifact hub.
	// +optional
	ReadmeURL string `json:"readmeUrl,omitempty"`
}

// OIDC holds oidc provider relevant information
type OIDC struct {
	// If true indicates that loft will act as an OIDC server
	Enabled bool `json:"enabled,omitempty"`

	// The clients that are allowed to request loft tokens
	Clients []OIDCClient `json:"clients,omitempty"`
}

// OIDCClient holds information about a client
type OIDCClient struct {
	// The client name
	Name string `json:"name,omitempty"`

	// The client id of the client
	ClientID string `json:"clientId,omitempty"`

	// The client secret of the client
	ClientSecret string `json:"clientSecret,omitempty"`

	// A registered set of redirect URIs. When redirecting from dex to the client, the URI
	// requested to redirect to MUST match one of these values, unless the client is "public".
	RedirectURIs []string `json:"redirectURIs"`
}

// Authentication holds authentication relevant information
type Authentication struct {
	// Password holds password authentication relevant information
	Password *AuthenticationPassword `json:"password,omitempty"`

	// OIDC holds oidc authentication configuration
	OIDC *AuthenticationOIDC `json:"oidc,omitempty"`
}

type AuthenticationPassword struct {
	// If true login via password is disabled
	Disabled bool `json:"disabled,omitempty"`
}

type AuthenticationOIDC struct {
	// IssuerURL is the URL the provider signs ID Tokens as. This will be the "iss"
	// field of all tokens produced by the provider and is used for configuration
	// discovery.
	//
	// The URL is usually the provider's URL without a path, for example
	// "https://accounts.google.com" or "https://login.salesforce.com".
	//
	// The provider must implement configuration discovery.
	// See: https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfig
	// +optional
	IssuerURL string `json:"issuerUrl,omitempty"`

	// ClientID the JWT must be issued for, the "sub" field. This plugin only trusts a single
	// client to ensure the plugin can be used with public providers.
	//
	// The plugin supports the "authorized party" OpenID Connect claim, which allows
	// specialized providers to issue tokens to a client for a different client.
	// See: https://openid.net/specs/openid-connect-core-1_0.html#IDToken
	// +optional
	ClientID string `json:"clientId,omitempty"`

	// ClientSecret to issue tokens from the OIDC provider
	// +optional
	ClientSecret string `json:"clientSecret,omitempty"`

	// Path to a PEM encoded root certificate of the provider. Optional
	// +optional
	CAFile string `json:"caFile,omitempty"`

	// UsernameClaim is the JWT field to use as the user's username.
	// +optional
	UsernameClaim string `json:"usernameClaim,omitempty"`

	// UsernamePrefix, if specified, causes claims mapping to username to be prefix with
	// the provided value. A value "oidc:" would result in usernames like "oidc:john".
	// +optional
	UsernamePrefix string `json:"usernamePrefix,omitempty"`

	// GroupsClaim, if specified, causes the OIDCAuthenticator to try to populate the user's
	// groups with an ID Token field. If the GroupsClaim field is present in an ID Token the value
	// must be a string or list of strings.
	// +optional
	GroupsClaim string `json:"groupsClaim,omitempty"`

	// GetUserInfo, if specified, tells the OIDCAuthenticator to try to populate the user's
	// information from the UserInfo.
	// +optional
	GetUserInfo bool `json:"getUserInfo,omitempty"`

	// GroupsPrefix, if specified, causes claims mapping to group names to be prefixed with the
	// value. A value "oidc:" would result in groups like "oidc:engineering" and "oidc:marketing".
	// +optional
	GroupsPrefix string `json:"groupsPrefix,omitempty"`

	// Type of the OIDC to show in the UI. Only for displaying purposes
	// +optional
	Type string `json:"type,omitempty"`

	// Cluster Account Templates that will be applied for users logging in through this authentication
	// +optional
	ClusterAccountTemplates []storagev1.UserClusterAccountTemplate `json:"clusterAccountTemplates,omitempty"`

	// A mapping between groups and cluster account templates. If the user has a certain group, the cluster
	// account template will be added during creation
	// +optional
	GroupClusterAccountTemplates map[string][]storagev1.UserClusterAccountTemplate `json:"groupClusterAccountTemplates,omitempty"`
}
