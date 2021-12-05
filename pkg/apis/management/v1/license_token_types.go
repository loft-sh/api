package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// License Token holds the license token information
// +k8s:openapi-gen=true
// +resource:path=licensetokens,rest=LicenseTokenREST
type LicenseToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LicenseTokenSpec   `json:"spec,omitempty"`
	Status LicenseTokenStatus `json:"status,omitempty"`
}

type LicenseTokenSpec struct {
}

type LicenseTokenStatus struct {
	// Token is the generated token
	// +optional
	Token string `json:"token,omitempty"`
}
