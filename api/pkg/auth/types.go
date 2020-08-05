package auth

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const GroupVersion = "authentication.loft.sh/v1"

type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Token string `json:"token"`
}

type AccessKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

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
	OIDC     *MethodOIDC     `json:"oidc,omitempty"`
	Password *MethodPassword `json:"password,omitempty"`
}

type MethodOIDC struct {
	// Indicates if the authentication method is enabled
	Enabled bool `json:"enabled,omitempty"`

	// Type of the OIDC to show in the UI. Only for displaying purposes
	Type string `json:"type,omitempty"`
}

type MethodPassword struct {
	// Indicates if the authentication method is enabled
	Enabled bool `json:"enabled,omitempty"`
}
