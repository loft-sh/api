package v1

import (
	configv1alpha1 "github.com/loft-sh/agentapi/v2/pkg/apis/kiosk/config/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterAccountTemplate holds the information
// +k8s:openapi-gen=true
type ClusterAccountTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterAccountTemplateSpec   `json:"spec,omitempty"`
	Status ClusterAccountTemplateStatus `json:"status,omitempty"`
}

func (a *ClusterAccountTemplate) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *ClusterAccountTemplate) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *ClusterAccountTemplate) GetAccess() []Access {
	return a.Spec.Access
}

func (a *ClusterAccountTemplate) SetAccess(access []Access) {
	a.Spec.Access = access
}

type ClusterAccountTemplateSpec struct {
	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Template is the account template that will be used to create a new account
	// +optional
	Template AccountTemplate `json:"accountTemplate,omitempty"`

	// Owns are resources that will be created in the target cluster for the account
	// Furthermore the account will be added as owner reference in this resources.
	// +optional
	Owns []runtime.RawExtension `json:"owns,omitempty"`

	// Selects the target clusters this account template should be applied to. An empty
	// cluster selector selects all clusters
	// +optional
	ClusterSelector *metav1.LabelSelector `json:"clusterSelector,omitempty"`

	// Clusters are the clusters this template should be applied on.
	// +optional
	Clusters []string `json:"clusters,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`
}

type AccountTemplate struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the account template spec
	// +optional
	Spec configv1alpha1.AccountSpec `json:"spec,omitempty"`
}

type ClusterAccountTemplateStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterAccountTemplateList contains a list of ClusterAccountTemplate
type ClusterAccountTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAccountTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterAccountTemplate{}, &ClusterAccountTemplateList{})
}
