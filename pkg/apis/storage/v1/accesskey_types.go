package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessKey holds the session information
// +k8s:openapi-gen=true
type AccessKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessKeySpec   `json:"spec,omitempty"`
	Status AccessKeyStatus `json:"status,omitempty"`
}

type AccessKeySpec struct {
	// The user this access key refers to
	User string `json:"user,omitempty"`

	// The actual access key that will be used as a bearer token
	Key string `json:"key,omitempty"`
	
	// Parent is used to share OIDC and external token information
	// with multiple access keys. Since copying an OIDC refresh token
	// would result in the other access keys becoming invalid after a refresh
	// parent allows access keys to share that information. 
	// 
	// The use case for this is primarily user generated access keys,
	// which will have the users current access key as parent if it contains
	// an OIDC token.
	// +optional
	Parent string `json:"parent,omitempty"`
	
	// If this field is true, the access key is still allowed to exist,
	// however will not work to access the api
	// +optional
	Disabled bool `json:"disabled,omitempty"`
	
	// The display name shown in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	
	// The time to life for this access key
	// +optional 
	TTL int64 `json:"ttl,omitempty"`
	
	// If this is specified, the time to life for this access key will
	// start after the lastActivity instead of creation timestamp
	// +optional
	TTLAfterLastActivity bool `json:"ttlAfterLastActivity,omitempty"`
	
	// The type of an access key, which basically describes if the access 
	// key is user managed or managed by loft itself.
	// +optional
	Type AccessKeyType `json:"type,omitempty"`
	
	// If available, contains information about the oidc login data for this 
	// access key
	// +optional
	OIDCLogin *AccessKeyOIDC `json:"oidcLogin,omitempty"`
	
	// If the token is a refresh token, contains information about it
	// +optional
	OIDCProvider *AccessKeyOIDCProvider `json:"oidcProvider,omitempty"`
}

type AccessKeyOIDCProvider struct {
	// ClientId the token was generated for
	// +optional
	ClientId string `json:"clientId,omitempty"`
	
	// Nonce to use
	// +optional
	Nonce string `json:"nonce,omitempty"`

	// RedirectUri to use
	// +optional
	RedirectUri string `json:"redirectUri,omitempty"`
	
	// Scopes to use
	// +optional
	Scopes string `json:"scopes,omitempty"`
}

type AccessKeyOIDC struct {
	// The current id token that was created during login
	// +optional
	IDToken []byte `json:"idToken,omitempty"`
	
	// The current access token that was created during login
	// +optional
	AccessToken []byte `json:"accessToken,omitempty"`
	
	// The current refresh token that was created during login
	// +optional
	RefreshToken []byte `json:"refreshToken,omitempty"`
	
	// The last time the id token was refreshed
	// +optional
	LastRefresh *metav1.Time `json:"lastRefresh,omitempty"`
}

// AccessKeyType describes the type of an access key
type AccessKeyType string

// These are the valid access key types
const (
	AccessKeyTypeNone             AccessKeyType = ""
	AccessKeyTypeLogin            AccessKeyType = "Login"
	AccessKeyTypeUser             AccessKeyType = "User"
	AccessKeyTypeReset            AccessKeyType = "Reset"
	AccessKeyTypeOIDCRefreshToken AccessKeyType = "OIDCRefreshToken"
)

// AccessKeyStatus holds the status of an access key
type AccessKeyStatus struct {
	// The last time this access key was used to access the api
	// +optional
	LastActivity *metav1.Time `json:"lastActivity,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessKeyList contains a list of AccessKey
type AccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessKey{}, &AccessKeyList{})
}
