package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +k8s:openapi-gen=true
// +resource:path=helmcharts,rest=HelmChartREST
type HelmChart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmChartSpec   `json:"spec,omitempty"`
	Status HelmChartStatus `json:"status,omitempty"`
}

type HelmChartSpec struct {
	// Repository is the repository name of this chart
	// +optional
	Repository HelmChartRepository `json:"repository,omitempty"`

	// Metadata provides information about a chart
	// +optional
	Metadata *Metadata `json:"metadata,omitempty"`

	// Versions holds all chart versions
	// +optional
	Versions []string `json:"versions,omitempty"`
}

type HelmChartStatus struct {
}

type HelmChartRepository struct {
	// Name is the name of the repository
	// +optional
	Name string `json:"name,omitempty"`

	// URL is the repository url
	// +optional
	URL string `json:"url,omitempty"`
}
