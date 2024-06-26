// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TranslateVClusterResourceNamesGetter has a method to return a TranslateVClusterResourceNameInterface.
// A group's client should implement this interface.
type TranslateVClusterResourceNamesGetter interface {
	TranslateVClusterResourceNames() TranslateVClusterResourceNameInterface
}

// TranslateVClusterResourceNameInterface has methods to work with TranslateVClusterResourceName resources.
type TranslateVClusterResourceNameInterface interface {
	Create(ctx context.Context, translateVClusterResourceName *v1.TranslateVClusterResourceName, opts metav1.CreateOptions) (*v1.TranslateVClusterResourceName, error)
	Update(ctx context.Context, translateVClusterResourceName *v1.TranslateVClusterResourceName, opts metav1.UpdateOptions) (*v1.TranslateVClusterResourceName, error)
	UpdateStatus(ctx context.Context, translateVClusterResourceName *v1.TranslateVClusterResourceName, opts metav1.UpdateOptions) (*v1.TranslateVClusterResourceName, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TranslateVClusterResourceName, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TranslateVClusterResourceNameList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TranslateVClusterResourceName, err error)
	TranslateVClusterResourceNameExpansion
}

// translateVClusterResourceNames implements TranslateVClusterResourceNameInterface
type translateVClusterResourceNames struct {
	client rest.Interface
}

// newTranslateVClusterResourceNames returns a TranslateVClusterResourceNames
func newTranslateVClusterResourceNames(c *ManagementV1Client) *translateVClusterResourceNames {
	return &translateVClusterResourceNames{
		client: c.RESTClient(),
	}
}

// Get takes name of the translateVClusterResourceName, and returns the corresponding translateVClusterResourceName object, and an error if there is any.
func (c *translateVClusterResourceNames) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TranslateVClusterResourceName, err error) {
	result = &v1.TranslateVClusterResourceName{}
	err = c.client.Get().
		Resource("translatevclusterresourcenames").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TranslateVClusterResourceNames that match those selectors.
func (c *translateVClusterResourceNames) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TranslateVClusterResourceNameList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TranslateVClusterResourceNameList{}
	err = c.client.Get().
		Resource("translatevclusterresourcenames").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested translateVClusterResourceNames.
func (c *translateVClusterResourceNames) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("translatevclusterresourcenames").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a translateVClusterResourceName and creates it.  Returns the server's representation of the translateVClusterResourceName, and an error, if there is any.
func (c *translateVClusterResourceNames) Create(ctx context.Context, translateVClusterResourceName *v1.TranslateVClusterResourceName, opts metav1.CreateOptions) (result *v1.TranslateVClusterResourceName, err error) {
	result = &v1.TranslateVClusterResourceName{}
	err = c.client.Post().
		Resource("translatevclusterresourcenames").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(translateVClusterResourceName).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a translateVClusterResourceName and updates it. Returns the server's representation of the translateVClusterResourceName, and an error, if there is any.
func (c *translateVClusterResourceNames) Update(ctx context.Context, translateVClusterResourceName *v1.TranslateVClusterResourceName, opts metav1.UpdateOptions) (result *v1.TranslateVClusterResourceName, err error) {
	result = &v1.TranslateVClusterResourceName{}
	err = c.client.Put().
		Resource("translatevclusterresourcenames").
		Name(translateVClusterResourceName.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(translateVClusterResourceName).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *translateVClusterResourceNames) UpdateStatus(ctx context.Context, translateVClusterResourceName *v1.TranslateVClusterResourceName, opts metav1.UpdateOptions) (result *v1.TranslateVClusterResourceName, err error) {
	result = &v1.TranslateVClusterResourceName{}
	err = c.client.Put().
		Resource("translatevclusterresourcenames").
		Name(translateVClusterResourceName.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(translateVClusterResourceName).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the translateVClusterResourceName and deletes it. Returns an error if one occurs.
func (c *translateVClusterResourceNames) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("translatevclusterresourcenames").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *translateVClusterResourceNames) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("translatevclusterresourcenames").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched translateVClusterResourceName.
func (c *translateVClusterResourceNames) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TranslateVClusterResourceName, err error) {
	result = &v1.TranslateVClusterResourceName{}
	err = c.client.Patch(pt).
		Resource("translatevclusterresourcenames").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
