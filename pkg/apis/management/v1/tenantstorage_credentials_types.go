package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TenantStorageCredentials is the S3 access to one tenant storage: where its bucket is and the
// tenant's S3 keys. Served by GET on the credentials subresource to callers that can see the
// storage and hold get on tenantstorages/credentials. Every read is audited.
// +subresource-request
type TenantStorageCredentials struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Endpoint is the S3 endpoint URL. Clients address the bucket path-style.
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// Bucket is the storage's bucket.
	// +optional
	Bucket string `json:"bucket,omitempty"`

	// AccessKey is the tenant's S3 access key.
	// +optional
	AccessKey string `json:"accessKey,omitempty"`

	// SecretKey is the tenant's S3 secret key.
	// +optional
	SecretKey string `json:"secretKey,omitempty"`

	// CAData is the PEM bundle that verifies the endpoint, empty when the endpoint needs none.
	// +optional
	CAData string `json:"caData,omitempty"`
}
