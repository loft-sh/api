package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type UserProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The new display name shown in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	// Username is the new username of the user
	// +optional
	Username string `json:"username,omitempty"`
	// Password is the new password of the user
	// +optional
	Password string `json:"password,omitempty"`
	// Email is the new email of the user
	// +optional
	Email string `json:"email,omitempty"`
	// Custom is custom information that should be saved of the user
	// +optional
	Custom string `json:"custom,omitempty"`
}
