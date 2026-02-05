package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BareMetalInstanceDiscover is used to discover systems via Redfish
// +subresource-request
type BareMetalInstanceDiscover struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BareMetalInstanceDiscoverSpec   `json:"spec,omitempty"`
	Status BareMetalInstanceDiscoverStatus `json:"status,omitempty"`
}

type BareMetalInstanceDiscoverSpec struct {
	// Endpoint is the Redfish API base URL (e.g., https://192.168.1.100:8443)
	Endpoint string `json:"endpoint,omitempty"`

	// Username for Redfish authentication
	Username string `json:"username,omitempty"`

	// Password for Redfish authentication
	Password string `json:"password,omitempty"`

	// InsecureSkipVerify skips TLS certificate verification
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
}

type BareMetalInstanceDiscoverStatus struct {
	// Systems contains the list of discovered systems
	Systems []DiscoveredSystem `json:"systems,omitempty"`

	// Error contains any error message from the discovery process
	Error string `json:"error,omitempty"`
}

type DiscoveredSystem struct {
	// ID is the system identifier
	ID string `json:"id,omitempty"`

	// Name is the human-readable system name
	Name string `json:"name,omitempty"`

	// URI is the Redfish URI for this system
	URI string `json:"uri,omitempty"`

	// NetworkInterfaces contains discovered network interfaces
	NetworkInterfaces []DiscoveredNetworkInterface `json:"networkInterfaces,omitempty"`

	// Error contains any error encountered while discovering this system
	Error string `json:"error,omitempty"`
}

type DiscoveredNetworkInterface struct {
	// ID is the interface identifier
	ID string `json:"id,omitempty"`

	// Name is the human-readable interface name
	Name string `json:"name,omitempty"`

	// MACAddress is the MAC address
	MACAddress string `json:"macAddress,omitempty"`

	// SpeedMbps is the link speed in Mbps
	SpeedMbps int `json:"speedMbps,omitempty"`

	// LinkStatus indicates link status (e.g., "LinkUp", "LinkDown")
	LinkStatus string `json:"linkStatus,omitempty"`

	// IPv4Addresses contains assigned IPv4 addresses
	IPv4Addresses []string `json:"ipv4Addresses,omitempty"`

	// IPv6Addresses contains assigned IPv6 addresses
	IPv6Addresses []string `json:"ipv6Addresses,omitempty"`
}
