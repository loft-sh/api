package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalSpaceConstraint holds the global space constraint information
// +k8s:openapi-gen=true
type GlobalSpaceConstraint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalSpaceConstraintSpec   `json:"spec,omitempty"`
	Status GlobalSpaceConstraintStatus `json:"status,omitempty"`
}

func (a *GlobalSpaceConstraint) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *GlobalSpaceConstraint) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *GlobalSpaceConstraint) GetAccess() []Access {
	return a.Spec.Access
}

func (a *GlobalSpaceConstraint) SetAccess(access []Access) {
	a.Spec.Access = access
}

type GlobalSpaceConstraintSpec struct {
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

	// ClusterAccessTemplate holds the cluster access template
	// +omitempty
	SpaceConstraintTemplate SpaceConstraintTemplate `json:"spaceConstraintTemplate,omitempty"`
}

type SpaceConstraintTemplate struct {
	// Metadata is the metadata of the space constraint object
	// +optional
	Metadata metav1.ObjectMeta `json:"metadata,omitempty"`

	// SpaceConstraintSpec holds the spec of the space constraint in the cluster
	// +optional
	SpaceConstraintSpec agentstoragev1.SpaceConstraintSpec `json:"spec,omitempty"`
}

// GlobalSpaceConstraintStatus holds the status of a user access
type GlobalSpaceConstraintStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalSpaceConstraintList contains a list of GlobalSpaceConstraint objects
type GlobalSpaceConstraintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalSpaceConstraint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GlobalSpaceConstraint{}, &GlobalSpaceConstraintList{})
}
