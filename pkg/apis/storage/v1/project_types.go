package v1

import (
	storagev1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ArgoIntegrationSynced storagev1.ConditionType = "ArgoIntegrationSynced"

	ArgoLastAppliedHashAnnotation                = "loft.sh/argo-integration-last-applied-hash"
	ArgoPreviousClusterAnnotation                = "loft.sh/argo-integration-previous-cluster"
	ArgoPreviousNamespaceAnnotation              = "loft.sh/argo-integration-previous-namespace"
	ArgoPreviousVirtualClusterInstanceAnnotation = "loft.sh/argo-integration-previous-virtualclusterinstance"
)

var (
	ProjectConditions []storagev1.ConditionType
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Project
// +k8s:openapi-gen=true
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectSpec   `json:"spec,omitempty"`
	Status ProjectStatus `json:"status,omitempty"`
}

func (a *Project) GetConditions() storagev1.Conditions {
	return a.Status.Conditions
}

func (a *Project) SetConditions(conditions storagev1.Conditions) {
	a.Status.Conditions = conditions
}

func (a *Project) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *Project) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *Project) GetAccess() []Access {
	return a.Spec.Access
}

func (a *Project) SetAccess(access []Access) {
	a.Spec.Access = access
}

type ProjectSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes an app
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Quotas define the quotas inside the project
	// +optional
	Quotas Quotas `json:"quotas,omitempty"`

	// AllowedClusters are target clusters that are allowed to target with
	// environments.
	// +optional
	AllowedClusters []AllowedCluster `json:"allowedClusters,omitempty"`

	// AllowedTemplates are the templates that are allowed to use in this
	// project.
	// +optional
	AllowedTemplates []AllowedTemplate `json:"allowedTemplates,omitempty"`

	// Members are the users and teams that are part of this project
	// +optional
	Members []Member `json:"members,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`

	// NamespacePattern specifies template patterns to use for creating each space or virtual cluster's namespace
	// +optional
	NamespacePattern *NamespacePattern `json:"namespacePattern,omitempty"`

	// ArgoIntegration holds information about ArgoCD Integration
	// +optional
	ArgoIntegration *ArgoIntegrationSpec `json:"argoCD,omitempty"`
}

type NamespacePattern struct {
	// Space holds the namespace pattern to use for space instances
	// +optional
	Space string `json:"space,omitempty"`

	// VirtualCluster holds the namespace pattern to use for virtual cluster instances
	// +optional
	VirtualCluster string `json:"virtualCluster,omitempty"`
}

type Quotas struct {
	// Project holds the quotas for the whole project
	// +optional
	Project map[string]string `json:"project,omitempty"`

	// User holds the quotas per user / team
	User map[string]string `json:"user,omitempty"`
}

var (
	SpaceTemplateKind          = "SpaceTemplate"
	VirtualClusterTemplateKind = "VirtualClusterTemplate"
)

type AllowedTemplate struct {
	// Kind of the template that is allowed. Currently only supports VirtualClusterTemplate & SpaceTemplate
	// +optional
	Kind string `json:"kind,omitempty"`

	// Group of the template that is allowed. Currently only supports storage.loft.sh
	// +optional
	Group string `json:"group,omitempty"`

	// Name of the template
	// +optional
	Name string `json:"name,omitempty"`

	// IsDefault specifies if the template should be used as a default
	// +optional
	IsDefault bool `json:"isDefault,omitempty"`
}

type Member struct {
	// Kind is the kind of the member. Currently either User or Team
	// +optional
	Kind string `json:"kind,omitempty"`

	// Group of the member. Currently only supports storage.loft.sh
	// +optional
	Group string `json:"group,omitempty"`

	// Name of the member
	// +optional
	Name string `json:"name,omitempty"`

	// ClusterRole is the assigned role for the above member
	ClusterRole string `json:"clusterRole,omitempty"`
}

type AllowedCluster struct {
	// Name is the name of the cluster that is allowed to create an environment in
	// +optional
	Name string `json:"name,omitempty"`
}

type ProjectStatus struct {
	// Quotas holds the quota status
	// +optional
	Quotas *QuotaStatus `json:"quotas,omitempty"`

	// Conditions holds several conditions the project might be in
	// +optional
	Conditions storagev1.Conditions `json:"conditions,omitempty"`
}

type QuotaStatus struct {
	// Project is the quota status for the whole project
	// +optional
	Project *QuotaStatusProject `json:"project,omitempty"`

	// User is the quota status for each user / team
	// +optional
	User *QuotaStatusUser `json:"user,omitempty"`
}

type QuotaStatusUser struct {
	// Limit is the amount limited per user / team
	// +optional
	Limit map[string]string `json:"limit,omitempty"`

	// Used is the used amount per user / team
	// +optional
	Used QuotaStatusUserUsed `json:"used,omitempty"`
}

type QuotaStatusUserUsed struct {
	// Users is a mapping of users to used resources
	// +optional
	Users map[string]map[string]string `json:"users,omitempty"`

	// Teams is a mapping of teams to used resources
	// +optional
	Teams map[string]map[string]string `json:"teams,omitempty"`
}

type QuotaStatusProject struct {
	// Limit is the amount limited
	// +optional
	Limit map[string]string `json:"limit,omitempty"`

	// Used is the amount currently used
	// +optional
	Used map[string]string `json:"used,omitempty"`
}

type ArgoIntegrationSpec struct {
	// Enabled indicates if the ArgoCD Integration is enabled for the project -- this knob only
	// enables the syncing of virtualclusters, but does not enable SSO integration or project
	// creation (see subsequent spec sections!).
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// Cluster defines the name of the cluster that ArgoCD is deployed into -- if not provided this
	// will default to 'loft-cluster'.
	// +optional
	Cluster string `json:"cluster,omitempty"`
	// VirtualClusterInstance defines the name of *virtual cluster* (instance) that ArgoCD is
	// deployed into. If provided, Cluster will be ignored and Loft will assume that ArgoCD is
	// running in the specified virtual cluster.
	// +optional
	VirtualClusterInstance string `json:"virtualClusterInstance,omitempty"`
	// Namespace defines the namespace in which ArgoCD is running in the cluster.
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// SSO defines single-sign-on related values for the ArgoCD Integration. Enabling SSO will allow
	// users to authenticate to ArgoCD via Loft.
	// +optional
	SSO *ArgoSSOSpec `json:"sso,omitempty"`
	// Project defines project related values for the ArgoCD Integration. Enabling Project
	// integration will cause Loft to generate and manage an ArgoCD appProject that corresponds to
	// the Loft Project.
	// +optional
	Project *ArgoProjectSpec `json:"project,omitempty"`
}

type ArgoSSOSpec struct {
	// Enabled indicates if the ArgoCD SSO Integration is enabled for this project. Enabling this
	// will cause Loft to configure SSO authentication via Loft in ArgoCD. If Projects are *not*
	// enabled, all users associated with this Project will be assigned either the 'read-only'
	// (default) role, *or* the roles set under the AssignedRoles field.
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// Host defines the ArgoCD host address that will be used for OIDC authentication between loft
	// and ArgoCD. If not specified OIDC integration will be skipped, but vclusters/spaces will
	// still be synced to ArgoCD.
	// +optional
	Host string `json:"host,omitempty"`
	// AssignedRoles is a list of roles to assign for users who authenticate via Loft -- by default
	// this will be the `read-only` role. If any roles are provided this will override the default
	// setting.
	// +optional
	AssignedRoles []string `json:"assignedRoles,omitempty"`
}

type ArgoProjectSpec struct {
	// Enabled indicates if the ArgoCD Project Integration is enabled for this project. Enabling
	// this will cause Loft to create an appProject in ArgoCD that is associated with the Loft
	// Project. When Project integration is enabled Loft will override the default assigned role
	// set in the SSO integration spec.
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// Metadata defines additional metadata to attach to the loft created project in ArgoCD.
	// +optional
	Metadata ArgoProjectSpecMetadata `json:"metadata,omitempty"`
	// SourceRepos is a list of source repositories to attach/allow on the project, if not specified
	// will be "*" indicating all source repositories.
	// +optional
	SourceRepos []string `json:"sourceRepos,omitempty"`
	// Roles is a list of roles that should be attached to the ArgoCD project. If roles are provided
	// no loft default roles will be set. If no roles are provided *and* SSO is enabled, loft will
	// configure sane default values.
	// +optional
	Roles []ArgoProjectRole `json:"roles,omitempty"`
}

type ArgoProjectSpecMetadata struct {
	// ExtraAnnotations are optional annotations that can be attached to the project in ArgoCD.
	// +optional
	ExtraAnnotations map[string]string `json:"extraAnnotations,omitempty"`
	// ExtraLabels are optional labels that can be attached to the project in ArgoCD.
	// +optional
	ExtraLabels map[string]string `json:"extraLabels,omitempty"`
	// Description to add to the ArgoCD project.
	// +optional
	Description string `json:"description,omitempty"`
}

type ArgoProjectRole struct {
	// Name of the ArgoCD role to attach to the project.
	Name string `json:"name,omitempty"`
	// Description to add to the ArgoCD project.
	// +optional
	Description string `json:"description,omitempty"`
	// Rules ist a list of policy rules to attach to the role.
	Rules []ArgoProjectPolicyRule `json:"rules,omitempty"`
	// Groups is a list of OIDC group names to bind to the role.
	Groups []string `json:"groups,omitempty"`
}

type ArgoProjectPolicyRule struct {
	// Action is one of "*", "get", "create", "update", "delete", "sync", or "override".
	// +optional
	Action string `json:"action,omitempty"`
	// Application is the ArgoCD project/repository to apply the rule to.
	// +optional
	Application string `json:"application,omitempty"`
	// Allow applies the "allow" permission to the rule, if allow is not set, the permission will
	// always be set to "deny".
	// +optional
	Allow bool `json:"permission,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectList contains a list of Project objects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}