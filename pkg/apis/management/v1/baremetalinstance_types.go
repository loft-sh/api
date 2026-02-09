package v1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BareMetalInstance provides a management interface for metal3 BareMetalHost resources
// +k8s:openapi-gen=true
// +resource:path=baremetalinstances,rest=BareMetalInstanceREST
// +subresource:request=BareMetalInstanceProvision,path=provision,kind=BareMetalInstanceProvision,rest=BareMetalInstanceProvisionREST
// +subresource:request=BareMetalInstanceDiscover,path=discover,kind=BareMetalInstanceDiscover,rest=BareMetalInstanceDiscoverREST
type BareMetalInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BareMetalInstanceSpec   `json:"spec,omitempty"`
	Status BareMetalInstanceStatus `json:"status,omitempty"`
}

// BareMetalInstanceSpec defines the desired state of a bare metal instance
type BareMetalInstanceSpec struct {
	// MAC address for boot identification
	MACAddress string `json:"macAddress,omitempty"`

	// BMC connection details
	BMC *BMCDetails `json:"bmc,omitempty"`

	// Image configuration
	Image *BareMetalImage `json:"image,omitempty"`

	// UserData is literal cloud-init user data (secret managed internally)
	UserData string `json:"userData,omitempty"`

	// NetworkData is literal cloud-init network config (secret managed internally)
	NetworkData string `json:"networkData,omitempty"`

	// RootDeviceHints for target disk selection
	RootDeviceHints *RootDeviceHints `json:"rootDeviceHints,omitempty"`

	// Network configuration (translated to BMH annotations)
	Network *NetworkConfig `json:"network,omitempty"`
}

// BareMetalImage defines the image to be provisioned
type BareMetalImage struct {
	URL          string `json:"url"`
	Checksum     string `json:"checksum"`
	ChecksumType string `json:"checksumType,omitempty"`
}

// BMCDetails defines BMC connection information
type BMCDetails struct {
	Address                        string `json:"address"`
	Username                       string `json:"username,omitempty"`
	Password                       string `json:"password,omitempty"`
	CredentialsName                string `json:"credentialsName,omitempty"`
	DisableCertificateVerification bool   `json:"disableCertificateVerification,omitempty"`
}

// RootDeviceHints defines hints for selecting the target disk
type RootDeviceHints struct {
	DeviceName string `json:"deviceName"`
}

// NetworkConfig defines network configuration for the instance
type NetworkConfig struct {
	// IP address in CIDR notation (e.g., "192.168.100.14/24")
	IP string `json:"ip"`
	// Gateway IP address (e.g., "192.168.100.1")
	Gateway string `json:"gateway,omitempty"`
	// DNS servers (e.g., ["10.0.0.2"])
	DNSServers []string `json:"dnsServers,omitempty"`
}

// BareMetalInstanceStatus defines the observed state of a bare metal instance
type BareMetalInstanceStatus struct {
	ProvisioningState string            `json:"provisioningState,omitempty"`
	ErrorMessage      string            `json:"errorMessage,omitempty"`
	// Hardware contains the hardware details as inline JSON (cpu, firmware, nics, storage, etc.)
	// This field is dynamically populated from the BMH status.hardware
	// Clients should unmarshal this to access the hardware information
	Hardware         json.RawMessage   `json:"hardware,omitempty"`
	PoweredOn        bool              `json:"poweredOn,omitempty"`
	OperationHistory *OperationHistory `json:"operationHistory,omitempty"`
}

// OperationHistory tracks the history of operations on the instance
type OperationHistory struct {
	Register    *BareMetalOperation `json:"register,omitempty"`
	Inspect     *BareMetalOperation `json:"inspect,omitempty"`
	Provision   *BareMetalOperation `json:"provision,omitempty"`
	Deprovision *BareMetalOperation `json:"deprovision,omitempty"`
}

// BareMetalOperation tracks the timing of an operation
type BareMetalOperation struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}
