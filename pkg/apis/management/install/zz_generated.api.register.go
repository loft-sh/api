// Code generated by generator. DO NOT EDIT.

package install

import (
	"github.com/loft-sh/api/pkg/apis/management"
	v1 "github.com/loft-sh/api/pkg/apis/management/v1"
	"github.com/loft-sh/apiserver/pkg/builders"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
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
		&management.ClusterCharts{},
		&management.ClusterDomain{},
		&management.ClusterMembers{},
		&management.ClusterPredefinedApps{},
		&management.ClusterReset{},
		&management.ClusterVirtualClusterDefaults{},
		&management.ClusterAccountTemplate{},
		&management.ClusterAccountTemplateList{},
		&management.ClusterConnect{},
		&management.ClusterConnectList{},
		&management.ClusterRole{},
		&management.ClusterRoleList{},
		&management.Config{},
		&management.ConfigList{},
		&management.DirectClusterEndpointToken{},
		&management.DirectClusterEndpointTokenList{},
		&management.Feature{},
		&management.FeatureList{},
		&management.Kiosk{},
		&management.KioskList{},
		&management.License{},
		&management.LicenseList{},
		&management.LicenseToken{},
		&management.LicenseTokenList{},
		&management.LoftUpgrade{},
		&management.LoftUpgradeList{},
		&management.OwnedAccessKey{},
		&management.OwnedAccessKeyList{},
		&management.PolicyViolation{},
		&management.PolicyViolationList{},
		&management.ResetAccessKey{},
		&management.ResetAccessKeyList{},
		&management.Self{},
		&management.SelfList{},
		&management.SelfSubjectAccessReview{},
		&management.SelfSubjectAccessReviewList{},
		&management.SharedSecret{},
		&management.SharedSecretList{},
		&management.SubjectAccessReview{},
		&management.SubjectAccessReviewList{},
		&management.Team{},
		&management.TeamList{},
		&management.TeamClusterRoles{},
		&management.TeamClusters{},
		&management.TeamSpaces{},
		&management.User{},
		&management.UserList{},
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
