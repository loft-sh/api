package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// License holds the license information
// +k8s:openapi-gen=true
// +resource:path=licenses,rest=LicenseREST
type License struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LicenseSpec   `json:"spec,omitempty"`
	Status LicenseStatus `json:"status,omitempty"`
}

type LicenseSpec struct {
}

type LicenseStatus struct {
	// +optional
	Instance string `json:"instance"`

	// +optional
	Info LicenseInfo `json:"info"`
}

type LicenseInfo struct {
	Announcement   string            `json:"announcement,omitempty"`
	License        string            `json:"license,omitempty"`
	CurrentTime    int64             `json:"currentTime"`
	ResourceLimits []ResourceLimit   `json:"resourceLimits,omitempty"`
	BlockRequests  []ResoureRequests `json:"blockRequests,omitempty"`
	Features       map[string]bool   `json:"features,omitempty"`
	Customer       CustomerInfo      `json:"customer,omitempty"`
	Subscription   SubscriptionInfo  `json:"subscription,omitempty"`
	Quantity       int64             `json:"quantity"`
	Plan           Plan              `json:"plan,omitempty"`
	Promotions     Promotions        `json:"promotions,omitempty"`
	Analytics      Analytics         `json:"analytics,omitempty"`
	Links          map[string]string `json:"links,omitempty"`
	BaseDomains    []string          `json:"baseDomains,omitempty"`
}

type CustomerInfo struct {
	Company      string `json:"company,omitempty"`
	Email        string `json:"email,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	City         string `json:"city,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
	Country      string `json:"country,omitempty"`
	Created      int64  `json:"created"`
}

type Analytics struct {
	Endpoint string            `json:"endpoint,omitempty"`
	Requests []ResoureRequests `json:"requests,omitempty"`
	Token    string            `json:"token,omitempty"`
}

type Promotions struct {
	Trial *TrialPromotion `json:"trial,omitempty"`
}

type TrialPromotion struct {
	Product     string `json:"product"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Plan struct {
	Price    int64       `json:"price,omitempty"`
	Currency string      `json:"currency"`
	Interval string      `json:"interval"`
	Product  PlanProduct `json:"product"`
}

type PlanProduct struct {
	Name      string `json:"name"`
	UnitLabel string `json:"unitLabel"`
}

type SubscriptionInfo struct {
	TrialEnd             int64                 `json:"trialEnd"`
	Status               string                `json:"status"`
	CurrentPeriodEnd     int64                 `json:"currentPeriodEnd"`
	NextInvoice          int64                 `json:"nextInvoice"`
	DefaultPaymentMethod *DefaultPaymentMethod `json:"defaultPaymentMethod,omitempty"`
	Created              int64                 `json:"created"`
}

type DefaultPaymentMethod struct {
	Card *DefaultPaymentMethodCard `json:"card"`
}

type DefaultPaymentMethodCard struct {
	Last4    string `json:"last4"`
	ExpMonth uint64 `json:"expMonth"`
	ExpYear  uint64 `json:"expYear"`
	Brand    string `json:"brand"`
	Funding  string `json:"funding"`
}

type ResourceLimit struct {
	Group             string            `json:"group,omitempty"`
	Version           string            `json:"version,omitempty"`
	Kind              string            `json:"kind,omitempty"`
	Limit             int64             `json:"limit,omitempty"`
	AcrossAllClusters bool              `json:"acrossAllClusters,omitempty"`
	BlockRequests     []ResoureRequests `json:"blockRequests,omitempty"`
}

type ResoureRequests struct {
	Verbs      []string `json:"verbs,omitempty"`
	Group      string   `json:"group,omitempty"`
	Resource   string   `json:"resource,omitempty"`
	Management bool     `json:"management,omitempty"`
}
