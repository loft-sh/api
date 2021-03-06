// Code generated by generator. DO NOT EDIT.

package install

import (
	"github.com/loft-sh/api/pkg/apis/cluster"
	v1 "github.com/loft-sh/api/pkg/apis/cluster/v1"
	"github.com/loft-sh/apiserver/pkg/builders"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func init() {
	Install(builders.Scheme)
}

func Install(scheme *runtime.Scheme) {
	utilruntime.Must(v1.AddToScheme(scheme))
	utilruntime.Must(cluster.AddToScheme(scheme))
	utilruntime.Must(addKnownTypes(scheme))
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(cluster.SchemeGroupVersion,
		&cluster.Account{},
		&cluster.AccountList{},
		&cluster.AccountClusterRoles{},
		&cluster.Context{},
		&cluster.ContextList{},
		&cluster.HelmRelease{},
		&cluster.HelmReleaseList{},
		&cluster.SleepModeConfig{},
		&cluster.SleepModeConfigList{},
	)
	return nil
}
