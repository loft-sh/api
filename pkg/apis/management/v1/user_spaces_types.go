package v1

import (
	clusterv1 "github.com/loft-sh/api/pkg/apis/cluster/v1"
	tenancyv1alpha1 "github.com/loft-sh/kiosk/pkg/apis/tenancy/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type UserSpaces struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spaces []ClusterSpace `json:"spaces,omitempty"`
}

type ClusterSpace struct {
	// Cluster is the cluster name of the cluster the space is in
	// +optional
	Cluster string `json:"cluster,omitempty"`

	// SleepModeConfig is the sleep mode config of the space
	// +optional
	SleepModeConfig *clusterv1.SleepModeConfig `json:"sleepModeConfig,omitempty"`

	// Owner describes the owner of the space. This can be either empty (nil), be a team or
	// an loft user. If the space has an account that does not belong to an user / team in loft
	// this is empty
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Space is the space object in the cluster
	Space tenancyv1alpha1.Space `json:"space,omitempty"`
}

type UserOrTeam struct {
	// User describes an user
	// +optional
	User *EntityInfo `json:"user,omitempty"`

	// Team describes a team
	// +optional
	Team *EntityInfo `json:"team,omitempty"`
}
