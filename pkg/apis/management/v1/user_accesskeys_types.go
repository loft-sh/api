package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type UserAccessKeys struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Username is the username the access keys are valid for
	// +optional
	Username string `json:"username,omitempty"`

	// Keys is an array of access keys from the user
	// +optional
	Keys []UserAccessKey `json:"keys,omitempty"`
}

// UserAccessKey defines an access key for an user
type UserAccessKey struct {
	// Name is the optional identifier of this access key
	// +optional
	Name string `json:"name,omitempty"`

	// Key is the actual access key
	Key string `json:"key,omitempty"`

	// Expires states when the access key will expire
	// +optional
	Expires int64 `json:"expires,omitempty"`

	// Timestamp when this access key was created
	// +optional
	CreationTimestamp int64 `json:"creationTimestamp,omitempty"`

	// Data holds additional access key values
	// +optional
	Data map[string]string `json:"data,omitempty"`
}
