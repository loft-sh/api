// Code generated by generator. DO NOT EDIT.

package cluster

import (
	"context"
	"fmt"

	storagev1 "github.com/loft-sh/api/pkg/apis/storage/v1"
	"github.com/loft-sh/api/pkg/managerfactory"
	"github.com/loft-sh/apiserver/pkg/builders"
	configv1alpha1 "github.com/loft-sh/kiosk/pkg/apis/config/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
)

type NewRESTFunc func(factory managerfactory.SharedManagerFactory) rest.Storage

var (
	ClusterAccountStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalAccount,
		func() runtime.Object { return &Account{} },     // Register versioned resource
		func() runtime.Object { return &AccountList{} }, // Register versioned resource list
		NewAccountREST,
	)
	NewAccountREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewAccountRESTFunc(Factory)
	}
	NewAccountRESTFunc    NewRESTFunc
	ClusterContextStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalContext,
		func() runtime.Object { return &Context{} },     // Register versioned resource
		func() runtime.Object { return &ContextList{} }, // Register versioned resource list
		NewContextREST,
	)
	NewContextREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewContextRESTFunc(Factory)
	}
	NewContextRESTFunc        NewRESTFunc
	ClusterHelmReleaseStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalHelmRelease,
		func() runtime.Object { return &HelmRelease{} },     // Register versioned resource
		func() runtime.Object { return &HelmReleaseList{} }, // Register versioned resource list
		NewHelmReleaseREST,
	)
	NewHelmReleaseREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewHelmReleaseRESTFunc(Factory)
	}
	NewHelmReleaseRESTFunc        NewRESTFunc
	ClusterSleepModeConfigStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalSleepModeConfig,
		func() runtime.Object { return &SleepModeConfig{} },     // Register versioned resource
		func() runtime.Object { return &SleepModeConfigList{} }, // Register versioned resource list
		NewSleepModeConfigREST,
	)
	NewSleepModeConfigREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewSleepModeConfigRESTFunc(Factory)
	}
	NewSleepModeConfigRESTFunc NewRESTFunc
	InternalAccount            = builders.NewInternalResource(
		"accounts",
		"Account",
		func() runtime.Object { return &Account{} },
		func() runtime.Object { return &AccountList{} },
	)
	InternalAccountStatus = builders.NewInternalResourceStatus(
		"accounts",
		"AccountStatus",
		func() runtime.Object { return &Account{} },
		func() runtime.Object { return &AccountList{} },
	)
	InternalAccountClusterRolesREST = builders.NewInternalSubresource(
		"accounts", "AccountClusterRoles", "clusterroles",
		func() runtime.Object { return &AccountClusterRoles{} },
	)
	NewAccountClusterRolesREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewAccountClusterRolesRESTFunc(Factory)
	}
	NewAccountClusterRolesRESTFunc NewRESTFunc
	InternalContext                = builders.NewInternalResource(
		"contexts",
		"Context",
		func() runtime.Object { return &Context{} },
		func() runtime.Object { return &ContextList{} },
	)
	InternalContextStatus = builders.NewInternalResourceStatus(
		"contexts",
		"ContextStatus",
		func() runtime.Object { return &Context{} },
		func() runtime.Object { return &ContextList{} },
	)
	InternalHelmRelease = builders.NewInternalResource(
		"helmreleases",
		"HelmRelease",
		func() runtime.Object { return &HelmRelease{} },
		func() runtime.Object { return &HelmReleaseList{} },
	)
	InternalHelmReleaseStatus = builders.NewInternalResourceStatus(
		"helmreleases",
		"HelmReleaseStatus",
		func() runtime.Object { return &HelmRelease{} },
		func() runtime.Object { return &HelmReleaseList{} },
	)
	InternalSleepModeConfig = builders.NewInternalResource(
		"sleepmodeconfigs",
		"SleepModeConfig",
		func() runtime.Object { return &SleepModeConfig{} },
		func() runtime.Object { return &SleepModeConfigList{} },
	)
	InternalSleepModeConfigStatus = builders.NewInternalResourceStatus(
		"sleepmodeconfigs",
		"SleepModeConfigStatus",
		func() runtime.Object { return &SleepModeConfig{} },
		func() runtime.Object { return &SleepModeConfigList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("cluster.loft.sh").WithKinds(
		InternalAccount,
		InternalAccountStatus,
		InternalAccountClusterRolesREST,
		InternalContext,
		InternalContextStatus,
		InternalHelmRelease,
		InternalHelmReleaseStatus,
		InternalSleepModeConfig,
		InternalSleepModeConfigStatus,
	)

	// Required by code generated by go2idl
	AddToScheme = (&runtime.SchemeBuilder{
		ApiVersion.SchemeBuilder.AddToScheme,
		RegisterDefaults,
	}).AddToScheme
	SchemeBuilder      = ApiVersion.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Account struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   AccountSpec
	Status AccountStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AccountClusterRoles struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	ClusterRoles []string
}

type AccountSpec struct {
	configv1alpha1.AccountSpec
}

type AccountStatus struct {
	configv1alpha1.AccountStatus
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Context struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ContextSpec
	Status ContextStatus
}

type ContextSpec struct {
}

type ContextStatus struct {
	Cluster storagev1.Cluster
	User    storagev1.User
	Teams   []storagev1.Team
}

type EpochInfo struct {
	Start int64
	Slept int64
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HelmRelease struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   HelmReleaseSpec
	Status HelmReleaseStatus
}

type HelmReleaseSpec struct {
	Chart                 storagev1.Chart
	Config                string
	InsecureSkipTlsVerify bool
}

type HelmReleaseStatus struct {
	Revision int
	Info     *storagev1.Info
	Metadata *storagev1.Metadata
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SleepModeConfig struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   SleepModeConfigSpec
	Status SleepModeConfigStatus
}

type SleepModeConfigSpec struct {
	ForceSleep         bool
	ForceSleepDuration *int64
	DeleteAllPods      bool
	DeleteAfter        int64
	SleepAfter         int64
}

type SleepModeConfigStatus struct {
	LastActivity        int64
	SleepingSince       int64
	ActiveConnections   int64
	CurrentEpoch        *EpochInfo
	LastEpoch           *EpochInfo
	SleptLastThirtyDays *float64
	SleptLastSevenDays  *float64
}

//
// Account Functions and Structs
//
// +k8s:deepcopy-gen=false
type AccountStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type AccountStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AccountList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Account
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AccountClusterRolesList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []AccountClusterRoles
}

func (Account) NewStatus() interface{} {
	return AccountStatus{}
}

func (pc *Account) GetStatus() interface{} {
	return pc.Status
}

func (pc *Account) SetStatus(s interface{}) {
	pc.Status = s.(AccountStatus)
}

func (pc *Account) GetSpec() interface{} {
	return pc.Spec
}

func (pc *Account) SetSpec(s interface{}) {
	pc.Spec = s.(AccountSpec)
}

func (pc *Account) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *Account) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc Account) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store Account.
// +k8s:deepcopy-gen=false
type AccountRegistry interface {
	ListAccounts(ctx context.Context, options *internalversion.ListOptions) (*AccountList, error)
	GetAccount(ctx context.Context, id string, options *metav1.GetOptions) (*Account, error)
	CreateAccount(ctx context.Context, id *Account) (*Account, error)
	UpdateAccount(ctx context.Context, id *Account) (*Account, error)
	DeleteAccount(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewAccountRegistry(sp builders.StandardStorageProvider) AccountRegistry {
	return &storageAccount{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageAccount struct {
	builders.StandardStorageProvider
}

func (s *storageAccount) ListAccounts(ctx context.Context, options *internalversion.ListOptions) (*AccountList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*AccountList), err
}

func (s *storageAccount) GetAccount(ctx context.Context, id string, options *metav1.GetOptions) (*Account, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*Account), nil
}

func (s *storageAccount) CreateAccount(ctx context.Context, object *Account) (*Account, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*Account), nil
}

func (s *storageAccount) UpdateAccount(ctx context.Context, object *Account) (*Account, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*Account), nil
}

func (s *storageAccount) DeleteAccount(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}

//
// Context Functions and Structs
//
// +k8s:deepcopy-gen=false
type ContextStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type ContextStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ContextList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Context
}

func (Context) NewStatus() interface{} {
	return ContextStatus{}
}

func (pc *Context) GetStatus() interface{} {
	return pc.Status
}

func (pc *Context) SetStatus(s interface{}) {
	pc.Status = s.(ContextStatus)
}

func (pc *Context) GetSpec() interface{} {
	return pc.Spec
}

func (pc *Context) SetSpec(s interface{}) {
	pc.Spec = s.(ContextSpec)
}

func (pc *Context) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *Context) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc Context) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store Context.
// +k8s:deepcopy-gen=false
type ContextRegistry interface {
	ListContexts(ctx context.Context, options *internalversion.ListOptions) (*ContextList, error)
	GetContext(ctx context.Context, id string, options *metav1.GetOptions) (*Context, error)
	CreateContext(ctx context.Context, id *Context) (*Context, error)
	UpdateContext(ctx context.Context, id *Context) (*Context, error)
	DeleteContext(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewContextRegistry(sp builders.StandardStorageProvider) ContextRegistry {
	return &storageContext{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageContext struct {
	builders.StandardStorageProvider
}

func (s *storageContext) ListContexts(ctx context.Context, options *internalversion.ListOptions) (*ContextList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ContextList), err
}

func (s *storageContext) GetContext(ctx context.Context, id string, options *metav1.GetOptions) (*Context, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*Context), nil
}

func (s *storageContext) CreateContext(ctx context.Context, object *Context) (*Context, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*Context), nil
}

func (s *storageContext) UpdateContext(ctx context.Context, object *Context) (*Context, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*Context), nil
}

func (s *storageContext) DeleteContext(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}

//
// HelmRelease Functions and Structs
//
// +k8s:deepcopy-gen=false
type HelmReleaseStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type HelmReleaseStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HelmReleaseList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []HelmRelease
}

func (HelmRelease) NewStatus() interface{} {
	return HelmReleaseStatus{}
}

func (pc *HelmRelease) GetStatus() interface{} {
	return pc.Status
}

func (pc *HelmRelease) SetStatus(s interface{}) {
	pc.Status = s.(HelmReleaseStatus)
}

func (pc *HelmRelease) GetSpec() interface{} {
	return pc.Spec
}

func (pc *HelmRelease) SetSpec(s interface{}) {
	pc.Spec = s.(HelmReleaseSpec)
}

func (pc *HelmRelease) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *HelmRelease) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc HelmRelease) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store HelmRelease.
// +k8s:deepcopy-gen=false
type HelmReleaseRegistry interface {
	ListHelmReleases(ctx context.Context, options *internalversion.ListOptions) (*HelmReleaseList, error)
	GetHelmRelease(ctx context.Context, id string, options *metav1.GetOptions) (*HelmRelease, error)
	CreateHelmRelease(ctx context.Context, id *HelmRelease) (*HelmRelease, error)
	UpdateHelmRelease(ctx context.Context, id *HelmRelease) (*HelmRelease, error)
	DeleteHelmRelease(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewHelmReleaseRegistry(sp builders.StandardStorageProvider) HelmReleaseRegistry {
	return &storageHelmRelease{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageHelmRelease struct {
	builders.StandardStorageProvider
}

func (s *storageHelmRelease) ListHelmReleases(ctx context.Context, options *internalversion.ListOptions) (*HelmReleaseList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*HelmReleaseList), err
}

func (s *storageHelmRelease) GetHelmRelease(ctx context.Context, id string, options *metav1.GetOptions) (*HelmRelease, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*HelmRelease), nil
}

func (s *storageHelmRelease) CreateHelmRelease(ctx context.Context, object *HelmRelease) (*HelmRelease, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*HelmRelease), nil
}

func (s *storageHelmRelease) UpdateHelmRelease(ctx context.Context, object *HelmRelease) (*HelmRelease, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*HelmRelease), nil
}

func (s *storageHelmRelease) DeleteHelmRelease(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}

//
// SleepModeConfig Functions and Structs
//
// +k8s:deepcopy-gen=false
type SleepModeConfigStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type SleepModeConfigStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SleepModeConfigList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []SleepModeConfig
}

func (SleepModeConfig) NewStatus() interface{} {
	return SleepModeConfigStatus{}
}

func (pc *SleepModeConfig) GetStatus() interface{} {
	return pc.Status
}

func (pc *SleepModeConfig) SetStatus(s interface{}) {
	pc.Status = s.(SleepModeConfigStatus)
}

func (pc *SleepModeConfig) GetSpec() interface{} {
	return pc.Spec
}

func (pc *SleepModeConfig) SetSpec(s interface{}) {
	pc.Spec = s.(SleepModeConfigSpec)
}

func (pc *SleepModeConfig) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *SleepModeConfig) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc SleepModeConfig) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store SleepModeConfig.
// +k8s:deepcopy-gen=false
type SleepModeConfigRegistry interface {
	ListSleepModeConfigs(ctx context.Context, options *internalversion.ListOptions) (*SleepModeConfigList, error)
	GetSleepModeConfig(ctx context.Context, id string, options *metav1.GetOptions) (*SleepModeConfig, error)
	CreateSleepModeConfig(ctx context.Context, id *SleepModeConfig) (*SleepModeConfig, error)
	UpdateSleepModeConfig(ctx context.Context, id *SleepModeConfig) (*SleepModeConfig, error)
	DeleteSleepModeConfig(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewSleepModeConfigRegistry(sp builders.StandardStorageProvider) SleepModeConfigRegistry {
	return &storageSleepModeConfig{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageSleepModeConfig struct {
	builders.StandardStorageProvider
}

func (s *storageSleepModeConfig) ListSleepModeConfigs(ctx context.Context, options *internalversion.ListOptions) (*SleepModeConfigList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*SleepModeConfigList), err
}

func (s *storageSleepModeConfig) GetSleepModeConfig(ctx context.Context, id string, options *metav1.GetOptions) (*SleepModeConfig, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*SleepModeConfig), nil
}

func (s *storageSleepModeConfig) CreateSleepModeConfig(ctx context.Context, object *SleepModeConfig) (*SleepModeConfig, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*SleepModeConfig), nil
}

func (s *storageSleepModeConfig) UpdateSleepModeConfig(ctx context.Context, object *SleepModeConfig) (*SleepModeConfig, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*SleepModeConfig), nil
}

func (s *storageSleepModeConfig) DeleteSleepModeConfig(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}
