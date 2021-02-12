package v1

import (
	configv1alpha1 "github.com/loft-sh/kiosk/pkg/apis/config/v1alpha1"
	tenancyv1alpha1 "github.com/loft-sh/kiosk/pkg/apis/tenancy/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// This file is just used as a collector for kiosk objects we want to generate stuff for

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User holds the user information
// +k8s:openapi-gen=true
// +resource:path=kiosk,rest=KioskREST
type Kiosk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KioskSpec   `json:"spec,omitempty"`
	Status KioskStatus `json:"status,omitempty"`
}

type KioskSpec struct {
	// tenancy.kiosk.sh
	Space   tenancyv1alpha1.Space   `json:"space,omitempty"`
	Account tenancyv1alpha1.Account `json:"account,omitempty"`

	// config.kiosk.sh
	ConfigAccount    configv1alpha1.Account          `json:"configAccount,omitempty"`
	AccountQuota     configv1alpha1.AccountQuota     `json:"accountQuota,omitempty"`
	Template         configv1alpha1.Template         `json:"template,omitempty"`
	TemplateInstance configv1alpha1.TemplateInstance `json:"templateInstance,omitempty"`
}

type KioskStatus struct {
}
