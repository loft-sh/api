package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HostValidationType is the protocol a group's settings are read with. The group carries a block
// of the same name saying where in that protocol the settings live.
// +enum
type HostValidationType string

const (
	// HostValidationTypeRedfish reads the settings over Redfish, from the resource named in the
	// group's redfish block.
	HostValidationTypeRedfish HostValidationType = "Redfish"
)

// HostValidationRedfishResource is a location relative to the target's Redfish system.
// +enum
type HostValidationRedfishResource string

const (
	// HostValidationRedfishResourceBios is the system's Bios resource, whose Attributes hold the
	// firmware settings and whose settings object holds the ones pending the next reboot.
	HostValidationRedfishResourceBios HostValidationRedfishResource = "Bios"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster

// HostValidationProfile names the settings a kind of host has to have. Node types select profiles
// by label, and a machine is picked for a node type only once it matches every profile that node
// type selects. It names neither machines nor a node provider: the controller resolves the backend
// per machine. Only platform admins manage it.
// +k8s:openapi-gen=true
type HostValidationProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HostValidationProfileSpec   `json:"spec,omitempty"`
	Status HostValidationProfileStatus `json:"status,omitempty"`
}

// HostValidationProfileSpec is what a host has to satisfy, grouped by where the settings live.
type HostValidationProfileSpec struct {
	// DisplayName is the name that should be displayed in the UI.
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes what the profile is intended for.
	// +optional
	Description string `json:"description,omitempty"`

	// OutOfBand holds the checks the controller reads from a BMC without running anything on the
	// machine, such as firmware settings.
	// +optional
	OutOfBand []HostValidationGroup `json:"outOfBand,omitempty"`
}

// HostValidationGroup says where a set of settings lives, and what each one has to be.
// +kubebuilder:validation:XValidation:rule="self.type != 'Redfish' || has(self.redfish)",message="a group of type Redfish has to say where its settings live: set redfish"
type HostValidationGroup struct {
	// Name identifies the group in the results. It is a DNS-1123 label, unique within the profile.
	Name string `json:"name"`

	// Target is the component whose BMC holds these settings.
	// +kubebuilder:validation:Enum=Host
	Target HostValidationTarget `json:"target"`

	// Type is the protocol the settings are read with, and names the block below that says where
	// in it they live.
	// +kubebuilder:validation:Enum=Redfish
	Type HostValidationType `json:"type"`

	// Redfish says where the settings live when Type is Redfish. Required then, refused otherwise.
	// +optional
	Redfish *HostValidationRedfish `json:"redfish,omitempty"`

	// Settings are the values this group requires. At least one.
	Settings []HostValidationSetting `json:"settings"`
}

// HostValidationRedfish locates a group's settings in Redfish.
type HostValidationRedfish struct {
	// Resource is the location relative to the target's Redfish system.
	// +kubebuilder:validation:Enum=Bios
	Resource HostValidationRedfishResource `json:"resource"`
}

// HostValidationSetting is one value a host has to have.
type HostValidationSetting struct {
	// Name identifies the setting in the results. It is a DNS-1123 label, unique within the group.
	Name string `json:"name"`

	// Attribute is the setting's name at the group's location. For a Redfish Bios group it is the
	// name from the BMC's attribute registry.
	Attribute string `json:"attribute"`

	// Value is what the attribute has to be. The controller renders the BMC's value before it
	// compares: a string as it is, a boolean as "true" or "false", a number in decimal. The
	// comparison is exact and case-sensitive.
	Value string `json:"value"`

	// Purpose says why the setting is required, for whoever reads the result.
	// +optional
	Purpose string `json:"purpose,omitempty"`
}

type HostValidationProfileStatus struct{}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HostValidationProfileList contains a list of HostValidationProfile.
type HostValidationProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostValidationProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HostValidationProfile{}, &HostValidationProfileList{})
}
