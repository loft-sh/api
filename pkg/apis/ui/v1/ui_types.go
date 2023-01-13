package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UISettings holds the loft ui configuration settings
// +k8s:openapi-gen=true
type UISettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UISettingsSpec   `json:"spec,omitempty"`
	Status UISettingsStatus `json:"status,omitempty"`
}

// UISettingsSpec holds the specification
type UISettingsSpec struct {
	// LogoURL is url pointing to the logo to use in the Loft UI, this path must be accessible from clients accessing
	// the Loft UI!
	// +optional
	LogoURL string `json:"logoURL,omitempty"`
	// LegalTemplate is a text (html) string containing the legal template to prompt to users when authenticating to
	// Loft
	// +optional
	LegalTemplate string `json:"legalTemplate,omitempty"`
	// PrimaryColor is the color value (ex: "#12345") to use as the primary color
	// +optional
	PrimaryColor string `json:"primaryColor,omitempty"`
	// SidebarColor is the color value (ex: "#12345") to use for the sidebar
	// +optional
	SidebarColor string `json:"sidebarColor,omitempty"`
	// AccentColor is the color value (ex: "#12345") to use for the accent
	// +optional
	AccentColor string `json:"accentColor,omitempty"`
}

// UISettingsStatus holds the status
type UISettingsStatus struct {
	// FeatureEnabled indicates if the UI white label feature is enabled or disabled
	FeatureEnabled bool `json:"featureEnabled"`
}
