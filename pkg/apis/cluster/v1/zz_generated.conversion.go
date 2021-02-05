// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	cluster "github.com/loft-sh/api/pkg/apis/cluster"
	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Account)(nil), (*cluster.Account)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Account_To_cluster_Account(a.(*Account), b.(*cluster.Account), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.Account)(nil), (*Account)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_Account_To_v1_Account(a.(*cluster.Account), b.(*Account), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountClusterRoles)(nil), (*cluster.AccountClusterRoles)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AccountClusterRoles_To_cluster_AccountClusterRoles(a.(*AccountClusterRoles), b.(*cluster.AccountClusterRoles), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.AccountClusterRoles)(nil), (*AccountClusterRoles)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_AccountClusterRoles_To_v1_AccountClusterRoles(a.(*cluster.AccountClusterRoles), b.(*AccountClusterRoles), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountClusterRolesList)(nil), (*cluster.AccountClusterRolesList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AccountClusterRolesList_To_cluster_AccountClusterRolesList(a.(*AccountClusterRolesList), b.(*cluster.AccountClusterRolesList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.AccountClusterRolesList)(nil), (*AccountClusterRolesList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_AccountClusterRolesList_To_v1_AccountClusterRolesList(a.(*cluster.AccountClusterRolesList), b.(*AccountClusterRolesList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountList)(nil), (*cluster.AccountList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AccountList_To_cluster_AccountList(a.(*AccountList), b.(*cluster.AccountList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.AccountList)(nil), (*AccountList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_AccountList_To_v1_AccountList(a.(*cluster.AccountList), b.(*AccountList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountSpec)(nil), (*cluster.AccountSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AccountSpec_To_cluster_AccountSpec(a.(*AccountSpec), b.(*cluster.AccountSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.AccountSpec)(nil), (*AccountSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_AccountSpec_To_v1_AccountSpec(a.(*cluster.AccountSpec), b.(*AccountSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountStatus)(nil), (*cluster.AccountStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AccountStatus_To_cluster_AccountStatus(a.(*AccountStatus), b.(*cluster.AccountStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.AccountStatus)(nil), (*AccountStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_AccountStatus_To_v1_AccountStatus(a.(*cluster.AccountStatus), b.(*AccountStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Context)(nil), (*cluster.Context)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Context_To_cluster_Context(a.(*Context), b.(*cluster.Context), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.Context)(nil), (*Context)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_Context_To_v1_Context(a.(*cluster.Context), b.(*Context), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ContextList)(nil), (*cluster.ContextList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ContextList_To_cluster_ContextList(a.(*ContextList), b.(*cluster.ContextList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.ContextList)(nil), (*ContextList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_ContextList_To_v1_ContextList(a.(*cluster.ContextList), b.(*ContextList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ContextSpec)(nil), (*cluster.ContextSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ContextSpec_To_cluster_ContextSpec(a.(*ContextSpec), b.(*cluster.ContextSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.ContextSpec)(nil), (*ContextSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_ContextSpec_To_v1_ContextSpec(a.(*cluster.ContextSpec), b.(*ContextSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ContextStatus)(nil), (*cluster.ContextStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ContextStatus_To_cluster_ContextStatus(a.(*ContextStatus), b.(*cluster.ContextStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.ContextStatus)(nil), (*ContextStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_ContextStatus_To_v1_ContextStatus(a.(*cluster.ContextStatus), b.(*ContextStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EpochInfo)(nil), (*cluster.EpochInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EpochInfo_To_cluster_EpochInfo(a.(*EpochInfo), b.(*cluster.EpochInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.EpochInfo)(nil), (*EpochInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_EpochInfo_To_v1_EpochInfo(a.(*cluster.EpochInfo), b.(*EpochInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*HelmRelease)(nil), (*cluster.HelmRelease)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_HelmRelease_To_cluster_HelmRelease(a.(*HelmRelease), b.(*cluster.HelmRelease), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.HelmRelease)(nil), (*HelmRelease)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_HelmRelease_To_v1_HelmRelease(a.(*cluster.HelmRelease), b.(*HelmRelease), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*HelmReleaseList)(nil), (*cluster.HelmReleaseList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_HelmReleaseList_To_cluster_HelmReleaseList(a.(*HelmReleaseList), b.(*cluster.HelmReleaseList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.HelmReleaseList)(nil), (*HelmReleaseList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_HelmReleaseList_To_v1_HelmReleaseList(a.(*cluster.HelmReleaseList), b.(*HelmReleaseList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*HelmReleaseSpec)(nil), (*cluster.HelmReleaseSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_HelmReleaseSpec_To_cluster_HelmReleaseSpec(a.(*HelmReleaseSpec), b.(*cluster.HelmReleaseSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.HelmReleaseSpec)(nil), (*HelmReleaseSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_HelmReleaseSpec_To_v1_HelmReleaseSpec(a.(*cluster.HelmReleaseSpec), b.(*HelmReleaseSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*HelmReleaseStatus)(nil), (*cluster.HelmReleaseStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_HelmReleaseStatus_To_cluster_HelmReleaseStatus(a.(*HelmReleaseStatus), b.(*cluster.HelmReleaseStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.HelmReleaseStatus)(nil), (*HelmReleaseStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_HelmReleaseStatus_To_v1_HelmReleaseStatus(a.(*cluster.HelmReleaseStatus), b.(*HelmReleaseStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SleepModeConfig)(nil), (*cluster.SleepModeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SleepModeConfig_To_cluster_SleepModeConfig(a.(*SleepModeConfig), b.(*cluster.SleepModeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.SleepModeConfig)(nil), (*SleepModeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_SleepModeConfig_To_v1_SleepModeConfig(a.(*cluster.SleepModeConfig), b.(*SleepModeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SleepModeConfigList)(nil), (*cluster.SleepModeConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SleepModeConfigList_To_cluster_SleepModeConfigList(a.(*SleepModeConfigList), b.(*cluster.SleepModeConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.SleepModeConfigList)(nil), (*SleepModeConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_SleepModeConfigList_To_v1_SleepModeConfigList(a.(*cluster.SleepModeConfigList), b.(*SleepModeConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SleepModeConfigSpec)(nil), (*cluster.SleepModeConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SleepModeConfigSpec_To_cluster_SleepModeConfigSpec(a.(*SleepModeConfigSpec), b.(*cluster.SleepModeConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.SleepModeConfigSpec)(nil), (*SleepModeConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_SleepModeConfigSpec_To_v1_SleepModeConfigSpec(a.(*cluster.SleepModeConfigSpec), b.(*SleepModeConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SleepModeConfigStatus)(nil), (*cluster.SleepModeConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SleepModeConfigStatus_To_cluster_SleepModeConfigStatus(a.(*SleepModeConfigStatus), b.(*cluster.SleepModeConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.SleepModeConfigStatus)(nil), (*SleepModeConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_SleepModeConfigStatus_To_v1_SleepModeConfigStatus(a.(*cluster.SleepModeConfigStatus), b.(*SleepModeConfigStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Account_To_cluster_Account(in *Account, out *cluster.Account, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_AccountSpec_To_cluster_AccountSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_AccountStatus_To_cluster_AccountStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Account_To_cluster_Account is an autogenerated conversion function.
func Convert_v1_Account_To_cluster_Account(in *Account, out *cluster.Account, s conversion.Scope) error {
	return autoConvert_v1_Account_To_cluster_Account(in, out, s)
}

func autoConvert_cluster_Account_To_v1_Account(in *cluster.Account, out *Account, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_cluster_AccountSpec_To_v1_AccountSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_cluster_AccountStatus_To_v1_AccountStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_cluster_Account_To_v1_Account is an autogenerated conversion function.
func Convert_cluster_Account_To_v1_Account(in *cluster.Account, out *Account, s conversion.Scope) error {
	return autoConvert_cluster_Account_To_v1_Account(in, out, s)
}

func autoConvert_v1_AccountClusterRoles_To_cluster_AccountClusterRoles(in *AccountClusterRoles, out *cluster.AccountClusterRoles, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.ClusterRoles = *(*[]string)(unsafe.Pointer(&in.ClusterRoles))
	return nil
}

// Convert_v1_AccountClusterRoles_To_cluster_AccountClusterRoles is an autogenerated conversion function.
func Convert_v1_AccountClusterRoles_To_cluster_AccountClusterRoles(in *AccountClusterRoles, out *cluster.AccountClusterRoles, s conversion.Scope) error {
	return autoConvert_v1_AccountClusterRoles_To_cluster_AccountClusterRoles(in, out, s)
}

func autoConvert_cluster_AccountClusterRoles_To_v1_AccountClusterRoles(in *cluster.AccountClusterRoles, out *AccountClusterRoles, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.ClusterRoles = *(*[]string)(unsafe.Pointer(&in.ClusterRoles))
	return nil
}

// Convert_cluster_AccountClusterRoles_To_v1_AccountClusterRoles is an autogenerated conversion function.
func Convert_cluster_AccountClusterRoles_To_v1_AccountClusterRoles(in *cluster.AccountClusterRoles, out *AccountClusterRoles, s conversion.Scope) error {
	return autoConvert_cluster_AccountClusterRoles_To_v1_AccountClusterRoles(in, out, s)
}

func autoConvert_v1_AccountClusterRolesList_To_cluster_AccountClusterRolesList(in *AccountClusterRolesList, out *cluster.AccountClusterRolesList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]cluster.AccountClusterRoles)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_AccountClusterRolesList_To_cluster_AccountClusterRolesList is an autogenerated conversion function.
func Convert_v1_AccountClusterRolesList_To_cluster_AccountClusterRolesList(in *AccountClusterRolesList, out *cluster.AccountClusterRolesList, s conversion.Scope) error {
	return autoConvert_v1_AccountClusterRolesList_To_cluster_AccountClusterRolesList(in, out, s)
}

func autoConvert_cluster_AccountClusterRolesList_To_v1_AccountClusterRolesList(in *cluster.AccountClusterRolesList, out *AccountClusterRolesList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]AccountClusterRoles)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_cluster_AccountClusterRolesList_To_v1_AccountClusterRolesList is an autogenerated conversion function.
func Convert_cluster_AccountClusterRolesList_To_v1_AccountClusterRolesList(in *cluster.AccountClusterRolesList, out *AccountClusterRolesList, s conversion.Scope) error {
	return autoConvert_cluster_AccountClusterRolesList_To_v1_AccountClusterRolesList(in, out, s)
}

func autoConvert_v1_AccountList_To_cluster_AccountList(in *AccountList, out *cluster.AccountList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]cluster.Account)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_AccountList_To_cluster_AccountList is an autogenerated conversion function.
func Convert_v1_AccountList_To_cluster_AccountList(in *AccountList, out *cluster.AccountList, s conversion.Scope) error {
	return autoConvert_v1_AccountList_To_cluster_AccountList(in, out, s)
}

func autoConvert_cluster_AccountList_To_v1_AccountList(in *cluster.AccountList, out *AccountList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Account)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_cluster_AccountList_To_v1_AccountList is an autogenerated conversion function.
func Convert_cluster_AccountList_To_v1_AccountList(in *cluster.AccountList, out *AccountList, s conversion.Scope) error {
	return autoConvert_cluster_AccountList_To_v1_AccountList(in, out, s)
}

func autoConvert_v1_AccountSpec_To_cluster_AccountSpec(in *AccountSpec, out *cluster.AccountSpec, s conversion.Scope) error {
	out.AccountSpec = in.AccountSpec
	return nil
}

// Convert_v1_AccountSpec_To_cluster_AccountSpec is an autogenerated conversion function.
func Convert_v1_AccountSpec_To_cluster_AccountSpec(in *AccountSpec, out *cluster.AccountSpec, s conversion.Scope) error {
	return autoConvert_v1_AccountSpec_To_cluster_AccountSpec(in, out, s)
}

func autoConvert_cluster_AccountSpec_To_v1_AccountSpec(in *cluster.AccountSpec, out *AccountSpec, s conversion.Scope) error {
	out.AccountSpec = in.AccountSpec
	return nil
}

// Convert_cluster_AccountSpec_To_v1_AccountSpec is an autogenerated conversion function.
func Convert_cluster_AccountSpec_To_v1_AccountSpec(in *cluster.AccountSpec, out *AccountSpec, s conversion.Scope) error {
	return autoConvert_cluster_AccountSpec_To_v1_AccountSpec(in, out, s)
}

func autoConvert_v1_AccountStatus_To_cluster_AccountStatus(in *AccountStatus, out *cluster.AccountStatus, s conversion.Scope) error {
	out.AccountStatus = in.AccountStatus
	return nil
}

// Convert_v1_AccountStatus_To_cluster_AccountStatus is an autogenerated conversion function.
func Convert_v1_AccountStatus_To_cluster_AccountStatus(in *AccountStatus, out *cluster.AccountStatus, s conversion.Scope) error {
	return autoConvert_v1_AccountStatus_To_cluster_AccountStatus(in, out, s)
}

func autoConvert_cluster_AccountStatus_To_v1_AccountStatus(in *cluster.AccountStatus, out *AccountStatus, s conversion.Scope) error {
	out.AccountStatus = in.AccountStatus
	return nil
}

// Convert_cluster_AccountStatus_To_v1_AccountStatus is an autogenerated conversion function.
func Convert_cluster_AccountStatus_To_v1_AccountStatus(in *cluster.AccountStatus, out *AccountStatus, s conversion.Scope) error {
	return autoConvert_cluster_AccountStatus_To_v1_AccountStatus(in, out, s)
}

func autoConvert_v1_Context_To_cluster_Context(in *Context, out *cluster.Context, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ContextSpec_To_cluster_ContextSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ContextStatus_To_cluster_ContextStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Context_To_cluster_Context is an autogenerated conversion function.
func Convert_v1_Context_To_cluster_Context(in *Context, out *cluster.Context, s conversion.Scope) error {
	return autoConvert_v1_Context_To_cluster_Context(in, out, s)
}

func autoConvert_cluster_Context_To_v1_Context(in *cluster.Context, out *Context, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_cluster_ContextSpec_To_v1_ContextSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_cluster_ContextStatus_To_v1_ContextStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_cluster_Context_To_v1_Context is an autogenerated conversion function.
func Convert_cluster_Context_To_v1_Context(in *cluster.Context, out *Context, s conversion.Scope) error {
	return autoConvert_cluster_Context_To_v1_Context(in, out, s)
}

func autoConvert_v1_ContextList_To_cluster_ContextList(in *ContextList, out *cluster.ContextList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]cluster.Context)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ContextList_To_cluster_ContextList is an autogenerated conversion function.
func Convert_v1_ContextList_To_cluster_ContextList(in *ContextList, out *cluster.ContextList, s conversion.Scope) error {
	return autoConvert_v1_ContextList_To_cluster_ContextList(in, out, s)
}

func autoConvert_cluster_ContextList_To_v1_ContextList(in *cluster.ContextList, out *ContextList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Context)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_cluster_ContextList_To_v1_ContextList is an autogenerated conversion function.
func Convert_cluster_ContextList_To_v1_ContextList(in *cluster.ContextList, out *ContextList, s conversion.Scope) error {
	return autoConvert_cluster_ContextList_To_v1_ContextList(in, out, s)
}

func autoConvert_v1_ContextSpec_To_cluster_ContextSpec(in *ContextSpec, out *cluster.ContextSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1_ContextSpec_To_cluster_ContextSpec is an autogenerated conversion function.
func Convert_v1_ContextSpec_To_cluster_ContextSpec(in *ContextSpec, out *cluster.ContextSpec, s conversion.Scope) error {
	return autoConvert_v1_ContextSpec_To_cluster_ContextSpec(in, out, s)
}

func autoConvert_cluster_ContextSpec_To_v1_ContextSpec(in *cluster.ContextSpec, out *ContextSpec, s conversion.Scope) error {
	return nil
}

// Convert_cluster_ContextSpec_To_v1_ContextSpec is an autogenerated conversion function.
func Convert_cluster_ContextSpec_To_v1_ContextSpec(in *cluster.ContextSpec, out *ContextSpec, s conversion.Scope) error {
	return autoConvert_cluster_ContextSpec_To_v1_ContextSpec(in, out, s)
}

func autoConvert_v1_ContextStatus_To_cluster_ContextStatus(in *ContextStatus, out *cluster.ContextStatus, s conversion.Scope) error {
	out.Cluster = in.Cluster
	out.User = in.User
	out.Teams = *(*[]storagev1.Team)(unsafe.Pointer(&in.Teams))
	return nil
}

// Convert_v1_ContextStatus_To_cluster_ContextStatus is an autogenerated conversion function.
func Convert_v1_ContextStatus_To_cluster_ContextStatus(in *ContextStatus, out *cluster.ContextStatus, s conversion.Scope) error {
	return autoConvert_v1_ContextStatus_To_cluster_ContextStatus(in, out, s)
}

func autoConvert_cluster_ContextStatus_To_v1_ContextStatus(in *cluster.ContextStatus, out *ContextStatus, s conversion.Scope) error {
	out.Cluster = in.Cluster
	out.User = in.User
	out.Teams = *(*[]storagev1.Team)(unsafe.Pointer(&in.Teams))
	return nil
}

// Convert_cluster_ContextStatus_To_v1_ContextStatus is an autogenerated conversion function.
func Convert_cluster_ContextStatus_To_v1_ContextStatus(in *cluster.ContextStatus, out *ContextStatus, s conversion.Scope) error {
	return autoConvert_cluster_ContextStatus_To_v1_ContextStatus(in, out, s)
}

func autoConvert_v1_EpochInfo_To_cluster_EpochInfo(in *EpochInfo, out *cluster.EpochInfo, s conversion.Scope) error {
	out.Start = in.Start
	out.Slept = in.Slept
	return nil
}

// Convert_v1_EpochInfo_To_cluster_EpochInfo is an autogenerated conversion function.
func Convert_v1_EpochInfo_To_cluster_EpochInfo(in *EpochInfo, out *cluster.EpochInfo, s conversion.Scope) error {
	return autoConvert_v1_EpochInfo_To_cluster_EpochInfo(in, out, s)
}

func autoConvert_cluster_EpochInfo_To_v1_EpochInfo(in *cluster.EpochInfo, out *EpochInfo, s conversion.Scope) error {
	out.Start = in.Start
	out.Slept = in.Slept
	return nil
}

// Convert_cluster_EpochInfo_To_v1_EpochInfo is an autogenerated conversion function.
func Convert_cluster_EpochInfo_To_v1_EpochInfo(in *cluster.EpochInfo, out *EpochInfo, s conversion.Scope) error {
	return autoConvert_cluster_EpochInfo_To_v1_EpochInfo(in, out, s)
}

func autoConvert_v1_HelmRelease_To_cluster_HelmRelease(in *HelmRelease, out *cluster.HelmRelease, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_HelmReleaseSpec_To_cluster_HelmReleaseSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_HelmReleaseStatus_To_cluster_HelmReleaseStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_HelmRelease_To_cluster_HelmRelease is an autogenerated conversion function.
func Convert_v1_HelmRelease_To_cluster_HelmRelease(in *HelmRelease, out *cluster.HelmRelease, s conversion.Scope) error {
	return autoConvert_v1_HelmRelease_To_cluster_HelmRelease(in, out, s)
}

func autoConvert_cluster_HelmRelease_To_v1_HelmRelease(in *cluster.HelmRelease, out *HelmRelease, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_cluster_HelmReleaseSpec_To_v1_HelmReleaseSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_cluster_HelmReleaseStatus_To_v1_HelmReleaseStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_cluster_HelmRelease_To_v1_HelmRelease is an autogenerated conversion function.
func Convert_cluster_HelmRelease_To_v1_HelmRelease(in *cluster.HelmRelease, out *HelmRelease, s conversion.Scope) error {
	return autoConvert_cluster_HelmRelease_To_v1_HelmRelease(in, out, s)
}

func autoConvert_v1_HelmReleaseList_To_cluster_HelmReleaseList(in *HelmReleaseList, out *cluster.HelmReleaseList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]cluster.HelmRelease)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_HelmReleaseList_To_cluster_HelmReleaseList is an autogenerated conversion function.
func Convert_v1_HelmReleaseList_To_cluster_HelmReleaseList(in *HelmReleaseList, out *cluster.HelmReleaseList, s conversion.Scope) error {
	return autoConvert_v1_HelmReleaseList_To_cluster_HelmReleaseList(in, out, s)
}

func autoConvert_cluster_HelmReleaseList_To_v1_HelmReleaseList(in *cluster.HelmReleaseList, out *HelmReleaseList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]HelmRelease)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_cluster_HelmReleaseList_To_v1_HelmReleaseList is an autogenerated conversion function.
func Convert_cluster_HelmReleaseList_To_v1_HelmReleaseList(in *cluster.HelmReleaseList, out *HelmReleaseList, s conversion.Scope) error {
	return autoConvert_cluster_HelmReleaseList_To_v1_HelmReleaseList(in, out, s)
}

func autoConvert_v1_HelmReleaseSpec_To_cluster_HelmReleaseSpec(in *HelmReleaseSpec, out *cluster.HelmReleaseSpec, s conversion.Scope) error {
	out.Chart = in.Chart
	out.Config = in.Config
	out.InsecureSkipTlsVerify = in.InsecureSkipTlsVerify
	return nil
}

// Convert_v1_HelmReleaseSpec_To_cluster_HelmReleaseSpec is an autogenerated conversion function.
func Convert_v1_HelmReleaseSpec_To_cluster_HelmReleaseSpec(in *HelmReleaseSpec, out *cluster.HelmReleaseSpec, s conversion.Scope) error {
	return autoConvert_v1_HelmReleaseSpec_To_cluster_HelmReleaseSpec(in, out, s)
}

func autoConvert_cluster_HelmReleaseSpec_To_v1_HelmReleaseSpec(in *cluster.HelmReleaseSpec, out *HelmReleaseSpec, s conversion.Scope) error {
	out.Chart = in.Chart
	out.Config = in.Config
	out.InsecureSkipTlsVerify = in.InsecureSkipTlsVerify
	return nil
}

// Convert_cluster_HelmReleaseSpec_To_v1_HelmReleaseSpec is an autogenerated conversion function.
func Convert_cluster_HelmReleaseSpec_To_v1_HelmReleaseSpec(in *cluster.HelmReleaseSpec, out *HelmReleaseSpec, s conversion.Scope) error {
	return autoConvert_cluster_HelmReleaseSpec_To_v1_HelmReleaseSpec(in, out, s)
}

func autoConvert_v1_HelmReleaseStatus_To_cluster_HelmReleaseStatus(in *HelmReleaseStatus, out *cluster.HelmReleaseStatus, s conversion.Scope) error {
	out.Revision = in.Revision
	out.Info = (*storagev1.Info)(unsafe.Pointer(in.Info))
	out.Metadata = (*storagev1.Metadata)(unsafe.Pointer(in.Metadata))
	return nil
}

// Convert_v1_HelmReleaseStatus_To_cluster_HelmReleaseStatus is an autogenerated conversion function.
func Convert_v1_HelmReleaseStatus_To_cluster_HelmReleaseStatus(in *HelmReleaseStatus, out *cluster.HelmReleaseStatus, s conversion.Scope) error {
	return autoConvert_v1_HelmReleaseStatus_To_cluster_HelmReleaseStatus(in, out, s)
}

func autoConvert_cluster_HelmReleaseStatus_To_v1_HelmReleaseStatus(in *cluster.HelmReleaseStatus, out *HelmReleaseStatus, s conversion.Scope) error {
	out.Revision = in.Revision
	out.Info = (*storagev1.Info)(unsafe.Pointer(in.Info))
	out.Metadata = (*storagev1.Metadata)(unsafe.Pointer(in.Metadata))
	return nil
}

// Convert_cluster_HelmReleaseStatus_To_v1_HelmReleaseStatus is an autogenerated conversion function.
func Convert_cluster_HelmReleaseStatus_To_v1_HelmReleaseStatus(in *cluster.HelmReleaseStatus, out *HelmReleaseStatus, s conversion.Scope) error {
	return autoConvert_cluster_HelmReleaseStatus_To_v1_HelmReleaseStatus(in, out, s)
}

func autoConvert_v1_SleepModeConfig_To_cluster_SleepModeConfig(in *SleepModeConfig, out *cluster.SleepModeConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_SleepModeConfigSpec_To_cluster_SleepModeConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_SleepModeConfigStatus_To_cluster_SleepModeConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_SleepModeConfig_To_cluster_SleepModeConfig is an autogenerated conversion function.
func Convert_v1_SleepModeConfig_To_cluster_SleepModeConfig(in *SleepModeConfig, out *cluster.SleepModeConfig, s conversion.Scope) error {
	return autoConvert_v1_SleepModeConfig_To_cluster_SleepModeConfig(in, out, s)
}

func autoConvert_cluster_SleepModeConfig_To_v1_SleepModeConfig(in *cluster.SleepModeConfig, out *SleepModeConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_cluster_SleepModeConfigSpec_To_v1_SleepModeConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_cluster_SleepModeConfigStatus_To_v1_SleepModeConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_cluster_SleepModeConfig_To_v1_SleepModeConfig is an autogenerated conversion function.
func Convert_cluster_SleepModeConfig_To_v1_SleepModeConfig(in *cluster.SleepModeConfig, out *SleepModeConfig, s conversion.Scope) error {
	return autoConvert_cluster_SleepModeConfig_To_v1_SleepModeConfig(in, out, s)
}

func autoConvert_v1_SleepModeConfigList_To_cluster_SleepModeConfigList(in *SleepModeConfigList, out *cluster.SleepModeConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]cluster.SleepModeConfig)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_SleepModeConfigList_To_cluster_SleepModeConfigList is an autogenerated conversion function.
func Convert_v1_SleepModeConfigList_To_cluster_SleepModeConfigList(in *SleepModeConfigList, out *cluster.SleepModeConfigList, s conversion.Scope) error {
	return autoConvert_v1_SleepModeConfigList_To_cluster_SleepModeConfigList(in, out, s)
}

func autoConvert_cluster_SleepModeConfigList_To_v1_SleepModeConfigList(in *cluster.SleepModeConfigList, out *SleepModeConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]SleepModeConfig)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_cluster_SleepModeConfigList_To_v1_SleepModeConfigList is an autogenerated conversion function.
func Convert_cluster_SleepModeConfigList_To_v1_SleepModeConfigList(in *cluster.SleepModeConfigList, out *SleepModeConfigList, s conversion.Scope) error {
	return autoConvert_cluster_SleepModeConfigList_To_v1_SleepModeConfigList(in, out, s)
}

func autoConvert_v1_SleepModeConfigSpec_To_cluster_SleepModeConfigSpec(in *SleepModeConfigSpec, out *cluster.SleepModeConfigSpec, s conversion.Scope) error {
	out.ForceSleep = in.ForceSleep
	out.ForceSleepDuration = (*int64)(unsafe.Pointer(in.ForceSleepDuration))
	out.DeleteAllPods = in.DeleteAllPods
	out.DeleteAfter = in.DeleteAfter
	out.SleepAfter = in.SleepAfter
	return nil
}

// Convert_v1_SleepModeConfigSpec_To_cluster_SleepModeConfigSpec is an autogenerated conversion function.
func Convert_v1_SleepModeConfigSpec_To_cluster_SleepModeConfigSpec(in *SleepModeConfigSpec, out *cluster.SleepModeConfigSpec, s conversion.Scope) error {
	return autoConvert_v1_SleepModeConfigSpec_To_cluster_SleepModeConfigSpec(in, out, s)
}

func autoConvert_cluster_SleepModeConfigSpec_To_v1_SleepModeConfigSpec(in *cluster.SleepModeConfigSpec, out *SleepModeConfigSpec, s conversion.Scope) error {
	out.ForceSleep = in.ForceSleep
	out.ForceSleepDuration = (*int64)(unsafe.Pointer(in.ForceSleepDuration))
	out.DeleteAllPods = in.DeleteAllPods
	out.DeleteAfter = in.DeleteAfter
	out.SleepAfter = in.SleepAfter
	return nil
}

// Convert_cluster_SleepModeConfigSpec_To_v1_SleepModeConfigSpec is an autogenerated conversion function.
func Convert_cluster_SleepModeConfigSpec_To_v1_SleepModeConfigSpec(in *cluster.SleepModeConfigSpec, out *SleepModeConfigSpec, s conversion.Scope) error {
	return autoConvert_cluster_SleepModeConfigSpec_To_v1_SleepModeConfigSpec(in, out, s)
}

func autoConvert_v1_SleepModeConfigStatus_To_cluster_SleepModeConfigStatus(in *SleepModeConfigStatus, out *cluster.SleepModeConfigStatus, s conversion.Scope) error {
	out.LastActivity = in.LastActivity
	out.SleepingSince = in.SleepingSince
	out.ActiveConnections = in.ActiveConnections
	out.CurrentEpoch = (*cluster.EpochInfo)(unsafe.Pointer(in.CurrentEpoch))
	out.LastEpoch = (*cluster.EpochInfo)(unsafe.Pointer(in.LastEpoch))
	out.SleptLastThirtyDays = (*float64)(unsafe.Pointer(in.SleptLastThirtyDays))
	out.SleptLastSevenDays = (*float64)(unsafe.Pointer(in.SleptLastSevenDays))
	return nil
}

// Convert_v1_SleepModeConfigStatus_To_cluster_SleepModeConfigStatus is an autogenerated conversion function.
func Convert_v1_SleepModeConfigStatus_To_cluster_SleepModeConfigStatus(in *SleepModeConfigStatus, out *cluster.SleepModeConfigStatus, s conversion.Scope) error {
	return autoConvert_v1_SleepModeConfigStatus_To_cluster_SleepModeConfigStatus(in, out, s)
}

func autoConvert_cluster_SleepModeConfigStatus_To_v1_SleepModeConfigStatus(in *cluster.SleepModeConfigStatus, out *SleepModeConfigStatus, s conversion.Scope) error {
	out.LastActivity = in.LastActivity
	out.SleepingSince = in.SleepingSince
	out.ActiveConnections = in.ActiveConnections
	out.CurrentEpoch = (*EpochInfo)(unsafe.Pointer(in.CurrentEpoch))
	out.LastEpoch = (*EpochInfo)(unsafe.Pointer(in.LastEpoch))
	out.SleptLastThirtyDays = (*float64)(unsafe.Pointer(in.SleptLastThirtyDays))
	out.SleptLastSevenDays = (*float64)(unsafe.Pointer(in.SleptLastSevenDays))
	return nil
}

// Convert_cluster_SleepModeConfigStatus_To_v1_SleepModeConfigStatus is an autogenerated conversion function.
func Convert_cluster_SleepModeConfigStatus_To_v1_SleepModeConfigStatus(in *cluster.SleepModeConfigStatus, out *SleepModeConfigStatus, s conversion.Scope) error {
	return autoConvert_cluster_SleepModeConfigStatus_To_v1_SleepModeConfigStatus(in, out, s)
}
