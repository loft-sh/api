package v1

import (
	configv1alpha1 "github.com/loft-sh/kiosk/pkg/apis/config/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type UserQuotas struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Quotas []ClusterQuota `json:"quotas,omitempty"`
}

type ClusterQuota struct {
	// Cluster is the cluster name of the cluster the space is in
	// +optional
	Cluster string `json:"cluster,omitempty"`

	// Owner describes the owner of the quota. This can be either empty (nil), be a team or
	// an loft user. If the space has an account that does not belong to an user / team in loft
	// this is empty
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Quota is the quota object in the cluster
	Quota configv1alpha1.AccountQuota `json:"quota,omitempty"`
}
