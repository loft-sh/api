package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:method=ListSpaces,verb=get,subresource=spaces,result=github.com/loft-sh/api/pkg/apis/management/v1.TeamSpaces
// +genclient:method=ListClusters,verb=get,subresource=clusters,result=github.com/loft-sh/api/pkg/apis/management/v1.TeamClusters
// +genclient:method=ListVirtualClusters,verb=get,subresource=virtualclusters,result=github.com/loft-sh/api/pkg/apis/management/v1.TeamVirtualClusters
// +genclient:method=ListAccessKeys,verb=get,subresource=accesskeys,result=github.com/loft-sh/api/pkg/apis/management/v1.TeamAccessKeys
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Team holds the team information
// +k8s:openapi-gen=true
// +resource:path=teams,rest=TeamREST
// +subresource:request=TeamSpaces,path=spaces,kind=TeamSpaces,rest=TeamSpacesREST
// +subresource:request=TeamClusters,path=clusters,kind=TeamClusters,rest=TeamClustersREST
// +subresource:request=TeamClusterRoles,path=clusterroles,kind=TeamClusterRoles,rest=TeamClusterRolesREST
// +subresource:request=TeamVirtualClusters,path=virtualclusters,kind=TeamVirtualClusters,rest=TeamVirtualClustersREST
// +subresource:request=TeamAccessKeys,path=accesskeys,kind=TeamAccessKeys,rest=TeamAccessKeysREST
type Team struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TeamSpec   `json:"spec,omitempty"`
	Status TeamStatus `json:"status,omitempty"`
}

type TeamSpec struct {
	storagev1.TeamSpec `json:",inline"`
}

type TeamStatus struct {
	storagev1.TeamStatus `json:",inline"`
}
