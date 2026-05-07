package renderer

import "testing"

func TestWithVersion_SetsVersion(t *testing.T) {
	opts := DefaultOptions()
	WithVersion("2.0.0")(&opts)

	if opts.Version != "2.0.0" {
		t.Errorf("expected Version '2.0.0', got %q", opts.Version)
	}
}

func TestWithVersion_EmptyIgnored(t *testing.T) {
	opts := DefaultOptions()
	opts.Version = "1.0.0"
	WithVersion("")(&opts)

	if opts.Version != "1.0.0" {
		t.Errorf("expected Version unchanged '1.0.0', got %q", opts.Version)
	}
}

func TestWithHostname_SetsHostname(t *testing.T) {
	opts := DefaultOptions()
	WithHostname("build-server")(&opts)

	if opts.Hostname != "build-server" {
		t.Errorf("expected Hostname 'build-server', got %q", opts.Hostname)
	}
}

func TestWithHostname_EmptyIgnored(t *testing.T) {
	opts := DefaultOptions()
	opts.Hostname = "existing-host"
	WithHostname("")(&opts)

	if opts.Hostname != "existing-host" {
		t.Errorf("expected Hostname unchanged, got %q", opts.Hostname)
	}
}

func TestWithInputFile_SetsPath(t *testing.T) {
	opts := DefaultOptions()
	WithInputFile("/reports/drift.json")(&opts)

	if opts.InputFile != "/reports/drift.json" {
		t.Errorf("expected InputFile '/reports/drift.json', got %q", opts.InputFile)
	}
}

func TestWithInputFile_EmptyAllowed(t *testing.T) {
	opts := DefaultOptions()
	opts.InputFile = "old.json"
	WithInputFile("")(&opts)

	if opts.InputFile != "" {
		t.Errorf("expected InputFile cleared, got %q", opts.InputFile)
	}
}
