// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// SelfSubjectAccessReviewsGetter has a method to return a SelfSubjectAccessReviewInterface.
// A group's client should implement this interface.
type SelfSubjectAccessReviewsGetter interface {
	SelfSubjectAccessReviews() SelfSubjectAccessReviewInterface
}

// SelfSubjectAccessReviewInterface has methods to work with SelfSubjectAccessReview resources.
type SelfSubjectAccessReviewInterface interface {
	Create(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.CreateOptions) (*v1.SelfSubjectAccessReview, error)
	Update(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.UpdateOptions) (*v1.SelfSubjectAccessReview, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.UpdateOptions) (*v1.SelfSubjectAccessReview, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SelfSubjectAccessReview, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SelfSubjectAccessReviewList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SelfSubjectAccessReview, err error)
	SelfSubjectAccessReviewExpansion
}

// selfSubjectAccessReviews implements SelfSubjectAccessReviewInterface
type selfSubjectAccessReviews struct {
	*gentype.ClientWithList[*v1.SelfSubjectAccessReview, *v1.SelfSubjectAccessReviewList]
}

// newSelfSubjectAccessReviews returns a SelfSubjectAccessReviews
func newSelfSubjectAccessReviews(c *ManagementV1Client) *selfSubjectAccessReviews {
	return &selfSubjectAccessReviews{
		gentype.NewClientWithList[*v1.SelfSubjectAccessReview, *v1.SelfSubjectAccessReviewList](
			"selfsubjectaccessreviews",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.SelfSubjectAccessReview { return &v1.SelfSubjectAccessReview{} },
			func() *v1.SelfSubjectAccessReviewList { return &v1.SelfSubjectAccessReviewList{} }),
	}
}
