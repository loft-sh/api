// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// SpaceInstancesGetter has a method to return a SpaceInstanceInterface.
// A group's client should implement this interface.
type SpaceInstancesGetter interface {
	SpaceInstances(namespace string) SpaceInstanceInterface
}

// SpaceInstanceInterface has methods to work with SpaceInstance resources.
type SpaceInstanceInterface interface {
	Create(ctx context.Context, spaceInstance *storagev1.SpaceInstance, opts metav1.CreateOptions) (*storagev1.SpaceInstance, error)
	Update(ctx context.Context, spaceInstance *storagev1.SpaceInstance, opts metav1.UpdateOptions) (*storagev1.SpaceInstance, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, spaceInstance *storagev1.SpaceInstance, opts metav1.UpdateOptions) (*storagev1.SpaceInstance, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*storagev1.SpaceInstance, error)
	List(ctx context.Context, opts metav1.ListOptions) (*storagev1.SpaceInstanceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *storagev1.SpaceInstance, err error)
	SpaceInstanceExpansion
}

// spaceInstances implements SpaceInstanceInterface
type spaceInstances struct {
	*gentype.ClientWithList[*storagev1.SpaceInstance, *storagev1.SpaceInstanceList]
}

// newSpaceInstances returns a SpaceInstances
func newSpaceInstances(c *StorageV1Client, namespace string) *spaceInstances {
	return &spaceInstances{
		gentype.NewClientWithList[*storagev1.SpaceInstance, *storagev1.SpaceInstanceList](
			"spaceinstances",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *storagev1.SpaceInstance { return &storagev1.SpaceInstance{} },
			func() *storagev1.SpaceInstanceList { return &storagev1.SpaceInstanceList{} },
		),
	}
}
