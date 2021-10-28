// Code generated by generator. DO NOT EDIT.

package apis

import (
	"github.com/loft-sh/api/pkg/apis/management"
	_ "github.com/loft-sh/api/pkg/apis/management/install" // Install the management group
	managementv1 "github.com/loft-sh/api/pkg/apis/management/v1"
	"github.com/loft-sh/api/pkg/apis/virtualcluster"
	_ "github.com/loft-sh/api/pkg/apis/virtualcluster/install" // Install the virtualcluster group
	virtualclusterv1 "github.com/loft-sh/api/pkg/apis/virtualcluster/v1"
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

var managementApiGroup = builders.NewApiGroupBuilder(
	"management.loft.sh",
	"github.com/loft-sh/api/pkg/apis/management").
	WithUnVersionedApi(management.ApiVersion).
	WithVersionedApis(
		managementv1.ApiVersion,
	).
	WithRootScopedKinds(
		"Announcement",
		"App",
		"Cluster",
		"ClusterAccountTemplate",
		"ClusterConnect",
		"Config",
		"DirectClusterEndpointToken",
		"Feature",
		"GlobalClusterAccess",
		"GlobalClusterRoleTemplate",
		"GlobalSpaceConstraint",
		"License",
		"LicenseToken",
		"LoftUpgrade",
		"OwnedAccessKey",
		"PolicyViolation",
		"ResetAccessKey",
		"Self",
		"SelfSubjectAccessReview",
		"SpaceTemplate",
		"SubjectAccessReview",
		"Task",
		"Team",
		"User",
		"VirtualClusterTemplate",
	)

func GetManagementAPIBuilder() *builders.APIGroupBuilder {
	return managementApiGroup
}

var virtualclusterApiGroup = builders.NewApiGroupBuilder(
	"virtualcluster.loft.sh",
	"github.com/loft-sh/api/pkg/apis/virtualcluster").
	WithUnVersionedApi(virtualcluster.ApiVersion).
	WithVersionedApis(
		virtualclusterv1.ApiVersion,
	).
	WithRootScopedKinds()

func GetVirtualclusterAPIBuilder() *builders.APIGroupBuilder {
	return virtualclusterApiGroup
}
