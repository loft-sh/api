package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// App holds the app information
// +k8s:openapi-gen=true
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppSpec   `json:"spec,omitempty"`
	Status AppStatus `json:"status,omitempty"`
}

func (a *App) GetAccess() []Access {
	return a.Spec.Access
}

func (a *App) SetAccess(access []Access) {
	a.Spec.Access = access
}

// AppSpec holds the specification
type AppSpec struct {
	// Description describes an app
	// +optional
	Description string `json:"description,omitempty"`

	// Icon holds an URL to the app icon
	// +optional
	Icon string `json:"icon,omitempty"`

	// RecommendedApp specifies where this app should show up as recommended app
	// +optional
	RecommendedApp []RecommendedApp `json:"recommendedApp,omitempty"`

	// manifest represents kubernetes resources that will be deployed into the target namespace
	// +optional
	Manifests string `json:"manifests,omitempty"`

	// helm defines the configuration for a helm deployment
	// +optional
	Helm *HelmConfiguration `json:"helm,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`
}

// HelmConfiguration holds the helm configuration
type HelmConfiguration struct {
	// Name of the chart to deploy
	Name string `json:"name"`

	// The additional helm values to use. Expected block string
	// +optional
	Values string `json:"values,omitempty"`

	// Version is the version of the chart to deploy
	// +optional
	Version string `json:"version,omitempty"`

	// The repo url to use
	// +optional
	RepoURL string `json:"repoUrl,omitempty"`

	// The username to use for the selected repository
	// +optional
	Username string `json:"username,omitempty"`

	// The password to use for the selected repository
	// +optional
	Password string `json:"password,omitempty"`

	// Determines if the remote location uses an insecure
	// TLS certificate.
	// +optional
	Insecure bool `json:"insecure,omitempty"`
}

// AppStatus holds the status
type AppStatus struct {
}

// RecommendedApp describes where an app can be displayed as recommended app
type RecommendedApp string

// Describe the status of a release
// NOTE: Make sure to update cmd/helm/status.go when adding or modifying any of these statuses.
const (
	// RecommendedAppCluster indicates that an app should be displayed as recommended app in the cluster view
	RecommendedAppCluster RecommendedApp = "cluster"
	// RecommendedAppSpace indicates that an app should be displayed as recommended app in the space view
	RecommendedAppSpace RecommendedApp = "space"
	// RecommendedAppVirtualCluster indicates that an app should be displayed as recommended app in the virtual cluster view
	RecommendedAppVirtualCluster RecommendedApp = "virtualcluster"
)

func (x RecommendedApp) String() string { return string(x) }

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppList contains a list of App
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

func init() {
	SchemeBuilder.Register(&App{}, &AppList{})
}
