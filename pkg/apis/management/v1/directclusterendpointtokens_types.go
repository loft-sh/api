package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DirectClusterEndpointToken holds the object information
// +k8s:openapi-gen=true
// +resource:path=directclusterendpointtokens,rest=DirectClusterEndpointTokenREST
type DirectClusterEndpointToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DirectClusterEndpointTokenSpec   `json:"spec,omitempty"`
	Status DirectClusterEndpointTokenStatus `json:"status,omitempty"`
}

// DirectClusterEndpointTokenSpec holds the object specification
type DirectClusterEndpointTokenSpec struct {
}

// DirectClusterEndpointTokenStatus holds the object status
type DirectClusterEndpointTokenStatus struct {
	// +optional
	Token string `json:"token,omitempty"`
}
