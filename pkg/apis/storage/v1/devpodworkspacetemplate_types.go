package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevPodWorkspaceTemplate holds the DevPodWorkspaceTemplate information
// +k8s:openapi-gen=true
type DevPodWorkspaceTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevPodWorkspaceTemplateSpec   `json:"spec,omitempty"`
	Status DevPodWorkspaceTemplateStatus `json:"status,omitempty"`
}

func (a *DevPodWorkspaceTemplate) GetVersions() []VersionAccessor {
	var retVersions []VersionAccessor
	for _, v := range a.Spec.Versions {
		b := v
		retVersions = append(retVersions, &b)
	}

	return retVersions
}

func (a *DevPodWorkspaceTemplateVersion) GetVersion() string {
	return a.Version
}

func (a *DevPodWorkspaceTemplate) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *DevPodWorkspaceTemplate) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *DevPodWorkspaceTemplate) GetAccess() []Access {
	return a.Spec.Access
}

func (a *DevPodWorkspaceTemplate) SetAccess(access []Access) {
	a.Spec.Access = access
}

// DevPodWorkspaceTemplateSpec holds the specification
type DevPodWorkspaceTemplateSpec struct {
	// DisplayName is the name that is shown in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes the virtual cluster template
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Parameters define additional app parameters that will set provider values
	// +optional
	Parameters []AppParameter `json:"parameters,omitempty"`

	// Provider holds the DevPod machine template
	Template DevPodWorkspaceTemplateDefinition `json:"template"`

	// Versions are different versions of the template that can be referenced as well
	// +optional
	Versions []DevPodWorkspaceTemplateVersion `json:"versions,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`
}

type DevPodWorkspaceTemplateDefinition struct {
	// Provider holds the DevPod provider configuration
	Provider DevPodWorkspaceProvider `json:"provider"`
}

type DevPodWorkspaceProvider struct {
	// Name is the name of the provider
	Name string `json:"name"`

	// Version is the provider version
	// +optional
	Version string `json:"version,omitempty"`

	// Source is the source the provider will be loaded from
	// +optional
	Source DevPodProviderSource `json:"source,omitempty"`

	// Options are the provider option values
	// +optional
	Options map[string]string `json:"options,omitempty"`
}

type DevPodProviderSource struct {
	// Github source for the provider
	Github string `json:"github,omitempty"`

	// File source for the provider
	File string `json:"file,omitempty"`

	// URL where the provider was downloaded from
	URL string `json:"url,omitempty"`
}

type DevPodWorkspaceTemplateVersion struct {
	// Template holds the DevPod template
	// +optional
	Template DevPodWorkspaceTemplateDefinition `json:"template,omitempty"`

	// Parameters define additional app parameters that will set provider values
	// +optional
	Parameters []AppParameter `json:"parameters,omitempty"`

	// Version is the version. Needs to be in X.X.X format.
	// +optional
	Version string `json:"version,omitempty"`
}

// DevPodWorkspaceTemplateStatus holds the status
type DevPodWorkspaceTemplateStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevPodWorkspaceTemplateList contains a list of DevPodWorkspaceTemplate
type DevPodWorkspaceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodWorkspaceTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DevPodWorkspaceTemplate{}, &DevPodWorkspaceTemplateList{})
}
