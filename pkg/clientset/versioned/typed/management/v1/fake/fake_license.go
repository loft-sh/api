// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	context "context"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	managementv1 "github.com/loft-sh/api/v4/pkg/clientset/versioned/typed/management/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gentype "k8s.io/client-go/gentype"
	testing "k8s.io/client-go/testing"
)

// fakeLicenses implements LicenseInterface
type fakeLicenses struct {
	*gentype.FakeClientWithList[*v1.License, *v1.LicenseList]
	Fake *FakeManagementV1
}

func newFakeLicenses(fake *FakeManagementV1) managementv1.LicenseInterface {
	return &fakeLicenses{
		gentype.NewFakeClientWithList[*v1.License, *v1.LicenseList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("licenses"),
			v1.SchemeGroupVersion.WithKind("License"),
			func() *v1.License { return &v1.License{} },
			func() *v1.LicenseList { return &v1.LicenseList{} },
			func(dst, src *v1.LicenseList) { dst.ListMeta = src.ListMeta },
			func(list *v1.LicenseList) []*v1.License { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.LicenseList, items []*v1.License) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}

// LicenseRequest takes the representation of a licenseRequest and creates it.  Returns the server's representation of the licenseRequest, and an error, if there is any.
func (c *fakeLicenses) LicenseRequest(ctx context.Context, licenseName string, licenseRequest *v1.LicenseRequest, opts metav1.CreateOptions) (result *v1.LicenseRequest, err error) {
	emptyResult := &v1.LicenseRequest{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateSubresourceActionWithOptions(c.Resource(), licenseName, "request", licenseRequest, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.LicenseRequest), err
}
