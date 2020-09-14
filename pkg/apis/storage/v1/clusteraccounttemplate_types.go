package v1

import (
	configv1alpha1 "github.com/kiosk-sh/kiosk/pkg/apis/config/v1alpha1"
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

type ClusterAccountTemplateSpec struct {
	// Template is the account template that will be used to create a new account
	// +optional
	Template AccountTemplate `json:"accountTemplate,omitempty"`

	// Owns are resources that will be created in the target cluster for the account
	// Furthermore the account will be added as owner reference in this resources.
	// +optional
	Owns []runtime.RawExtension `json:"owns,omitempty"`

	// Selects the target clusters this account template should be applied to. An empty
	// cluster selector selects all clusters
	ClusterSelector metav1.LabelSelector `json:"clusterSelector,omitempty"`
}

type AccountTemplate struct {
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
