package v1

import (
	storagev1 "github.com/loft-sh/agentapi/v3/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	InstanceScheduled            storagev1.ConditionType = "Scheduled"
	InstanceTemplateSynced       storagev1.ConditionType = "TemplateSynced"
	InstanceTemplateResolved     storagev1.ConditionType = "TemplateResolved"
	InstanceSpaceSynced          storagev1.ConditionType = "SpaceSynced"
	InstanceSpaceReady           storagev1.ConditionType = "SpaceReady"
	InstanceVirtualClusterSynced storagev1.ConditionType = "VirtualClusterSynced"
	InstanceVirtualClusterReady  storagev1.ConditionType = "VirtualClusterReady"
)

var (
	VirtualClusterConditions = []storagev1.ConditionType{
		InstanceScheduled,
		InstanceTemplateResolved,
		InstanceSpaceSynced,
		InstanceSpaceReady,
		InstanceVirtualClusterSynced,
		InstanceVirtualClusterReady,
	}
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualClusterInstance
// +k8s:openapi-gen=true
type VirtualClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualClusterInstanceSpec   `json:"spec,omitempty"`
	Status VirtualClusterInstanceStatus `json:"status,omitempty"`
}

func (a *VirtualClusterInstance) GetConditions() storagev1.Conditions {
	return a.Status.Conditions
}

func (a *VirtualClusterInstance) SetConditions(conditions storagev1.Conditions) {
	a.Status.Conditions = conditions
}

func (a *VirtualClusterInstance) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *VirtualClusterInstance) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *VirtualClusterInstance) GetAccess() []Access {
	return a.Spec.Access
}

func (a *VirtualClusterInstance) SetAccess(access []Access) {
	a.Spec.Access = access
}

type VirtualClusterInstanceSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes a virtual cluster instance
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// TemplateRef holds the virtual cluster template reference
	// +optional
	TemplateRef *TemplateRef `json:"templateRef,omitempty"`

	// Template is the inline template to use for virtual cluster creation. This is mutually
	// exclusive with templateRef.
	// +optional
	Template *VirtualClusterTemplateDefinition `json:"template,omitempty"`

	// ClusterRef is the reference to the connected cluster holding
	// this virtual cluster
	// +optional
	ClusterRef VirtualClusterClusterRef `json:"clusterRef,omitempty"`

	// Parameters are values to pass to the template
	// +optional
	Parameters string `json:"parameters,omitempty"`

	// ExtraAccessRules defines extra rules which users and teams should have which access to the virtual
	// cluster.
	// +optional
	ExtraAccessRules []storagev1.InstanceAccessRule `json:"extraAccessRules,omitempty"`

	// Access to the virtual cluster object itself
	// +optional
	Access []Access `json:"access,omitempty"`
}

type VirtualClusterInstanceStatus struct {
	// Phase describes the current phase the virtual cluster instance is in
	// +optional
	Phase InstancePhase `json:"phase,omitempty"`

	// Reason describes the reason in machine-readable form why the cluster is in the current
	// phase
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message describes the reason in human-readable form why the cluster is in the current
	// phase
	// +optional
	Message string `json:"message,omitempty"`

	// Conditions holds several conditions the virtual cluster might be in
	// +optional
	Conditions storagev1.Conditions `json:"conditions,omitempty"`

	// VirtualClusterObjects are the objects that were applied within the virtual cluster itself
	// +optional
	VirtualClusterObjects *storagev1.ObjectsStatus `json:"virtualClusterObjects,omitempty"`

	// SpaceObjects are the objects that were applied within the virtual cluster space
	// +optional
	SpaceObjects *storagev1.ObjectsStatus `json:"spaceObjects,omitempty"`

	// VirtualCluster is the template rendered with all the parameters
	// +optional
	VirtualCluster *VirtualClusterTemplateDefinition `json:"virtualCluster,omitempty"`

	// IgnoreReconciliation tells the controller to ignore reconciliation for this instance -- this
	// is primarily used when migrating virtual cluster instances from project to project; this
	// prevents a situation where there are two virtual cluster instances representing the same
	// virtual cluster which could cause issues with concurrent reconciliations of the same object.
	// Once the virtual cluster instance has been cloned and placed into the new project, this
	// (the "old") virtual cluster instance can safely be deleted.
	IgnoreReconciliation bool `json:"ignoreReconciliation,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualClusterInstanceList contains a list of VirtualClusterInstance objects
type VirtualClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualClusterInstance{}, &VirtualClusterInstanceList{})
}