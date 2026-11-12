package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/v5/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// StorageProtocol is one way tenants reach a tier's data. A tier may offer several, for example
// file and object access to the same data.
type StorageProtocol string

const (
	// StorageProtocolPOSIX is the vendor's native file client, mounted on machines and served to
	// clusters through the vendor CSI driver.
	StorageProtocolPOSIX StorageProtocol = "posix"
	StorageProtocolNFS   StorageProtocol = "nfs"
	StorageProtocolSMB   StorageProtocol = "smb"
	StorageProtocolS3    StorageProtocol = "s3"
	StorageProtocolBlock StorageProtocol = "block"
)

// KnownStorageProtocols lists every protocol a profile may name. Which of them a provider serves is
// up to the provider.
var KnownStorageProtocols = []StorageProtocol{StorageProtocolPOSIX, StorageProtocolNFS, StorageProtocolSMB, StorageProtocolS3, StorageProtocolBlock}

// StorageProviderKind names the vendor behind a storage connection.
type StorageProviderKind string

const (
	StorageProviderWeka StorageProviderKind = "weka"
	StorageProviderVast StorageProviderKind = "vast"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Connector",type="string",JSONPath=".spec.connectorRef"
// +kubebuilder:printcolumn:name="Protocols",type="string",JSONPath=".spec.protocols"
// +kubebuilder:subresource:status

// StorageProfile is a tier of storage volume: which storage system it is provisioned on and the
// vendor's own object model, written as YAML and passed to the vendor API as JSON when a
// StorageVolume is created from it. The platform fills names, sizes and credentials and nothing
// else, so a new vendor setting is a new line in a profile, not a platform change.
// +k8s:openapi-gen=true
type StorageProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageProfileSpec   `json:"spec,omitempty"`
	Status StorageProfileStatus `json:"status,omitempty"`
}

func (a *StorageProfile) GetConditions() agentstoragev1.Conditions {
	return a.Status.Conditions
}

func (a *StorageProfile) SetConditions(conditions agentstoragev1.Conditions) {
	a.Status.Conditions = conditions
}

func (a *StorageProfile) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *StorageProfile) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *StorageProfile) GetAccess() []Access {
	return a.Spec.Access
}

func (a *StorageProfile) SetAccess(access []Access) {
	a.Spec.Access = access
}

// StorageProfileSpec defines a storage tier.
type StorageProfileSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description is a short summary of the tier, shown next to the display name in the admin
	// profile picker and list. Keep the storage vendor out of it, as out of the profile name.
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`

	// ConnectorRef names the storage Connector (loft.sh/connector-type: storage) the tier is
	// provisioned on. Immutable.
	ConnectorRef string `json:"connectorRef"`

	// Protocols is how tenants reach the tier's data: posix (the vendor's native file client), nfs,
	// smb, s3 or block, one or several for the same data. WEKA serves posix and s3, VAST nfs. s3 is
	// only valid next to the file protocol: it adds a bucket over the storage's volume.
	// +optional
	Protocols []StorageProtocol `json:"protocols,omitempty"`

	// Weka is the WEKA object model for this tier. Set when the connector's provider is weka.
	// +optional
	Weka *WekaStorageSection `json:"weka,omitempty"`

	// Vast is the VAST object model for this tier. Set when the connector's provider is vast.
	// +optional
	Vast *VastStorageSection `json:"vast,omitempty"`
}

// VastStorageSection is VAST's object model as YAML, as a profile carries it. Each document is sent
// as the JSON body of the matching VMS create call after the platform fills the keys it owns
// (names, tenant_id, policy_id, the path, the quota size and credentials); unknown keys pass through
// untouched. Capacity keys may use the template "{{ .storage.size }}" and may be written as
// Kubernetes quantities. At creation the tenant-level documents are copied into the tenant's
// StorageAccount and the volume-level ones into the StorageVolume.
type VastStorageSection struct {
	VastTenantSection `json:",inline"`
	VastVolumeSection `json:",inline"`
}

// VastTenantSection holds the documents that describe the vendor tenant, its VIP pool and its
// identities, one set per platform tenant on a VAST cluster.
type VastTenantSection struct {
	// Tenant is the body of POST /tenants/ minus name.
	// +optional
	Tenant *runtime.RawExtension `json:"tenant,omitempty"`

	// VIPPool is the body of POST /vippools/ for the tenant's pool minus name, tenant_id,
	// ip_ranges and subnet_cidr, which the platform allocates from the connector's VIP range. The
	// role defaults to PROTOCOLS.
	// +optional
	VIPPool *runtime.RawExtension `json:"vipPool,omitempty"`

	// CSIRole is the body of POST /roles/ for the tenant-scoped role the platform and the CSI
	// managers hold, minus name and tenant_id. permissions_list defaults to view_logical,
	// create_logical, edit_logical and delete_logical.
	// +optional
	CSIRole *runtime.RawExtension `json:"csiRole,omitempty"`

	// CSIManager is the body of POST /managers/ for the identity tenant clusters use, minus
	// username, password, roles and tenant_id. user_type defaults to TENANT_ADMIN, without which
	// VAST refuses a manager in a tenant.
	// +optional
	CSIManager *runtime.RawExtension `json:"csiManager,omitempty"`
}

// VastVolumeSection holds the documents that describe one volume, a view with its policy and
// quota on VAST.
type VastVolumeSection struct {
	// ViewPolicy is the body of POST /viewpolicies/ minus name and tenant_id. flavor defaults to
	// NFS. Listing hosts in nfs_no_squash without setting nfs_root_squash clears nfs_root_squash,
	// whose default of all hosts would otherwise win.
	// +optional
	ViewPolicy *runtime.RawExtension `json:"viewPolicy,omitempty"`

	// View is the body of POST /views/ minus name, path, policy_id and tenant_id. The path is
	// /<tenant>/<storage>; protocols defaults to NFS and create_dir to true.
	// +optional
	View *runtime.RawExtension `json:"view,omitempty"`

	// Quota is the body of POST /quotas/ minus name, path and tenant_id. hard_limit defaults to
	// the storage size.
	// +optional
	Quota *runtime.RawExtension `json:"quota,omitempty"`

	// Mount configures the NFS mounts machines perform at boot. Platform side, not sent to VAST.
	// +optional
	Mount *VastStorageMount `json:"mount,omitempty"`
}

// TenantPart is a copy of the tenant-level documents, nil for a nil section.
func (s *VastStorageSection) TenantPart() *VastTenantSection {
	if s == nil {
		return nil
	}
	return &VastTenantSection{Tenant: copyRaw(s.Tenant), VIPPool: copyRaw(s.VIPPool), CSIRole: copyRaw(s.CSIRole), CSIManager: copyRaw(s.CSIManager)}
}

// VolumePart is a copy of the volume-level documents, nil for a nil section.
func (s *VastStorageSection) VolumePart() *VastVolumeSection {
	if s == nil {
		return nil
	}
	out := &VastVolumeSection{ViewPolicy: copyRaw(s.ViewPolicy), View: copyRaw(s.View), Quota: copyRaw(s.Quota)}
	if s.Mount != nil {
		mount := *s.Mount
		mount.Options = append([]string(nil), s.Mount.Options...)
		out.Mount = &mount
	}
	return out
}

// VastStorageMount configures the NFS mounts of machines.
type VastStorageMount struct {
	// Options replace the default NFS mount options vers=3,proto=tcp, for example vers=4.1 or
	// proto=rdma,port=20049.
	// +optional
	Options []string `json:"options,omitempty"`
}

// WekaStorageSection is WEKA's object model as YAML, as a profile carries it. Each document is sent
// as the JSON body of the matching create call after the platform fills the keys it owns (names,
// sizes and credentials); unknown keys pass through untouched. Capacity keys may use the templates
// "{{ .storage.size }}" (the StorageVolume size) and "{{ .tenant.size }}" (the sum of the tenant's
// StorageVolume sizes) and may be written with units, which WEKA parses itself. At creation the
// tenant-level documents are copied into the tenant's StorageAccount and the volume-level ones into
// the StorageVolume.
type WekaStorageSection struct {
	WekaTenantSection `json:",inline"`
	WekaVolumeSection `json:",inline"`
}

// WekaTenantSection holds the documents that describe the vendor tenant and its identities, one
// set per platform tenant on a WEKA cluster.
type WekaTenantSection struct {
	// Tenant is the body of POST /tenants minus name, username and password.
	// +optional
	Tenant *runtime.RawExtension `json:"tenant,omitempty"`

	// CSIUser is the body of POST /users for the identity tenant cluster node plugins log in with,
	// minus username and password. The role defaults to CSI, the least csi-wekafs accepts.
	// +optional
	CSIUser *runtime.RawExtension `json:"csiUser,omitempty"`

	// MountUser is the body of POST /users for the mount-only identity machines mount with, minus
	// username and password. The role defaults to Regular, which a mount token needs.
	// +optional
	MountUser *runtime.RawExtension `json:"mountUser,omitempty"`

	// S3User is the body of POST /users for the tenant's S3 identity, minus username, password,
	// role (always S3) and org. posix_uid and posix_gid, the owner of objects written over S3,
	// default to 0.
	// +optional
	S3User *runtime.RawExtension `json:"s3User,omitempty"`
}

// WekaVolumeSection holds the documents that describe one volume, a filesystem on WEKA.
type WekaVolumeSection struct {
	// Filesystem is the body of POST /fileSystems minus name.
	// +optional
	Filesystem *runtime.RawExtension `json:"filesystem,omitempty"`

	// Mount configures the wekafs mounts machines perform at boot. Platform side, not sent to WEKA.
	// +optional
	Mount *WekaStorageMount `json:"mount,omitempty"`
}

// TenantPart is a copy of the tenant-level documents, nil for a nil section.
func (s *WekaStorageSection) TenantPart() *WekaTenantSection {
	if s == nil {
		return nil
	}
	return &WekaTenantSection{Tenant: copyRaw(s.Tenant), CSIUser: copyRaw(s.CSIUser), MountUser: copyRaw(s.MountUser), S3User: copyRaw(s.S3User)}
}

// VolumePart is a copy of the volume-level documents, nil for a nil section.
func (s *WekaStorageSection) VolumePart() *WekaVolumeSection {
	if s == nil {
		return nil
	}
	out := &WekaVolumeSection{Filesystem: copyRaw(s.Filesystem)}
	if s.Mount != nil {
		mount := *s.Mount
		mount.Options = append([]string(nil), s.Mount.Options...)
		out.Mount = &mount
	}
	return out
}

func copyRaw(in *runtime.RawExtension) *runtime.RawExtension {
	if in == nil {
		return nil
	}
	return &runtime.RawExtension{Raw: append([]byte(nil), in.Raw...)}
}

// WekaStorageMount configures the mounts of machines.
type WekaStorageMount struct {
	// Options are filesystem-level options passed to mount -o, for example sync_on_close. Client-mode
	// options such as net or num_cores belong to the machine and are set on its node type.
	// +optional
	Options []string `json:"options,omitempty"`

	// SkipAgentInstall leaves a machine without the WEKA agent alone instead of installing it from a
	// backend, for images that ship the client.
	// +optional
	SkipAgentInstall bool `json:"skipAgentInstall,omitempty"`
}

// StorageProfileStatus holds what the platform observed about a profile.
type StorageProfileStatus struct {
	// Conditions holds several conditions the profile might be in
	// +optional
	Conditions agentstoragev1.Conditions `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageProfileList contains a list of StorageProfile
type StorageProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageProfile{}, &StorageProfileList{})
}
