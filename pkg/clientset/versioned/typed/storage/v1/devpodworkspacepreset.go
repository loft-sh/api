// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// DevPodWorkspacePresetsGetter has a method to return a DevPodWorkspacePresetInterface.
// A group's client should implement this interface.
type DevPodWorkspacePresetsGetter interface {
	DevPodWorkspacePresets() DevPodWorkspacePresetInterface
}

// DevPodWorkspacePresetInterface has methods to work with DevPodWorkspacePreset resources.
type DevPodWorkspacePresetInterface interface {
	Create(ctx context.Context, devPodWorkspacePreset *v1.DevPodWorkspacePreset, opts metav1.CreateOptions) (*v1.DevPodWorkspacePreset, error)
	Update(ctx context.Context, devPodWorkspacePreset *v1.DevPodWorkspacePreset, opts metav1.UpdateOptions) (*v1.DevPodWorkspacePreset, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DevPodWorkspacePreset, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DevPodWorkspacePresetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DevPodWorkspacePreset, err error)
	DevPodWorkspacePresetExpansion
}

// devPodWorkspacePresets implements DevPodWorkspacePresetInterface
type devPodWorkspacePresets struct {
	*gentype.ClientWithList[*v1.DevPodWorkspacePreset, *v1.DevPodWorkspacePresetList]
}

// newDevPodWorkspacePresets returns a DevPodWorkspacePresets
func newDevPodWorkspacePresets(c *StorageV1Client) *devPodWorkspacePresets {
	return &devPodWorkspacePresets{
		gentype.NewClientWithList[*v1.DevPodWorkspacePreset, *v1.DevPodWorkspacePresetList](
			"devpodworkspacepresets",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.DevPodWorkspacePreset { return &v1.DevPodWorkspacePreset{} },
			func() *v1.DevPodWorkspacePresetList { return &v1.DevPodWorkspacePresetList{} }),
	}
}
