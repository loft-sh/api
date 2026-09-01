package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SlurmInstanceTopology holds the Slurm network topology (switches or blocks and
// the nodes beneath them) retrieved from the tenant cluster by running
// scontrol show topoconf --json on the Slurm controller.
// +subresource-request
type SlurmInstanceTopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status SlurmInstanceTopologyStatus `json:"status,omitempty"`
}

// SlurmInstanceTopologyStatus is the observed topology state.
type SlurmInstanceTopologyStatus struct {
	// Supported is true when the tenant cluster's Slurm exposes topology
	// configuration (Slurm 25.05+ with a topology plugin configured). When false,
	// Configs is empty and Message describes why.
	Supported bool `json:"supported"`

	// Message optionally describes why topology data could not be retrieved (for
	// example the Slurm version does not support topology export, or no topology
	// plugin is configured).
	// +optional
	Message string `json:"message,omitempty"`

	// Configs are the topology configurations reported by Slurm.
	// +optional
	Configs []SlurmTopologyConfig `json:"configs,omitempty"`
}

// SlurmTopologyConfig is a single topology configuration. Exactly one of Tree,
// Block or Flat describes the layout, mirroring Slurm's TopologyPlugin.
type SlurmTopologyConfig struct {
	// Name is the arbitrary topology name defined in topology.conf.
	// +optional
	Name string `json:"name,omitempty"`
	// ClusterDefault is true when this topology is used outside the context of
	// partitions.
	// +optional
	ClusterDefault bool `json:"clusterDefault,omitempty"`
	// Tree is set when the topology uses the topology/tree plugin.
	// +optional
	Tree *SlurmTopologyTree `json:"tree,omitempty"`
	// Block is set when the topology uses the topology/block plugin.
	// +optional
	Block *SlurmTopologyBlock `json:"block,omitempty"`
	// Flat is true when the topology uses the topology/flat plugin.
	// +optional
	Flat bool `json:"flat,omitempty"`
}

// SlurmTopologyTree is the topology/tree layout: a set of switches, each with
// child switches and/or leaf nodes.
type SlurmTopologyTree struct {
	// Switches are the switch definitions making up the tree.
	// +optional
	Switches []SlurmTopologySwitch `json:"switches,omitempty"`
}

// SlurmTopologySwitch is a single switch in a tree topology.
type SlurmTopologySwitch struct {
	// Switch is the arbitrary, Slurm-internal switch name.
	Switch string `json:"switch"`
	// Children is the hostlist expression of child switches (for example s[0-1]).
	// Either Children or Nodes is set.
	// +optional
	Children string `json:"children,omitempty"`
	// Nodes is the hostlist expression of leaf nodes beneath the switch (for
	// example slinky-[0-1]). Either Children or Nodes is set.
	// +optional
	Nodes string `json:"nodes,omitempty"`
}

// SlurmTopologyBlock is the topology/block layout: a set of blocks sized by
// BlockSizes.
type SlurmTopologyBlock struct {
	// BlockSizes is the planning base block size alongside any higher-level block
	// sizes that are enforced.
	// +optional
	BlockSizes []int32 `json:"blockSizes,omitempty"`
	// Blocks are the block definitions.
	// +optional
	Blocks []SlurmTopologyBlockEntry `json:"blocks,omitempty"`
}

// SlurmTopologyBlockEntry is a single block in a block topology.
type SlurmTopologyBlockEntry struct {
	// Block is the arbitrary, Slurm-internal block name.
	Block string `json:"block"`
	// Nodes is the hostlist expression of nodes in the block.
	// +optional
	Nodes string `json:"nodes,omitempty"`
}
