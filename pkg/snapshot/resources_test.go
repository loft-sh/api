package snapshot

import (
	"encoding/json"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNewSnapshotRequestResources(t *testing.T) {
	options := &Options{
		Type: "s3",
		S3: S3Options{
			Bucket: "bucket",
			Key:    "snapshots/test.tar.gz",
		},
		IncludeVolumes: true,
	}

	secret, err := NewSnapshotOptionsSecret("vcluster-ns", "my-vcluster", options)
	if err != nil {
		t.Fatal(err)
	}
	if secret.Namespace != "vcluster-ns" {
		t.Fatalf("expected namespace vcluster-ns, got %q", secret.Namespace)
	}
	if secret.GenerateName != "my-vcluster-snapshot-request-" {
		t.Fatalf("expected generate name my-vcluster-snapshot-request-, got %q", secret.GenerateName)
	}
	if secret.Labels[SnapshotRequestLabel] != "" {
		t.Fatalf("expected snapshot request label")
	}
	if secret.Labels[VClusterNameLabel] != "my-vcluster" {
		t.Fatalf("expected vCluster name label my-vcluster, got %q", secret.Labels[VClusterNameLabel])
	}
	if secret.Labels[VClusterNamespaceLabel] != "vcluster-ns" {
		t.Fatalf("expected vCluster namespace label vcluster-ns, got %q", secret.Labels[VClusterNamespaceLabel])
	}

	var storedOptions Options
	err = json.Unmarshal(secret.Data[OptionsKey], &storedOptions)
	if err != nil {
		t.Fatal(err)
	}
	if storedOptions.GetURL() != "s3://bucket/snapshots/test.tar.gz" {
		t.Fatalf("expected stored options URL, got %q", storedOptions.GetURL())
	}

	now := metav1.Now()
	request := NewRequest("my-vcluster-snapshot-request-abcde", now, options)
	configMap, err := NewSnapshotRequestConfigMap("vcluster-ns", "my-vcluster", request)
	if err != nil {
		t.Fatal(err)
	}
	configMap.Name = request.Name

	if configMap.Name != request.Name {
		t.Fatalf("expected configmap name %q, got %q", request.Name, configMap.Name)
	}
	if configMap.Namespace != "vcluster-ns" {
		t.Fatalf("expected configmap namespace vcluster-ns, got %q", configMap.Namespace)
	}
	if configMap.Labels[SnapshotRequestLabel] != "" {
		t.Fatalf("expected snapshot request label")
	}

	storedRequest, err := UnmarshalRequest(configMap)
	if err != nil {
		t.Fatal(err)
	}
	if storedRequest.Name != request.Name {
		t.Fatalf("expected request name %q, got %q", request.Name, storedRequest.Name)
	}
	if storedRequest.Spec.URL != "s3://bucket/snapshots/test.tar.gz" {
		t.Fatalf("expected request URL, got %q", storedRequest.Spec.URL)
	}
	if !storedRequest.Spec.IncludeVolumes {
		t.Fatalf("expected include volumes")
	}
}

func TestUnmarshalOptionsAcceptsRestoreRequestLabel(t *testing.T) {
	options := &Options{
		Type: "s3",
		S3: S3Options{
			Bucket: "bucket",
			Key:    "snapshots/restore.tar.gz",
		},
	}

	secret, err := NewOptionsSecret(RestoreRequestLabel, "vcluster-ns", "my-vcluster", options)
	if err != nil {
		t.Fatal(err)
	}

	storedOptions, err := UnmarshalOptions(secret)
	if err != nil {
		t.Fatal(err)
	}
	if storedOptions.GetURL() != "s3://bucket/snapshots/restore.tar.gz" {
		t.Fatalf("expected stored options URL, got %q", storedOptions.GetURL())
	}
}
