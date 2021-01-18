package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +k8s:openapi-gen=true
// +resource:path=helmreleases,rest=HelmReleaseREST
type HelmRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmReleaseSpec   `json:"spec,omitempty"`
	Status HelmReleaseStatus `json:"status,omitempty"`
}

type HelmReleaseSpec struct {
	// Chart holds information about a chart
	Chart storagev1.Chart `json:"chart,omitempty"`

	// Config is the set of extra Values added to the chart.
	// These values override the default values inside of the chart.
	// +optional
	Config string `json:"config,omitempty"`

	// If tls certificate checks for the chart download should be skipped
	// +optional
	InsecureSkipTlsVerify bool `json:"insecureSkipTlsVerify,omitempty"`
}

type HelmReleaseStatus struct {
	// Revision is an int which represents the revision of the release.
	Revision int `json:"version,omitempty"`

	// Info provides information about a release
	// +optional
	Info *storagev1.Info `json:"info,omitempty"`

	// Metadata provides information about a chart
	// +optional
	Metadata *storagev1.Metadata `json:"metadata,omitempty"`
}
