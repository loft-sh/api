package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/v5/pkg/apis/loft/storage/v1"
	storagev1 "github.com/loft-sh/api/v5/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Connector",type="string",JSONPath=".spec.connectorRef"
// +kubebuilder:printcolumn:name="Protocols",type="string",JSONPath=".spec.protocols"
// +kubebuilder:subresource:status

// StorageProfile holds a tier of storage volume.
// +k8s:openapi-gen=true
// +resource:path=storageprofiles,rest=StorageProfileREST
type StorageProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageProfileSpec   `json:"spec,omitempty"`
	Status StorageProfileStatus `json:"status,omitempty"`
}

// StorageProfileSpec defines spec of a storage profile.
type StorageProfileSpec struct {
	storagev1.StorageProfileSpec `json:",inline"`
}

type StorageProfileStatus struct {
	storagev1.StorageProfileStatus `json:",inline"`
}

func (a *StorageProfile) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *StorageProfile) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *StorageProfile) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *StorageProfile) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}

func (a *StorageProfile) GetConditions() agentstoragev1.Conditions {
	return a.Status.Conditions
}

func (a *StorageProfile) SetConditions(conditions agentstoragev1.Conditions) {
	a.Status.Conditions = conditions
}
