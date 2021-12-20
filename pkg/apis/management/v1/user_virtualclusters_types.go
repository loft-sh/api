package v1

import (
	clusterv1 "github.com/loft-sh/agentapi/pkg/apis/loft/cluster/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type UserVirtualClusters struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	VirtualClusters []ClusterVirtualCluster `json:"virtualClusters,omitempty"`
}

type ClusterVirtualCluster struct {
	// Cluster is the cluster name of the cluster the space is in
	// +optional
	Cluster string `json:"cluster,omitempty"`

	// SleepModeConfig is the sleep mode config of the space
	// +optional
	SleepModeConfig *clusterv1.SleepModeConfig `json:"sleepModeConfig,omitempty"`

	// VirtualCluster is the VirtualCluster object in the cluster
	VirtualCluster clusterv1.VirtualCluster `json:"virtualCluster,omitempty"`
}
