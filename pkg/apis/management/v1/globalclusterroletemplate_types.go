package v1

import (
	clusterv1 "github.com/loft-sh/agentapi/pkg/apis/loft/cluster/v1"
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalClusterRoleTemplate holds the globalClusterRoleTemplate information
// +k8s:openapi-gen=true
// +resource:path=globalclusterroletemplates,rest=GlobalClusterRoleTemplateREST
type GlobalClusterRoleTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalClusterRoleTemplateSpec   `json:"spec,omitempty"`
	Status GlobalClusterRoleTemplateStatus `json:"status,omitempty"`
}

// GlobalClusterRoleTemplateSpec holds the specification
type GlobalClusterRoleTemplateSpec struct {
	storagev1.GlobalClusterRoleTemplateSpec `json:",inline"`
}

// GlobalClusterRoleTemplateStatus holds the status
type GlobalClusterRoleTemplateStatus struct {
	storagev1.GlobalClusterRoleTemplateStatus `json:",inline"`

	// +optional
	Clusters []*clusterv1.EntityInfo `json:"clusters,omitempty"`
}

func (a *GlobalClusterRoleTemplate) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *GlobalClusterRoleTemplate) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *GlobalClusterRoleTemplate) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *GlobalClusterRoleTemplate) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}
