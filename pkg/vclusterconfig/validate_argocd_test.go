package vclusterconfig

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func validArgoCDIntegration() *ArgoCDIntegration {
	return &ArgoCDIntegration{Enabled: true, Connector: "my-connector"}
}

func argoCDDeploy(apps ...ArgoCDApplication) *ArgoCDDeploy {
	return &ArgoCDDeploy{Applications: apps}
}

func argoCDApp(name, displayName string) ArgoCDApplication {
	return ArgoCDApplication{
		Name:        name,
		DisplayName: displayName,
		Template:    &ArgoCDApplicationTemplate{Name: "tmpl"},
	}
}

func TestValidateArgoCD(t *testing.T) {
	fldPath := field.NewPath("spec")

	tests := []struct {
		name        string
		integration *ArgoCDIntegration
		deploy      *ArgoCDDeploy
		wantErrs    int
	}{
		{
			name:        "nil deploy — no errors",
			integration: validArgoCDIntegration(),
			deploy:      nil,
			wantErrs:    0,
		},
		{
			name:        "empty applications — no errors",
			integration: validArgoCDIntegration(),
			deploy:      &ArgoCDDeploy{},
			wantErrs:    0,
		},
		{
			name:        "valid — name only",
			integration: validArgoCDIntegration(),
			deploy:      argoCDDeploy(argoCDApp("nginx", "")),
			wantErrs:    0,
		},
		{
			name:        "valid — displayName only",
			integration: validArgoCDIntegration(),
			deploy:      argoCDDeploy(argoCDApp("", "Nginx")),
			wantErrs:    0,
		},
		{
			name:        "valid — name and displayName",
			integration: validArgoCDIntegration(),
			deploy:      argoCDDeploy(argoCDApp("nginx", "Nginx")),
			wantErrs:    0,
		},
		{
			name:        "error — neither name nor displayName",
			integration: validArgoCDIntegration(),
			deploy:      argoCDDeploy(argoCDApp("", "")),
			wantErrs:    1,
		},
		{
			name:        "error — whitespace-only name treated as absent",
			integration: validArgoCDIntegration(),
			deploy:      argoCDDeploy(argoCDApp("  ", "")),
			wantErrs:    1,
		},
		{
			name:        "error — duplicate names",
			integration: validArgoCDIntegration(),
			deploy:      argoCDDeploy(argoCDApp("nginx", ""), argoCDApp("nginx", "other")),
			wantErrs:    1,
		},
		{
			name:        "error — duplicate displayNames",
			integration: validArgoCDIntegration(),
			deploy:      argoCDDeploy(argoCDApp("", "Nginx"), argoCDApp("other", "Nginx")),
			wantErrs:    1,
		},
		{
			name:        "error — nil integration produces two errors",
			integration: nil,
			deploy:      argoCDDeploy(argoCDApp("nginx", "")),
			wantErrs:    2,
		},
		{
			name:        "error — integration disabled",
			integration: &ArgoCDIntegration{Enabled: false, Connector: "c"},
			deploy:      argoCDDeploy(argoCDApp("nginx", "")),
			wantErrs:    1,
		},
		{
			name:        "error — missing connector",
			integration: &ArgoCDIntegration{Enabled: true},
			deploy:      argoCDDeploy(argoCDApp("nginx", "")),
			wantErrs:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := ValidateArgoCD(fldPath, tt.integration, tt.deploy)
			if len(errs) != tt.wantErrs {
				t.Errorf("ValidateArgoCD() returned %d errors, want %d: %v", len(errs), tt.wantErrs, errs)
			}
		})
	}
}
