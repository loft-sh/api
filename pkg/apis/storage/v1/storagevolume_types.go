package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/v5/pkg/apis/loft/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	// StorageVolumeConditionTypeProvisioned is true once the vendor tenant, its credentials and the
	// volume exist.
	StorageVolumeConditionTypeProvisioned = "Provisioned"
	// StorageVolumeConditionTypeSynced is true when the vendor section has been written in full;
	// false names the keys the vendor refused on update, so their change was not applied.
	StorageVolumeConditionTypeSynced = "Synced"
	// StorageVolumeConditionTypeObjectReady is true once the storage's bucket and the tenant's S3
	// identity exist. Only present when the profile lists s3; file access does not depend on it.
	StorageVolumeConditionTypeObjectReady = "ObjectReady"
)

var StorageVolumeConditions = []agentstoragev1.ConditionType{
	StorageVolumeConditionTypeProvisioned,
	StorageVolumeConditionTypeSynced,
	StorageVolumeConditionTypeObjectReady,
}

// StorageVolumePhase is where a StorageVolume is in its lifecycle.
type StorageVolumePhase string

const (
	StorageVolumePhasePending   StorageVolumePhase = "Pending"
	StorageVolumePhaseAvailable StorageVolumePhase = "Available"
	StorageVolumePhaseFailed    StorageVolumePhase = "Failed"
)

// StorageVolumeMountsProperty is the NodeClaim property that lists the StorageVolumes a machine
// mounts at boot as comma separated <storage>:<path> pairs, for example
// "datasets:/mnt/datasets,scratch:/mnt/scratch". The <storage>=<path> form is accepted too.
const StorageVolumeMountsProperty = "vcluster.com/storage-mounts"

// StorageVolumeClientOptionsProperty carries a machine's WEKA client-mode mount options, comma
// separated in the mount option grammar, for example "num_cores=1" or "net=udp". They set up the
// client container the host's first mount creates (its NIC, dedicated cores and pinned memory),
// which every later mount and csi-wekafs reuse, so they belong to the machine and not to the storage
// tier. Only client-mode keys are accepted; a storage's own options stay in its mount section. VAST
// mounts over NFS, which has no client container, so its mount script ignores the property.
const StorageVolumeClientOptionsProperty = "vcluster.com/storage-client-options"

const (
	// StorageVolumeMountLabel is the label a worker that mounts storage volume puts on its own
	// Node once it joined, StorageVolumeMountSucceeded or StorageVolumeMountFailed. A workload
	// that needs the mounts selects nodes labelled succeeded.
	StorageVolumeMountLabel = "storage.vcluster.com/mount"
	// StorageVolumeMountMessageAnnotation carries the mount step's last output lines on the Node.
	StorageVolumeMountMessageAnnotation = "storage.vcluster.com/mount-message"

	StorageVolumeMountSucceeded = "succeeded"
	StorageVolumeMountFailed    = "failed"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Profile",type="string",JSONPath=".spec.profile"
// +kubebuilder:printcolumn:name="Size",type="string",JSONPath=".spec.size"
// +kubebuilder:subresource:status

// StorageVolume is storage of one tier assigned to one tenant. A platform admin creates it with a
// profile and a size and assigns it with the tenant.vcluster.com/exclusive-to label, like a
// NetworkEnvironment. The first StorageVolume of a tenant on a storage system creates the tenant's
// StorageAccount there unless an admin created it first; every StorageVolume is one volume (a
// filesystem on WEKA) the tenant's clusters and machines mount.
// +k8s:openapi-gen=true
type StorageVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageVolumeSpec   `json:"spec,omitempty"`
	Status StorageVolumeStatus `json:"status,omitempty"`
}

func (a *StorageVolume) GetConditions() agentstoragev1.Conditions {
	return a.Status.Conditions
}

func (a *StorageVolume) SetConditions(conditions agentstoragev1.Conditions) {
	a.Status.Conditions = conditions
}

func (a *StorageVolume) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *StorageVolume) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *StorageVolume) GetAccess() []Access {
	return a.Spec.Access
}

func (a *StorageVolume) SetAccess(access []Access) {
	a.Spec.Access = access
}

// StorageVolumeSpec defines the storage a tenant gets.
type StorageVolumeSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`

	// Profile names the StorageProfile the storage is created from. Immutable.
	Profile string `json:"profile"`

	// Size is the capacity of the volume. Growing it grows the vendor quota first and the volume
	// after; shrinking is passed to the vendor, which refuses it when used data does not fit.
	Size resource.Quantity `json:"size"`

	// Weka is this storage's copy of the profile's volume-level WEKA documents, filled from the
	// profile when the storage is created and editable per customer from then on, so a profile
	// change only affects storages created after it. Empty means the profile's documents apply.
	// Tenant-level documents live on the tenant's StorageAccount.
	// +optional
	Weka *WekaVolumeSection `json:"weka,omitempty"`

	// Vast is this storage's copy of the profile's volume-level VAST documents, filled and edited
	// like Weka.
	// +optional
	Vast *VastVolumeSection `json:"vast,omitempty"`

	// Import adopts a volume that exists already instead of creating it, and creates the tenant's
	// StorageAccount in import mode when it does not exist yet. The platform never deletes what it
	// did not create.
	// +optional
	Import bool `json:"import,omitempty"`

	// Encrypted creates the volume encrypted at rest. On WEKA the client encrypts on the host and
	// a KMS has to wrap the key; without one the storage is refused. It wins over an encrypted key in
	// the vendor section, and the two have to agree at creation. Set at creation only: a later
	// change is reported by the Synced condition and never applied.
	// +optional
	Encrypted bool `json:"encrypted,omitempty"`
}

// StorageVolumeStatus holds what the platform observed on the vendor side.
type StorageVolumeStatus struct {
	// Phase is the lifecycle phase.
	// +optional
	Phase StorageVolumePhase `json:"phase,omitempty"`

	// Reason is a machine readable reason for the phase.
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message is a human readable message for the phase.
	// +optional
	Message string `json:"message,omitempty"`

	// Conditions holds several conditions the storage might be in
	// +optional
	Conditions agentstoragev1.Conditions `json:"conditions,omitempty"`

	// Provider is the vendor behind the storage, resolved from the profile's connector.
	// +optional
	Provider StorageProviderKind `json:"provider,omitempty"`

	// ConnectorRef is the connector the storage was provisioned on, resolved from the profile.
	// +optional
	ConnectorRef string `json:"connectorRef,omitempty"`

	// StorageAccount names the StorageAccount that owns the vendor tenant and identities the storage
	// lives in.
	// +optional
	StorageAccount string `json:"storageAccount,omitempty"`

	// Volume identifies the vendor-side objects that back the storage.
	// +optional
	Volume *StorageVolumeVendor `json:"volume,omitempty"`

	// Capacity is the last observed size, used and available capacity of the volume.
	// +optional
	Capacity *StorageVolumeCapacity `json:"capacity,omitempty"`

	// Applied holds the volume document last written, so a change in the spec can be diffed against
	// what the vendor has.
	// +optional
	Applied *StorageVolumeApplied `json:"applied,omitempty"`

	// CredentialsSecret names the Secret in the platform namespace that holds the vendor tenant's
	// administrative, CSI and mount identities. It belongs to the StorageAccount and is shared by the
	// StorageVolumes of one tenant on one storage system.
	// +optional
	CredentialsSecret string `json:"credentialsSecret,omitempty"`

	// Encrypted is whether the vendor reports the volume as encrypted at rest.
	// +optional
	Encrypted *bool `json:"encrypted,omitempty"`

	// Object is the S3 bucket over the volume, recorded once object access is ready. It stays
	// recorded for the delete when object access fails, and goes once a profile that drops s3 took
	// the bucket out of the tenant's policy; a tenant reads it only while ObjectReady is True. Only
	// for a storage whose profile lists s3.
	// +optional
	Object *StorageVolumeObject `json:"object,omitempty"`
}

// StorageVolumeObject is where a storage's data is reached over S3.
type StorageVolumeObject struct {
	// Endpoint is the S3 endpoint URL the connector names.
	// +optional
	Endpoint string `json:"endpoint,omitempty"`
	// Bucket is the bucket over the storage's volume.
	// +optional
	Bucket string `json:"bucket,omitempty"`
	// PathStyle is true when clients address the bucket in the path rather than the host name.
	// +optional
	PathStyle bool `json:"pathStyle,omitempty"`
}

// StorageVolumeVendor identifies the vendor objects behind a StorageVolume.
type StorageVolumeVendor struct {
	// Tenant is the vendor-side tenant name.
	// +optional
	Tenant string `json:"tenant,omitempty"`
	// TenantUID is the vendor-side tenant identifier.
	// +optional
	TenantUID string `json:"tenantUID,omitempty"`
	// Name is the vendor-side volume name, a filesystem on WEKA and the view path on VAST.
	// +optional
	Name string `json:"name,omitempty"`
	// UID is the vendor-side volume identifier.
	// +optional
	UID string `json:"uid,omitempty"`
	// CreatedByPlatform is false for an imported volume, which the platform never deletes.
	// +optional
	CreatedByPlatform bool `json:"createdByPlatform,omitempty"`
}

// StorageVolumeCapacity is the observed capacity of a volume.
type StorageVolumeCapacity struct {
	// +optional
	Size *resource.Quantity `json:"size,omitempty"`
	// +optional
	Used *resource.Quantity `json:"used,omitempty"`
	// +optional
	Available *resource.Quantity `json:"available,omitempty"`
}

// StorageVolumeApplied holds the vendor document last written.
type StorageVolumeApplied struct {
	// Volume is the volume document as written, platform-owned keys included.
	// +optional
	Volume *runtime.RawExtension `json:"volume,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageVolumeList contains a list of StorageVolume
type StorageVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageVolume `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageVolume{}, &StorageVolumeList{})
}
