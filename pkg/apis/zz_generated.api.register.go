// Code generated by generator. DO NOT EDIT.

package apis

import (
	"github.com/loft-sh/api/v3/pkg/apis/management"
	_ "github.com/loft-sh/api/v3/pkg/apis/management/install" // Install the management group
	managementv1 "github.com/loft-sh/api/v3/pkg/apis/management/v1"
	"github.com/loft-sh/api/v3/pkg/apis/virtualcluster"
	_ "github.com/loft-sh/api/v3/pkg/apis/virtualcluster/install" // Install the virtualcluster group
	virtualclusterv1 "github.com/loft-sh/api/v3/pkg/apis/virtualcluster/v1"
	"github.com/loft-sh/apiserver/pkg/builders"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	localSchemeBuilder = runtime.SchemeBuilder{
		managementv1.AddToScheme,
		virtualclusterv1.AddToScheme,
	}
	AddToScheme = localSchemeBuilder.AddToScheme
)

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetManagementAPIBuilder(),
		GetVirtualclusterAPIBuilder(),
	}
}

func GetManagementAPIBuilder() *builders.APIGroupBuilder {
	return builders.NewApiGroupBuilder(
		"management.loft.sh",
		"github.com/loft-sh/api/v3/pkg/apis/management").
		WithUnVersionedApi(management.ApiVersion).
		WithVersionedApis(
			managementv1.ApiVersion,
		).
		WithRootScopedKinds(
			"AgentAuditEvent",
			"Announcement",
			"App",
			"Backup",
			"Cluster",
			"ClusterAccess",
			"ClusterConnect",
			"ClusterRoleTemplate",
			"Config",
			"ConvertVirtualClusterConfig",
			"DevPodWorkspaceTemplate",
			"DirectClusterEndpointToken",
			"Event",
			"Feature",
			"IngressAuthToken",
			"License",
			"LicenseToken",
			"LoftUpgrade",
			"OwnedAccessKey",
			"Project",
			"RedirectToken",
			"ResetAccessKey",
			"Runner",
			"Self",
			"SelfSubjectAccessReview",
			"SpaceConstraint",
			"SpaceTemplate",
			"SubjectAccessReview",
			"Task",
			"Team",
			"User",
			"VirtualClusterTemplate",
		)
}
func GetVirtualclusterAPIBuilder() *builders.APIGroupBuilder {
	return builders.NewApiGroupBuilder(
		"virtualcluster.loft.sh",
		"github.com/loft-sh/api/v3/pkg/apis/virtualcluster").
		WithUnVersionedApi(virtualcluster.ApiVersion).
		WithVersionedApis(
			virtualclusterv1.ApiVersion,
		).
		WithRootScopedKinds()
}
