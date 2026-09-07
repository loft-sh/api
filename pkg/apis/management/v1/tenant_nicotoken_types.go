package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TenantNICoToken holds the request and response for a short-lived,
// user-scoped NICo token for the tenant's onboarded org.
// +subresource-request
type TenantNICoToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status TenantNICoTokenStatus `json:"status,omitempty"`
}

type TenantNICoTokenStatus struct {
	// Token is the platform-signed NICo bearer token.
	// +optional
	Token string `json:"token,omitempty"`

	// Org is the NICo organization the token is scoped to.
	// +optional
	Org string `json:"org,omitempty"`

	// Endpoint is the NICo REST API endpoint the token is minted for.
	// +optional
	Endpoint string `json:"endpoint,omitempty"`
}
