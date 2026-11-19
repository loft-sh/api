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
	// smb, s3 or block, one or several for the same data. WEKA serves posix today.
	// +optional
	Protocols []StorageProtocol `json:"protocols,omitempty"`

	// Weka is the WEKA object model for this tier. Set when the connector's provider is weka.
	// +optional
	Weka *WekaStorageSection `json:"weka,omitempty"`
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

	// CSIUser is the body of POST /users for the identity tenant clusters mount with, minus
	// username and password. The role defaults to CSI.
	// +optional
	CSIUser *runtime.RawExtension `json:"csiUser,omitempty"`

	// MountUser is the body of POST /users for the identity machines mount with, minus username and
	// password. The role defaults to Regular, which a mount token needs.
	// +optional
	MountUser *runtime.RawExtension `json:"mountUser,omitempty"`
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
	return &WekaTenantSection{Tenant: copyRaw(s.Tenant), CSIUser: copyRaw(s.CSIUser), MountUser: copyRaw(s.MountUser)}
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
	// Options are passed to mount -o, for example net=udp for a host without DPDK. They only matter
	// for the first mount on a host without a running client container.
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
