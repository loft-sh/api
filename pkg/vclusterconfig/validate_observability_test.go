package vclusterconfig

import (
	"reflect"
	"testing"

	"github.com/loft-sh/api/v5/pkg/vclusterconfig/constants"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestValidateObservability(t *testing.T) {
	fldPath := field.NewPath("spec")

	tests := []struct {
		name        string
		integration *ObservabilityIntegration
		wantFields  []string
	}{
		{
			name:        "nil integration: no errors",
			integration: nil,
			wantFields:  nil,
		},
		{
			name:        "disabled: no errors even without connector",
			integration: &ObservabilityIntegration{Enabled: false},
			wantFields:  nil,
		},
		{
			name:        "valid: enabled with connector, no overrides (happy path)",
			integration: &ObservabilityIntegration{Enabled: true, Connector: "fleet-obs"},
			wantFields:  nil,
		},
		{
			name:        "error: enabled without connector",
			integration: &ObservabilityIntegration{Enabled: true},
			wantFields:  []string{"spec.integrations.observability.connector"},
		},
		{
			name:        "error: enabled with a whitespace-only connector",
			integration: &ObservabilityIntegration{Enabled: true, Connector: "   "},
			wantFields:  []string{"spec.integrations.observability.connector"},
		},
		{
			name: "valid: enabled with connector and override target",
			integration: &ObservabilityIntegration{
				Enabled:       true,
				Connector:     "fleet-obs",
				GatewaySecret: &GatewaySecret{Namespace: "my-obs", Name: "metrics-writer"},
			},
			wantFields: nil,
		},
		{
			name: "error: override entry missing namespace and name",
			integration: &ObservabilityIntegration{
				Enabled:       true,
				Connector:     "fleet-obs",
				GatewaySecret: &GatewaySecret{},
			},
			wantFields: []string{
				"spec.integrations.observability.gatewaySecret.namespace",
				"spec.integrations.observability.gatewaySecret.name",
			},
		},
		{
			name: "error: override entry missing name only",
			integration: &ObservabilityIntegration{
				Enabled:       true,
				Connector:     "fleet-obs",
				GatewaySecret: &GatewaySecret{Namespace: "my-obs"},
			},
			wantFields: []string{"spec.integrations.observability.gatewaySecret.name"},
		},
		{
			name: "error: override entry with a whitespace-only namespace and name",
			integration: &ObservabilityIntegration{
				Enabled:       true,
				Connector:     "fleet-obs",
				GatewaySecret: &GatewaySecret{Namespace: " ", Name: "\t"},
			},
			wantFields: []string{
				"spec.integrations.observability.gatewaySecret.namespace",
				"spec.integrations.observability.gatewaySecret.name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := ValidateObservability(fldPath, tt.integration)
			if len(errs) != len(tt.wantFields) {
				t.Fatalf("ValidateObservability() returned %d errors, want %d: %v", len(errs), len(tt.wantFields), errs)
			}
			for i, err := range errs {
				if err.Field != tt.wantFields[i] {
					t.Errorf("error[%d].Field = %q, want %q", i, err.Field, tt.wantFields[i])
				}
			}
		})
	}
}

func TestObservabilityDeliveryTarget(t *testing.T) {
	defaultTarget := &GatewaySecret{
		Namespace: constants.DefaultObservabilityNamespace,
		Name:      constants.DefaultMetricsWriterSecretName,
	}

	tests := []struct {
		name        string
		integration *ObservabilityIntegration
		want        *GatewaySecret
	}{
		{
			name:        "nil receiver: no target",
			integration: nil,
			want:        nil,
		},
		{
			name:        "disabled: no target",
			integration: &ObservabilityIntegration{Enabled: false},
			want:        nil,
		},
		{
			name: "disabled with override: no target",
			integration: &ObservabilityIntegration{
				Enabled:       false,
				GatewaySecret: &GatewaySecret{Namespace: "a", Name: "x"},
			},
			want: nil,
		},
		{
			name:        "enabled without a connector: no target",
			integration: &ObservabilityIntegration{Enabled: true},
			want:        nil,
		},
		{
			name:        "enabled with a whitespace-only connector: no target",
			integration: &ObservabilityIntegration{Enabled: true, Connector: "   "},
			want:        nil,
		},
		{
			name: "enabled with a whitespace-only connector and an override: no target",
			integration: &ObservabilityIntegration{
				Enabled:       true,
				Connector:     " ",
				GatewaySecret: &GatewaySecret{Namespace: "a", Name: "x"},
			},
			want: nil,
		},
		{
			name:        "enabled without override: default target",
			integration: &ObservabilityIntegration{Enabled: true, Connector: "fleet-obs"},
			want:        defaultTarget,
		},
		{
			name: "enabled with override: returned verbatim",
			integration: &ObservabilityIntegration{
				Enabled:       true,
				Connector:     "fleet-obs",
				GatewaySecret: &GatewaySecret{Namespace: "a", Name: "x"},
			},
			want: &GatewaySecret{Namespace: "a", Name: "x"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.integration.DeliveryTarget()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeliveryTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
