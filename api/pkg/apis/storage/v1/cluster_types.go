package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User holds the user information
// +k8s:openapi-gen=true
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// ClusterSpec holds the user specification
type ClusterSpec struct {
	// Holds a reference to a secret that holds the kube config to access this cluster
	// +optional
	Config SecretRef `json:"config,omitempty"`

	// Local specifies if it is the local cluster that should be connected, when this is specified, config is optional
	// +optional
	Local bool `json:"local,omitempty"`

	// The namespace where the cluster components will be installed in
	// +optional
	ManagementNamespace string `json:"managementNamespace,omitempty"`

	// If specified this name is displayed in the UI instead of the metadata name
	// +optional
	DisplayName string `json:"displayName,omitempty"`
}

// ClusterStatus holds the user status
type ClusterStatus struct {
	// +optional
	Phase ClusterStatusPhase `json:"phase,omitempty"`

	// +optional
	Reason string `json:"reason,omitempty"`

	// +optional
	Message string `json:"message,omitempty"`
}

// ClusterStatusPhase describes the phase of a cluster
type ClusterStatusPhase string

// These are the valid admin account types
const (
	ClusterStatusPhaseInitializing ClusterStatusPhase = ""
	ClusterStatusPhaseInitialized  ClusterStatusPhase = "Initialized"
	ClusterStatusPhaseFailed       ClusterStatusPhase = "Failed"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
