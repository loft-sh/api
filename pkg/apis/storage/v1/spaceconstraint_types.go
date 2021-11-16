package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpaceConstraint holds the global space constraint information
// +k8s:openapi-gen=true
type SpaceConstraint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpaceConstraintSpec   `json:"spec,omitempty"`
	Status SpaceConstraintStatus `json:"status,omitempty"`
}

func (a *SpaceConstraint) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *SpaceConstraint) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *SpaceConstraint) GetAccess() []Access {
	return a.Spec.Access
}

func (a *SpaceConstraint) SetAccess(access []Access) {
	a.Spec.Access = access
}

type SpaceConstraintSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes a space constraint object
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Clusters are the clusters this template should be applied on.
	// +optional
	Clusters []string `json:"clusters,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`

	// LocalSpaceConstraintTemplate holds the space constraint template
	// +omitempty
	LocalSpaceConstraintTemplate LocalSpaceConstraintTemplate `json:"localSpaceConstraintTemplate,omitempty"`
}

type LocalSpaceConstraintTemplate struct {
	// Metadata is the metadata of the space constraint object
	// +optional
	Metadata metav1.ObjectMeta `json:"metadata,omitempty"`

	// LocalSpaceConstraintSpec holds the spec of the space constraint in the cluster
	// +optional
	LocalSpaceConstraintSpec agentstoragev1.LocalSpaceConstraintSpec `json:"spec,omitempty"`
}

// SpaceConstraintStatus holds the status of a user access
type SpaceConstraintStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpaceConstraintList contains a list of SpaceConstraint objects
type SpaceConstraintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpaceConstraint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpaceConstraint{}, &SpaceConstraintList{})
}
