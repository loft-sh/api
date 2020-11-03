package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualCluster holds the information
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type VirtualCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualClusterSpec   `json:"spec,omitempty"`
	Status VirtualClusterStatus `json:"status,omitempty"`
}

type VirtualClusterSpec struct {
	// The helm release configuration for the virtual cluster. This is optional, but
	// when filled, loft will deploy the specified chart for the given
	// +optional
	HelmRelease *VirtualClusterHelmRelease `json:"helmRelease,omitempty"`

	// A label selector to select the virtual cluster pod to route
	// incoming requests to.
	// +optional
	Pod *PodSelector `json:"pod,omitempty"`

	// A reference to the cluster admin kube config. This is needed for
	// the cli & ui to access the virtual clusters
	// +optional
	KubeConfigRef *SecretRef `json:"kubeConfigRef,omitempty"`
}

type VirtualClusterHelmRelease struct {
	// infos about what chart to deploy
	// +optional
	Chart VirtualClusterHelmChart `json:"chart,omitempty"`

	// the values for the given chart
	// +optional
	Values string `json:"values,omitempty"`
}

type VirtualClusterHelmChart struct {
	// the name of the helm chart
	// +optional
	Name string `json:"name,omitempty"`

	// the repo of the helm chart
	// +optional
	Repo string `json:"repo,omitempty"`

	// the version of the helm chart to use
	// +optional
	Version string `json:"version,omitempty"`
}

type PodSelector struct {
	// A label selector to select the virtual cluster pod to route
	// incoming requests to.
	// +optional
	Selector metav1.LabelSelector `json:"podSelector,omitempty"`

	// The port of the pod to route to
	// +optional
	Port *int `json:"port,omitempty"`
}

// VirtualClusterStatus holds the status of a virtual cluster
type VirtualClusterStatus struct {
	// the status of the helm release that was used to deploy the virtual cluster
	// +optional
	HelmRelease *VirtualClusterHelmReleaseStatus `json:"helmRelease,omitempty"`
}

type VirtualClusterHelmReleaseStatus struct {
	// +optional
	Phase VirtualClusterHelmReleaseStatusPhase `json:"phase,omitempty"`

	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`

	// +optional
	Reason string `json:"reason,omitempty"`

	// +optional
	Message string `json:"message,omitempty"`

	// the release that was deployed
	// +optional
	Release VirtualClusterHelmRelease `json:"release,omitempty"`
}

// VirtualClusterHelmReleaseStatusPhase describes the phase of a virtual cluster helm release
type VirtualClusterHelmReleaseStatusPhase string

// These are the valid admin account types
const (
	VirtualClusterHelmReleaseStatusNotDeployed VirtualClusterHelmReleaseStatusPhase = ""
	VirtualClusterHelmReleaseStatusDeployed    VirtualClusterHelmReleaseStatusPhase = "Deployed"
	VirtualClusterHelmReleaseStatusFailed      VirtualClusterHelmReleaseStatusPhase = "Failed"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualClusterList contains a list of User
type VirtualClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualCluster{}, &VirtualClusterList{})
}
