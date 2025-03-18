// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	managementv1 "github.com/loft-sh/api/v4/pkg/clientset/versioned/typed/management/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeDevPodWorkspacePresets implements DevPodWorkspacePresetInterface
type fakeDevPodWorkspacePresets struct {
	*gentype.FakeClientWithList[*v1.DevPodWorkspacePreset, *v1.DevPodWorkspacePresetList]
	Fake *FakeManagementV1
}

func newFakeDevPodWorkspacePresets(fake *FakeManagementV1) managementv1.DevPodWorkspacePresetInterface {
	return &fakeDevPodWorkspacePresets{
		gentype.NewFakeClientWithList[*v1.DevPodWorkspacePreset, *v1.DevPodWorkspacePresetList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("devpodworkspacepresets"),
			v1.SchemeGroupVersion.WithKind("DevPodWorkspacePreset"),
			func() *v1.DevPodWorkspacePreset { return &v1.DevPodWorkspacePreset{} },
			func() *v1.DevPodWorkspacePresetList { return &v1.DevPodWorkspacePresetList{} },
			func(dst, src *v1.DevPodWorkspacePresetList) { dst.ListMeta = src.ListMeta },
			func(list *v1.DevPodWorkspacePresetList) []*v1.DevPodWorkspacePreset {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.DevPodWorkspacePresetList, items []*v1.DevPodWorkspacePreset) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
