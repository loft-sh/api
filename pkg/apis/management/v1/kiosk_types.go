package v1

import (
	configv1alpha1 "github.com/loft-sh/agentapi/v2/pkg/apis/kiosk/config/v1alpha1"
	clusterv1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/cluster/v1"
	storagev1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/storage/v1"
	"github.com/loft-sh/jspolicy/pkg/apis/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// This file is just used as a collector for kiosk objects we want to generate stuff for

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Kiosk holds the kiosk types
// +k8s:openapi-gen=true
// +resource:path=kiosk,rest=KioskREST
type Kiosk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KioskSpec   `json:"spec,omitempty"`
	Status KioskStatus `json:"status,omitempty"`
}

type KioskSpec struct {
	// config.kiosk.sh
	ConfigAccount    configv1alpha1.Account          `json:"configAccount,omitempty"`
	Template         configv1alpha1.Template         `json:"template,omitempty"`
	TemplateInstance configv1alpha1.TemplateInstance `json:"templateInstance,omitempty"`

	// policy.loft.sh
	JsPolicy           v1beta1.JsPolicy           `json:"jsPolicy,omitempty"`
	JsPolicyBundle     v1beta1.JsPolicyBundle     `json:"jsPolicyBundle,omitempty"`
	JsPolicyViolations v1beta1.JsPolicyViolations `json:"jsPolicyViolations,omitempty"`

	// cluster.loft.sh
	HelmRelease        clusterv1.HelmRelease        `json:"helmRelease,omitempty"`
	SleepModeConfig    clusterv1.SleepModeConfig    `json:"sleepModeConfig,omitempty"`
	Space              clusterv1.Space              `json:"space,omitempty"`
	VirtualCluster     clusterv1.VirtualCluster     `json:"virtualCluster,omitempty"`
	LocalClusterAccess clusterv1.LocalClusterAccess `json:"localClusterAccess,omitempty"`
	ClusterQuota       clusterv1.ClusterQuota       `json:"clusterQuota,omitempty"`
	ChartInfo          clusterv1.ChartInfo          `json:"chartInfo,omitempty"`

	// storage.loft.sh
	SpaceConstraint       storagev1.LocalSpaceConstraint     `json:"localSpaceConstraint,omitempty"`
	StorageClusterAccess  storagev1.LocalClusterAccess       `json:"localStorageClusterAccess,omitempty"`
	ClusterRoleTemplate   storagev1.LocalClusterRoleTemplate `json:"localClusterRoleTemplate,omitempty"`
	StorageClusterQuota   storagev1.ClusterQuota             `json:"storageClusterQuota,omitempty"`
	StorageVirtualCluster storagev1.VirtualCluster           `json:"storageVirtualCluster,omitempty"`
	LocalUser             storagev1.LocalUser                `json:"localUser,omitempty"`
	LocalTeam             storagev1.LocalTeam                `json:"localTeam,omitempty"`
}

type KioskStatus struct {
}
