package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterRoleTemplate holds the global role template information
// +k8s:openapi-gen=true
type ClusterRoleTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterRoleTemplateSpec   `json:"spec,omitempty"`
	Status ClusterRoleTemplateStatus `json:"status,omitempty"`
}

func (a *ClusterRoleTemplate) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *ClusterRoleTemplate) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *ClusterRoleTemplate) GetAccess() []Access {
	return a.Spec.Access
}

func (a *ClusterRoleTemplate) SetAccess(access []Access) {
	a.Spec.Access = access
}

type ClusterRoleTemplateSpec struct {
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

	// LocalClusterRoleTemplate holds the cluster role template
	// +omitempty
	LocalClusterRoleTemplate LocalClusterRoleTemplate `json:"localClusterRoleTemplate,omitempty"`
}

type LocalClusterRoleTemplate struct {
	// Metadata is the metadata of the cluster role template object
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	Metadata metav1.ObjectMeta `json:"metadata,omitempty"`

	// LocalClusterRoleTemplateSpec holds the spec of the cluster role template in the cluster
	// +optional
	LocalClusterRoleTemplateSpec agentstoragev1.LocalClusterRoleTemplateSpec `json:"spec,omitempty"`
}

// ClusterRoleTemplateStatus holds the status of a user access
type ClusterRoleTemplateStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterRoleTemplateList contains a list of ClusterRoleTemplate objects
type ClusterRoleTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRoleTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterRoleTemplate{}, &ClusterRoleTemplateList{})
}
