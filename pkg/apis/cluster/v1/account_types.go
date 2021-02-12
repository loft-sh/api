package v1

import (
	configv1alpha1 "github.com/loft-sh/kiosk/pkg/apis/config/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User holds the user information
// +k8s:openapi-gen=true
// +resource:path=accounts,rest=AccountREST
// +subresource:request=AccountClusterRoles,path=clusterroles,kind=AccountClusterRoles,rest=AccountClusterRolesREST
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccountSpec   `json:"spec,omitempty"`
	Status AccountStatus `json:"status,omitempty"`
}

type AccountSpec struct {
	configv1alpha1.AccountSpec `json:",inline"`
}

type AccountStatus struct {
	configv1alpha1.AccountStatus `json:",inline"`
}
