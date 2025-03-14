// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	storagev1 "github.com/loft-sh/api/v4/pkg/clientset/versioned/typed/storage/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeProjects implements ProjectInterface
type fakeProjects struct {
	*gentype.FakeClientWithList[*v1.Project, *v1.ProjectList]
	Fake *FakeStorageV1
}

func newFakeProjects(fake *FakeStorageV1) storagev1.ProjectInterface {
	return &fakeProjects{
		gentype.NewFakeClientWithList[*v1.Project, *v1.ProjectList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("projects"),
			v1.SchemeGroupVersion.WithKind("Project"),
			func() *v1.Project { return &v1.Project{} },
			func() *v1.ProjectList { return &v1.ProjectList{} },
			func(dst, src *v1.ProjectList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ProjectList) []*v1.Project { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.ProjectList, items []*v1.Project) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
