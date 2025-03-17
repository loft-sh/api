package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevPodWorkspaceTemplate holds the DevPodWorkspaceTemplate information
// +k8s:openapi-gen=true
type DevPodWorkspaceTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevPodWorkspaceTemplateSpec   `json:"spec,omitempty"`
	Status DevPodWorkspaceTemplateStatus `json:"status,omitempty"`
}

func (a *DevPodWorkspaceTemplate) GetVersions() []VersionAccessor {
	var retVersions []VersionAccessor
	for _, v := range a.Spec.Versions {
		b := v
		retVersions = append(retVersions, &b)
	}

	return retVersions
}

func (a *DevPodWorkspaceTemplateVersion) GetVersion() string {
	return a.Version
}

func (a *DevPodWorkspaceTemplate) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *DevPodWorkspaceTemplate) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *DevPodWorkspaceTemplate) GetAccess() []Access {
	return a.Spec.Access
}

func (a *DevPodWorkspaceTemplate) SetAccess(access []Access) {
	a.Spec.Access = access
}

// DevPodWorkspaceTemplateSpec holds the specification
type DevPodWorkspaceTemplateSpec struct {
	// DisplayName is the name that is shown in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes the virtual cluster template
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Parameters define additional app parameters that will set provider values
	// +optional
	Parameters []AppParameter `json:"parameters,omitempty"`

	// Template holds the DevPod workspace template
	Template DevPodWorkspaceTemplateDefinition `json:"template,omitempty"`

	// Versions are different versions of the template that can be referenced as well
	// +optional
	Versions []DevPodWorkspaceTemplateVersion `json:"versions,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`
}

type DevPodWorkspaceTemplateDefinition struct {
	// Kubernetes holds the definition for kubernetes based workspaces
	Kubernetes *DevPodWorkspaceKubernetesSpec `json:"kubernetes,omitempty"`

	// WorkspaceEnv are environment variables that should be available within the created workspace.
	// +optional
	WorkspaceEnv map[string]DevPodProviderOption `json:"workspaceEnv,omitempty"`

	// InstanceTemplate holds the workspace instance template
	// +optional
	InstanceTemplate DevPodWorkspaceInstanceTemplateDefinition `json:"instanceTemplate,omitempty"`

	// CredentialForwarding specifies controls for how workspaces created by this template forward credentials into the workspace
	// +optional
	CredentialForwarding *CredentialForwarding `json:"credentialForwarding,omitempty"`

	// Provider holds the legacy VM provider configuration
	//
	// Deprecated: use fields on template instead
	// +optional
	Provider *DevPodWorkspaceProvider `json:"provider,omitempty"`
}

type DevPodWorkspaceKubernetesSpec struct {
	// Pod holds the definition for workspace pod.
	//
	// Defaults will be applied for fields that aren't specified.
	// +optional
	Pod *DevPodWorkspacePodTemplate `json:"pod,omitempty"`

	// VolumeClaim holds the definition for the main workspace persistent volume.
	// This volume is guaranteed to exist for the lifespan of the workspace.
	//
	// Defaults will be applied for fields that aren't specified.
	// +optional
	VolumeClaim *DevPodWorkspaceVolumeClaimTemplate `json:"volumeClaim,omitempty"`

	// PodTimeout specifies a maximum duration to wait for the workspace pod to start up before failing.
	// Default: 10m
	// +optional
	PodTimeout string `json:"podTimeout,omitempty"`

	// NodeArchitecture specifies the node architecture the workspace image will be built for.
	// Only necessary if you need to build workspace images on the fly in the kubernetes cluster and your cluster is mixed architecture.
	// +optional
	NodeArchitecure string `json:"nodeArchitecture,omitempty"`

	// SpaceTemplateRef is a reference to the space that should get created for this DevPod.
	// If this is specified, the kubernetes provider will be selected automatically.
	// +optional
	SpaceTemplateRef *TemplateRef `json:"spaceTemplateRef,omitempty"`

	// SpaceTemplate is the inline template for a space that should get created for this DevPod.
	// If this is specified, the kubernetes provider will be selected automatically.
	// +optional
	SpaceTemplate *SpaceTemplateDefinition `json:"spaceTemplate,omitempty"`

	// VirtualClusterTemplateRef is a reference to the virtual cluster that should get created for this DevPod.
	// If this is specified, the kubernetes provider will be selected automatically.
	// +optional
	VirtualClusterTemplateRef *TemplateRef `json:"virtualClusterTemplateRef,omitempty"`

	// VirtualClusterTemplate is the inline template for a virtual cluster that should get created for this DevPod.
	// If this is specified, the kubernetes provider will be selected automatically.
	// +optional
	VirtualClusterTemplate *VirtualClusterTemplateDefinition `json:"virtualClusterTemplate,omitempty"`
}

// DevPodWorkspacePodTemplate is a less restrictive PodTemplate
type DevPodWorkspacePodTemplate struct {
	// The pods metadata
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	TemplateMetadata `json:"metadata,omitempty"`

	Spec DevPodWorkspacePodTemplateSpec `json:"spec,omitempty"`
}

type DevPodWorkspacePodResourceRequirements struct {
	// Limits describes the maximum amount of compute resources allowed.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// +optional
	Limits map[corev1.ResourceName]string `json:"limits,omitempty" protobuf:"bytes,1,rep,name=limits"`
	// Requests describes the minimum amount of compute resources required.
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
	// otherwise to an implementation-defined value. Requests cannot exceed Limits.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// +optional
	Requests map[corev1.ResourceName]string `json:"requests,omitempty" protobuf:"bytes,2,rep,name=requests"`
	// Claims lists the names of resources, defined in spec.resourceClaims,
	// that are used by this container.
	//
	// This is an alpha field and requires enabling the
	// DynamicResourceAllocation feature gate.
	//
	// This field is immutable. It can only be set for containers.
	//
	// +listType=map
	// +listMapKey=name
	// +featureGate=DynamicResourceAllocation
	// +optional
	Claims []corev1.ResourceClaim `json:"claims,omitempty" protobuf:"bytes,3,opt,name=claims"`
}

// A single application container that you want to run within a pod.
type DevPodWorkspaceContainer struct {
	// Name of the container specified as a DNS_LABEL.
	// Each container in a pod must have a unique name (DNS_LABEL).
	// Cannot be updated.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Container image name.
	// More info: https://kubernetes.io/docs/concepts/containers/images
	// This field is optional to allow higher level config management to default or override
	// container images in workload controllers like Deployments and StatefulSets.
	// +optional
	Image string `json:"image,omitempty" protobuf:"bytes,2,opt,name=image"`
	// Entrypoint array. Not executed within a shell.
	// The container image's ENTRYPOINT is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
	// produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
	// of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	// +listType=atomic
	Command []string `json:"command,omitempty" protobuf:"bytes,3,rep,name=command"`
	// Arguments to the entrypoint.
	// The container image's CMD is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
	// produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
	// of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	// +listType=atomic
	Args []string `json:"args,omitempty" protobuf:"bytes,4,rep,name=args"`
	// Container's working directory.
	// If not specified, the container runtime's default will be used, which
	// might be configured in the container image.
	// Cannot be updated.
	// +optional
	WorkingDir string `json:"workingDir,omitempty" protobuf:"bytes,5,opt,name=workingDir"`
	// List of ports to expose from the container. Not specifying a port here
	// DOES NOT prevent that port from being exposed. Any port which is
	// listening on the default "0.0.0.0" address inside a container will be
	// accessible from the network.
	// Modifying this array with strategic merge patch may corrupt the data.
	// For more information See https://github.com/kubernetes/kubernetes/issues/108255.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=containerPort
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=containerPort
	// +listMapKey=protocol
	Ports []corev1.ContainerPort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports"`
	// List of sources to populate environment variables in the container.
	// The keys defined within a source must be a C_IDENTIFIER. All invalid keys
	// will be reported as an event when the container is starting. When a key exists in multiple
	// sources, the value associated with the last source will take precedence.
	// Values defined by an Env with a duplicate key will take precedence.
	// Cannot be updated.
	// +optional
	// +listType=atomic
	EnvFrom []corev1.EnvFromSource `json:"envFrom,omitempty" protobuf:"bytes,19,rep,name=envFrom"`
	// List of environment variables to set in the container.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Env []corev1.EnvVar `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"`
	// Compute Resources required by this container.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// +optional
	Resources DevPodWorkspacePodResourceRequirements `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"`
	// Resources resize policy for the container.
	// +featureGate=InPlacePodVerticalScaling
	// +optional
	// +listType=atomic
	ResizePolicy []corev1.ContainerResizePolicy `json:"resizePolicy,omitempty" protobuf:"bytes,23,rep,name=resizePolicy"`
	// RestartPolicy defines the restart behavior of individual containers in a pod.
	// This field may only be set for init containers, and the only allowed value is "Always".
	// For non-init containers or when this field is not specified,
	// the restart behavior is defined by the Pod's restart policy and the container type.
	// Setting the RestartPolicy as "Always" for the init container will have the following effect:
	// this init container will be continually restarted on
	// exit until all regular containers have terminated. Once all regular
	// containers have completed, all init containers with restartPolicy "Always"
	// will be shut down. This lifecycle differs from normal init containers and
	// is often referred to as a "sidecar" container. Although this init
	// container still starts in the init container sequence, it does not wait
	// for the container to complete before proceeding to the next init
	// container. Instead, the next init container starts immediately after this
	// init container is started, or after any startupProbe has successfully
	// completed.
	// +featureGate=SidecarContainers
	// +optional
	RestartPolicy *corev1.ContainerRestartPolicy `json:"restartPolicy,omitempty" protobuf:"bytes,24,opt,name=restartPolicy,casttype=ContainerRestartPolicy"`
	// Pod volumes to mount into the container's filesystem.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=mountPath
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=mountPath
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty" patchStrategy:"merge" patchMergeKey:"mountPath" protobuf:"bytes,9,rep,name=volumeMounts"`
	// volumeDevices is the list of block devices to be used by the container.
	// +patchMergeKey=devicePath
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=devicePath
	// +optional
	VolumeDevices []corev1.VolumeDevice `json:"volumeDevices,omitempty" patchStrategy:"merge" patchMergeKey:"devicePath" protobuf:"bytes,21,rep,name=volumeDevices"`
	// Periodic probe of container liveness.
	// Container will be restarted if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	LivenessProbe *corev1.Probe `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`
	// Periodic probe of container service readiness.
	// Container will be removed from service endpoints if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	ReadinessProbe *corev1.Probe `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`
	// StartupProbe indicates that the Pod has successfully initialized.
	// If specified, no other probes are executed until this completes successfully.
	// If this probe fails, the Pod will be restarted, just as if the livenessProbe failed.
	// This can be used to provide different probe parameters at the beginning of a Pod's lifecycle,
	// when it might take a long time to load data or warm a cache, than during steady-state operation.
	// This cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	StartupProbe *corev1.Probe `json:"startupProbe,omitempty" protobuf:"bytes,22,opt,name=startupProbe"`
	// Actions that the management system should take in response to container lifecycle events.
	// Cannot be updated.
	// +optional
	Lifecycle *corev1.Lifecycle `json:"lifecycle,omitempty" protobuf:"bytes,12,opt,name=lifecycle"`
	// Optional: Path at which the file to which the container's termination message
	// will be written is mounted into the container's filesystem.
	// Message written is intended to be brief final status, such as an assertion failure message.
	// Will be truncated by the node if greater than 4096 bytes. The total message length across
	// all containers will be limited to 12kb.
	// Defaults to /dev/termination-log.
	// Cannot be updated.
	// +optional
	TerminationMessagePath string `json:"terminationMessagePath,omitempty" protobuf:"bytes,13,opt,name=terminationMessagePath"`
	// Indicate how the termination message should be populated. File will use the contents of
	// terminationMessagePath to populate the container status message on both success and failure.
	// FallbackToLogsOnError will use the last chunk of container log output if the termination
	// message file is empty and the container exited with an error.
	// The log output is limited to 2048 bytes or 80 lines, whichever is smaller.
	// Defaults to File.
	// Cannot be updated.
	// +optional
	TerminationMessagePolicy corev1.TerminationMessagePolicy `json:"terminationMessagePolicy,omitempty" protobuf:"bytes,20,opt,name=terminationMessagePolicy,casttype=TerminationMessagePolicy"`
	// Image pull policy.
	// One of Always, Never, IfNotPresent.
	// Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	// +optional
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty" protobuf:"bytes,14,opt,name=imagePullPolicy,casttype=PullPolicy"`
	// SecurityContext defines the security options the container should be run with.
	// If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	// +optional
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty" protobuf:"bytes,15,opt,name=securityContext"`

	// Variables for interactive containers, these have very specialized use-cases (e.g. debugging)
	// and shouldn't be used for general purpose containers.

	// Whether this container should allocate a buffer for stdin in the container runtime. If this
	// is not set, reads from stdin in the container will always result in EOF.
	// Default is false.
	// +optional
	Stdin bool `json:"stdin,omitempty" protobuf:"varint,16,opt,name=stdin"`
	// Whether the container runtime should close the stdin channel after it has been opened by
	// a single attach. When stdin is true the stdin stream will remain open across multiple attach
	// sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the
	// first client attaches to stdin, and then remains open and accepts data until the client disconnects,
	// at which time stdin is closed and remains closed until the container is restarted. If this
	// flag is false, a container processes that reads from stdin will never receive an EOF.
	// Default is false
	// +optional
	StdinOnce bool `json:"stdinOnce,omitempty" protobuf:"varint,17,opt,name=stdinOnce"`
	// Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.
	// Default is false.
	// +optional
	TTY bool `json:"tty,omitempty" protobuf:"varint,18,opt,name=tty"`
}

// DevPodWorkspacePodTemplateSpec is a less restrictive PodSpec
type DevPodWorkspacePodTemplateSpec struct {
	// List of volumes that can be mounted by containers belonging to the pod.
	// +optional
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// List of initialization containers belonging to the pod.
	// +optional
	InitContainers []DevPodWorkspaceContainer `json:"initContainers,omitempty"`

	// List of containers belonging to the pod.
	// +optional
	Containers []DevPodWorkspaceContainer `json:"containers,omitempty"`

	// Restart policy for all containers within the pod.
	// +optional
	RestartPolicy corev1.RestartPolicy `json:"restartPolicy,omitempty"`

	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
	// +optional
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty"`

	// Optional duration in seconds the pod may be active on the node relative to
	// +optional
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty"`

	// Set DNS policy for the pod.
	// +optional
	DNSPolicy corev1.DNSPolicy `json:"dnsPolicy,omitempty"`

	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// ServiceAccountName is the name of the ServiceAccount to use to run this pod.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
	// +optional
	AutomountServiceAccountToken *bool `json:"automountServiceAccountToken,omitempty"`

	// NodeName indicates in which node this pod is scheduled.
	// +optional
	NodeName string `json:"nodeName,omitempty"`

	// Host networking requested for this pod. Use the host's network namespace.
	// +optional
	HostNetwork bool `json:"hostNetwork,omitempty"`

	// Use the host's pid namespace.
	// +optional
	HostPID bool `json:"hostPID,omitempty"`

	// Use the host's ipc namespace.
	// +optional
	HostIPC bool `json:"hostIPC,omitempty"`

	// Share a single process namespace between all of the containers in a pod.
	// +optional
	ShareProcessNamespace *bool `json:"shareProcessNamespace,omitempty"`

	// SecurityContext holds pod-level security attributes and common container settings.
	// +optional
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`

	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
	// +optional
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`

	// Specifies the hostname of the Pod
	// +optional
	Hostname string `json:"hostname,omitempty"`

	// If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>".
	// +optional
	Subdomain string `json:"subdomain,omitempty"`

	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// If specified, the pod will be dispatched by specified scheduler.
	// If not specified, the pod will be dispatched by default scheduler.
	// +optional
	SchedulerName string `json:"schedulerName,omitempty"`

	// If specified, the pod's tolerations.
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts
	// +optional
	HostAliases []corev1.HostAlias `json:"hostAliases,omitempty"`

	// If specified, indicates the pod's priority.
	// +optional
	PriorityClassName string `json:"priorityClassName,omitempty"`

	// +optional
	Priority *int32 `json:"priority,omitempty"`

	// Specifies the DNS parameters of a pod.
	// +optional
	DNSConfig *corev1.PodDNSConfig `json:"dnsConfig,omitempty"`

	// If specified, all readiness gates will be evaluated for pod readiness.
	// +optional
	ReadinessGates []corev1.PodReadinessGate `json:"readinessGates,omitempty"`

	// RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod
	// +optional
	RuntimeClassName *string `json:"runtimeClassName,omitempty"`

	// EnableServiceLinks indicates whether information about services should be injected into pod's
	// environment variables, matching the syntax of Docker links.
	// +optional
	EnableServiceLinks *bool `json:"enableServiceLinks,omitempty"`

	// PreemptionPolicy is the Policy for preempting pods with lower priority.
	// +optional
	PreemptionPolicy *corev1.PreemptionPolicy `json:"preemptionPolicy,omitempty"`

	// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass.
	// +optional
	Overhead corev1.ResourceList `json:"overhead,omitempty"`

	// TopologySpreadConstraints describes how a group of pods ought to spread across topology domains.
	// +optional
	TopologySpreadConstraints []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`

	// If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default).
	// In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname).
	// In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\Tcpip\\Parameters to FQDN.
	// +optional
	SetHostnameAsFQDN *bool `json:"setHostnameAsFQDN,omitempty"`

	// Specifies the OS of the containers in the pod.
	// +optional
	OS *corev1.PodOS `json:"os,omitempty"`

	// Use the host's user namespace.
	// +optional
	HostUsers *bool `json:"hostUsers,omitempty"`

	// SchedulingGates is an opaque list of values that if specified will block scheduling the pod.
	// If schedulingGates is not empty, the pod will stay in the SchedulingGated state and the
	// scheduler will not attempt to schedule the pod.
	// +optional
	SchedulingGates []corev1.PodSchedulingGate `json:"schedulingGates,omitempty"`

	// ResourceClaims defines which ResourceClaims must be allocated
	// and reserved before the Pod is allowed to start. The resources
	// will be made available to those containers which consume them
	// by name.
	// +optional
	ResourceClaims []corev1.PodResourceClaim `json:"resourceClaims,omitempty"`

	// Resources is the total amount of CPU and Memory resources required by all
	// containers in the pod. It supports specifying Requests and Limits for
	// "cpu" and "memory" resource names only. ResourceClaims are not supported.
	// +optional
	Resources *DevPodWorkspacePodResourceRequirements `json:"resources,omitempty"`
}

type DevPodWorkspaceVolumeClaimTemplate struct {
	// The pods metadata
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	TemplateMetadata `json:"metadata,omitempty"`

	Spec corev1.PersistentVolumeClaimSpec `json:"spec,omitempty"`
}

// +enum
type GitCloneStrategy string

// WARN: Need to match https://github.com/loft-sh/devpod/pkg/git/clone.go
const (
	FullCloneStrategy     GitCloneStrategy = ""
	BloblessCloneStrategy GitCloneStrategy = "blobless"
	TreelessCloneStrategy GitCloneStrategy = "treeless"
	ShallowCloneStrategy  GitCloneStrategy = "shallow"
)

type DevPodWorkspaceProvider struct {
	// Name is the name of the provider. This can also be an url.
	// +optional
	Name string `json:"name,omitempty"`

	// Options are the provider option values
	// +optional
	Options map[string]DevPodProviderOption `json:"options,omitempty"`

	// Env are environment options to set when using the provider.
	// +optional
	Env map[string]DevPodProviderOption `json:"env,omitempty"`
}

type DevPodWorkspaceInstanceTemplateDefinition struct {
	// The workspace instance metadata
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	TemplateMetadata `json:"metadata,omitempty"`
}

type DevPodProviderOption struct {
	// Value of this option.
	// +optional
	Value string `json:"value,omitempty"`

	// ValueFrom specifies a secret where this value should be taken from.
	// +optional
	ValueFrom *DevPodProviderOptionFrom `json:"valueFrom,omitempty"`
}

type DevPodProviderOptionFrom struct {
	// ProjectSecretRef is the project secret to use for this value.
	// +optional
	ProjectSecretRef *corev1.SecretKeySelector `json:"projectSecretRef,omitempty"`

	// SharedSecretRef is the shared secret to use for this value.
	// +optional
	SharedSecretRef *corev1.SecretKeySelector `json:"sharedSecretRef,omitempty"`
}

type DevPodProviderSource struct {
	// Github source for the provider
	Github string `json:"github,omitempty"`

	// File source for the provider
	File string `json:"file,omitempty"`

	// URL where the provider was downloaded from
	URL string `json:"url,omitempty"`
}

type DevPodWorkspaceTemplateVersion struct {
	// Template holds the DevPod template
	// +optional
	Template DevPodWorkspaceTemplateDefinition `json:"template,omitempty"`

	// Parameters define additional app parameters that will set provider values
	// +optional
	Parameters []AppParameter `json:"parameters,omitempty"`

	// Version is the version. Needs to be in X.X.X format.
	// +optional
	Version string `json:"version,omitempty"`
}

type CredentialForwarding struct {
	// Docker specifies controls for how workspaces created by this template forward docker credentials
	// +optional
	Docker *DockerCredentialForwarding `json:"docker,omitempty"`

	// Git specifies controls for how workspaces created by this template forward git credentials
	// +optional
	Git *GitCredentialForwarding `json:"git,omitempty"`
}

type DockerCredentialForwarding struct {
	// Disabled prevents all workspaces created by this template from forwarding credentials into the workspace
	// +optional
	Disabled bool `json:"disabled,omitempty"`
}

type GitCredentialForwarding struct {
	// Disabled prevents all workspaces created by this template from forwarding credentials into the workspace
	// +optional
	Disabled bool `json:"disabled,omitempty"`
}

// DevPodWorkspaceTemplateStatus holds the status
type DevPodWorkspaceTemplateStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevPodWorkspaceTemplateList contains a list of DevPodWorkspaceTemplate
type DevPodWorkspaceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodWorkspaceTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DevPodWorkspaceTemplate{}, &DevPodWorkspaceTemplateList{})
}
