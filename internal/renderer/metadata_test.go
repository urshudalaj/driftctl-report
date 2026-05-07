package renderer

import (
	"strings"
	"testing"
	"time"
)

func TestBuildMetadata_DefaultVersion(t *testing.T) {
	opts := DefaultOptions()
	m := buildMetadata(opts)

	if !strings.HasSuffix(m.GeneratedBy, "dev") {
		t.Errorf("expected GeneratedBy to end with 'dev', got %q", m.GeneratedBy)
	}
}

func TestBuildMetadata_CustomVersion(t *testing.T) {
	opts := DefaultOptions()
	opts.Version = "1.2.3"
	m := buildMetadata(opts)

	if !strings.Contains(m.GeneratedBy, "1.2.3") {
		t.Errorf("expected GeneratedBy to contain version, got %q", m.GeneratedBy)
	}
}

func TestBuildMetadata_DefaultHostname(t *testing.T) {
	opts := DefaultOptions()
	m := buildMetadata(opts)

	if m.Hostname != "unknown" {
		t.Errorf("expected default hostname 'unknown', got %q", m.Hostname)
	}
}

func TestBuildMetadata_CustomHostname(t *testing.T) {
	opts := DefaultOptions()
	opts.Hostname = "ci-runner-01"
	m := buildMetadata(opts)

	if m.Hostname != "ci-runner-01" {
		t.Errorf("expected hostname 'ci-runner-01', got %q", m.Hostname)
	}
}

func TestBuildMetadata_GeneratedAtIsRFC3339(t *testing.T) {
	opts := DefaultOptions()
	m := buildMetadata(opts)

	_, err := time.Parse(time.RFC3339, m.GeneratedAt)
	if err != nil {
		t.Errorf("GeneratedAt %q is not valid RFC3339: %v", m.GeneratedAt, err)
	}
}

func TestBuildMetadata_InputFile(t *testing.T) {
	opts := DefaultOptions()
	opts.InputFile = "/tmp/drift.json"
	m := buildMetadata(opts)

	if m.InputFile != "/tmp/drift.json" {
		t.Errorf("expected InputFile '/tmp/drift.json', got %q", m.InputFile)
	}
}

func TestBuildMetadata_GoVersionNonEmpty(t *testing.T) {
	opts := DefaultOptions()
	m := buildMetadata(opts)

	if m.GoVersion == "" {
		t.Error("expected non-empty GoVersion")
	}
}
