package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectSecret holds the Project Secret information
// +k8s:openapi-gen=true
// +resource:path=projectsecrets,rest=ProjectSecretREST
type ProjectSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectSecretSpec   `json:"spec,omitempty"`
	Status ProjectSecretStatus `json:"status,omitempty"`
}

// ProjectSecretSpec holds the specification
type ProjectSecretSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes a Project secret
	// +optional
	Description string `json:"description,omitempty"`

	// Data contains the secret data. Each key must consist of alphanumeric
	// characters, '-', '_' or '.'. The serialized form of the secret data is a
	// base64 encoded string, representing the arbitrary (possibly non-string)
	// data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
	// +optional
	Data map[string][]byte `json:"data,omitempty"`
}

// ProjectSecretStatus holds the status
type ProjectSecretStatus struct {
}
