package auth

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const GroupVersion = "authentication.loft.sh/v1"

// OIDCTokenRequest is used by the /auth/oidc/token route
type OIDCTokenRequest struct {
	Token       string `json:"token,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}

// OIDCRefreshRequest is used by the /auth/oidc/refresh route
type OIDCRefreshRequest struct {
	RefreshToken string `json:"refreshToken,omitempty"`
}

// TokenRequest is used by the /auth/token route
type TokenRequest struct {
	Key string `json:"key,omitempty"`
}

// PasswordLoginRequest is used by the /auth/password/login route
type PasswordLoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Token string `json:"token"`
}

type AccessKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	User      string `json:"user"`
	Username  string `json:"username"`
	AccessKey string `json:"accessKey"`
}

type OIDCRedirect struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Redirect string `json:"redirect,omitempty"`
}

type OIDCToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	IDToken      string `json:"idToken"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Info struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Methods InfoMethods `json:"methods,omitempty"`
}

type InfoMethods struct {
	SSO      *MethodSSO      `json:"sso,omitempty"`
	Password *MethodPassword `json:"password,omitempty"`
}

type MethodSSO struct {
	// Indicates if the authentication method is enabled
	Enabled bool `json:"enabled,omitempty"`

	// Type of the SSO to show in the UI. Only for displaying purposes
	Type string `json:"type,omitempty"`

	// LoginEndpoint is the path the UI will request a login url from
	LoginEndpoint string `json:"loginEndpoint,omitempty"`
}

type MethodPassword struct {
	// Indicates if the authentication method is enabled
	Enabled bool `json:"enabled,omitempty"`
}
