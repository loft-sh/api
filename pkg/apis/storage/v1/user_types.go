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

	// A reference to the users access keys
	// +optional
	AccessKeysRef *SecretRef `json:"accessKeysRef,omitempty"`

	// A reference to the users access keys
	// +optional
	CodesRef *SecretRef `json:"codesRef,omitempty"`

	// ClusterAccountTemplates that should be applied for the user
	// +optional
	ClusterAccountTemplates []UserClusterAccountTemplate `json:"clusterAccountTemplates,omitempty"`
}

type UserClusterAccountTemplate struct {
	// Name of the cluster account template to apply
	// +optional
	Name string `json:"name,omitempty"`

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

	// ClusterAccountTemplates holds information about which cluster account templates were applied
	// +optional
	ClusterAccountTemplates []UserClusterAccountTemplateStatus `json:"clusterAccountTemplates,omitempty"`
}

type UserClusterAccountTemplateStatus struct {
	// Name of the cluster account template that was applied
	// +optional
	Name string `json:"name,omitempty"`

	// Clusters holds the cluster on which this template was applied
	// +optional
	Clusters []ClusterAccountTemplateClusterStatus `json:"clusters,omitempty"`

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

type ClusterAccountTemplateClusterStatus struct {
	// Name of the cluster where the cluster account template was applied
	// +optional
	Name string `json:"name,omitempty"`

	// Status holds the status of the account in the target cluster
	// +optional
	Status ClusterAccountTemplateClusterStatusPhase `json:"phase,omitempty"`

	// Reason describes why loft couldn't sync the account with a machine readable identifier
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message describes why loft couldn't sync the account in human language
	// +optional
	Message string `json:"message,omitempty"`
}

// ClusterAccountTemplateClusterStatusPhase describes the phase of a cluster account template
type ClusterAccountTemplateClusterStatusPhase string

// These are the valid admin account types
const (
	ClusterAccountTemplateClusterStatusPhaseCreated ClusterAccountTemplateClusterStatusPhase = "Created"
	ClusterAccountTemplateClusterStatusPhaseSkipped ClusterAccountTemplateClusterStatusPhase = "Skipped"
	ClusterAccountTemplateClusterStatusPhaseFailed  ClusterAccountTemplateClusterStatusPhase = "Failed"
)

// ClusterAccountTemplateStatusPhase describes the phase of a cluster account template
type ClusterAccountTemplateStatusPhase string

// These are the valid admin account types
const (
	ClusterAccountTemplateStatusPhaseCompleted ClusterAccountTemplateStatusPhase = "Completed"
	ClusterAccountTemplateStatusPhaseFailed    ClusterAccountTemplateStatusPhase = "Failed"
)

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

	// Accounts is the account name of the user in the cluster
	// +optional
	Accounts []string `json:"accounts,omitempty"`
}

// AccountClusterStatusPhase describes the phase of a cluster
type AccountClusterStatusPhase string

// These are the valid admin account types
const (
	AccountClusterStatusPhaseSynced AccountClusterStatusPhase = "Synced"
	AccountClusterStatusPhaseFailed AccountClusterStatusPhase = "Failed"
)

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
