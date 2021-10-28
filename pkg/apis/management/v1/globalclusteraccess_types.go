package v1

import (
	clusterv1 "github.com/loft-sh/agentapi/pkg/apis/loft/cluster/v1"
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalClusterAccess holds the globalClusterAccess information
// +k8s:openapi-gen=true
// +resource:path=globalclusteraccesses,rest=GlobalClusterAccessREST
type GlobalClusterAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalClusterAccessSpec   `json:"spec,omitempty"`
	Status GlobalClusterAccessStatus `json:"status,omitempty"`
}

// GlobalClusterAccessSpec holds the specification
type GlobalClusterAccessSpec struct {
	storagev1.GlobalClusterAccessSpec `json:",inline"`
}

// GlobalClusterAccessStatus holds the status
type GlobalClusterAccessStatus struct {
	storagev1.GlobalClusterAccessStatus `json:",inline"`

	// +optional
	Clusters []*clusterv1.EntityInfo `json:"clusters,omitempty"`

	// +optional
	Users []*clusterv1.UserOrTeam `json:"users,omitempty"`

	// +optional
	Teams []*clusterv1.EntityInfo `json:"teams,omitempty"`

	// +optional
	SpaceConstraint *clusterv1.EntityInfo `json:"spaceConstraint,omitempty"`
}

func (a *GlobalClusterAccess) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *GlobalClusterAccess) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *GlobalClusterAccess) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *GlobalClusterAccess) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}
