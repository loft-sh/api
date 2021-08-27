package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Self holds information about the currently logged in user
// +k8s:openapi-gen=true
// +resource:path=selves,rest=SelfREST
type Self struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SelfSpec   `json:"spec,omitempty"`
	Status SelfStatus `json:"status,omitempty"`
}

type SelfSpec struct {
}

type SelfStatus struct {
	// The name of the currently logged in user
	// +optional
	User string `json:"user,omitempty"`

	// The name of the currently logged in team
	// +optional
	Team string `json:"team,omitempty"`

	// The name of the currently used access key
	// +optional
	AccessKey string `json:"accessKey,omitempty"`

	// The type of the currently used access key
	// +optional
	AccessKeyType storagev1.AccessKeyType `json:"accessKeyType,omitempty"`

	// The subject of the currently logged in user
	// +optional
	Subject string `json:"subject,omitempty"`

	// The groups of the currently logged in user
	// +optional
	Groups []string `json:"groups,omitempty"`
}
