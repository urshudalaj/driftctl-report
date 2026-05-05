package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// DriftReport represents the top-level driftctl JSON output structure.
type DriftReport struct {
	Summary   Summary    `json:"summary"`
	Managed   []Resource `json:"managed"`
	Unmanaged []Resource `json:"unmanaged"`
	Deleted   []Resource `json:"deleted"`
	Differences []Difference `json:"differences"`
}

// Summary holds aggregate counts from a driftctl scan.
type Summary struct {
	TotalResources  int     `json:"total_resources"`
	TotalDrifted    int     `json:"total_drifted"`
	TotalUnmanaged  int     `json:"total_unmanaged"`
	TotalDeleted    int     `json:"total_deleted"`
	TotalManaged    int     `json:"total_managed"`
	Coverage        float64 `json:"coverage"`
}

// Resource represents a single infrastructure resource.
type Resource struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Difference captures a drift between IaC state and real infrastructure.
type Difference struct {
	Res     Resource      `json:"res"`
	Changes []FieldChange `json:"changes"`
}

// FieldChange describes a single changed field within a resource.
type FieldChange struct {
	ComputedFields []string `json:"computed"`
	Type           string   `json:"type"`
	Path           string   `json:"path"`
	Before         any      `json:"before"`
	After          any      `json:"after"`
}

// ParseFile reads and parses a driftctl JSON report from the given file path.
func ParseFile(path string) (*DriftReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening report file: %w", err)
	}
	defer f.Close()
	return Parse(f)
}

// Parse decodes a driftctl JSON report from an io.Reader.
func Parse(r io.Reader) (*DriftReport, error) {
	var report DriftReport
	dec := json.NewDecoder(r)
	if err := dec.Decode(&report); err != nil {
		return nil, fmt.Errorf("decoding driftctl JSON: %w", err)
	}
	return &report, nil
}
