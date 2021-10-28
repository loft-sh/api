package v1

import (
	clusterv1 "github.com/loft-sh/agentapi/pkg/apis/loft/cluster/v1"
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalSpaceConstraint holds the globalSpaceConstraint information
// +k8s:openapi-gen=true
// +resource:path=globalspaceconstraints,rest=GlobalSpaceConstraintREST
type GlobalSpaceConstraint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalSpaceConstraintSpec   `json:"spec,omitempty"`
	Status GlobalSpaceConstraintStatus `json:"status,omitempty"`
}

// GlobalSpaceConstraintSpec holds the specification
type GlobalSpaceConstraintSpec struct {
	storagev1.GlobalSpaceConstraintSpec `json:",inline"`
}

// GlobalSpaceConstraintStatus holds the status
type GlobalSpaceConstraintStatus struct {
	storagev1.GlobalSpaceConstraintStatus `json:",inline"`

	// +optional
	ClusterRole *clusterv1.EntityInfo `json:"clusterRole,omitempty"`

	// +optional
	Clusters []*clusterv1.EntityInfo `json:"clusters,omitempty"`
}

func (a *GlobalSpaceConstraint) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *GlobalSpaceConstraint) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *GlobalSpaceConstraint) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *GlobalSpaceConstraint) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}
