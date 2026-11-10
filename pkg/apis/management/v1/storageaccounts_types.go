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
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Connector",type="string",JSONPath=".spec.connectorRef"
// +kubebuilder:printcolumn:name="Reserved",type="string",JSONPath=".status.quota.reserved"
// +kubebuilder:subresource:status

// StorageAccount holds one tenant's presence on one storage system.
// +k8s:openapi-gen=true
// +resource:path=storageaccounts,rest=StorageAccountREST,statusRest=StorageAccountStatusREST
type StorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageAccountSpec   `json:"spec,omitempty"`
	Status StorageAccountStatus `json:"status,omitempty"`
}

// StorageAccountSpec defines spec of a storage account.
type StorageAccountSpec struct {
	storagev1.StorageAccountSpec `json:",inline"`
}

type StorageAccountStatus struct {
	storagev1.StorageAccountStatus `json:",inline"`
}

func (a *StorageAccount) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *StorageAccount) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *StorageAccount) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *StorageAccount) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}

func (a *StorageAccount) GetConditions() agentstoragev1.Conditions {
	return a.Status.Conditions
}

func (a *StorageAccount) SetConditions(conditions agentstoragev1.Conditions) {
	a.Status.Conditions = conditions
}
