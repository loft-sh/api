package v1

import (
	storagev1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualClusterTemplate holds the virtualClusterTemplate information
// +k8s:openapi-gen=true
type VirtualClusterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualClusterTemplateSpec   `json:"spec,omitempty"`
	Status VirtualClusterTemplateStatus `json:"status,omitempty"`
}

func (a *VirtualClusterTemplate) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *VirtualClusterTemplate) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *VirtualClusterTemplate) GetAccess() []Access {
	return a.Spec.Access
}

func (a *VirtualClusterTemplate) SetAccess(access []Access) {
	a.Spec.Access = access
}

// VirtualClusterTemplateSpec holds the specification
type VirtualClusterTemplateSpec struct {
	// DisplayName is the name that is shown in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes the virtual cluster template
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Template holds the virtual cluster template
	// +optional
	Template VirtualClusterTemplateDefinition `json:"template,omitempty"`

	// SpaceTemplate to use to create the virtual cluster space if it does not exist
	// +optional
	SpaceTemplateRef *VirtualClusterTemplateSpaceTemplateRef `json:"spaceTemplateRef,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`
}

type VirtualClusterTemplateSpaceTemplateRef struct {
	// Name of the space template
	// +optional
	Name string `json:"name,omitempty"`
}

type VirtualClusterTemplateDefinition struct {
	// The virtual cluster metadata
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The virtual cluster access
	// +optional
	Access *storagev1.VirtualClusterAccess `json:"access,omitempty"`

	// The helm release configuration for the virtual cluster.
	// +optional
	HelmRelease *storagev1.VirtualClusterHelmRelease `json:"helmRelease,omitempty"`

	// Objects are Kubernetes style yamls that should get deployed into the virtual cluster
	// +optional
	Objects string `json:"objects,omitempty"`

	// Apps specifies the apps that should get deployed by this template
	// +optional
	Apps []VirtualClusterAppReference `json:"apps,omitempty"`
}

type VirtualClusterAppReference struct {
	// Name of the target app
	// +optional
	Name string `json:"name,omitempty"`

	// Namespace specifies in which target namespace the app should
	// get deployed in
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// ReleaseName is the name of the app release
	// +optional
	ReleaseName string `json:"releaseName,omitempty"`
}

// VirtualClusterTemplateStatus holds the status
type VirtualClusterTemplateStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualClusterTemplateList contains a list of VirtualClusterTemplate
type VirtualClusterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualClusterTemplate{}, &VirtualClusterTemplateList{})
}
