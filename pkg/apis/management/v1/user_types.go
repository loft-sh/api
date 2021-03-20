package v1

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:method=GetProfile,verb=get,subresource=profile,result=github.com/loft-sh/api/pkg/apis/management/v1.UserProfile
// +genclient:method=UpdateProfile,verb=create,subresource=profile,input=github.com/loft-sh/api/pkg/apis/management/v1.UserProfile,result=github.com/loft-sh/api/pkg/apis/management/v1.UserProfile
// +genclient:method=ListSpaces,verb=get,subresource=spaces,result=github.com/loft-sh/api/pkg/apis/management/v1.UserSpaces
// +genclient:method=ListQuotas,verb=get,subresource=quotas,result=github.com/loft-sh/api/pkg/apis/management/v1.UserQuotas
// +genclient:method=ListTeams,verb=get,subresource=teams,result=github.com/loft-sh/api/pkg/apis/management/v1.UserTeams
// +genclient:method=ListClusters,verb=get,subresource=clusters,result=github.com/loft-sh/api/pkg/apis/management/v1.UserClusters
// +genclient:method=ListVirtualClusters,verb=get,subresource=virtualclusters,result=github.com/loft-sh/api/pkg/apis/management/v1.UserVirtualClusters
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User holds the user information
// +k8s:openapi-gen=true
// +resource:path=users,rest=UserREST
// +subresource:request=UserSpaces,path=spaces,kind=UserSpaces,rest=UserSpacesREST
// +subresource:request=UserTeams,path=teams,kind=UserTeams,rest=UserTeamsREST
// +subresource:request=UserClusters,path=clusters,kind=UserClusters,rest=UserClustersREST
// +subresource:request=UserProfile,path=profile,kind=UserProfile,rest=UserProfileREST
// +subresource:request=UserClusterRoles,path=clusterroles,kind=UserClusterRoles,rest=UserClusterRolesREST
// +subresource:request=UserQuotas,path=quotas,kind=UserQuotas,rest=UserQuotasREST
// +subresource:request=UserVirtualClusters,path=virtualclusters,kind=UserVirtualClusters,rest=UserVirtualClustersREST
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserSpec   `json:"spec,omitempty"`
	Status UserStatus `json:"status,omitempty"`
}

type UserSpec struct {
	storagev1.UserSpec `json:",inline"`
}

// UserStatus holds the status of an user
type UserStatus struct {
	storagev1.UserStatus `json:",inline"`
}
