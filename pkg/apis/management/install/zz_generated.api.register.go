// Code generated by generator. DO NOT EDIT.

package install

import (
	"github.com/loft-sh/api/pkg/apis/management"
	v1 "github.com/loft-sh/api/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

func init() {
	Install(builders.Scheme)
}

func Install(scheme *runtime.Scheme) {
	utilruntime.Must(v1.AddToScheme(scheme))
	utilruntime.Must(management.AddToScheme(scheme))
	utilruntime.Must(addKnownTypes(scheme))
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(management.SchemeGroupVersion,
		&management.Announcement{},
		&management.AnnouncementList{},
		&management.Cluster{},
		&management.ClusterList{},
		&management.ClusterMembers{},
		&management.ClusterVirtualClusterDefaults{},
		&management.ClusterConnect{},
		&management.ClusterConnectList{},
		&management.ClusterRole{},
		&management.ClusterRoleList{},
		&management.Config{},
		&management.ConfigList{},
		&management.Feature{},
		&management.FeatureList{},
		&management.Kiosk{},
		&management.KioskList{},
		&management.License{},
		&management.LicenseList{},
		&management.LicenseToken{},
		&management.LicenseTokenList{},
		&management.SelfSubjectAccessReview{},
		&management.SelfSubjectAccessReviewList{},
		&management.SubjectAccessReview{},
		&management.SubjectAccessReviewList{},
		&management.Team{},
		&management.TeamList{},
		&management.TeamClusterRoles{},
		&management.TeamClusters{},
		&management.TeamSpaces{},
		&management.User{},
		&management.UserList{},
		&management.UserAccessKeys{},
		&management.UserClusterRoles{},
		&management.UserClusters{},
		&management.UserProfile{},
		&management.UserQuotas{},
		&management.UserSpaces{},
		&management.UserTeams{},
		&management.UserVirtualClusters{},
	)
	return nil
}
