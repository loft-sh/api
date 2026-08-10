package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TenantLabel is the canonical label key binding a resource to a Tenant.
// Set on metadata.labels by the management apiserver REST handler on
// Create; Admin users can modify, but invisible and immutable to tenants
// after creation. Absent = operator-global.
const TenantLabel = "tenant.platform.vcluster.com/name"

// TenantImpersonationExtra is the user.Info.Extra key carrying the
// impersonated tenant override. Travels on the wire as
// Impersonate-Extra-tenant (case-insensitive per HTTP; the apiserver
// impersonation filter lowercases on read). Kubernetes' built-in
// impersonation filter authorizes the impersonate verb on
// userextras/tenant before our handlers see the request.
const TenantImpersonationExtra = "tenant"

// TenantConditionNICoOnboarded is set on a Tenant that opts into NICo (via the
// nico.vcluster.com/org annotation) once its NICo tenant org has been
// materialized. It is False with a reason while onboarding is pending or
// failing.
const TenantConditionNICoOnboarded agentstoragev1.ConditionType = "NICoOnboarded"

// TenantConditionNICoNodeProviderFound reports whether the Tenant's NICo
// NodeProvider could be resolved. It is False when the referenced NodeProvider
// is missing, is not a NICo provider, or when the Tenant resolves to more than
// one NICo provider.
const TenantConditionNICoNodeProviderFound agentstoragev1.ConditionType = "NICoNodeProviderFound"

// TenantConditionNICoSiteCredentialsAvailable reports whether the Tenant's
// credentials for the NICo endpoint could be loaded. False carries the reason.
const TenantConditionNICoSiteCredentialsAvailable agentstoragev1.ConditionType = "NICoSiteCredentialsAvailable"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Tenant is a customer-scoped envelope sitting between Global and
// Project. It is optional: installs with no Tenant objects behave
// exactly as today.
// +k8s:openapi-gen=true
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TenantSpec   `json:"spec,omitempty"`
	Status TenantStatus `json:"status,omitempty"`
}

func (a *Tenant) GetConditions() agentstoragev1.Conditions {
	return a.Status.Conditions
}

func (a *Tenant) SetConditions(conditions agentstoragev1.Conditions) {
	a.Status.Conditions = conditions
}

func (a *Tenant) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *Tenant) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *Tenant) GetAccess() []Access {
	return a.Spec.Access
}

func (a *Tenant) SetAccess(access []Access) {
	a.Spec.Access = access
}

type TenantSpec struct {
	// DisplayName is the name that should be displayed in the UI.
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes this Tenant.
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this Tenant. Access carries RBAC for the
	// Tenant resource itself (transformed to Roles and RoleBindings),
	// governing operator-side delegation — which Platform Operator users
	// may read or edit this Tenant. Tenant membership for humans is
	// carried as a loft:tenant:<name> group claim on User/Team, not by
	// Access.
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Access holds the access rights for users and teams on the Tenant CR
	// itself.
	// +optional
	Access []Access `json:"access,omitempty"`

	// Hosts are the hostnames that map to this tenant. Used for SSO
	// bootstrap, UI branding, and per-request tenant resolution.
	// +optional
	Hosts []HostBinding `json:"hosts,omitempty"`

	// Authentication holds per-tenant SSO connector configuration. Mirrors
	// Config.Status.Authentication — same Go type — so the platform's
	// existing connector machinery can dispatch from a Tenant's value.
	// +optional
	Authentication *Authentication `json:"auth,omitempty"`

	// ResourceAllowances controls, per management.loft.sh kind, which
	// admin-owned resources this tenant may see and use, and how (see
	// ScopeMode). Each entry overrides the platform's shipped default for the
	// matched instances. Multiple entries per resource are allowed (e.g. a
	// co-held pool plus one leased instance); a name-matched entry wins over a
	// kind-wide (no-resourceNames) entry. A resource with no entry falls to
	// the shipped per-kind default, then to the catch-all. ScopeUnscoped is
	// status-only and is rejected here.
	// +optional
	ResourceAllowances []ResourceAllowance `json:"resourceAllowances,omitempty"`

	// ResourceQuotas caps how many of a resource this tenant may hold or
	// consume, aggregated across all the tenant's projects. Parity with
	// Project quotas, not a replacement (Projects keep their per-project
	// quotas; the tenant quota is an outer bound).
	// +optional
	ResourceQuotas []ResourceQuota `json:"resourceQuotas,omitempty"`
}

// ScopeMode is the per-tenant treatment of a management.loft.sh kind's
// admin-owned (unlabeled) instances. A tenant's own-labeled instances are
// always read-write, and another tenant's instances are always hidden,
// independent of scope.
type ScopeMode string

const (
	// ScopeDisabled denies the kind to the tenant at the authorizer:
	// `kubectl auth can-i` returns allowed:false and all verbs are Forbidden.
	ScopeDisabled ScopeMode = "disabled"
	// ScopeOwned shows only the tenant's own-labeled instances (read-write);
	// admin-owned instances are hidden.
	ScopeOwned ScopeMode = "owned"
	// ScopeGranted additionally shows the matched admin-owned instances
	// read-only and usable, co-held across tenants. The §3 Granted Resource
	// pattern.
	ScopeGranted ScopeMode = "granted"
	// ScopeLeased is ScopeGranted with cross-Tenant exclusivity: at most one
	// tenant may hold a given instance. The §3 Leased Resource pattern.
	ScopeLeased ScopeMode = "leased"
	// ScopeUnscoped means the kind is not tenant-scoped: admin-owned instances
	// follow plain RBAC. It is the catch-all default and is reported in
	// status; it is NOT valid in Spec.ResourceAllowances.
	ScopeUnscoped ScopeMode = "unscoped"
)

// ResourceAllowance scopes one management.loft.sh kind (optionally narrowed to
// specific instances) for a tenant. Used in Spec (operator intent) and in
// Status (resolved effective scope).
type ResourceAllowance struct {
	// Resource is the lowercase plural name of a management.loft.sh resource
	// (e.g. "projects", "clusters", "virtualclustertemplates"). The group is
	// always management.loft.sh, so there is no apiGroup field.
	Resource string `json:"resource"`

	// Scope is the treatment applied to this resource for the tenant.
	Scope ScopeMode `json:"scope,omitempty"`

	// ResourceNames narrows the entry to specific admin-owned instances by
	// name. Empty or ["*"] applies to the whole kind. Only meaningful for
	// granted/leased.
	// +optional
	ResourceNames []string `json:"resourceNames,omitempty"`
}

// ResourceQuota caps consumption of one management.loft.sh resource for a
// tenant. Keys in the Tenant/User maps are conditions relative to the resource
// — "total", "active", "!active", "template=<name>", "type=<name>",
// "provider=<name>" — and values are integer counts.
type ResourceQuota struct {
	// Resource is the lowercase plural name of the counted management.loft.sh
	// resource (e.g. "virtualclusterinstances", "nodeclaims").
	Resource string `json:"resource"`

	// Tenant caps usage aggregated across all the tenant's projects.
	// +optional
	Tenant map[string]string `json:"tenant,omitempty"`

	// User caps usage per individual user or team.
	// +optional
	User map[string]string `json:"user,omitempty"`
}

// HostBinding binds a hostname to this Tenant for routing and SSO
// resolution.
type HostBinding struct {
	// Hostname is the DNS name the platform will treat as belonging to
	// this Tenant (e.g. acme.platform.example.com).
	Hostname string `json:"hostname"`
}

// TenantStatus surfaces reconciler-managed state.
type TenantStatus struct {
	// Conditions describes the current observed conditions of the Tenant.
	// +optional
	Conditions agentstoragev1.Conditions `json:"conditions,omitempty"`

	// ObservedGeneration is the generation last observed by the
	// reconciler.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// ResourceAllowances is the resolved effective scope for every scopable
	// management.loft.sh kind for this tenant: the shipped per-kind defaults
	// merged with Spec.ResourceAllowances. Each entry's Scope is the effective
	// scope and may be ScopeUnscoped. Recomputed each reconcile from discovery
	// + defaults + spec; observability only (the apiserver computes effective
	// scope live at request time). The leased entries are the per-Tenant
	// projection of the cross-Tenant exclusivity index.
	// +optional
	ResourceAllowances []ResourceAllowance `json:"resourceAllowances,omitempty"`

	// ResourceQuotas reports usage against the Spec.ResourceQuotas caps,
	// aggregated across the tenant's projects.
	// +optional
	ResourceQuotas []ResourceQuotaStatus `json:"resourceQuotas,omitempty"`

	// NICo reports the NICo tenant org materialized for this Tenant, set once
	// the Tenant opts into NICo via the nico.vcluster.com/org annotation.
	// +optional
	NICo *TenantNICoStatus `json:"nico,omitempty"`
}

// ResourceQuotaStatus reports limit-vs-used for one resource's quota.
type ResourceQuotaStatus struct {
	// Resource is the counted management.loft.sh resource.
	Resource string `json:"resource"`

	// Tenant reports the tenant-aggregate limit and used counts.
	// +optional
	Tenant *QuotaUsage `json:"tenant,omitempty"`

	// User reports the per-user/team limit and used counts.
	// +optional
	User *QuotaUsage `json:"user,omitempty"`
}

// QuotaUsage pairs configured limits with observed usage; both maps are keyed
// by the same condition keys as the corresponding ResourceQuota.
type QuotaUsage struct {
	// Limit echoes the configured caps (condition key -> count).
	// +optional
	Limit map[string]string `json:"limit,omitempty"`

	// Used is the observed usage (condition key -> count).
	// +optional
	Used map[string]string `json:"used,omitempty"`
}

// TenantNICoStatus reports the NICo tenant org materialized for a Tenant and the
// NICo provider it is bound to, and is set by the platform. A Tenant is bound to
// a single NICo provider, and so to a single NICo site.
type TenantNICoStatus struct {
	// Org is the NICo tenant org this Tenant maps to.
	// +optional
	Org string `json:"org,omitempty"`

	// TenantID is the id of the materialized NICo tenant.
	// +optional
	TenantID string `json:"tenantId,omitempty"`

	// NodeProvider is the platform NodeProvider this Tenant uses.
	// +optional
	NodeProvider string `json:"nodeProvider,omitempty"`

	// Endpoint is the NICo REST API endpoint taken from that NodeProvider.
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// ProviderOrg is the NICo provider organization.
	// +optional
	ProviderOrg string `json:"providerOrg,omitempty"`

	// InsecureSkipTLSVerify reports whether TLS verification is skipped when
	// reaching the NICo endpoint.
	// +optional
	InsecureSkipTLSVerify bool `json:"insecureSkipTLSVerify,omitempty"`

	// TenantAccountID is the id of this Tenant's NICo TenantAccount.
	// +optional
	TenantAccountID string `json:"tenantAccountId,omitempty"`

	// TenantAccountManaged reports whether the platform owns the NICo
	// TenantAccount and removes it during cleanup.
	// +optional
	TenantAccountManaged bool `json:"tenantAccountManaged,omitempty"`

	// IPBlockID is the id of the tenant-scoped NICo IPBlock from which VPC
	// prefixes are carved. Set by the tenant onboarding flow (ENGNODE-604).
	// +optional
	IPBlockID string `json:"ipBlockId,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TenantList contains a list of Tenant objects.
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tenant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tenant{}, &TenantList{})
}
