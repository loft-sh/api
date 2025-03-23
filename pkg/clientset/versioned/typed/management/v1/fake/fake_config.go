// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	managementv1 "github.com/loft-sh/api/v4/pkg/clientset/versioned/typed/management/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeConfigs implements ConfigInterface
type fakeConfigs struct {
	*gentype.FakeClientWithList[*v1.Config, *v1.ConfigList]
	Fake *FakeManagementV1
}

func newFakeConfigs(fake *FakeManagementV1) managementv1.ConfigInterface {
	return &fakeConfigs{
		gentype.NewFakeClientWithList[*v1.Config, *v1.ConfigList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("configs"),
			v1.SchemeGroupVersion.WithKind("Config"),
			func() *v1.Config { return &v1.Config{} },
			func() *v1.ConfigList { return &v1.ConfigList{} },
			func(dst, src *v1.ConfigList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ConfigList) []*v1.Config { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.ConfigList, items []*v1.Config) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
