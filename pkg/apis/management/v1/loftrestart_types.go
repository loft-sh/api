package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LoftRestart holds the restart information
// +k8s:openapi-gen=true
// +resource:path=loftrestarts,rest=LoftRestartREST
type LoftRestart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoftRestartSpec   `json:"spec,omitempty"`
	Status LoftRestartStatus `json:"status,omitempty"`
}

type LoftRestartSpec struct {
}

type LoftRestartStatus struct {
}
