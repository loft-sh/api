package openapi

import "testing"

func TestAppInstanceDestinationNamespaceDocumentation(t *testing.T) {
	tests := map[string]struct {
		description string
		want        string
	}{
		"connected cluster": {
			description: schema_pkg_apis_storage_v1_AppInstanceDestinationCluster(nil).Schema.Properties["namespace"].Description,
			want:        "Namespace in the cluster the helm release is deployed into. If empty, uses the app's default namespace only for initial release resolution. Existing instances retain their recorded release coordinates.",
		},
		"virtual cluster": {
			description: schema_pkg_apis_storage_v1_AppInstanceDestinationVirtualCluster(nil).Schema.Properties["namespace"].Description,
			want:        "Namespace the helm release is deployed into. Only used when target is vCluster; for the host target the release is always deployed into the virtual cluster's host namespace. If empty, uses the app's default namespace only for initial release resolution. Existing instances retain their recorded release coordinates.",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.description != tt.want {
				t.Errorf("namespace description = %q, want %q", tt.description, tt.want)
			}
		})
	}
}
