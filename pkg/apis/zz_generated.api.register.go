// Code generated by generator. DO NOT EDIT.

package apis

import (
	"github.com/loft-sh/api/pkg/apis/cluster"
	_ "github.com/loft-sh/api/pkg/apis/cluster/install" // Install the cluster group
	clusterv1 "github.com/loft-sh/api/pkg/apis/cluster/v1"
	"github.com/loft-sh/api/pkg/apis/management"
	_ "github.com/loft-sh/api/pkg/apis/management/install" // Install the management group
	managementv1 "github.com/loft-sh/api/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

var (
	localSchemeBuilder = runtime.SchemeBuilder{
		clusterv1.AddToScheme,
		managementv1.AddToScheme,
	}
	AddToScheme = localSchemeBuilder.AddToScheme
)

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetClusterAPIBuilder(),
		GetManagementAPIBuilder(),
	}
}

var clusterApiGroup = builders.NewApiGroupBuilder(
	"cluster.loft.sh",
	"github.com/loft-sh/api/pkg/apis/cluster").
	WithUnVersionedApi(cluster.ApiVersion).
	WithVersionedApis(
		clusterv1.ApiVersion,
	).
	WithRootScopedKinds(
		"Account",
		"Context",
	)

func GetClusterAPIBuilder() *builders.APIGroupBuilder {
	return clusterApiGroup
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
		"Cluster",
		"ClusterConnect",
		"ClusterRole",
		"Config",
		"Feature",
		"License",
		"LicenseToken",
		"SelfSubjectAccessReview",
		"SubjectAccessReview",
		"Team",
		"User",
	)

func GetManagementAPIBuilder() *builders.APIGroupBuilder {
	return managementApiGroup
}
