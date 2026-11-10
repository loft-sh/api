package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/v5/pkg/apis/loft/storage/v1"
	storagev1 "github.com/loft-sh/api/v5/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:method=GetCredentials,verb=get,subresource=credentials,result=github.com/loft-sh/api/v5/pkg/apis/management/v1.TenantStorageCredentials
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Profile",type="string",JSONPath=".spec.profile"
// +kubebuilder:printcolumn:name="Size",type="string",JSONPath=".spec.size"
// +kubebuilder:subresource:status

// TenantStorage holds storage of one tier assigned to one tenant.
// +k8s:openapi-gen=true
// +resource:path=tenantstorages,rest=TenantStorageREST,statusRest=TenantStorageStatusREST
// +subresource:request=TenantStorageCredentials,path=credentials,kind=TenantStorageCredentials,rest=TenantStorageCredentialsREST
type TenantStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TenantStorageSpec   `json:"spec,omitempty"`
	Status TenantStorageStatus `json:"status,omitempty"`
}

// TenantStorageSpec defines spec of a tenant storage.
type TenantStorageSpec struct {
	storagev1.TenantStorageSpec `json:",inline"`
}

type TenantStorageStatus struct {
	storagev1.TenantStorageStatus `json:",inline"`
}

func (a *TenantStorage) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *TenantStorage) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *TenantStorage) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *TenantStorage) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}

func (a *TenantStorage) GetConditions() agentstoragev1.Conditions {
	return a.Status.Conditions
}

func (a *TenantStorage) SetConditions(conditions agentstoragev1.Conditions) {
	a.Status.Conditions = conditions
}
