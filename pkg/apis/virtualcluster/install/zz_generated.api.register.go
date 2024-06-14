// Code generated by generator. DO NOT EDIT.

package install

import (
	"github.com/loft-sh/api/v4/pkg/apis/virtualcluster"
	v1 "github.com/loft-sh/api/v4/pkg/apis/virtualcluster/v1"
	"github.com/loft-sh/apiserver/pkg/builders"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func init() {
	Install(builders.Scheme)
}

func Install(scheme *runtime.Scheme) {
	utilruntime.Must(v1.AddToScheme(scheme))
	utilruntime.Must(virtualcluster.AddToScheme(scheme))
	utilruntime.Must(addKnownTypes(scheme))
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(virtualcluster.SchemeGroupVersion,
		&virtualcluster.HelmRelease{},
		&virtualcluster.HelmReleaseList{},
	)
	return nil
}
