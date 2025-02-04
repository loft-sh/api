// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	managementv1 "github.com/loft-sh/api/v4/pkg/clientset/versioned/typed/management/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeSpaceInstances implements SpaceInstanceInterface
type fakeSpaceInstances struct {
	*gentype.FakeClientWithList[*v1.SpaceInstance, *v1.SpaceInstanceList]
	Fake *FakeManagementV1
}

func newFakeSpaceInstances(fake *FakeManagementV1, namespace string) managementv1.SpaceInstanceInterface {
	return &fakeSpaceInstances{
		gentype.NewFakeClientWithList[*v1.SpaceInstance, *v1.SpaceInstanceList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("spaceinstances"),
			v1.SchemeGroupVersion.WithKind("SpaceInstance"),
			func() *v1.SpaceInstance { return &v1.SpaceInstance{} },
			func() *v1.SpaceInstanceList { return &v1.SpaceInstanceList{} },
			func(dst, src *v1.SpaceInstanceList) { dst.ListMeta = src.ListMeta },
			func(list *v1.SpaceInstanceList) []*v1.SpaceInstance { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.SpaceInstanceList, items []*v1.SpaceInstance) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
