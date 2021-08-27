package v1

import (
	configv1alpha1 "github.com/loft-sh/agentapi/pkg/apis/kiosk/config/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type ClusterMembers struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Teams holds all the teams that have access to the cluster
	Teams []ClusterMember `json:"teams,omitempty"`

	// Users holds all the users that have access to the cluster
	Users []ClusterMember `json:"users,omitempty"`
}

type ClusterMember struct {
	// Info about the user or team
	// +optional
	Info EntityInfo `json:"info,omitempty"`

	// The account of the member in the cluster
	// +optional
	Account *configv1alpha1.Account `json:"account,omitempty"`
}

type EntityInfo struct {
	// Name is the kubernetes name of the object
	Name string `json:"name,omitempty"`

	// The display name shown in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// The username that is used to login
	// +optional
	Username string `json:"username,omitempty"`

	// The users email address
	// +optional
	Email string `json:"email,omitempty"`

	// The user subject
	// +optional
	Subject string `json:"subject,omitempty"`
}
