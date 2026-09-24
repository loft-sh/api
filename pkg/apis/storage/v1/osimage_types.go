package v1

import (
	"slices"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OSImage holds the information of machine networks
// +k8s:openapi-gen=true
type OSImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OSImageSpec   `json:"spec,omitempty"`
	Status OSImageStatus `json:"status,omitempty"`
}

func (a *OSImage) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *OSImage) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *OSImage) GetAccess() []Access {
	return a.Spec.Access
}

func (a *OSImage) SetAccess(access []Access) {
	a.Spec.Access = access
}

type OSImageSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes an OS image
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Access holds the access rights for users and teams
	// +optional
	Access []Access `json:"access,omitempty"`

	// Properties is the configuration for the OS image
	// +optional
	Properties map[string]string `json:"properties,omitempty"`

	// ConnectorRef names the image store connector holding this image's blob.
	// Empty means a properties-only image. Immutable once set.
	// +optional
	ConnectorRef string `json:"connectorRef,omitempty"`

	// Format is the on-disk format of the image blob and part of the object key.
	// Immutable once set.
	// +optional
	Format OSImageFormat `json:"format,omitempty"`
}

// OSImageFormat is the on-disk format of an OS image blob.
// +kubebuilder:validation:Enum=qcow2
type OSImageFormat string

const OSImageFormatQCOW2 OSImageFormat = "qcow2"

const (
	// OSImageUploadSessionIDAnnotation holds the multipart upload ID of the open session.
	OSImageUploadSessionIDAnnotation = "machines.vcluster.com/os-image-upload-session-id"

	// OSImageUploadDeclaredSizeAnnotation holds the byte size the session was opened for.
	OSImageUploadDeclaredSizeAnnotation = "machines.vcluster.com/os-image-upload-declared-size"

	// OSImageUploadLastInitiatedAtAnnotation holds the RFC3339 time of the last upload call.
	OSImageUploadLastInitiatedAtAnnotation = "machines.vcluster.com/os-image-upload-last-initiated-at"

	// OSImageUploadCompletedAtAnnotation holds the RFC3339 time the object was completed.
	OSImageUploadCompletedAtAnnotation = "machines.vcluster.com/os-image-upload-completed-at"

	// OSImageUploadChecksumAnnotation holds the sha256 the client declared at finalize.
	OSImageUploadChecksumAnnotation = "machines.vcluster.com/os-image-upload-checksum"
)

// OSImageUploadInitiateAnnotations is what opening a session writes.
var OSImageUploadInitiateAnnotations = []string{
	OSImageUploadSessionIDAnnotation,
	OSImageUploadDeclaredSizeAnnotation,
	OSImageUploadLastInitiatedAtAnnotation,
}

// OSImageUploadFinalizeAnnotations is what completing the object writes.
var OSImageUploadFinalizeAnnotations = []string{
	OSImageUploadCompletedAtAnnotation,
	OSImageUploadChecksumAnnotation,
}

// OSImageUploadSessionAnnotations is every annotation one upload session writes, so cleanup
// clears this set rather than a list of its own.
var OSImageUploadSessionAnnotations = slices.Concat(
	OSImageUploadInitiateAnnotations,
	OSImageUploadFinalizeAnnotations,
)

type OSImageStatus struct{}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OSImageList contains a list of OSImages
type OSImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OSImage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OSImage{}, &OSImageList{})
}
