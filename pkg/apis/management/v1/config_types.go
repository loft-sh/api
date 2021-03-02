package v1

import (
	auditv1 "github.com/loft-sh/api/pkg/apis/audit/v1"
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
	
	// Audit holds audit configuration
	// +optional
	Audit *Audit `json:"audit,omitempty"`
}

// ConfigStatus holds the status
type ConfigStatus struct {
}

// Audit holds the audit configuration options for loft. Changing any options will require a loft restart
// to take effect.
type Audit struct {
	// If audit is enabled and incoming api requests will be logged based on the supplied policy.  
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	
	// The audit policy to use and log requests. By default loft will not log anything
	Policy AuditPolicy `json:"policy,omitempty"`
	
	// The path where to save the audit log files. This is required if audit is enabled. Backup log files will 
	// be retained in the same directory.
	Path string `json:"path,omitempty"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	// +optional
	MaxAge int `json:"maxAge,omitempty"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	// +optional
	MaxBackups int `json:"maxBackups,omitempty"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	// +optional
	MaxSize int `json:"maxSize,omitempty"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	// +optional
	Compress bool `json:"compress,omitempty"`
}

// AuditPolicy describes the audit policy to use for loft
type AuditPolicy struct {
	// Rules specify the audit Level a request should be recorded at.
	// A request may match multiple rules, in which case the FIRST matching rule is used.
	// The default audit level is None, but can be overridden by a catch-all rule at the end of the list.
	// PolicyRules are strictly ordered.
	Rules []AuditPolicyRule `json:"rules,omitempty"`

	// OmitStages is a list of stages for which no events are created. Note that this can also
	// be specified per rule in which case the union of both are omitted.
	// +optional
	OmitStages []auditv1.Stage `json:"omitStages,omitempty"`
}

// AuditPolicy describes a policy for auditing
type AuditPolicyRule struct {
	// The Level that requests matching this rule are recorded at.
	Level auditv1.Level `json:"level"`

	// The users (by authenticated user name) this rule applies to.
	// An empty list implies every user.
	// +optional
	Users []string `json:"users,omitempty"`
	// The user groups this rule applies to. A user is considered matching
	// if it is a member of any of the UserGroups.
	// An empty list implies every user group.
	// +optional
	UserGroups []string `json:"userGroups,omitempty"`

	// The verbs that match this rule.
	// An empty list implies every verb.
	// +optional
	Verbs []string `json:"verbs,omitempty"`

	// Rules can apply to API resources (such as "pods" or "secrets"),
	// non-resource URL paths (such as "/api"), or neither, but not both.
	// If neither is specified, the rule is treated as a default for all URLs.

	// Resources that this rule matches. An empty list implies all kinds in all API groups.
	// +optional
	Resources []GroupResources `json:"resources,omitempty"`
	// Namespaces that this rule matches.
	// The empty string "" matches non-namespaced resources.
	// An empty list implies every namespace.
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`

	// NonResourceURLs is a set of URL paths that should be audited.
	// *s are allowed, but only as the full, final step in the path.
	// Examples:
	//  "/metrics" - Log requests for apiserver metrics
	//  "/healthz*" - Log all health checks
	// +optional
	NonResourceURLs []string `json:"nonResourceURLs,omitempty"`

	// OmitStages is a list of stages for which no events are created. Note that this can also
	// be specified policy wide in which case the union of both are omitted.
	// An empty list means no restrictions will apply.
	// +optional
	OmitStages []auditv1.Stage `json:"omitStages,omitempty" protobuf:"bytes,8,rep,name=omitStages"`

	// RequestTargets is a list of request targets for which events are created. 
	// An empty list implies every request.
	// +optional
	RequestTargets []auditv1.RequestTarget `json:"requestTargets,omitempty"`
	
	// Clusters that this rule matches. Only applies to cluster requests.
	// If this is set, no events for non cluster requests will be created.
	// An empty list means no restrictions will apply.
	// +optional
	Clusters []string `json:"clusters,omitempty"` 
}

// GroupResources represents resource kinds in an API group.
type GroupResources struct {
	// Group is the name of the API group that contains the resources.
	// The empty string represents the core API group.
	// +optional
	Group string `json:"group,omitempty" protobuf:"bytes,1,opt,name=group"`
	// Resources is a list of resources this rule applies to.
	//
	// For example:
	// 'pods' matches pods.
	// 'pods/log' matches the log subresource of pods.
	// '*' matches all resources and their subresources.
	// 'pods/*' matches all subresources of pods.
	// '*/scale' matches all scale subresources.
	//
	// If wildcard is present, the validation rule will ensure resources do not
	// overlap with each other.
	//
	// An empty list implies all resources and subresources in this API groups apply.
	// +optional
	Resources []string `json:"resources,omitempty" protobuf:"bytes,2,rep,name=resources"`
	// ResourceNames is a list of resource instance names that the policy matches.
	// Using this field requires Resources to be specified.
	// An empty list implies that every instance of the resource is matched.
	// +optional
	ResourceNames []string `json:"resourceNames,omitempty" protobuf:"bytes,3,rep,name=resourceNames"`
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
	GroupClusterAccountTemplates []AuthenticationGroupClusterAccountTemplate `json:"groupClusterAccountTemplates,omitempty"`
}

type AuthenticationGroupClusterAccountTemplate struct {
	// Group is the name of the group that should be matched
	Group string `json:"group"`

	// Cluster Account Templates that will be applied for users logging in through this authentication
	// +optional
	ClusterAccountTemplates []storagev1.UserClusterAccountTemplate `json:"clusterAccountTemplates,omitempty"`
}
