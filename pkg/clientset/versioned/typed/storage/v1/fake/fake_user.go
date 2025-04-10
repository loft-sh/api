// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	storagev1 "github.com/loft-sh/api/v4/pkg/clientset/versioned/typed/storage/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeUsers implements UserInterface
type fakeUsers struct {
	*gentype.FakeClientWithList[*v1.User, *v1.UserList]
	Fake *FakeStorageV1
}

func newFakeUsers(fake *FakeStorageV1) storagev1.UserInterface {
	return &fakeUsers{
		gentype.NewFakeClientWithList[*v1.User, *v1.UserList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("users"),
			v1.SchemeGroupVersion.WithKind("User"),
			func() *v1.User { return &v1.User{} },
			func() *v1.UserList { return &v1.UserList{} },
			func(dst, src *v1.UserList) { dst.ListMeta = src.ListMeta },
			func(list *v1.UserList) []*v1.User { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.UserList, items []*v1.User) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
