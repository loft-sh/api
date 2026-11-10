package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/v5/pkg/apis/loft/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	// StorageAccountConditionTypeProvisioned is true once the vendor tenant and its identities exist.
	StorageAccountConditionTypeProvisioned = "Provisioned"
	// StorageAccountConditionTypeSynced is true when the tenant document has been written in full;
	// false names the keys the vendor only accepts at creation and cannot change afterwards.
	StorageAccountConditionTypeSynced = "Synced"
	// StorageAccountConditionTypeKMSConfigured is true when a KMS outside the storage system wraps
	// the keys of the tenant's encrypted storages; the message names its type and scope. Storages
	// that ask for encryption without one are refused.
	StorageAccountConditionTypeKMSConfigured = "KMSConfigured"
)

var StorageAccountConditions = []agentstoragev1.ConditionType{
	StorageAccountConditionTypeProvisioned,
	StorageAccountConditionTypeSynced,
	StorageAccountConditionTypeKMSConfigured,
}

// StorageAccountPhase is where a StorageAccount is in its lifecycle.
type StorageAccountPhase string

const (
	StorageAccountPhasePending StorageAccountPhase = "Pending"
	StorageAccountPhaseReady   StorageAccountPhase = "Ready"
	StorageAccountPhaseFailed  StorageAccountPhase = "Failed"
)

// StorageAccountDeletionPolicy says what happens on the vendor side when a StorageAccount is deleted.
type StorageAccountDeletionPolicy string

const (
	// StorageAccountDeletionPolicyDelete removes the vendor tenant, its identities and the
	// credentials Secret once no TenantStorage uses the tenant any more. The default.
	StorageAccountDeletionPolicyDelete StorageAccountDeletionPolicy = "Delete"
	// StorageAccountDeletionPolicyRetain leaves the vendor tenant, its identities and the credentials
	// Secret in place, so the tenant can be imported again later.
	StorageAccountDeletionPolicyRetain StorageAccountDeletionPolicy = "Retain"
)

// StorageAccountName names the StorageAccount the platform creates for a tenant on a connector. A
// tenant name cannot contain "--", so splitting at the first "--" is unambiguous and no two tenant
// and connector pairs share a name.
func StorageAccountName(tenant, connector string) string {
	return tenant + "--" + connector
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Connector",type="string",JSONPath=".spec.connectorRef"
// +kubebuilder:printcolumn:name="Reserved",type="string",JSONPath=".status.quota.reserved"
// +kubebuilder:printcolumn:name="Storages",type="integer",JSONPath=".status.quota.storages"
// +kubebuilder:subresource:status

// StorageAccount is one platform tenant's presence on one storage system: the vendor tenant, the
// identities the platform minted for it and the quota it reserves there, which is the sum of the
// sizes of the tenant's TenantStorages on that system. The platform creates it with the tenant's
// first TenantStorage on a connector unless an admin created it first, and it is assigned to the
// tenant by the tenant.vcluster.com/exclusive-to label like a NetworkEnvironment. The vendor tenant
// lives as long as the StorageAccount does.
// +k8s:openapi-gen=true
type StorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageAccountSpec   `json:"spec,omitempty"`
	Status StorageAccountStatus `json:"status,omitempty"`
}

func (a *StorageAccount) GetConditions() agentstoragev1.Conditions {
	return a.Status.Conditions
}

func (a *StorageAccount) SetConditions(conditions agentstoragev1.Conditions) {
	a.Status.Conditions = conditions
}

func (a *StorageAccount) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *StorageAccount) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *StorageAccount) GetAccess() []Access {
	return a.Spec.Access
}

func (a *StorageAccount) SetAccess(access []Access) {
	a.Spec.Access = access
}

// StorageAccountSpec defines a tenant's presence on one storage system.
type StorageAccountSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`

	// ConnectorRef names the storage Connector (loft.sh/connector-type: storage) the tenant lives
	// on. Immutable. A platform tenant has at most one StorageAccount per connector.
	ConnectorRef string `json:"connectorRef"`

	// Profile names the StorageProfile whose tenant-level documents seed Weka or Vast at creation
	// when both are empty. Informational afterwards.
	// +optional
	Profile string `json:"profile,omitempty"`

	// Weka is this tenant's copy of the tenant-level WEKA documents, filled from the profile when
	// the StorageAccount is created and editable per customer from then on.
	// +optional
	Weka *WekaTenantSection `json:"weka,omitempty"`

	// Vast is this tenant's copy of the tenant-level VAST documents, filled from the profile when
	// the StorageAccount is created and editable per customer from then on.
	// +optional
	Vast *VastTenantSection `json:"vast,omitempty"`

	// DeletionPolicy says whether deleting the StorageAccount deletes the vendor tenant (Delete, the
	// default) or leaves it in place (Retain). A tenant the platform imported is never deleted.
	// +optional
	DeletionPolicy StorageAccountDeletionPolicy `json:"deletionPolicy,omitempty"`

	// Import adopts a vendor tenant that exists already instead of creating one. Its administrative
	// identity has to be in the credentials Secret already.
	// +optional
	Import bool `json:"import,omitempty"`
}

// StorageAccountStatus holds what the platform observed on the vendor side.
type StorageAccountStatus struct {
	// Phase is the lifecycle phase.
	// +optional
	Phase StorageAccountPhase `json:"phase,omitempty"`

	// Reason is a machine readable reason for the phase.
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message is a human readable message for the phase.
	// +optional
	Message string `json:"message,omitempty"`

	// Conditions holds several conditions the tenant might be in
	// +optional
	Conditions agentstoragev1.Conditions `json:"conditions,omitempty"`

	// Provider is the vendor behind the connector.
	// +optional
	Provider StorageProviderKind `json:"provider,omitempty"`

	// Tenant identifies the vendor tenant.
	// +optional
	Tenant *StorageAccountVendor `json:"tenant,omitempty"`

	// Applied holds the tenant document last written, platform-owned keys included except
	// credentials, so a change in the spec can be diffed against what the vendor has.
	// +optional
	Applied *runtime.RawExtension `json:"applied,omitempty"`

	// Quota is the reservation the platform keeps on the vendor tenant.
	// +optional
	Quota *StorageAccountQuota `json:"quota,omitempty"`

	// CredentialsSecret names the Secret in the platform namespace that holds the administrative,
	// CSI and mount identities of the vendor tenant.
	// +optional
	CredentialsSecret string `json:"credentialsSecret,omitempty"`

	// Identities are the usernames the platform minted; the passwords stay in the Secret.
	// +optional
	Identities *StorageAccountIdentities `json:"identities,omitempty"`
}

// StorageAccountVendor identifies the vendor tenant behind a StorageAccount.
type StorageAccountVendor struct {
	// Name is the vendor-side tenant name.
	// +optional
	Name string `json:"name,omitempty"`
	// UID is the vendor-side tenant identifier.
	// +optional
	UID string `json:"uid,omitempty"`
	// CreatedByPlatform is false for an imported tenant, which the platform never deletes.
	// +optional
	CreatedByPlatform bool `json:"createdByPlatform,omitempty"`
}

// StorageAccountQuota is the reservation the platform keeps on the vendor tenant.
type StorageAccountQuota struct {
	// Reserved is the sum of the sizes of the tenant's TenantStorages on this connector, the value
	// written to the vendor's quota keys.
	// +optional
	Reserved *resource.Quantity `json:"reserved,omitempty"`
	// Storages is how many TenantStorages the sum covers.
	// +optional
	Storages int32 `json:"storages,omitempty"`
}

// StorageAccountIdentities are the usernames of the identities the platform minted.
type StorageAccountIdentities struct {
	// Admin is the vendor tenant's administrative identity, which the platform operates with. It
	// never reaches a tenant cluster.
	// +optional
	Admin string `json:"admin,omitempty"`
	// CSI is the tenant's CSI identity, which tenant cluster node plugins log in with: the CSI
	// driver refuses a mount-only one.
	// +optional
	CSI string `json:"csi,omitempty"`
	// Mount is the mount-only identity machines mount with.
	// +optional
	Mount string `json:"mount,omitempty"`
	// S3 is the tenant's S3 identity, minted with the first storage that offers s3. Its keys stay in
	// the Secret.
	// +optional
	S3 string `json:"s3,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageAccountList contains a list of StorageAccount
type StorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccount `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageAccount{}, &StorageAccountList{})
}
