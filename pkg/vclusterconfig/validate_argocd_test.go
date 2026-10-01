package vclusterconfig

import (
	"strings"
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

// TestValidateArgoCDClusterMetadata checks that cluster metadata is validated
// even when no applications are set.
func TestValidateArgoCDClusterMetadata(t *testing.T) {
	tests := []struct {
		name       string
		metadata   ArgoCDClusterMetadata
		wantErrs   int
		wantSubstr string
	}{
		{
			name:     "valid metadata passes",
			metadata: ArgoCDClusterMetadata{Labels: map[string]string{"env": "prod"}, Annotations: map[string]string{"owner": "team"}},
		},
		{
			name:       "invalid label key is rejected",
			metadata:   ArgoCDClusterMetadata{Labels: map[string]string{"not a key": "prod"}},
			wantErrs:   1,
			wantSubstr: "integrations.argoCD.cluster.metadata.labels",
		},
		{
			name:       "invalid label value is rejected",
			metadata:   ArgoCDClusterMetadata{Labels: map[string]string{"env": "not a value!"}},
			wantErrs:   1,
			wantSubstr: "integrations.argoCD.cluster.metadata.labels",
		},
		{
			name:       "invalid annotation key is rejected",
			metadata:   ArgoCDClusterMetadata{Annotations: map[string]string{"not a key": "x"}},
			wantErrs:   1,
			wantSubstr: "integrations.argoCD.cluster.metadata.annotations",
		},
		{
			name:     "free form annotation value is allowed",
			metadata: ArgoCDClusterMetadata{Annotations: map[string]string{"owner": "anything at all!"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			integration := &ArgoCDIntegration{
				Enabled:   true,
				Connector: "my-argo",
				Cluster:   &ArgoCDIntegrationCluster{Metadata: test.metadata},
			}

			// deploy is nil on purpose: metadata must still be validated.
			errs := ValidateArgoCD(field.NewPath("spec"), integration, nil)
			if len(errs) != test.wantErrs {
				t.Fatalf("got %d errors, want %d: %v", len(errs), test.wantErrs, errs)
			}
			if test.wantSubstr != "" && !strings.Contains(errs.ToAggregate().Error(), test.wantSubstr) {
				t.Errorf("error %q does not mention %q", errs.ToAggregate(), test.wantSubstr)
			}
		})
	}
}

// TestValidateArgoCDWithoutClusterBlockIsUnchanged checks that an integration with no
// cluster block and no applications still has no errors.
func TestValidateArgoCDWithoutClusterBlockIsUnchanged(t *testing.T) {
	integration := &ArgoCDIntegration{Enabled: true, Connector: "my-argo"}
	if errs := ValidateArgoCD(field.NewPath("spec"), integration, nil); len(errs) != 0 {
		t.Errorf("got %v, want no errors", errs)
	}
}
