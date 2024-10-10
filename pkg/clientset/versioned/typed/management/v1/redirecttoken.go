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

// RedirectTokensGetter has a method to return a RedirectTokenInterface.
// A group's client should implement this interface.
type RedirectTokensGetter interface {
	RedirectTokens() RedirectTokenInterface
}

// RedirectTokenInterface has methods to work with RedirectToken resources.
type RedirectTokenInterface interface {
	Create(ctx context.Context, redirectToken *v1.RedirectToken, opts metav1.CreateOptions) (*v1.RedirectToken, error)
	Update(ctx context.Context, redirectToken *v1.RedirectToken, opts metav1.UpdateOptions) (*v1.RedirectToken, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, redirectToken *v1.RedirectToken, opts metav1.UpdateOptions) (*v1.RedirectToken, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.RedirectToken, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RedirectTokenList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RedirectToken, err error)
	RedirectTokenExpansion
}

// redirectTokens implements RedirectTokenInterface
type redirectTokens struct {
	*gentype.ClientWithList[*v1.RedirectToken, *v1.RedirectTokenList]
}

// newRedirectTokens returns a RedirectTokens
func newRedirectTokens(c *ManagementV1Client) *redirectTokens {
	return &redirectTokens{
		gentype.NewClientWithList[*v1.RedirectToken, *v1.RedirectTokenList](
			"redirecttokens",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.RedirectToken { return &v1.RedirectToken{} },
			func() *v1.RedirectTokenList { return &v1.RedirectTokenList{} }),
	}
}
