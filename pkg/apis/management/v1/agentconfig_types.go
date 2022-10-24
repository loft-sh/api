package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AgentConfig holds the agent loft configuration
// +k8s:openapi-gen=true
type AgentConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AgentConfigSpec   `json:"spec,omitempty"`
	Status AgentConfigStatus `json:"status,omitempty"`
}

// AgentConfigSpec holds the specification
type AgentConfigSpec struct {
	// Cluster is the name of the cluster the agent is running in
	// +optional
	Cluster string `json:"cluster,omitempty"`

	// Instance is the instance identifier of the cluster agent
	// +optional
	Instance string `json:"instance,omitempty"`

	// Analytics config for the agent
	// +optional
	Analytics *Analytics `json:"analytics,omitempty"`

	// DefaultImageRegistry defines if we should prefix the virtual cluster image
	// +optional
	DefaultImageRegistry string `json:"defaultImageRegistry,omitempty"`

	// TokenCaCert is the certificate authority the Loft tokens will
	// be signed with
	// +optional
	TokenCaCert []byte `json:"tokenCaCert,omitempty"`

	// LoftHost defines the host for the agent's loft instance
	// +optional
	LoftHost string `json:"loftHost,omitempty"`

	// LoftAPIHost defines the host for the loft api. If empty, Loft will
	// create an ssh tunnel to the agent pod.
	// +optional
	LoftAPIHost string `json:"loftAPIHost,omitempty"`

	// LoftAPIKey defines the api key the agent should use to connect to the
	// loft api server.
	// +optional
	LoftAPIKey string `json:"loftAPIKey,omitempty"`
}

// AgentConfigStatus holds the status, which is the parsed raw config
type AgentConfigStatus struct {
}
