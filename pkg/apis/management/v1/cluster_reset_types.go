package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type ClusterReset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Kiosk bool `json:"kiosk,omitempty"`

	// +optional
	RBAC bool `json:"rbac,omitempty"`

	// +optional
	Templates bool `json:"templates,omitempty"`

	// +optional
	VirtualCluster bool `json:"virtualCluster,omitempty"`

	// +optional
	DirectClusterEndpoint bool `json:"directClusterEndpoint,omitempty"`
}
