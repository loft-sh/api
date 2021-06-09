// Code generated by generator. DO NOT EDIT.

package v1

import (
	"github.com/loft-sh/api/pkg/apis/management"
	"github.com/loft-sh/apiserver/pkg/builders"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func addKnownTypes(scheme *runtime.Scheme) error {
	// TODO this will get cleaned up with the scheme types are fixed
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Announcement{},
		&AnnouncementList{},
		&Cluster{},
		&ClusterList{},
		&ClusterCharts{},
		&ClusterDomain{},
		&ClusterMembers{},
		&ClusterPredefinedApps{},
		&ClusterReset{},
		&ClusterVirtualClusterDefaults{},
		&ClusterAccountTemplate{},
		&ClusterAccountTemplateList{},
		&ClusterConnect{},
		&ClusterConnectList{},
		&ClusterRole{},
		&ClusterRoleList{},
		&Config{},
		&ConfigList{},
		&Feature{},
		&FeatureList{},
		&Kiosk{},
		&KioskList{},
		&License{},
		&LicenseList{},
		&LicenseToken{},
		&LicenseTokenList{},
		&LoftUpgrade{},
		&LoftUpgradeList{},
		&OwnedAccessKey{},
		&OwnedAccessKeyList{},
		&PolicyViolation{},
		&PolicyViolationList{},
		&ResetAccessKey{},
		&ResetAccessKeyList{},
		&Self{},
		&SelfList{},
		&SelfSubjectAccessReview{},
		&SelfSubjectAccessReviewList{},
		&SharedSecret{},
		&SharedSecretList{},
		&SubjectAccessReview{},
		&SubjectAccessReviewList{},
		&Team{},
		&TeamList{},
		&TeamClusterRoles{},
		&TeamClusters{},
		&TeamSpaces{},
		&User{},
		&UserList{},
		&UserClusterRoles{},
		&UserClusters{},
		&UserProfile{},
		&UserQuotas{},
		&UserSpaces{},
		&UserTeams{},
		&UserVirtualClusters{},
	)
	return nil
}

var (
	ApiVersion = builders.NewApiVersion("management.loft.sh", "v1").WithResources(
		management.ManagementAnnouncementStorage,
		management.ManagementClusterStorage,
		builders.NewApiResourceWithStorage(
			management.InternalClusterChartsREST,
			func() runtime.Object { return &ClusterCharts{} }, // Register versioned resource
			nil,
			management.NewClusterChartsREST),
		builders.NewApiResourceWithStorage(
			management.InternalClusterDomainREST,
			func() runtime.Object { return &ClusterDomain{} }, // Register versioned resource
			nil,
			management.NewClusterDomainREST),
		builders.NewApiResourceWithStorage(
			management.InternalClusterMembersREST,
			func() runtime.Object { return &ClusterMembers{} }, // Register versioned resource
			nil,
			management.NewClusterMembersREST),
		builders.NewApiResourceWithStorage(
			management.InternalClusterPredefinedAppsREST,
			func() runtime.Object { return &ClusterPredefinedApps{} }, // Register versioned resource
			nil,
			management.NewClusterPredefinedAppsREST),
		builders.NewApiResourceWithStorage(
			management.InternalClusterResetREST,
			func() runtime.Object { return &ClusterReset{} }, // Register versioned resource
			nil,
			management.NewClusterResetREST),
		builders.NewApiResourceWithStorage(
			management.InternalClusterVirtualClusterDefaultsREST,
			func() runtime.Object { return &ClusterVirtualClusterDefaults{} }, // Register versioned resource
			nil,
			management.NewClusterVirtualClusterDefaultsREST),
		management.ManagementClusterAccountTemplateStorage,
		management.ManagementClusterConnectStorage,
		management.ManagementClusterRoleStorage,
		management.ManagementConfigStorage,
		management.ManagementFeatureStorage,
		management.ManagementKioskStorage,
		management.ManagementLicenseStorage,
		management.ManagementLicenseTokenStorage,
		management.ManagementLoftUpgradeStorage,
		management.ManagementOwnedAccessKeyStorage,
		management.ManagementPolicyViolationStorage,
		management.ManagementResetAccessKeyStorage,
		management.ManagementSelfStorage,
		management.ManagementSelfSubjectAccessReviewStorage,
		management.ManagementSharedSecretStorage,
		management.ManagementSubjectAccessReviewStorage,
		management.ManagementTeamStorage,
		builders.NewApiResourceWithStorage(
			management.InternalTeamClusterRolesREST,
			func() runtime.Object { return &TeamClusterRoles{} }, // Register versioned resource
			nil,
			management.NewTeamClusterRolesREST),
		builders.NewApiResourceWithStorage(
			management.InternalTeamClustersREST,
			func() runtime.Object { return &TeamClusters{} }, // Register versioned resource
			nil,
			management.NewTeamClustersREST),
		builders.NewApiResourceWithStorage(
			management.InternalTeamSpacesREST,
			func() runtime.Object { return &TeamSpaces{} }, // Register versioned resource
			nil,
			management.NewTeamSpacesREST),
		management.ManagementUserStorage,
		builders.NewApiResourceWithStorage(
			management.InternalUserClusterRolesREST,
			func() runtime.Object { return &UserClusterRoles{} }, // Register versioned resource
			nil,
			management.NewUserClusterRolesREST),
		builders.NewApiResourceWithStorage(
			management.InternalUserClustersREST,
			func() runtime.Object { return &UserClusters{} }, // Register versioned resource
			nil,
			management.NewUserClustersREST),
		builders.NewApiResourceWithStorage(
			management.InternalUserProfileREST,
			func() runtime.Object { return &UserProfile{} }, // Register versioned resource
			nil,
			management.NewUserProfileREST),
		builders.NewApiResourceWithStorage(
			management.InternalUserQuotasREST,
			func() runtime.Object { return &UserQuotas{} }, // Register versioned resource
			nil,
			management.NewUserQuotasREST),
		builders.NewApiResourceWithStorage(
			management.InternalUserSpacesREST,
			func() runtime.Object { return &UserSpaces{} }, // Register versioned resource
			nil,
			management.NewUserSpacesREST),
		builders.NewApiResourceWithStorage(
			management.InternalUserTeamsREST,
			func() runtime.Object { return &UserTeams{} }, // Register versioned resource
			nil,
			management.NewUserTeamsREST),
		builders.NewApiResourceWithStorage(
			management.InternalUserVirtualClustersREST,
			func() runtime.Object { return &UserVirtualClusters{} }, // Register versioned resource
			nil,
			management.NewUserVirtualClustersREST),
	)

	// Required by code generated by go2idl
	AddToScheme = (&runtime.SchemeBuilder{
		ApiVersion.SchemeBuilder.AddToScheme,
		RegisterDefaults,
		RegisterConversions,
		addKnownTypes,
		func(scheme *runtime.Scheme) error {
			metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
			return nil
		},
	}).AddToScheme

	SchemeBuilder      = ApiVersion.SchemeBuilder
	localSchemeBuilder = SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AnnouncementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Announcement `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterChartsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterCharts `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterDomain `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterMembersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterMembers `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterPredefinedAppsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterPredefinedApps `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterResetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterReset `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterVirtualClusterDefaultsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterVirtualClusterDefaults `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterAccountTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAccountTemplate `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterConnectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterConnect `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRole `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FeatureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Feature `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KioskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kiosk `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []License `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LicenseTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicenseToken `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LoftUpgradeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoftUpgrade `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OwnedAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OwnedAccessKey `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PolicyViolationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyViolation `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ResetAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResetAccessKey `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SelfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Self `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SelfSubjectAccessReviewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SelfSubjectAccessReview `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SharedSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedSecret `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SubjectAccessReviewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubjectAccessReview `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Team `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TeamClusterRolesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamClusterRoles `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TeamClustersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamClusters `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TeamSpacesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamSpaces `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserClusterRolesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserClusterRoles `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserClustersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserClusters `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserProfile `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserQuotasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserQuotas `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserSpacesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserSpaces `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserTeamsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserTeams `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserVirtualClustersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserVirtualClusters `json:"items"`
}
