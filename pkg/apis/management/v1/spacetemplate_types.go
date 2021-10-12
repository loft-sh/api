package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpaceTemplate holds the information
// +k8s:openapi-gen=true
// +resource:path=spacetemplates,rest=SpaceTemplateREST
type SpaceTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpaceTemplateSpec   `json:"spec,omitempty"`
	Status SpaceTemplateStatus `json:"status,omitempty"`
}

// SpaceTemplateSpec holds the specification
type SpaceTemplateSpec struct {
	storagev1.SpaceTemplateSpec `json:",inline"`
}

// SpaceTemplateStatus holds the status
type SpaceTemplateStatus struct {
	storagev1.SpaceTemplateStatus `json:",inline"`
}