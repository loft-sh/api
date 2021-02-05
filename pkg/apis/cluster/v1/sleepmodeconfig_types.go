package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SleepModeConfig holds the sleepmode information
// +k8s:openapi-gen=true
// +resource:path=sleepmodeconfigs,rest=SleepModeConfigREST
type SleepModeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SleepModeConfigSpec   `json:"spec,omitempty"`
	Status SleepModeConfigStatus `json:"status,omitempty"`
}

type SleepModeConfigSpec struct {
	// If force sleep is true the space will sleep
	// +optional
	ForceSleep bool `json:"forceSleep,omitempty"`

	// If force sleep duration is set, this will force the space to sleep
	// for the given duration. It also implies that forceSleep is true.
	// During this period loft will also block certain requests to that space.
	// If this is set to 0, it means the space will sleep until it is manually
	// woken up via the cli or ui.
	// +optional
	ForceSleepDuration *int64 `json:"forceSleepDuration,omitempty"`

	// If true will delete all pods on sleep in the space regardless of
	// if they have a parent set
	// +optional
	DeleteAllPods bool `json:"deleteAllPods,omitempty"`

	// DeleteAfter specifies after how many seconds of inactivity the space should be deleted
	// +optional
	DeleteAfter int64 `json:"deleteAfter,omitempty"`

	// SleepAfter specifies after how many seconds of inactivity the space should sleep
	// +optional
	SleepAfter int64 `json:"sleepAfter,omitempty"`
}

type SleepModeConfigStatus struct {
	// LastActivity indicates the last activity in the namespace
	// +optional
	LastActivity int64 `json:"lastActivity,omitempty"`

	// SleepingSince specifies since when the space is sleeping (if this is not specified, loft assumes the space is not sleeping)
	// +optional
	SleepingSince int64 `json:"sleepingSince,omitempty"`

	// ActiveConnections is the amount of open connections to this space
	// +optional
	ActiveConnections int64 `json:"activeConnections,omitempty"`

	// Optional info that indicates how long the space was sleeping in the current epoch
	// +optional
	CurrentEpoch *EpochInfo `json:"currentEpoch,omitempty"`

	// Optional info that indicates how long the space was sleeping in the current epoch
	// +optional
	LastEpoch *EpochInfo `json:"lastEpoch,omitempty"`

	// This is a calculated field that will be returned but not saved and describes the percentage since the space
	// was created or the last 30 days the space has slept
	// +optional
	SleptLastThirtyDays *float64 `json:"sleptLastThirtyDays,omitempty"`

	// This is a calculated field that will be returned but not saved and describes the percentage since the space
	// was created or the last 7 days the space has slept
	// +optional
	SleptLastSevenDays *float64 `json:"sleptLastSevenDays,omitempty"`
}

// EpochInfo holds information about how long the space was sleeping in the epoch
type EpochInfo struct {
	// Timestamp when the epoch has started
	// +optional
	Start int64 `json:"start,omitempty"`
	// Amount of milliseconds the space has slept in the epoch
	// +optional
	Slept int64 `json:"slept,omitempty"`
}
