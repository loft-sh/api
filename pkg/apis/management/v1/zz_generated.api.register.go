// Code generated by generator. DO NOT EDIT.

package v1

import (
	"github.com/loft-sh/api/v4/pkg/apis/management"
	"github.com/loft-sh/apiserver/pkg/builders"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func addKnownTypes(scheme *runtime.Scheme) error {
	// TODO this will get cleaned up with the scheme types are fixed
	scheme.AddKnownTypes(SchemeGroupVersion,
		&AgentAuditEvent{},
		&AgentAuditEventList{},
		&Announcement{},
		&AnnouncementList{},
		&App{},
		&AppList{},
		&AppCredentials{},
		&Backup{},
		&BackupList{},
		&BackupApply{},
		&Cluster{},
		&ClusterList{},
		&ClusterAccessKey{},
		&ClusterAgentConfig{},
		&ClusterCharts{},
		&ClusterDomain{},
		&ClusterMemberAccess{},
		&ClusterMembers{},
		&ClusterReset{},
		&ClusterVirtualClusterDefaults{},
		&ClusterAccess{},
		&ClusterAccessList{},
		&ClusterRoleTemplate{},
		&ClusterRoleTemplateList{},
		&Config{},
		&ConfigList{},
		&ConvertVirtualClusterConfig{},
		&ConvertVirtualClusterConfigList{},
		&DevPodEnvironmentTemplate{},
		&DevPodEnvironmentTemplateList{},
		&DevPodWorkspaceInstance{},
		&DevPodWorkspaceInstanceList{},
		&DevPodDeleteOptions{},
		&DevPodStatusOptions{},
		&DevPodSshOptions{},
		&DevPodWorkspaceInstanceState{},
		&DevPodStopOptions{},
		&DevPodWorkspaceInstanceTroubleshoot{},
		&DevPodUpOptions{},
		&DevPodWorkspacePreset{},
		&DevPodWorkspacePresetList{},
		&DevPodWorkspaceTemplate{},
		&DevPodWorkspaceTemplateList{},
		&DirectClusterEndpointToken{},
		&DirectClusterEndpointTokenList{},
		&Event{},
		&EventList{},
		&Feature{},
		&FeatureList{},
		&IngressAuthToken{},
		&IngressAuthTokenList{},
		&Kiosk{},
		&KioskList{},
		&License{},
		&LicenseList{},
		&LicenseRequest{},
		&LicenseToken{},
		&LicenseTokenList{},
		&LoftUpgrade{},
		&LoftUpgradeList{},
		&OIDCClient{},
		&OIDCClientList{},
		&OwnedAccessKey{},
		&OwnedAccessKeyList{},
		&Project{},
		&ProjectList{},
		&ProjectChartInfo{},
		&ProjectCharts{},
		&ProjectClusters{},
		&ProjectImportSpace{},
		&ProjectMembers{},
		&ProjectMigrateSpaceInstance{},
		&ProjectMigrateVirtualClusterInstance{},
		&ProjectRunners{},
		&ProjectTemplates{},
		&ProjectSecret{},
		&ProjectSecretList{},
		&RedirectToken{},
		&RedirectTokenList{},
		&RegisterVirtualCluster{},
		&RegisterVirtualClusterList{},
		&ResetAccessKey{},
		&ResetAccessKeyList{},
		&Runner{},
		&RunnerList{},
		&RunnerAccessKey{},
		&RunnerConfig{},
		&Self{},
		&SelfList{},
		&SelfSubjectAccessReview{},
		&SelfSubjectAccessReviewList{},
		&SharedSecret{},
		&SharedSecretList{},
		&SpaceInstance{},
		&SpaceInstanceList{},
		&SpaceTemplate{},
		&SpaceTemplateList{},
		&SubjectAccessReview{},
		&SubjectAccessReviewList{},
		&Task{},
		&TaskList{},
		&TaskLog{},
		&Team{},
		&TeamList{},
		&TeamAccessKeys{},
		&TeamClusters{},
		&TranslateVClusterResourceName{},
		&TranslateVClusterResourceNameList{},
		&User{},
		&UserList{},
		&UserAccessKeys{},
		&UserClusters{},
		&UserPermissions{},
		&UserProfile{},
		&VirtualClusterInstance{},
		&VirtualClusterInstanceList{},
		&VirtualClusterAccessKey{},
		&VirtualClusterExternalDatabase{},
		&VirtualClusterInstanceKubeConfig{},
		&VirtualClusterInstanceLog{},
		&VirtualClusterTemplate{},
		&VirtualClusterTemplateList{},
	)
	return nil
}

var (
	ApiVersion = builders.NewApiVersion("management.loft.sh", "v1").WithResources(
		management.ManagementAgentAuditEventStorage,
		management.ManagementAnnouncementStorage,
		management.ManagementAppStorage,
		builders.NewApiResourceWithStorage(
			management.InternalAppCredentialsREST,
			func() runtime.Object { return &AppCredentials{} }, // Register versioned resource
			nil,
			management.NewAppCredentialsREST,
		),
		management.ManagementBackupStorage,
		builders.NewApiResourceWithStorage(
			management.InternalBackupApplyREST,
			func() runtime.Object { return &BackupApply{} }, // Register versioned resource
			nil,
			management.NewBackupApplyREST,
		),
		management.ManagementClusterStorage,
		builders.NewApiResourceWithStorage(
			management.InternalClusterStatus,
			func() runtime.Object { return &Cluster{} },     // Register versioned resource
			func() runtime.Object { return &ClusterList{} }, // Register versioned resource list
			management.NewClusterStatusREST),
		builders.NewApiResourceWithStorage(
			management.InternalClusterAccessKeyREST,
			func() runtime.Object { return &ClusterAccessKey{} }, // Register versioned resource
			nil,
			management.NewClusterAccessKeyREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalClusterAgentConfigREST,
			func() runtime.Object { return &ClusterAgentConfig{} }, // Register versioned resource
			nil,
			management.NewClusterAgentConfigREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalClusterChartsREST,
			func() runtime.Object { return &ClusterCharts{} }, // Register versioned resource
			nil,
			management.NewClusterChartsREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalClusterDomainREST,
			func() runtime.Object { return &ClusterDomain{} }, // Register versioned resource
			nil,
			management.NewClusterDomainREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalClusterMemberAccessREST,
			func() runtime.Object { return &ClusterMemberAccess{} }, // Register versioned resource
			nil,
			management.NewClusterMemberAccessREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalClusterMembersREST,
			func() runtime.Object { return &ClusterMembers{} }, // Register versioned resource
			nil,
			management.NewClusterMembersREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalClusterResetREST,
			func() runtime.Object { return &ClusterReset{} }, // Register versioned resource
			nil,
			management.NewClusterResetREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalClusterVirtualClusterDefaultsREST,
			func() runtime.Object { return &ClusterVirtualClusterDefaults{} }, // Register versioned resource
			nil,
			management.NewClusterVirtualClusterDefaultsREST,
		),
		management.ManagementClusterAccessStorage,
		management.ManagementClusterRoleTemplateStorage,
		management.ManagementConfigStorage,
		management.ManagementConvertVirtualClusterConfigStorage,
		management.ManagementDevPodEnvironmentTemplateStorage,
		management.ManagementDevPodWorkspaceInstanceStorage,
		builders.NewApiResourceWithStorage(
			management.InternalDevPodDeleteOptionsREST,
			func() runtime.Object { return &DevPodDeleteOptions{} }, // Register versioned resource
			nil,
			management.NewDevPodDeleteOptionsREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalDevPodStatusOptionsREST,
			func() runtime.Object { return &DevPodStatusOptions{} }, // Register versioned resource
			nil,
			management.NewDevPodStatusOptionsREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalDevPodSshOptionsREST,
			func() runtime.Object { return &DevPodSshOptions{} }, // Register versioned resource
			nil,
			management.NewDevPodSshOptionsREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalDevPodWorkspaceInstanceStateREST,
			func() runtime.Object { return &DevPodWorkspaceInstanceState{} }, // Register versioned resource
			nil,
			management.NewDevPodWorkspaceInstanceStateREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalDevPodStopOptionsREST,
			func() runtime.Object { return &DevPodStopOptions{} }, // Register versioned resource
			nil,
			management.NewDevPodStopOptionsREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalDevPodWorkspaceInstanceTroubleshootREST,
			func() runtime.Object { return &DevPodWorkspaceInstanceTroubleshoot{} }, // Register versioned resource
			nil,
			management.NewDevPodWorkspaceInstanceTroubleshootREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalDevPodUpOptionsREST,
			func() runtime.Object { return &DevPodUpOptions{} }, // Register versioned resource
			nil,
			management.NewDevPodUpOptionsREST,
		),
		management.ManagementDevPodWorkspacePresetStorage,
		management.ManagementDevPodWorkspaceTemplateStorage,
		management.ManagementDirectClusterEndpointTokenStorage,
		management.ManagementEventStorage,
		management.ManagementFeatureStorage,
		builders.NewApiResourceWithStorage(
			management.InternalFeatureStatus,
			func() runtime.Object { return &Feature{} },     // Register versioned resource
			func() runtime.Object { return &FeatureList{} }, // Register versioned resource list
			management.NewFeatureStatusREST),
		management.ManagementIngressAuthTokenStorage,
		management.ManagementKioskStorage,
		management.ManagementLicenseStorage,
		builders.NewApiResourceWithStorage(
			management.InternalLicenseRequestREST,
			func() runtime.Object { return &LicenseRequest{} }, // Register versioned resource
			nil,
			management.NewLicenseRequestREST,
		),
		management.ManagementLicenseTokenStorage,
		management.ManagementLoftUpgradeStorage,
		management.ManagementOIDCClientStorage,
		management.ManagementOwnedAccessKeyStorage,
		management.ManagementProjectStorage,
		builders.NewApiResourceWithStorage(
			management.InternalProjectStatus,
			func() runtime.Object { return &Project{} },     // Register versioned resource
			func() runtime.Object { return &ProjectList{} }, // Register versioned resource list
			management.NewProjectStatusREST),
		builders.NewApiResourceWithStorage(
			management.InternalProjectChartInfoREST,
			func() runtime.Object { return &ProjectChartInfo{} }, // Register versioned resource
			nil,
			management.NewProjectChartInfoREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalProjectChartsREST,
			func() runtime.Object { return &ProjectCharts{} }, // Register versioned resource
			nil,
			management.NewProjectChartsREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalProjectClustersREST,
			func() runtime.Object { return &ProjectClusters{} }, // Register versioned resource
			nil,
			management.NewProjectClustersREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalProjectImportSpaceREST,
			func() runtime.Object { return &ProjectImportSpace{} }, // Register versioned resource
			nil,
			management.NewProjectImportSpaceREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalProjectMembersREST,
			func() runtime.Object { return &ProjectMembers{} }, // Register versioned resource
			nil,
			management.NewProjectMembersREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalProjectMigrateSpaceInstanceREST,
			func() runtime.Object { return &ProjectMigrateSpaceInstance{} }, // Register versioned resource
			nil,
			management.NewProjectMigrateSpaceInstanceREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalProjectMigrateVirtualClusterInstanceREST,
			func() runtime.Object { return &ProjectMigrateVirtualClusterInstance{} }, // Register versioned resource
			nil,
			management.NewProjectMigrateVirtualClusterInstanceREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalProjectRunnersREST,
			func() runtime.Object { return &ProjectRunners{} }, // Register versioned resource
			nil,
			management.NewProjectRunnersREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalProjectTemplatesREST,
			func() runtime.Object { return &ProjectTemplates{} }, // Register versioned resource
			nil,
			management.NewProjectTemplatesREST,
		),
		management.ManagementProjectSecretStorage,
		management.ManagementRedirectTokenStorage,
		management.ManagementRegisterVirtualClusterStorage,
		management.ManagementResetAccessKeyStorage,
		management.ManagementRunnerStorage,
		builders.NewApiResourceWithStorage(
			management.InternalRunnerStatus,
			func() runtime.Object { return &Runner{} },     // Register versioned resource
			func() runtime.Object { return &RunnerList{} }, // Register versioned resource list
			management.NewRunnerStatusREST),
		builders.NewApiResourceWithStorage(
			management.InternalRunnerAccessKeyREST,
			func() runtime.Object { return &RunnerAccessKey{} }, // Register versioned resource
			nil,
			management.NewRunnerAccessKeyREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalRunnerConfigREST,
			func() runtime.Object { return &RunnerConfig{} }, // Register versioned resource
			nil,
			management.NewRunnerConfigREST,
		),
		management.ManagementSelfStorage,
		management.ManagementSelfSubjectAccessReviewStorage,
		management.ManagementSharedSecretStorage,
		management.ManagementSpaceInstanceStorage,
		management.ManagementSpaceTemplateStorage,
		management.ManagementSubjectAccessReviewStorage,
		management.ManagementTaskStorage,
		builders.NewApiResourceWithStorage(
			management.InternalTaskLogREST,
			func() runtime.Object { return &TaskLog{} }, // Register versioned resource
			nil,
			management.NewTaskLogREST,
		),
		management.ManagementTeamStorage,
		builders.NewApiResourceWithStorage(
			management.InternalTeamAccessKeysREST,
			func() runtime.Object { return &TeamAccessKeys{} }, // Register versioned resource
			nil,
			management.NewTeamAccessKeysREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalTeamClustersREST,
			func() runtime.Object { return &TeamClusters{} }, // Register versioned resource
			nil,
			management.NewTeamClustersREST,
		),
		management.ManagementTranslateVClusterResourceNameStorage,
		management.ManagementUserStorage,
		builders.NewApiResourceWithStorage(
			management.InternalUserAccessKeysREST,
			func() runtime.Object { return &UserAccessKeys{} }, // Register versioned resource
			nil,
			management.NewUserAccessKeysREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalUserClustersREST,
			func() runtime.Object { return &UserClusters{} }, // Register versioned resource
			nil,
			management.NewUserClustersREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalUserPermissionsREST,
			func() runtime.Object { return &UserPermissions{} }, // Register versioned resource
			nil,
			management.NewUserPermissionsREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalUserProfileREST,
			func() runtime.Object { return &UserProfile{} }, // Register versioned resource
			nil,
			management.NewUserProfileREST,
		),
		management.ManagementVirtualClusterInstanceStorage,
		builders.NewApiResourceWithStorage(
			management.InternalVirtualClusterAccessKeyREST,
			func() runtime.Object { return &VirtualClusterAccessKey{} }, // Register versioned resource
			nil,
			management.NewVirtualClusterAccessKeyREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalVirtualClusterExternalDatabaseREST,
			func() runtime.Object { return &VirtualClusterExternalDatabase{} }, // Register versioned resource
			nil,
			management.NewVirtualClusterExternalDatabaseREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalVirtualClusterInstanceKubeConfigREST,
			func() runtime.Object { return &VirtualClusterInstanceKubeConfig{} }, // Register versioned resource
			nil,
			management.NewVirtualClusterInstanceKubeConfigREST,
		),
		builders.NewApiResourceWithStorage(
			management.InternalVirtualClusterInstanceLogREST,
			func() runtime.Object { return &VirtualClusterInstanceLog{} }, // Register versioned resource
			nil,
			management.NewVirtualClusterInstanceLogREST,
		),
		management.ManagementVirtualClusterTemplateStorage,
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

type AgentAuditEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentAuditEvent `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AnnouncementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Announcement `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCredentialsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppCredentials `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BackupApplyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupApply `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAccessKey `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterAgentConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAgentConfig `json:"items"`
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

type ClusterMemberAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterMemberAccess `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterMembersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterMembers `json:"items"`
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

type ClusterAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAccess `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterRoleTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRoleTemplate `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ConvertVirtualClusterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConvertVirtualClusterConfig `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodEnvironmentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodEnvironmentTemplate `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodWorkspaceInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodWorkspaceInstance `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodDeleteOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodDeleteOptions `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodStatusOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodStatusOptions `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodSshOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodSshOptions `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodWorkspaceInstanceStateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodWorkspaceInstanceState `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodStopOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodStopOptions `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodWorkspaceInstanceTroubleshootList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodWorkspaceInstanceTroubleshoot `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodUpOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodUpOptions `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodWorkspacePresetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodWorkspacePreset `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DevPodWorkspaceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodWorkspaceTemplate `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DirectClusterEndpointTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectClusterEndpointToken `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Event `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FeatureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Feature `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IngressAuthTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IngressAuthToken `json:"items"`
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

type LicenseRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicenseRequest `json:"items"`
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

type OIDCClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OIDCClient `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OwnedAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OwnedAccessKey `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectChartInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectChartInfo `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectChartsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectCharts `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectClustersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectClusters `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectImportSpaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectImportSpace `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectMembersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectMembers `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectMigrateSpaceInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectMigrateSpaceInstance `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectMigrateVirtualClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectMigrateVirtualClusterInstance `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectRunnersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectRunners `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectTemplatesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectTemplates `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectSecret `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RedirectTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedirectToken `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RegisterVirtualClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegisterVirtualCluster `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ResetAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResetAccessKey `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RunnerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Runner `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RunnerAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunnerAccessKey `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RunnerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunnerConfig `json:"items"`
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

type SpaceInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpaceInstance `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SpaceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpaceTemplate `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SubjectAccessReviewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubjectAccessReview `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Task `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TaskLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TaskLog `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Team `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TeamAccessKeysList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamAccessKeys `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TeamClustersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamClusters `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TranslateVClusterResourceNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TranslateVClusterResourceName `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserAccessKeysList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserAccessKeys `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserClustersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserClusters `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserPermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPermissions `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserProfile `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterInstance `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualClusterAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterAccessKey `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualClusterExternalDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterExternalDatabase `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualClusterInstanceKubeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterInstanceKubeConfig `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualClusterInstanceLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterInstanceLog `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualClusterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterTemplate `json:"items"`
}
