package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualClusterTemplate holds the information
// +k8s:openapi-gen=true
// +resource:path=virtualclustertemplates,rest=VirtualClusterTemplateREST
type VirtualClusterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualClusterTemplateSpec   `json:"spec,omitempty"`
	Status VirtualClusterTemplateStatus `json:"status,omitempty"`
}

// VirtualClusterTemplateSpec holds the specification
type VirtualClusterTemplateSpec struct {
	storagev1.VirtualClusterTemplateSpec `json:",inline"`
}

// VirtualClusterTemplateStatus holds the status
type VirtualClusterTemplateStatus struct {
	storagev1.VirtualClusterTemplateStatus `json:",inline"`
}
