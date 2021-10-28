package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalClusterRoleTemplate holds the global role template information
// +k8s:openapi-gen=true
type GlobalClusterRoleTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalClusterRoleTemplateSpec   `json:"spec,omitempty"`
	Status GlobalClusterRoleTemplateStatus `json:"status,omitempty"`
}

func (a *GlobalClusterRoleTemplate) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *GlobalClusterRoleTemplate) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *GlobalClusterRoleTemplate) GetAccess() []Access {
	return a.Spec.Access
}

func (a *GlobalClusterRoleTemplate) SetAccess(access []Access) {
	a.Spec.Access = access
}

type GlobalClusterRoleTemplateSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes a cluster role template object
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Clusters are the clusters this template should be applied on.
	// +optional
	Clusters []string `json:"clusters,omitempty"`

	// Management defines if this cluster role should be created in the management instance.
	// +optional
	Management bool `json:"management,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`

	// ClusterRoleTemplate holds the cluster role template
	// +omitempty
	ClusterRoleTemplate ClusterRoleTemplate `json:"clusterRoleTemplate,omitempty"`
}

type ClusterRoleTemplate struct {
	// Metadata is the metadata of the cluster role template object
	// +optional
	Metadata metav1.ObjectMeta `json:"metadata,omitempty"`

	// ClusterRoleTemplateSpec holds the spec of the cluster role template in the cluster
	// +optional
	ClusterRoleTemplateSpec agentstoragev1.ClusterRoleTemplateSpec `json:"spec,omitempty"`
}

// GlobalClusterRoleTemplateStatus holds the status of a user access
type GlobalClusterRoleTemplateStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalClusterRoleTemplateList contains a list of GlobalClusterRoleTemplate objects
type GlobalClusterRoleTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalClusterRoleTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GlobalClusterRoleTemplate{}, &GlobalClusterRoleTemplateList{})
}
