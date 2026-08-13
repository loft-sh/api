package vclusterconfig

import (
	"testing"

	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestSupportsStandaloneSnapshots(t *testing.T) {
	tests := []struct {
		version string
		want    bool
	}{
		{version: "", want: false},
		{version: "0.30.0", want: false},
		{version: "0.36.5", want: false},
		{version: "0.36.9", want: false},
		// every 0.37 build qualifies, prerelease or not: they are one development stream
		{version: "0.37.0-alpha.0", want: true},
		{version: "0.37.0-alpha.1", want: true},
		{version: "0.37.0-beta.1", want: true},
		{version: "0.37.0-next.internal.8", want: true},
		{version: "0.37.0", want: true},
		{version: "0.37.2", want: true},
		{version: "0.38.0-alpha.1", want: true},
		// a registering vCluster reports its own version string, so the "v" may already be there
		{version: "v0.36.5", want: false},
		{version: "v0.37.0", want: true},
		{version: "v0.38.0", want: true},
		{version: "V0.37.0", want: true},
		{version: " 0.37.0 ", want: true},
		{version: "not-a-version", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.version, func(t *testing.T) {
			release := storagev1.VirtualClusterHelmRelease{
				Chart: storagev1.VirtualClusterHelmChart{Version: tt.version},
			}
			if got := SupportsStandaloneSnapshots(release); got != tt.want {
				t.Errorf("SupportsStandaloneSnapshots(%q) = %v, want %v", tt.version, got, tt.want)
			}
		})
	}
}

func TestValidateStandaloneSnapshots(t *testing.T) {
	fldPath := field.NewPath("spec")
	secret := &SnapshotSecretCredential{SecretName: "creds"}

	tests := []struct {
		name       string
		snapshots  *Snapshots
		wantFields []string
	}{
		{
			name:      "nil snapshots: no errors",
			snapshots: nil,
		},
		{
			name:      "auto snapshots without storage: no errors",
			snapshots: &Snapshots{Auto: &SnapshotsAuto{}},
		},
		{
			name:      "s3 with a credential secret is allowed",
			snapshots: standaloneSnapshots(&SnapshotStorage{Type: "s3", S3: SnapshotStorageS3{Credential: secret}}),
		},
		{
			name:      "oci with a credential secret is allowed",
			snapshots: standaloneSnapshots(&SnapshotStorage{Type: "oci", OCI: SnapshotStorageOCI{Credential: secret}}),
		},
		{
			// an inline username means the credentials are in the config itself, so the platform
			// has everything it needs without a Secret
			name:      "oci with an inline username is allowed",
			snapshots: standaloneSnapshots(&SnapshotStorage{Type: "oci", OCI: SnapshotStorageOCI{Username: "robot"}}),
		},
		{
			name:      "azure with a credential secret is allowed",
			snapshots: standaloneSnapshots(&SnapshotStorage{Type: "azure", Azure: SnapshotStorageAzure{Credential: secret}}),
		},
		{
			// the volume lives inside the tenant, which the platform can neither list nor prune
			name:       "container storage is rejected",
			snapshots:  standaloneSnapshots(&SnapshotStorage{Type: "container"}),
			wantFields: []string{"spec.snapshots.auto.storage.type"},
		},
		{
			// no credential means the tenant would rely on an instance profile or workload
			// identity, which the platform process does not have
			name:       "s3 without a credential is rejected",
			snapshots:  standaloneSnapshots(&SnapshotStorage{Type: "s3"}),
			wantFields: []string{"spec.snapshots.auto.storage.s3.credential"},
		},
		{
			name:       "oci without a credential or username is rejected",
			snapshots:  standaloneSnapshots(&SnapshotStorage{Type: "oci"}),
			wantFields: []string{"spec.snapshots.auto.storage.oci.credential"},
		},
		{
			name:       "azure without a credential is rejected",
			snapshots:  standaloneSnapshots(&SnapshotStorage{Type: "azure"}),
			wantFields: []string{"spec.snapshots.auto.storage.azure.credential"},
		},
		{
			// the blob URL is the snapshot location, so it is written into the tenant's request
			// ConfigMap; a SAS token there would persist a credential in the tenant in plaintext
			name: "azure with a SAS token in the blob URL is rejected",
			snapshots: standaloneSnapshots(&SnapshotStorage{Type: "azure", Azure: SnapshotStorageAzure{
				Credential: secret,
				BlobURL:    "https://acct.blob.core.windows.net/backups?sv=2022-11-02&ss=b&sig=REDACTED%3D",
			}}),
			wantFields: []string{"spec.snapshots.auto.storage.azure.blobUrl"},
		},
		{
			name: "azure with a plain blob URL is allowed",
			snapshots: standaloneSnapshots(&SnapshotStorage{Type: "azure", Azure: SnapshotStorageAzure{
				Credential: secret,
				BlobURL:    "https://acct.blob.core.windows.net/backups",
			}}),
		},
		{
			// the credential rule runs first, so a URL carrying a token does not mask the missing
			// Secret; both are wrong and the Secret is the one to fix
			name: "azure with a SAS token but no credential reports the credential",
			snapshots: standaloneSnapshots(&SnapshotStorage{Type: "azure", Azure: SnapshotStorageAzure{
				BlobURL: "https://acct.blob.core.windows.net/backups?sig=REDACTED%3D",
			}}),
			wantFields: []string{"spec.snapshots.auto.storage.azure.credential"},
		},
		{
			// a SAS token is only a problem for standalone; pod-based resolves the URL itself and
			// nothing writes it into a ConfigMap, so the same rule must not fire on other backends
			name: "s3 url with a sig query is not treated as azure",
			snapshots: standaloneSnapshots(&SnapshotStorage{Type: "s3", S3: SnapshotStorageS3{
				Url:        "s3://my-bucket/snapshots?sig=whatever",
				Credential: secret,
			}}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := ValidateStandaloneSnapshots(fldPath, tt.snapshots)
			if len(errs) != len(tt.wantFields) {
				t.Fatalf("got %d errors %v, want %d on %v", len(errs), errs, len(tt.wantFields), tt.wantFields)
			}
			for i, want := range tt.wantFields {
				if errs[i].Field != want {
					t.Errorf("error %d on field %q, want %q", i, errs[i].Field, want)
				}
			}
		})
	}
}

func standaloneSnapshots(storage *SnapshotStorage) *Snapshots {
	return &Snapshots{Auto: &SnapshotsAuto{Storage: storage}}
}
