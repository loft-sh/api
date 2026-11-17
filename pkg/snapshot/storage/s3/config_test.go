package s3

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/go-logr/logr"
	snapshotapi "github.com/loft-sh/api/v4/pkg/snapshot"
)

// TestWithStaticCredentials pins that credentials stay on the config. The platform builds one store per
// instance in a shared process, so credentials must never reach the process environment.
func TestWithStaticCredentials(t *testing.T) {
	cfg, err := newConfigBuilder(logr.Discard()).WithStaticCredentials("id", "secret", "token").Build()
	if err != nil {
		t.Fatalf("build config: %v", err)
	}

	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("retrieve credentials: %v", err)
	}
	if creds.AccessKeyID != "id" || creds.SecretAccessKey != "secret" || creds.SessionToken != "token" {
		t.Errorf("got credentials %q/%q/%q, want id/secret/token", creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken)
	}

	for _, key := range []string{"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_SESSION_TOKEN"} {
		if value := os.Getenv(key); value != "" {
			t.Errorf("%s leaked into the process environment as %q", key, value)
		}
	}
}

// TestWithStaticCredentialsNone: no credentials is how a caller asks for the default chain, so it must
// stay a no-op rather than an error.
func TestWithStaticCredentialsNone(t *testing.T) {
	t.Setenv("AWS_ACCESS_KEY_ID", "ambient-id")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "ambient-secret")

	cfg, err := newConfigBuilder(logr.Discard()).WithStaticCredentials("", "", "").Build()
	if err != nil {
		t.Fatalf("build config: %v", err)
	}

	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("retrieve credentials: %v", err)
	}
	if creds.AccessKeyID != "ambient-id" {
		t.Errorf("got access key %q, want the default chain to resolve it", creds.AccessKeyID)
	}
}

// TestWithStaticCredentialsPartialFailsClosed: a half-filled credential is always a mistake, and falling
// back would resolve the running process's own identity against a caller-chosen bucket.
func TestWithStaticCredentialsPartialFailsClosed(t *testing.T) {
	t.Setenv("AWS_ACCESS_KEY_ID", "ambient-id")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "ambient-secret")

	for _, tt := range []struct {
		name            string
		accessKeyID     string
		secretAccessKey string
	}{
		{name: "only the access key id", accessKeyID: "id"},
		{name: "only the secret access key", secretAccessKey: "secret"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newConfigBuilder(logr.Discard()).WithStaticCredentials(tt.accessKeyID, tt.secretAccessKey, "").Build()
			if err == nil {
				t.Fatal("expected a partial credential to fail rather than resolve the ambient identity")
			}
			if !strings.Contains(err.Error(), "incomplete") {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// TestInitLeavesIRSAEnvAloneWithStaticCredentials pins that Init does not touch the process environment
// when it has credentials to pin. WithCredentialsFile clears the IRSA variables process-wide, which in a
// shared process disables IRSA for every other AWS client; the platform builds a store per instance in
// its own process, so that must not happen on a path that never needed the file.
func TestInitLeavesIRSAEnvAloneWithStaticCredentials(t *testing.T) {
	// the ambient fallback inside WithCredentialsFile is what makes this reachable without any
	// credentials file being configured
	t.Setenv("AWS_SHARED_CREDENTIALS_FILE", "/tmp/does-not-need-to-exist")
	t.Setenv("AWS_WEB_IDENTITY_TOKEN_FILE", "/var/run/secrets/token")
	t.Setenv("AWS_ROLE_ARN", "arn:aws:iam::123456789012:role/platform")
	t.Setenv("AWS_ROLE_SESSION_NAME", "platform")

	store := NewStore(logr.Discard())
	_ = store.Init(&snapshotapi.S3Options{
		Region:          "eu-west-1",
		AccessKeyID:     "id",
		SecretAccessKey: "secret",
	})

	for _, key := range []string{"AWS_WEB_IDENTITY_TOKEN_FILE", "AWS_ROLE_ARN", "AWS_ROLE_SESSION_NAME"} {
		if os.Getenv(key) == "" {
			t.Errorf("%s was cleared process-wide, which disables IRSA for every other client", key)
		}
	}
}

// TestInitStillReadsCredentialsFileWithoutStaticCredentials pins the other half: with nothing to pin, the
// credentials file is still the way a caller supplies credentials, so that path is unchanged.
func TestInitStillReadsCredentialsFileWithoutStaticCredentials(t *testing.T) {
	// without this the SDK falls through to EC2 IMDS and the test waits on a network timeout
	t.Setenv("AWS_EC2_METADATA_DISABLED", "true")
	t.Setenv("AWS_SHARED_CREDENTIALS_FILE", "/tmp/does-not-need-to-exist")
	t.Setenv("AWS_WEB_IDENTITY_TOKEN_FILE", "/var/run/secrets/token")

	store := NewStore(logr.Discard())
	_ = store.Init(&snapshotapi.S3Options{Region: "eu-west-1"})

	if os.Getenv("AWS_WEB_IDENTITY_TOKEN_FILE") != "" {
		t.Error("the credentials-file path still clears IRSA, so it was not consulted")
	}
}
