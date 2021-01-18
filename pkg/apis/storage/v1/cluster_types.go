package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User holds the user information
// +k8s:openapi-gen=true
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// ClusterSpec holds the user specification
type ClusterSpec struct {
	// Holds a reference to a secret that holds the kube config to access this cluster
	// +optional
	Config SecretRef `json:"config,omitempty"`

	// Local specifies if it is the local cluster that should be connected, when this is specified, config is optional
	// +optional
	Local bool `json:"local,omitempty"`

	// The namespace where the cluster components will be installed in
	// +optional
	ManagementNamespace string `json:"managementNamespace,omitempty"`

	// If specified this name is displayed in the UI instead of the metadata name
	// +optional
	DisplayName string `json:"displayName,omitempty"`
}

// ClusterStatus holds the user status
type ClusterStatus struct {
	// +optional
	Phase ClusterStatusPhase `json:"phase,omitempty"`

	// +optional
	Reason string `json:"reason,omitempty"`

	// +optional
	Message string `json:"message,omitempty"`
}

// ClusterStatusPhase describes the phase of a cluster
type ClusterStatusPhase string

// These are the valid admin account types
const (
	ClusterStatusPhaseInitializing ClusterStatusPhase = ""
	ClusterStatusPhaseInitialized  ClusterStatusPhase = "Initialized"
	ClusterStatusPhaseFailed       ClusterStatusPhase = "Failed"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}

type HelmChart struct {
	// Metadata provides information about a chart
	// +optional
	Metadata Metadata `json:"metadata,omitempty"`

	// Versions holds all chart versions
	// +optional
	Versions []string `json:"versions,omitempty"`

	// Repository is the repository name of this chart
	// +optional
	Repository HelmChartRepository `json:"repository,omitempty"`
}

type HelmChartRepository struct {
	// Name is the name of the repository
	// +optional
	Name string `json:"name,omitempty"`

	// URL is the repository url
	// +optional
	URL string `json:"url,omitempty"`
}

// Chart describes a chart
type Chart struct {
	// Name is the chart name in the repository
	Name string `json:"name,omitempty"`

	// Version is the chart version in the repository
	// +optional
	Version string `json:"version,omitempty"`

	// RepoURL is the repo url where the chart can be found
	// +optional
	RepoURL string `json:"repoURL,omitempty"`

	// The username that is required for this repository
	// +optional
	Username string `json:"username,omitempty"`

	// The password that is required for this repository
	// +optional
	Password string `json:"password,omitempty"`
}

// Info describes release information.
type Info struct {
	// FirstDeployed is when the release was first deployed.
	// +optional
	FirstDeployed metav1.Time `json:"first_deployed,omitempty"`
	// LastDeployed is when the release was last deployed.
	// +optional
	LastDeployed metav1.Time `json:"last_deployed,omitempty"`
	// Deleted tracks when this object was deleted.
	// +optional
	Deleted metav1.Time `json:"deleted"`
	// Description is human-friendly "log entry" about this release.
	// +optional
	Description string `json:"description,omitempty"`
	// Status is the current state of the release
	// +optional
	Status Status `json:"status,omitempty"`
	// Contains the rendered templates/NOTES.txt if available
	// +optional
	Notes string `json:"notes,omitempty"`
}

// Status is the status of a release
type Status string

// Describe the status of a release
// NOTE: Make sure to update cmd/helm/status.go when adding or modifying any of these statuses.
const (
	// StatusUnknown indicates that a release is in an uncertain state.
	StatusUnknown Status = "unknown"
	// StatusDeployed indicates that the release has been pushed to Kubernetes.
	StatusDeployed Status = "deployed"
	// StatusUninstalled indicates that a release has been uninstalled from Kubernetes.
	StatusUninstalled Status = "uninstalled"
	// StatusSuperseded indicates that this release object is outdated and a newer one exists.
	StatusSuperseded Status = "superseded"
	// StatusFailed indicates that the release was not successfully deployed.
	StatusFailed Status = "failed"
	// StatusUninstalling indicates that a uninstall operation is underway.
	StatusUninstalling Status = "uninstalling"
	// StatusPendingInstall indicates that an install operation is underway.
	StatusPendingInstall Status = "pending-install"
	// StatusPendingUpgrade indicates that an upgrade operation is underway.
	StatusPendingUpgrade Status = "pending-upgrade"
	// StatusPendingRollback indicates that an rollback operation is underway.
	StatusPendingRollback Status = "pending-rollback"
)

func (x Status) String() string { return string(x) }

// Maintainer describes a Chart maintainer.
type Maintainer struct {
	// Name is a user name or organization name
	// +optional
	Name string `json:"name,omitempty"`
	// Email is an optional email address to contact the named maintainer
	// +optional
	Email string `json:"email,omitempty"`
	// URL is an optional URL to an address for the named maintainer
	// +optional
	URL string `json:"url,omitempty"`
}

// Metadata for a Chart file. This models the structure of a Chart.yaml file.
type Metadata struct {
	// The name of the chart
	// +optional
	Name string `json:"name,omitempty"`
	// The URL to a relevant project page, git repo, or contact person
	// +optional
	Home string `json:"home,omitempty"`
	// Source is the URL to the source code of this chart
	// +optional
	Sources []string `json:"sources,omitempty"`
	// A SemVer 2 conformant version string of the chart
	// +optional
	Version string `json:"version,omitempty"`
	// A one-sentence description of the chart
	// +optional
	Description string `json:"description,omitempty"`
	// A list of string keywords
	// +optional
	Keywords []string `json:"keywords,omitempty"`
	// A list of name and URL/email address combinations for the maintainer(s)
	// +optional
	Maintainers []*Maintainer `json:"maintainers,omitempty"`
	// The URL to an icon file.
	// +optional
	Icon string `json:"icon,omitempty"`
	// The API Version of this chart.
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`
	// The condition to check to enable chart
	// +optional
	Condition string `json:"condition,omitempty"`
	// The tags to check to enable chart
	// +optional
	Tags string `json:"tags,omitempty"`
	// The version of the application enclosed inside of this chart.
	// +optional
	AppVersion string `json:"appVersion,omitempty"`
	// Whether or not this chart is deprecated
	// +optional
	Deprecated bool `json:"deprecated,omitempty"`
	// Annotations are additional mappings uninterpreted by Helm,
	// made available for inspection by other applications.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// KubeVersion is a SemVer constraint specifying the version of Kubernetes required.
	// +optional
	KubeVersion string `json:"kubeVersion,omitempty"`
	// Specifies the chart type: application or library
	// +optional
	Type string `json:"type,omitempty"`
}
