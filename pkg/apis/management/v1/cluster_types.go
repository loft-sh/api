package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:method=ResetCluster,verb=create,subresource=reset,input=github.com/loft-sh/api/pkg/apis/management/v1.ClusterReset,result=github.com/loft-sh/api/pkg/apis/management/v1.ClusterReset
// +genclient:method=ListMembers,verb=get,subresource=members,result=github.com/loft-sh/api/pkg/apis/management/v1.ClusterMembers
// +genclient:method=ListVirtualClusterDefaults,verb=get,subresource=virtualclusterdefaults,result=github.com/loft-sh/api/pkg/apis/management/v1.ClusterVirtualClusterDefaults
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User holds the user information
// +k8s:openapi-gen=true
// +resource:path=clusters,rest=ClusterREST
// +subresource:request=ClusterReset,path=reset,kind=ClusterReset,rest=ClusterResetREST
// +subresource:request=ClusterMembers,path=members,kind=ClusterMembers,rest=ClusterMembersREST
// +subresource:request=ClusterVirtualClusterDefaults,path=virtualclusterdefaults,kind=ClusterVirtualClusterDefaults,rest=ClusterVirtualClusterDefaultsREST
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// ClusterSpec holds the specification
type ClusterSpec struct {
	storagev1.ClusterSpec `json:",inline"`
}

// ClusterStatus holds the status
type ClusterStatus struct {
	storagev1.ClusterStatus `json:",inline"`
}
