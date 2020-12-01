package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SharedSecret holds the secret information
// +k8s:openapi-gen=true
type SharedSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SharedSecretSpec   `json:"spec,omitempty"`
	Status SharedSecretStatus `json:"status,omitempty"`
}

// SharedSecretSpec holds the specification
type SharedSecretSpec struct {
	// Data contains the secret data. Each key must consist of alphanumeric
	// characters, '-', '_' or '.'. The serialized form of the secret data is a
	// base64 encoded string, representing the arbitrary (possibly non-string)
	// data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
	// +optional
	Data map[string][]byte `json:"data,omitempty"`

	// Access holds the access rights for users and teams which will be transformed
	// to Roles and RoleBindings
	// +optional
	Access []SharedSecretAccess `json:"access,omitempty"`
}

// SharedSecretAccess describes the access to a secret
type SharedSecretAccess struct {
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. VerbAll represents all kinds.
	Verbs []string `json:"verbs"`

	// Users specifies which users should be able to access this secret with the aforementioned verbs
	// +optional
	Users []string `json:"users,omitempty"`

	// Teams specifies which teams should be able to access this secret with the aforementioned verbs
	// +optional
	Teams []string `json:"teams,omitempty"`
}

// SharedSecretStatus holds the status
type SharedSecretStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SharedSecretList contains a list of SharedSecret
type SharedSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SharedSecret{}, &SharedSecretList{})
}
