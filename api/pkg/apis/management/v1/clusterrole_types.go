package v1

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User holds the user information
// +k8s:openapi-gen=true
// +resource:path=clusterroles,rest=ClusterRoleREST
type ClusterRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterRoleSpec   `json:"spec,omitempty"`
	Status ClusterRoleStatus `json:"status,omitempty"`
}

// ClusterRoleSpec holds the user specification
type ClusterRoleSpec struct {
	// Rules holds all the PolicyRules for this ClusterRole
	// +optional
	Rules []rbacv1.PolicyRule `json:"rules"`

	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole.
	// If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be
	// stomped by the controller.
	// +optional
	AggregationRule *rbacv1.AggregationRule `json:"aggregationRule,omitempty"`
}

// ClusterRoleStatus holds the user status
type ClusterRoleStatus struct {
}
