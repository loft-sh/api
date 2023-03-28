package install

import (
	"github.com/loft-sh/api/v2/pkg/apis/management"
	v1 "github.com/loft-sh/api/v2/pkg/apis/management/v1"
	"github.com/loft-sh/apiserver/pkg/builders"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func init() {
	InstallOptions(builders.Scheme)
	InstallOptions(builders.ParameterScheme)
	utilruntime.Must(v1.RegisterConversions(builders.ParameterScheme))
}

func InstallOptions(scheme *runtime.Scheme) {
	utilruntime.Must(v1.InstallOptions(scheme))
	utilruntime.Must(addKnownOptionsTypes(scheme))
}

func addKnownOptionsTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(management.SchemeGroupVersion,
		&management.TaskLogOptions{},
		&management.UserSpacesOptions{},
		&management.UserVirtualClustersOptions{},
		&management.UserQuotasOptions{},
	)
	return nil
}
