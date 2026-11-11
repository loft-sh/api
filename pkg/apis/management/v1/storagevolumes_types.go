package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/v5/pkg/apis/loft/storage/v1"
	storagev1 "github.com/loft-sh/api/v5/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:method=GetCredentials,verb=get,subresource=credentials,result=github.com/loft-sh/api/v5/pkg/apis/management/v1.StorageVolumeCredentials
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Profile",type="string",JSONPath=".spec.profile"
// +kubebuilder:printcolumn:name="Size",type="string",JSONPath=".spec.size"
// +kubebuilder:subresource:status

// StorageVolume holds storage of one tier assigned to one tenant.
// +k8s:openapi-gen=true
// +resource:path=storagevolumes,rest=StorageVolumeREST,statusRest=StorageVolumeStatusREST
// +subresource:request=StorageVolumeCredentials,path=credentials,kind=StorageVolumeCredentials,rest=StorageVolumeCredentialsREST
type StorageVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageVolumeSpec   `json:"spec,omitempty"`
	Status StorageVolumeStatus `json:"status,omitempty"`
}

// StorageVolumeSpec defines spec of a storage volume.
type StorageVolumeSpec struct {
	storagev1.StorageVolumeSpec `json:",inline"`
}

type StorageVolumeStatus struct {
	storagev1.StorageVolumeStatus `json:",inline"`
}

func (a *StorageVolume) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *StorageVolume) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *StorageVolume) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *StorageVolume) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}

func (a *StorageVolume) GetConditions() agentstoragev1.Conditions {
	return a.Status.Conditions
}

func (a *StorageVolume) SetConditions(conditions agentstoragev1.Conditions) {
	a.Status.Conditions = conditions
}
