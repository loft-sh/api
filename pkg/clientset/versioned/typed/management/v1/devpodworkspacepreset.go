// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
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
	Create(ctx context.Context, devPodWorkspacePreset *managementv1.DevPodWorkspacePreset, opts metav1.CreateOptions) (*managementv1.DevPodWorkspacePreset, error)
	Update(ctx context.Context, devPodWorkspacePreset *managementv1.DevPodWorkspacePreset, opts metav1.UpdateOptions) (*managementv1.DevPodWorkspacePreset, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, devPodWorkspacePreset *managementv1.DevPodWorkspacePreset, opts metav1.UpdateOptions) (*managementv1.DevPodWorkspacePreset, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*managementv1.DevPodWorkspacePreset, error)
	List(ctx context.Context, opts metav1.ListOptions) (*managementv1.DevPodWorkspacePresetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *managementv1.DevPodWorkspacePreset, err error)
	DevPodWorkspacePresetExpansion
}

// devPodWorkspacePresets implements DevPodWorkspacePresetInterface
type devPodWorkspacePresets struct {
	*gentype.ClientWithList[*managementv1.DevPodWorkspacePreset, *managementv1.DevPodWorkspacePresetList]
}

// newDevPodWorkspacePresets returns a DevPodWorkspacePresets
func newDevPodWorkspacePresets(c *ManagementV1Client) *devPodWorkspacePresets {
	return &devPodWorkspacePresets{
		gentype.NewClientWithList[*managementv1.DevPodWorkspacePreset, *managementv1.DevPodWorkspacePresetList](
			"devpodworkspacepresets",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *managementv1.DevPodWorkspacePreset { return &managementv1.DevPodWorkspacePreset{} },
			func() *managementv1.DevPodWorkspacePresetList { return &managementv1.DevPodWorkspacePresetList{} },
		),
	}
}
