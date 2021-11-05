package v1

import (
	clusterv1 "github.com/loft-sh/agentapi/pkg/apis/loft/cluster/v1"
	storagev1 "github.com/loft-sh/agentapi/pkg/apis/loft/storage/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +k8s:openapi-gen=true
type Task struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TaskSpec   `json:"spec,omitempty"`
	Status TaskStatus `json:"status,omitempty"`
}

// GetConditions returns the set of conditions for this object.
func (a *Task) GetConditions() storagev1.Conditions {
	return a.Status.Conditions
}

// SetConditions sets the conditions on this object.
func (a *Task) SetConditions(conditions storagev1.Conditions) {
	a.Status.Conditions = conditions
}

func (a *Task) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *Task) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *Task) GetAccess() []Access {
	return a.Spec.Access
}

func (a *Task) SetAccess(access []Access) {
	a.Spec.Access = access
}

type TaskSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Target where this task should get executed
	// +optional
	Target Target `json:"target,omitempty"`

	// Task defines the task to execute
	// +optional
	Task TaskDefinition `json:"task,omitempty"`
}

type TaskDefinition struct {
	// ApplyTask executes a kubectl apply
	// +optional
	ApplyTask *ApplyTask `json:"apply,omitempty"`

	// HelmTask executes a helm command
	// +optional
	HelmTask *HelmTask `json:"helm,omitempty"`

	// SpaceCreationTask creates a new space
	// +optional
	SpaceCreationTask *SpaceCreationTask `json:"spaceCreation,omitempty"`

	// VirtualClusterCreatioTask creates a new virtual cluster
	// +optional
	VirtualClusterCreationTask *VirtualClusterCreationTask `json:"virtualClusterCreation,omitempty"`
}

type ApplyTask struct {
	// Manifests are the manifests to apply
	Manifests string `json:"manifests,omitempty"`

	// Args are extra arguments used to apply the manifests
	// +optional
	Args []string `json:"args,omitempty"`
}

type VirtualClusterCreationTask struct {
	// The virtual cluster metadata
	// +optional
	Metadata metav1.ObjectMeta `json:"metadata,omitempty"`

	// The helm release configuration for the virtual cluster.
	// +optional
	HelmRelease *storagev1.VirtualClusterHelmRelease `json:"helmRelease,omitempty"`

	// Wait defines if the task should wait until the virtual cluster is ready
	// +optional
	Wait bool `json:"wait,omitempty"`

	// Apps specifies the apps that should get deployed by this template
	// +optional
	Apps []VirtualClusterCreationAppReference `json:"apps,omitempty"`

	// SpaceCreationTask creates a new space if defined for the virtual cluster
	// +optional
	SpaceCreationTask *SpaceCreationTask `json:"spaceCreation,omitempty"`
}

type VirtualClusterCreationAppReference struct {
	// Name of the target app
	// +optional
	Name string `json:"name,omitempty"`

	// Namespace specifies in which target namespace the app should
	// get deployed in
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// ReleaseName is the name of the app release
	// +optional
	ReleaseName string `json:"releaseName,omitempty"`

	// Parameters are additional helm chart values that will get merged
	// with config and are then used to deploy the helm chart.
	// +optional
	Parameters string `json:"parameters,omitempty"`
}

type SpaceCreationTask struct {
	// Metadata of the space
	// +optional
	Metadata metav1.ObjectMeta `json:"metadata,omitempty"`

	// Owner defines the space owner
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Apps specifies the apps that should get deployed by this template
	// +optional
	Apps []SpaceCreationAppReference `json:"apps,omitempty"`
}

type SpaceCreationAppReference struct {
	// Name of the target app
	// +optional
	Name string `json:"name,omitempty"`

	// ReleaseName of the target app
	// +optional
	ReleaseName string `json:"releaseName,omitempty"`

	// Parameters are additional helm values that are merged
	// with the original helm values.
	// +optional
	Parameters string `json:"parameters,omitempty"`
}

type HelmTask struct {
	// Release holds the release information
	// +optional
	Release HelmTaskRelease `json:"release,omitempty"`

	// StreamContainer can be used to stream a containers logs instead of the helm output.
	// +optional
	StreamContainer *StreamContainer `json:"streamContainer,omitempty"`

	// Type is the task type. Defaults to Upgrade
	// +optional
	Type HelmTaskType `json:"type,omitempty"`

	// RollbackRevision is the revision to rollback to
	// +optional
	RollbackRevision string `json:"rollbackRevision,omitempty"`

	// Args are extra args to pass to helm
	// +optional
	Args []string `json:"args,omitempty"`
}

type HelmTaskRelease struct {
	// Name is the name of the release
	Name string `json:"name,omitempty"`

	// Namespace of the release, if empty will use the target namespace
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Labels are additional labels for the helm release.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Config is the helm config to use to deploy the release
	// +optional
	Config clusterv1.HelmReleaseSpec `json:"config,omitempty"`
}

type StreamContainer struct {
	// Label selector for pods. The newest matching pod will be used to stream logs from
	// +optional
	Selector metav1.LabelSelector `json:"selector" protobuf:"bytes,2,opt,name=selector"`

	// Container is the container name to use
	// +optional
	Container string `json:"container,omitempty"`
}

type TaskStatus struct {
	// Started determines if the task was started
	// +optional
	Started bool `json:"started,omitempty"`

	// Conditions holds several conditions the virtual cluster might be in
	// +optional
	Conditions storagev1.Conditions `json:"conditions,omitempty"`

	// PodPhase describes the phase this task is in
	// +optional
	PodPhase corev1.PodPhase `json:"podPhase,omitempty"`

	// ObservedGeneration is the latest generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// ContainerState describes the container state of the task
	// +optional
	ContainerState *corev1.ContainerStatus `json:"containerState,omitempty"`
}

// Common ConditionTypes used by Cluster API objects.
const (
	// TaskStartedCondition defines the task started condition type that summarizes the operational state of the virtual cluster API object.
	TaskStartedCondition storagev1.ConditionType = "TaskStarted"
)

// HelmTaskType describes the type of a task
type HelmTaskType string

// These are the valid admin account types
const (
	HelmTaskTypeInstall  HelmTaskType = "Install"
	HelmTaskTypeUpgrade  HelmTaskType = "Upgrade"
	HelmTaskTypeDelete   HelmTaskType = "Delete"
	HelmTaskTypeRollback HelmTaskType = "Rollback"
)

type Target struct {
	// Cluster defines a connected cluster as target
	// +optional
	Cluster *TargetCluster `json:"cluster,omitempty"`

	// VirtualCluster defines a virtual cluster as target
	// +optional
	VirtualCluster *TargetVirtualCluster `json:"virtualCluster,omitempty"`
}

type TargetCluster struct {
	// Cluster is the cluster where the task should get executed
	// +optional
	Cluster string `json:"cluster,omitempty"`

	// Namespace is the namespace where the task should get executed
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

type TargetVirtualCluster struct {
	// Cluster is the cluster where the virtual cluster lies
	// +optional
	Cluster string `json:"cluster,omitempty"`

	// Namespace is the namespace where the virtual cluster is located
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Name of the virtual cluster
	// +optional
	Name string `json:"name,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TaskList contains a list of Task
type TaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Task `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Task{}, &TaskList{})
}
