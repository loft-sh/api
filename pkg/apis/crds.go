package apis

import (
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	"github.com/loft-sh/kiosk/pkg/store/crd"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var (
	// TypeDefinitions to create the appropriate crds
	TypeDefinitions = []*crd.TypeDefinition{
		&crd.TypeDefinition{
			GVK:      storagev1.SchemeGroupVersion.WithKind("User"),
			Singular: "user",
			Plural:   "users",
			Scope:    apiextensionsv1.ClusterScoped,
		},
		&crd.TypeDefinition{
			GVK:      storagev1.SchemeGroupVersion.WithKind("Team"),
			Singular: "team",
			Plural:   "teams",
			Scope:    apiextensionsv1.ClusterScoped,
		},
		&crd.TypeDefinition{
			GVK:      storagev1.SchemeGroupVersion.WithKind("Cluster"),
			Singular: "cluster",
			Plural:   "clusters",
			Scope:    apiextensionsv1.ClusterScoped,
		},
		&crd.TypeDefinition{
			GVK:      storagev1.SchemeGroupVersion.WithKind("SharedSecret"),
			Singular: "sharedsecret",
			Plural:   "sharedsecrets",
			Scope:    apiextensionsv1.NamespaceScoped,
		},
		&crd.TypeDefinition{
			GVK:      storagev1.SchemeGroupVersion.WithKind("AccessKey"),
			Singular: "accesskey",
			Plural:   "accesskeys",
			Scope:    apiextensionsv1.ClusterScoped,
		},
		&crd.TypeDefinition{
			GVK:      storagev1.SchemeGroupVersion.WithKind("ClusterAccountTemplate"),
			Singular: "clusteraccounttemplate",
			Plural:   "clusteraccounttemplates",
			Scope:    apiextensionsv1.ClusterScoped,
		},
	}
)
