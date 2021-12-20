package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// App holds the information
// +k8s:openapi-gen=true
// +resource:path=apps,rest=AppREST
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppSpec   `json:"spec,omitempty"`
	Status AppStatus `json:"status,omitempty"`
}

// AppSpec holds the specification
type AppSpec struct {
	storagev1.AppSpec `json:",inline"`
}

// AppStatus holds the status
type AppStatus struct {
	storagev1.AppStatus `json:",inline"`
}
