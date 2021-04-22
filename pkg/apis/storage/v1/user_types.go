package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User holds the user information
// +k8s:openapi-gen=true
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserSpec   `json:"spec,omitempty"`
	Status UserStatus `json:"status,omitempty"`
}

type UserSpec struct {
	// The display name shown in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// The username that is used to login
	Username string `json:"username,omitempty"`

	// The URL to an icon that should be shown for the user
	// +optional
	Icon string `json:"icon,omitempty"`

	// The users email address
	// +optional
	Email string `json:"email,omitempty"`

	// The user subject as presented by the token
	Subject string `json:"subject,omitempty"`

	// The groups the user has access to
	// +optional
	Groups []string `json:"groups,omitempty"`

	// The extras the user has
	// +optional
	Extras map[string][]string `json:"extras,omitempty"`

	// A reference to the user password
	// +optional
	PasswordRef *SecretRef `json:"passwordRef,omitempty"`

	// Deprecated: Use the Access Key CRD instead
	// A reference to the users access keys
	// +optional
	AccessKeysRef *SecretRef `json:"accessKeysRef,omitempty"`

	// A reference to the users access keys
	// +optional
	CodesRef *SecretRef `json:"codesRef,omitempty"`

	// ImagePullSecrets holds secret references to image pull
	// secrets the user has access to.
	// +optional
	ImagePullSecrets []*KindSecretRef `json:"imagePullSecrets,omitempty"`

	// ClusterAccountTemplates that should be applied for the user
	// +optional
	ClusterAccountTemplates []UserClusterAccountTemplate `json:"clusterAccountTemplates,omitempty"`

	// TokenGeneration can be used to invalidate all user tokens
	// +optional
	TokenGeneration int64 `json:"tokenGeneration,omitempty"`

	// If disabled is true, an user will not be able to login anymore. All other user resources
	// are unaffected and other users can still interact with this user
	// +optional
	Disabled bool `json:"disabled,omitempty"`
}

type UserClusterAccountTemplate struct {
	// Name of the cluster account template to apply
	// +optional
	Name string `json:"name,omitempty"`

	// Sync defines if Loft should sync changes to the cluster account template
	// to the cluster accounts and create new accounts if new clusters match the templates.
	// +optional
	Sync bool `json:"sync,omitempty"`

	// AccountName is the name of the account that should
	// be created. Defaults to the user or team kubernetes name.
	// +optional
	AccountName string `json:"accountName,omitempty"`
}

// UserStatus holds the status of an user
type UserStatus struct {
	// Clusters holds information about which clusters the user has accounts in
	// +optional
	Clusters []AccountClusterStatus `json:"clusters,omitempty"`
}

// AccountClusterStatus holds the status of an account in a cluster
type AccountClusterStatus struct {
	// Status holds the status of the account in the target cluster
	// +optional
	Status AccountClusterStatusPhase `json:"phase,omitempty"`

	// Reason describes why loft couldn't sync the account with a machine readable identifier
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message describes why loft couldn't sync the account in human language
	// +optional
	Message string `json:"message,omitempty"`

	// Cluster is the cluster name of the user in the cluster
	// +optional
	Cluster string `json:"cluster,omitempty"`

	// AccountsClusterTemplate status is the status of the account cluster template that was used
	// to create the cluster account
	// +optional
	AccountsClusterTemplate []AccountClusterTemplateStatus `json:"accountsClusterTemplateStatus,omitempty"`

	// Accounts is the account name of the user in the cluster
	// +optional
	Accounts []string `json:"accounts,omitempty"`
}

type AccountClusterTemplateStatus struct {
	// Name is the name of the cluster account template
	// +optional
	Name string `json:"name,omitempty"`
	
	// Account is the name of the account in the cluster
	// +optional
	Account string `json:"account,omitempty"`

	// AccountTemplateHash is the hash of the account template that was applied
	// +optional
	AccountTemplateHash string `json:"accountTemplateHash,omitempty"`

	// OwnsHash is the hash of the owns part of the cluster account template that was
	// applied
	// +optional
	OwnsHash string `json:"ownsHash,omitempty"`

	// Status holds the status of the account in the target cluster
	// +optional
	Status ClusterAccountTemplateStatusPhase `json:"phase,omitempty"`

	// Reason describes why loft couldn't sync the account with a machine readable identifier
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message describes why loft couldn't sync the account in human language
	// +optional
	Message string `json:"message,omitempty"`
}

// ClusterAccountTemplateStatusPhase describes the phase of a cluster
type ClusterAccountTemplateStatusPhase string

// These are the valid admin account types
const (
	ClusterAccountTemplateStatusPhaseSynced ClusterAccountTemplateStatusPhase = "Synced"
	ClusterAccountTemplateStatusPhaseFailedAccount ClusterAccountTemplateStatusPhase = "FailedAccount"
	ClusterAccountTemplateStatusPhaseFailedObjects ClusterAccountTemplateStatusPhase = "FailedObjects"
)

// AccountClusterStatusPhase describes the phase of a cluster
type AccountClusterStatusPhase string

// These are the valid admin account types
const (
	AccountClusterStatusPhaseSynced AccountClusterStatusPhase = "Synced"
	AccountClusterStatusPhaseFailed AccountClusterStatusPhase = "Failed"
)

// KindSecretRef is the reference to a secret containing the user password
type KindSecretRef struct {
	// APIGroup is the api group of the secret
	APIGroup string `json:"apiGroup,omitempty"`
	// Kind is the kind of the secret
	Kind string `json:"kind,omitempty"`
	// +optional
	SecretName string `json:"secretName,omitempty"`
	// +optional
	SecretNamespace string `json:"secretNamespace,omitempty"`
	// +optional
	Key string `json:"key,omitempty"`
}

// SecretRef is the reference to a secret containing the user password
type SecretRef struct {
	// +optional
	SecretName string `json:"secretName,omitempty"`
	// +optional
	SecretNamespace string `json:"secretNamespace,omitempty"`
	// +optional
	Key string `json:"key,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserList contains a list of User
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
