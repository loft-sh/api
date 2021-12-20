package management

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserSpacesOptions struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster where to retrieve spaces from
	// +optional
	Cluster []string `json:"cluster,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserVirtualClustersOptions struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster where to retrieve virtual clusters from
	// +optional
	Cluster []string `json:"cluster,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserQuotasOptions struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster where to retrieve quotas from
	// +optional
	Cluster []string `json:"cluster,omitempty"`
}
