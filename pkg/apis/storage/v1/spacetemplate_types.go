package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpaceTemplate holds the space template information
// +k8s:openapi-gen=true
type SpaceTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpaceTemplateSpec   `json:"spec,omitempty"`
	Status SpaceTemplateStatus `json:"status,omitempty"`
}

func (a *SpaceTemplate) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *SpaceTemplate) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *SpaceTemplate) GetAccess() []Access {
	return a.Spec.Access
}

func (a *SpaceTemplate) SetAccess(access []Access) {
	a.Spec.Access = access
}

// SpaceTemplateSpec holds the specification
type SpaceTemplateSpec struct {
	// DisplayName is the name that is shown in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes the space template
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Template holds the space template
	// +optional
	Template SpaceTemplateDefinition `json:"template,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`
}

type SpaceTemplateDefinition struct {
	// The space metadata
	// +optional
	Metadata metav1.ObjectMeta `json:"metadata,omitempty"`

	// Apps specifies the apps that should get deployed by this template
	// +optional
	Apps []SpaceAppReference `json:"apps,omitempty"`
}

type SpaceAppReference struct {
	// Name of the target app
	// +optional
	Name string `json:"name,omitempty"`

	// ReleaseName of the target app
	// +optional
	ReleaseName string `json:"releaseName,omitempty"`
}

// SpaceTemplateStatus holds the status
type SpaceTemplateStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpaceTemplateList contains a list of SpaceTemplate
type SpaceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpaceTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpaceTemplate{}, &SpaceTemplateList{})
}
