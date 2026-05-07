package renderer

import (
	"runtime"
	"time"
)

// ReportMetadata holds contextual information embedded in every generated report.
type ReportMetadata struct {
	// GeneratedAt is the UTC timestamp when the report was rendered.
	GeneratedAt string
	// GeneratedBy is the tool name and version string.
	GeneratedBy string
	// GoVersion is the Go runtime version used to build the binary.
	GoVersion string
	// Hostname is the machine name that produced the report, if available.
	Hostname string
	// InputFile is the path to the driftctl JSON file that was parsed.
	InputFile string
}

// buildMetadata constructs a ReportMetadata value from the current runtime
// environment and the supplied options.
func buildMetadata(opts Options) ReportMetadata {
	hostname := opts.Hostname
	if hostname == "" {
		hostname = "unknown"
	}

	version := opts.Version
	if version == "" {
		version = "dev"
	}

	return ReportMetadata{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		GeneratedBy: "driftctl-report " + version,
		GoVersion:   runtime.Version(),
		Hostname:    hostname,
		InputFile:   opts.InputFile,
	}
}
