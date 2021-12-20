package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +k8s:conversion-gen:explicit-from=net/url.Values
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserSpacesOptions struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster where to retrieve spaces from
	// +optional
	Cluster []string `json:"cluster,omitempty"`
}

// +k8s:conversion-gen:explicit-from=net/url.Values
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserVirtualClustersOptions struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster where to retrieve virtual clusters from
	// +optional
	Cluster []string `json:"cluster,omitempty"`
}

// +k8s:conversion-gen:explicit-from=net/url.Values
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserQuotasOptions struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster where to retrieve quotas from
	// +optional
	Cluster []string `json:"cluster,omitempty"`
}

func InstallOptions(scheme *runtime.Scheme) error {
	return addKnownOptionsTypes(scheme)
}

func addKnownOptionsTypes(scheme *runtime.Scheme) error {
	// TODO this will get cleaned up with the scheme types are fixed
	scheme.AddKnownTypes(SchemeGroupVersion, &UserSpacesOptions{}, &UserVirtualClustersOptions{}, &UserQuotasOptions{})
	return nil
}
