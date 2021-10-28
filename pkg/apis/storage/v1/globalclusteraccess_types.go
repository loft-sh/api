package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalClusterAccess holds the global cluster access information
// +k8s:openapi-gen=true
type GlobalClusterAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalClusterAccessSpec   `json:"spec,omitempty"`
	Status GlobalClusterAccessStatus `json:"status,omitempty"`
}

func (a *GlobalClusterAccess) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *GlobalClusterAccess) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *GlobalClusterAccess) GetAccess() []Access {
	return a.Spec.Access
}

func (a *GlobalClusterAccess) SetAccess(access []Access) {
	a.Spec.Access = access
}

type GlobalClusterAccessSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes a cluster access object
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
	ClusterAccessTemplate ClusterAccessTemplate `json:"clusterAccessTemplate,omitempty"`
}

type ClusterAccessTemplate struct {
	// Metadata is the metadata of the cluster access object
	// +optional
	Metadata metav1.ObjectMeta `json:"metadata,omitempty"`

	// ClusterAccessSpec holds the spec of the cluster access in the cluster
	// +optional
	ClusterAccessSpec agentstoragev1.ClusterAccessSpec `json:"spec,omitempty"`
}

// GlobalClusterAccessStatus holds the status of a user access
type GlobalClusterAccessStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalClusterAccessList contains a list of GlobalClusterAccess objects
type GlobalClusterAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalClusterAccess `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GlobalClusterAccess{}, &GlobalClusterAccessList{})
}
