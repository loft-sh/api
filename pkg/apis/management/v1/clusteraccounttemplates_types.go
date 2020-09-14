package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User holds the user information
// +k8s:openapi-gen=true
// +resource:path=clusteraccounttemplates,rest=ClusterAccountTemplateREST
type ClusterAccountTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterAccountTemplateSpec   `json:"spec,omitempty"`
	Status ClusterAccountTemplateStatus `json:"status,omitempty"`
}

// ClusterAccountTemplateSpec holds the specification
type ClusterAccountTemplateSpec struct {
	storagev1.ClusterAccountTemplateSpec `json:",inline"`
}

// ClusterAccountTemplateStatus holds the status
type ClusterAccountTemplateStatus struct {
	storagev1.ClusterAccountTemplateStatus `json:",inline"`
}
