package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/v5/pkg/apis/loft/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MachineConditionTypeSynced indicates the machine state was confirmed
	// against the provider's inventory on the last sync attempt.
	MachineConditionTypeSynced agentstoragev1.ConditionType = "Synced"

	// MachineNodeClaimAnnotation names the NodeClaim ("<namespace>/<name>")
	// currently consuming this machine. Set by the inventory sync when the
	// provider-side consumer matches a platform NodeClaim; absent when the
	// machine is free or consumed outside the platform.
	MachineNodeClaimAnnotation = "machines.vcluster.com/node-claim"

	// MachineProviderLabelsAnnotation records keys the inventory sync last wrote. It is present, possibly
	// empty, once the sync has run, which is how selectors tell an unsynced Machine from one without
	// provider labels.
	MachineProviderLabelsAnnotation = "machines.vcluster.com/provider-labels"

	// MachineSourceLabel names the inventory a machine was recorded from before
	// the provider knew about it. A machine without it is a plain mirror of the
	// provider inventory and disappears with it; a machine carrying it is owned
	// by that source's sync, which is the only thing allowed to delete it, and
	// the provider sync only fills in its status once the provider lists it.
	MachineSourceLabel = "machines.vcluster.com/source"

	// MachineSourceNetBox is the MachineSourceLabel value for machines imported
	// from a NetBox inventory through NodeProviderMetal3NetBox.
	MachineSourceNetBox = "netbox"

	// MachineConditionTypeRegistered is set on sourced machines: True once the
	// provider holds a record for the machine (metal3: the BareMetalHost
	// exists), False with the reason while it cannot be created.
	MachineConditionTypeRegistered agentstoragev1.ConditionType = "Registered"

	// MachineConditionTypeSourceAttached is set on sourced machines: True while
	// the source still selects the machine, False with reason Orphaned once it
	// no longer does (the NetBox device lost its tag or was deleted) but the
	// machine was kept because it is provisioned or held by a NodeClaim.
	// Releasing it is then an operator's decision, not the sync's.
	MachineConditionTypeSourceAttached agentstoragev1.ConditionType = "SourceAttached"

	// MachineConditionTypeHostValidated summarizes the checks recorded in
	// status.hostValidation: the in-band attempt, which runs tests on the
	// machine, and the check against the host validation profiles its node types
	// select, which reads its BMC. Its reason is one of the
	// MachineHostValidatedReason constants, the worst of the checks that have a
	// result, and its message is that check's. It is absent while no check has a
	// result. The gates read the result of the check they need from
	// status.hostValidation, not this condition.
	MachineConditionTypeHostValidated agentstoragev1.ConditionType = "HostValidated"

	// MachineReservedForAnnotation reserves an idle machine for one taker, so
	// validation and a tenant NodeClaim never take the same machine:
	// MachineReservedForValidation, or MachineReservedForNodeClaimPrefix followed
	// by the NodeClaim's UID. Only the platform writes it, with an optimistic
	// update.
	MachineReservedForAnnotation = "machines.vcluster.com/reserved-for"

	// MachineReservedForValidation is the MachineReservedForAnnotation value
	// while the platform validates the machine.
	MachineReservedForValidation = "validation"

	// MachineReservedForNodeClaimPrefix prefixes the UID of the NodeClaim a
	// machine is reserved for in MachineReservedForAnnotation.
	MachineReservedForNodeClaimPrefix = "nodeclaim/"

	// MachineHandoverDeadlineAnnotation is when a reservation for a NodeClaim
	// stops waiting for the machine to leave the validation cluster, in RFC
	// 3339. Only the platform writes it.
	MachineHandoverDeadlineAnnotation = "machines.vcluster.com/handover-deadline"

	// MachineRerunHostValidationAnnotation asks for a new validation attempt. The
	// platform removes it once the new attempt is recorded in status.hostValidation.inBand.
	MachineRerunHostValidationAnnotation = "machines.vcluster.com/rerun-host-validation"
)

// Reasons of the HostValidated condition, from the worst down. Where a check
// stands in detail is in its result in status.hostValidation.
const (
	// MachineHostValidatedReasonFailed is set, as False, when a check failed: an
	// in-band attempt TestFailed, JoinFailed or SetupFailed, or a profile was
	// Mismatched, ReadFailed or Unsupported.
	MachineHostValidatedReasonFailed = "Failed"
	// MachineHostValidatedReasonRunning is set, as Unknown, while a check runs: an
	// in-band test, or the read of a profile.
	MachineHostValidatedReasonRunning = "Running"
	// MachineHostValidatedReasonPending is set, as Unknown, while a check waits to
	// run: an in-band attempt Joining or Starting, or a profile that matches once
	// the machine reboots (RebootRequired).
	MachineHostValidatedReasonPending = "Pending"
	// MachineHostValidatedReasonSucceeded is set, as True, once every check that
	// has a result passed.
	MachineHostValidatedReasonSucceeded = "Succeeded"
)

// MachineInBandValidationResult is where a machine's in-band attempt stands.
// TestFailed blames the machine, while JoinFailed and SetupFailed blame the
// setup, so an unreachable validation cluster never reads as bad hardware.
// +enum
type MachineInBandValidationResult string

const (
	// MachineInBandValidationJoining is while the validation NodeClaim exists and
	// its node has not joined the validation cluster.
	MachineInBandValidationJoining MachineInBandValidationResult = "Joining"
	// MachineInBandValidationStarting is once the node joined and before the
	// first test runs.
	MachineInBandValidationStarting MachineInBandValidationResult = "Starting"
	// MachineInBandValidationRunning is while a test's Certification or Job runs
	// on the node.
	MachineInBandValidationRunning MachineInBandValidationResult = "Running"
	// MachineInBandValidationPassed is once every test passed.
	MachineInBandValidationPassed MachineInBandValidationResult = "Passed"
	// MachineInBandValidationTestFailed is when a test failed on the node: its
	// Certification failed and recorded the node, or its Job exited non-zero or
	// ran past its timeout. The message is the reason recorded for the node, or
	// the last line the Job printed.
	MachineInBandValidationTestFailed MachineInBandValidationResult = "TestFailed"
	// MachineInBandValidationJoinFailed is when the validation NodeClaim failed
	// or the node did not join before the join deadline.
	MachineInBandValidationJoinFailed MachineInBandValidationResult = "JoinFailed"
	// MachineInBandValidationSetupFailed is when the tests could not run: the
	// suite is missing or not understood, the Certification kind is not served in
	// the validation cluster, a test's image cannot be pulled, a Certification
	// failed without recording the node, the setup deadline passed, or a test pod
	// was lost without an exit code.
	MachineInBandValidationSetupFailed MachineInBandValidationResult = "SetupFailed"
)

// HostValidationTarget is the component whose BMC holds a group's settings. It
// tells the host validation controller which BMC to ask, and so which serial to
// route the read by.
// +enum
type HostValidationTarget string

const (
	// HostValidationTargetHost is the machine's own BMC, routed by
	// status.hardware.serial.
	HostValidationTargetHost HostValidationTarget = "Host"
)

// HostValidationResult is the verdict of one host validation profile on a
// machine, in MachineHostValidationProfileResult: whether the machine may serve
// a node type that requires the profile. It is the first of the values below
// that the read or any of the profile's settings calls for, which is the order
// they are declared in.
// +enum
type HostValidationResult string

const (
	// HostValidationResultUnsupported means nothing can read this machine's
	// settings: its node provider is not NICo, or it has no BMC proxy endpoint.
	// The setup is at fault, not the host, and only a change to the node
	// provider clears it.
	HostValidationResultUnsupported HostValidationResult = "Unsupported"
	// HostValidationResultReadFailed means the proxy or the BMC returned an
	// error, or the BMC reported another serial number. The setup is at fault,
	// not the host. The controller reads again on its own.
	HostValidationResultReadFailed HostValidationResult = "ReadFailed"
	// HostValidationResultMismatched means a setting has another value, or the
	// BMC does not report it at all. The host is at fault, so it is not read
	// again on its own: an admin fixes the setting and asks for a re-check.
	HostValidationResultMismatched HostValidationResult = "Mismatched"
	// HostValidationResultRebootRequired means every setting that differs
	// already has the required value staged, so the machine matches the profile
	// once it reboots. The platform never reboots a machine, NICo or an admin
	// does. The controller reads again on its own.
	HostValidationResultRebootRequired HostValidationResult = "RebootRequired"
	// HostValidationResultReading means the check was requested and has not
	// finished. A result left at Reading by a restart is checked again.
	HostValidationResultReading HostValidationResult = "Reading"
	// HostValidationResultMatched means every setting has the value the profile
	// requires.
	HostValidationResultMatched HostValidationResult = "Matched"
)

// HostValidationSettingResult is what the BMC reported for one setting, in
// MachineHostValidationSettingResult. Matched and Mismatched mean for the setting
// what they mean for a whole profile. Staged and Missing say more than a
// profile can, so a profile reads Missing as Mismatched and Staged as
// RebootRequired.
// +enum
type HostValidationSettingResult string

const (
	// HostValidationSettingResultMatched means the BMC reports the value the
	// profile requires.
	HostValidationSettingResultMatched HostValidationSettingResult = "Matched"
	// HostValidationSettingResultMismatched means the BMC reports another value
	// and has none staged, so the setting stays as it is until an admin changes
	// it.
	HostValidationSettingResultMismatched HostValidationSettingResult = "Mismatched"
	// HostValidationSettingResultStaged means the BMC reports another value and
	// the required one staged in the resource's settings object, which takes
	// effect at the next reboot.
	HostValidationSettingResultStaged HostValidationSettingResult = "Staged"
	// HostValidationSettingResultMissing means the BMC does not report the
	// setting's attribute at all, which usually means the profile describes
	// another class of hardware.
	HostValidationSettingResultMissing HostValidationSettingResult = "Missing"
)

// Results of one test in MachineValidationTestResult.
const (
	MachineValidationTestRunning = "Running"
	MachineValidationTestPassed  = "Passed"
	MachineValidationTestFailed  = "Failed"
)

// Kinds of object a test runs as, in MachineValidationTestResult.
const (
	// MachineValidationTestKindCertification is an NVIDIA Cluster Readiness
	// Engine Certification, the default kind of test.
	MachineValidationTestKindCertification = "Certification"
	// MachineValidationTestKindJob is a Kubernetes Job, for a test the
	// Certifications do not cover.
	MachineValidationTestKindJob = "Job"
)

// Types of MachineValidationArtifact, media types of its data once
// decompressed.
const (
	// MachineValidationArtifactTypeCertificationReport is the report of an
	// NVIDIA Cluster Readiness Engine Certification, as the JSON `nvcrectl
	// certification report --results-file` writes, without the sections that
	// only several nodes, repeated runs or diagnose mode fill.
	MachineValidationArtifactTypeCertificationReport = "application/vnd.nvidia.nvcre.certification-report+json"

	// MachineValidationArtifactTypeJobReport is the report of a Job test as
	// JSON: the Job's name, result (PASSED, FAILED or RUNNING), failureReason
	// and runtimeSeconds, and in containers, for each container of its pods,
	// the pod, nodeName, name, init, exitCode, reason and message of how it
	// ended or why it waits, the end of its log in log (at most 800 lines and
	// 32 KiB) and logError when the log could not be read.
	MachineValidationArtifactTypeJobReport = "application/vnd.vcluster.machine-validation.job-report+json"
)

// MachineConditions lists the conditions summarized into the Ready condition.
var MachineConditions = []agentstoragev1.ConditionType{
	MachineConditionTypeSynced,
	MachineConditionTypeRegistered,
}

// MachinePhase is the machine's state as the provider sees it, normalized
// across providers. It says nothing about platform NodeClaims: which claim
// holds a machine is answered by the node-claim annotation.
type MachinePhase string

const (
	// MachinePhasePending means the machine is recorded in the platform from an
	// inventory source (see MachineSourceLabel) but the provider does not list
	// it yet: the BareMetalHost is still to be created, or was rejected. Reason
	// and Message say which.
	MachinePhasePending MachinePhase = "Pending"
	// MachinePhaseAvailable means the provider reports the machine as
	// inventoried and free. For metal3 this is the host's reported state, so a
	// host a consumer has just asked for reads Available until BMO starts on it.
	MachinePhaseAvailable MachinePhase = "Available"
	// MachinePhaseProvisioned means the provider handed the machine to a
	// consumer: metal3 reports the host provisioned or externally provisioned,
	// NICo assigned it to an instance. Whether that consumer is one of this
	// platform's NodeClaims is answered by the node-claim annotation.
	MachinePhaseProvisioned MachinePhase = "Provisioned"
	// MachinePhaseProvisioning means the provider is bringing the machine up:
	// registering, inspecting, preparing, imaging. The provider's own word for
	// the step is in Reason, which for metal3 can be its literal "provisioning"
	// as well as "inspecting" or "preparing".
	MachinePhaseProvisioning MachinePhase = "Provisioning"
	// MachinePhaseDeprovisioning means the provider is tearing the machine
	// down: wiping it after release or powering it off for deletion. Reason
	// carries the provider's step, for metal3 "deprovisioning", "powering off
	// before delete" or "deleting"; NICo reports InUse after the instance is
	// gone.
	MachinePhaseDeprovisioning MachinePhase = "Deprovisioning"
	// MachinePhaseMaintenance means the machine was deliberately taken out of
	// service (metal3 detached, NICo Maintenance or Decommissioned).
	MachinePhaseMaintenance MachinePhase = "Maintenance"
	// MachinePhaseError means the provider reports a fault. Reason and Message
	// carry the provider's own words for it.
	MachinePhaseError MachinePhase = "Error"
	// MachinePhaseUnknown means the provider could not be reached or reports no
	// state at all, so the mirrored state may be stale.
	MachinePhaseUnknown MachinePhase = "Unknown"
)

// MachinePowerState is the normalized power state of a machine. Empty means
// the provider does not report power at all.
type MachinePowerState string

const (
	MachinePowerStateOn  MachinePowerState = "On"
	MachinePowerStateOff MachinePowerState = "Off"
	// MachinePowerStateUnknown means the provider reports a state that maps
	// to neither On nor Off.
	MachinePowerStateUnknown MachinePowerState = "Unknown"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Provider",type="string",JSONPath=".spec.providerRef"
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Power",type="string",JSONPath=".status.powerState"
// +kubebuilder:printcolumn:name="Product",type="string",JSONPath=".status.hardware.product"
// +kubebuilder:printcolumn:name="Last Sync",type="date",JSONPath=".status.lastSyncTime"

// Machine is a read-only mirror of one physical machine in a bare metal
// provider's inventory. It is created, updated, and deleted exclusively by the
// platform's inventory sync; users cannot write it. Its name is
// "<providerRef>.<provider-side machine id>" (the NodeType naming convention):
// metal3 uses the BareMetalHost name, NICo the machine id.
// +k8s:openapi-gen=true
type Machine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="machine spec is read-only"
	Spec   MachineSpec   `json:"spec,omitempty"`
	Status MachineStatus `json:"status,omitempty"`
}

func (a *Machine) GetConditions() agentstoragev1.Conditions {
	return a.Status.Conditions
}

func (a *Machine) SetConditions(conditions agentstoragev1.Conditions) {
	a.Status.Conditions = conditions
}

// MachineSpec holds the mirror's identity. The machine id lives in the
// object name; the spec only duplicates the provider reference for indexing.
type MachineSpec struct {
	// DisplayName is shown in the UI; providers fill it from their native
	// naming (metal3: the BareMetalHost name, NICo: the machine hostname).
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// ProviderRef is the NodeProvider whose inventory this machine mirrors.
	ProviderRef string `json:"providerRef,omitempty"`
}

// MachineStatus is the observed provider-side state of the machine.
type MachineStatus struct {
	// Phase is the machine's state at the provider, normalized across providers.
	// It says nothing about platform NodeClaims: the node-claim annotation
	// answers which claim, if any, holds the machine.
	// +optional
	Phase MachinePhase `json:"phase,omitempty"`

	// Reason is the provider's own name for the state behind the phase, kept
	// verbatim: a metal3 provisioning state or errorType, a NICo machine status
	// or health alert id.
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message is what the provider says about it, when it says anything: a
	// metal3 errorMessage, a NICo alert or maintenance message.
	// +optional
	Message string `json:"message,omitempty"`

	// Conditions holds the sync conditions of the machine.
	// +optional
	Conditions agentstoragev1.Conditions `json:"conditions,omitempty"`

	// Hardware is the machine's discovered hardware inventory.
	// +optional
	Hardware *MachineHardware `json:"hardware,omitempty"`

	// System is the machine's observed runtime state (assigned addressing and
	// similar), as opposed to the physical inventory in Hardware.
	// +optional
	System *MachineSystem `json:"system,omitempty"`

	// PowerState is the machine's normalized power state: On, Off, or Unknown
	// when the provider reports a state that maps to neither (e.g. an error).
	// Empty when the provider does not report power at all.
	// +optional
	PowerState MachinePowerState `json:"powerState,omitempty"`

	// LastSyncTime is when this state was last confirmed against the provider.
	// +optional
	LastSyncTime metav1.Time `json:"lastSyncTime,omitempty"`

	// HostValidation records the checks the platform ran on the machine, which
	// the HostValidated condition was decided from. The inventory sync keeps it,
	// as the provider knows nothing of it, but a machine the sync recreates
	// because its identity changed starts without it.
	// +optional
	HostValidation *MachineHostValidationStatus `json:"hostValidation,omitempty"`
}

// MachineHostValidationStatus holds the results of a machine's validation: the
// in-band attempt, which runs a workload on the machine, and the last check
// against its host validation profiles, which reads its BMC without running
// anything on it.
type MachineHostValidationStatus struct {
	// InBand is the machine's last in-band attempt, which runs the tests of its
	// NodeProvider's validation tier on it. It gates assigning the machine to a
	// tenant and using it for an instance.
	// +optional
	InBand *MachineInBandValidationStatus `json:"inBand,omitempty"`

	// ReadTime is when the settings behind the profile results were read. The
	// retry of a RebootRequired or ReadFailed profile counts from it.
	// +optional
	ReadTime *metav1.Time `json:"readTime,omitempty"`

	// Profiles holds one result per host validation profile the machine was
	// checked against. A profile added to a node type after this check has no
	// entry, and does not count as matched. They gate picking the machine for a
	// node type.
	// +optional
	Profiles []MachineHostValidationProfileResult `json:"profiles,omitempty"`
}

// MachineInBandValidationStatus is a machine's in-band attempt.
type MachineInBandValidationStatus struct {
	// Result is where the attempt stands. Only Passed lets the machine be
	// assigned to a tenant or used for an instance on a provider with a
	// validation tier.
	// +optional
	Result MachineInBandValidationResult `json:"result,omitempty"`

	// Message says what the attempt is doing, or why it ended as it did.
	// +optional
	Message string `json:"message,omitempty"`

	// Tier is the validation tier the attempt ran for.
	// +optional
	Tier string `json:"tier,omitempty"`

	// AttemptID identifies the attempt. The Certifications and Jobs of any other
	// attempt are ignored.
	// +optional
	AttemptID string `json:"attemptID,omitempty"`

	// SuiteHash is the hash of the test suite the attempt ran. A changed suite
	// does not invalidate a pass, so this says which suite a pass was for.
	// +optional
	SuiteHash string `json:"suiteHash,omitempty"`

	// ClusterUID is the UID of the validation tenant cluster the attempt ran in.
	// +optional
	ClusterUID string `json:"clusterUID,omitempty"`

	// NodeName is the node the machine joined the validation cluster as.
	// +optional
	NodeName string `json:"nodeName,omitempty"`

	// StartTime is when the attempt started.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// FinishTime is when the attempt ended, unset while it runs.
	// +optional
	FinishTime *metav1.Time `json:"finishTime,omitempty"`

	// Tests holds the result of each test the attempt reached, in suite order.
	// +optional
	Tests []MachineValidationTestResult `json:"tests,omitempty"`

	// Rejoin tracks putting a machine that passed back into the validation
	// cluster after it was released. A rejoin runs no tests.
	// +optional
	Rejoin *MachineValidationRejoin `json:"rejoin,omitempty"`

	// LeavingSince is when the platform started taking the machine out of the
	// validation cluster for a tenant that assigned or reserved it. Unset once
	// its validation NodeClaim is gone.
	// +optional
	LeavingSince *metav1.Time `json:"leavingSince,omitempty"`
}

// MachineHostValidationProfileResult is what one host validation profile
// decided for a machine.
type MachineHostValidationProfileResult struct {
	// Name is the host validation profile this result is for.
	Name string `json:"name"`

	// Result is the profile's result.
	Result HostValidationResult `json:"result"`

	// Message says what the result means, naming the first setting that did not
	// match and both of its values, or what went wrong with the read.
	// +optional
	Message string `json:"message,omitempty"`

	// SpecHash is the hash of the profile's spec this result was decided for. A
	// profile that changed since reads differently from one that did not, without
	// the result itself being touched before the next check.
	// +optional
	SpecHash string `json:"specHash,omitempty"`

	// OutOfBand mirrors the profile's out-of-band groups, one result each.
	// +optional
	OutOfBand []MachineHostValidationGroupResult `json:"outOfBand,omitempty"`
}

// MachineHostValidationGroupResult is the result of one group of a profile,
// and names the BMC it was read from.
type MachineHostValidationGroupResult struct {
	// Name is the group's name in the profile.
	Name string `json:"name"`

	// Target is the component whose BMC was read.
	Target HostValidationTarget `json:"target"`

	// Serial is the serial number the read was routed by. A group result exists
	// only for a group the controller tried to read, and the serial it routed by
	// is known before it sends anything, so this is always set. A machine with no
	// serial at all cannot be routed to and fails the whole profile as ReadFailed
	// with no group results.
	Serial string `json:"serial"`

	// Settings holds one result per setting of the group.
	// +optional
	Settings []MachineHostValidationSettingResult `json:"settings,omitempty"`
}

// MachineHostValidationSettingResult is what one setting of a group read as.
type MachineHostValidationSettingResult struct {
	// Name is the setting's name in the profile.
	Name string `json:"name"`

	// Expected is the value the profile requires.
	Expected string `json:"expected"`

	// Current is the value the BMC reports, rendered as a string. Empty when the
	// attribute is missing or nothing was read.
	// +optional
	Current string `json:"current,omitempty"`

	// Staged is the value the BMC holds for the next reboot, from the Redfish
	// settings object of the resource the attribute lives in. Empty when the BMC
	// stages nothing for it.
	// +optional
	Staged string `json:"staged,omitempty"`

	// Result is the setting's result.
	Result HostValidationSettingResult `json:"result"`
}

// MachineValidationTestResult is the result of one test of a validation attempt.
type MachineValidationTestResult struct {
	// Name is the test's name in the suite.
	Name string `json:"name"`

	// Result is Running, Passed or Failed.
	Result string `json:"result"`

	// Message is what the test found: the reason a failed Certification recorded
	// for the node, or the last line a Job printed.
	// +optional
	Message string `json:"message,omitempty"`

	// Kind is what the test ran as in the validation cluster, Certification or
	// Job.
	// +optional
	Kind string `json:"kind,omitempty"`

	// StartTime is when the test's Certification or Job was created.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// FinishTime is when the test's result was recorded, unset while it runs.
	// +optional
	FinishTime *metav1.Time `json:"finishTime,omitempty"`

	// ObjectName is the test's Certification or Job, in the namespace of the
	// suite. It is cleared once the object is deleted, a day after the attempt
	// finished.
	// +optional
	ObjectName string `json:"objectName,omitempty"`

	// Report is the test's report, recorded with its result: a Certification's
	// NCRE report, or how a Job's containers ended with the end of their logs.
	// It stays after the Certification or Job is deleted.
	// +optional
	Report *MachineValidationArtifact `json:"report,omitempty"`
}

// MachineValidationArtifact is a file a validation test produced.
type MachineValidationArtifact struct {
	// Type is the media type of Data once decompressed, one of the
	// MachineValidationArtifactType constants.
	Type string `json:"type"`

	// Data is the artifact compressed with gzip, base64 encoded in JSON.
	Data []byte `json:"data"`
}

// MachineValidationRejoin is the state of a machine's rejoin to the validation
// cluster.
type MachineValidationRejoin struct {
	// StartTime is when the running rejoin started, unset while none runs.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// LastError is why the last rejoin failed.
	// +optional
	LastError string `json:"lastError,omitempty"`

	// LastFailureTime is when the last rejoin failed. The next one waits for a
	// backoff from it.
	// +optional
	LastFailureTime *metav1.Time `json:"lastFailureTime,omitempty"`
}

// MachineSystem is the machine's observed runtime state.
type MachineSystem struct {
	// Network is the machine's network runtime state.
	// +optional
	Network *MachineNetwork `json:"network,omitempty"`
}

// MachineNetwork is the machine's network runtime state.
type MachineNetwork struct {
	// Interfaces lists the machine's network interfaces and their current
	// addressing.
	// +optional
	Interfaces []MachineNetworkInterface `json:"interfaces,omitempty"`
}

// MachineNetworkInterface is one network interface of a machine.
type MachineNetworkInterface struct {
	// Name is the interface name when the provider reports it (e.g. enp1s0).
	// +optional
	Name string `json:"name,omitempty"`

	// MAC is the interface's hardware address.
	// +optional
	MAC string `json:"mac,omitempty"`

	// Primary marks the interface the machine's traffic goes through (NICo:
	// the primary interface; metal3: the PXE/provisioning interface).
	// +optional
	Primary bool `json:"primary,omitempty"`

	// IPs are the addresses currently assigned to the interface; empty for a
	// machine waiting unprovisioned in the pool.
	// +optional
	IPs []string `json:"ips,omitempty"`
}

// MachineHardware is the normalized hardware inventory of a machine.
type MachineHardware struct {
	// CPU aggregates all CPU capabilities of the machine (sockets, cores and
	// threads summed across every CPU entry in Capabilities).
	// +optional
	CPU *CPUResource `json:"cpu,omitempty"`

	// Memory aggregates all Memory capabilities of the machine (capacity
	// summed across every Memory entry in Capabilities).
	// +optional
	Memory *resource.Quantity `json:"memory,omitempty"`

	// Storage aggregates all Storage capabilities of the machine (capacity
	// summed across every Storage entry in Capabilities).
	// +optional
	Storage *resource.Quantity `json:"storage,omitempty"`

	// Capabilities lists the machine's discovered devices, one entry per
	// device model.
	// +optional
	Capabilities []MachineCapability `json:"capabilities,omitempty"`

	// Vendor is the machine's hardware vendor when the provider reports it.
	// +optional
	Vendor string `json:"vendor,omitempty"`

	// Product is the machine's product name when the provider reports it.
	// +optional
	Product string `json:"product,omitempty"`

	// Serial is the machine's serial number when the provider reports it.
	// +optional
	Serial string `json:"serial,omitempty"`
}

// MachineCapabilityType is the kind of device a MachineCapability describes.
type MachineCapabilityType string

const (
	MachineCapabilityCPU        MachineCapabilityType = "CPU"
	MachineCapabilityMemory     MachineCapabilityType = "Memory"
	MachineCapabilityGPU        MachineCapabilityType = "GPU"
	MachineCapabilityDPU        MachineCapabilityType = "DPU"
	MachineCapabilityNetwork    MachineCapabilityType = "Network"
	MachineCapabilityStorage    MachineCapabilityType = "Storage"
	MachineCapabilityInfiniBand MachineCapabilityType = "InfiniBand"
)

// MachineCapability describes one device model present in a machine.
type MachineCapability struct {
	// Type is the device kind.
	Type MachineCapabilityType `json:"type"`

	// Name is the device model name.
	// +optional
	Name string `json:"name,omitempty"`

	// Vendor is the device vendor.
	// +optional
	Vendor string `json:"vendor,omitempty"`

	// Count is how many devices of this model the machine has.
	// +optional
	Count int64 `json:"count,omitempty"`

	// Resources holds the device resources this capability provides, per
	// device kind (CPU topology for CPU entries, capacity for the rest).
	// +optional
	Resources *Resources `json:"resources,omitempty"`
}

// Resources describes what one capability entry provides. Which fields are
// set depends on the capability type.
type Resources struct {
	// Capacity is the per-device capacity when the provider reports one
	// (Memory size, GPU memory, disk size).
	// +optional
	Capacity *resource.Quantity `json:"capacity,omitempty"`

	// CPU topology fields, inlined; only set on CPU capabilities.
	CPUResource `json:",inline"`
}

// CPUResource describes CPU topology. On a CPU capability entry it covers all
// packages of that entry's model together (sockets is the package count,
// cores/threads are totals across them); MachineHardware.CPU is the sum over
// all CPU entries. Providers fill what they report (metal3: architecture and
// logical CPU count as threads; NICo: packages with per-package cores, and
// threads when scout provides them); zero means "not reported".
type CPUResource struct {
	// Architecture is the CPU architecture (e.g. x86_64, aarch64).
	// +optional
	Architecture string `json:"architecture,omitempty"`

	// Sockets is the number of CPU packages.
	// +optional
	Sockets int64 `json:"sockets,omitempty"`

	// Cores is the number of physical cores.
	// +optional
	Cores int64 `json:"cores,omitempty"`

	// Threads is the number of logical CPUs.
	// +optional
	Threads int64 `json:"threads,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineList contains a list of Machine
type MachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Machine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Machine{}, &MachineList{})
}
