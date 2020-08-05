package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +k8s:openapi-gen=true
// +resource:path=contexts,rest=ContextREST
type Context struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContextSpec   `json:"spec,omitempty"`
	Status ContextStatus `json:"status,omitempty"`
}

type ContextSpec struct {
}

type ContextStatus struct {
	// The current cluster the context was requested
	Cluster storagev1.Cluster `json:"cluster,omitempty"`

	// The current user the context was requested
	User storagev1.User `json:"user,omitempty"`

	// Teams the requesting user is part of
	Teams []storagev1.Team `json:"teams,omitempty"`
}
