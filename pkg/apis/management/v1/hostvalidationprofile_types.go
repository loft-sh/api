package v1

import (
	storagev1 "github.com/loft-sh/api/v5/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster

// HostValidationProfile exposes the settings a kind of host has to have to platform admins.
// +k8s:openapi-gen=true
// +resource:path=hostvalidationprofiles,rest=HostValidationProfileREST
type HostValidationProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HostValidationProfileSpec   `json:"spec,omitempty"`
	Status HostValidationProfileStatus `json:"status,omitempty"`
}

type HostValidationProfileSpec struct {
	storagev1.HostValidationProfileSpec `json:",inline"`
}

type HostValidationProfileStatus struct {
	storagev1.HostValidationProfileStatus `json:",inline"`
}
