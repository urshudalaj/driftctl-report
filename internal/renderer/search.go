package renderer

import (
	"strings"

	"github.com/snyk/driftctl-report/internal/parser"
)

// SearchOptions configures how resource search is performed.
type SearchOptions struct {
	// Query is the case-insensitive substring to match against resource IDs and types.
	Query string
}

// searchResources filters managed and unmanaged resources whose ID or type
// contains the query string (case-insensitive). Diff resources are matched on
// their resource ID.
func searchResources(report parser.DriftctlReport, opts SearchOptions) parser.DriftctlReport {
	if opts.Query == "" {
		return report
	}
	q := strings.ToLower(opts.Query)

	report.Summary.Managed = filterByQuery(report.Summary.Managed, q)
	report.Summary.Unmanaged = filterByQuery(report.Summary.Unmanaged, q)
	report.Summary.Missing = filterByQuery(report.Summary.Missing, q)
	report.Differences = filterDiffByQuery(report.Differences, q)

	return report
}

func filterByQuery(resources []parser.Resource, q string) []parser.Resource {
	var out []parser.Resource
	for _, r := range resources {
		if strings.Contains(strings.ToLower(r.ID), q) ||
			strings.Contains(strings.ToLower(r.Type), q) {
			out = append(out, r)
		}
	}
	return out
}

func filterDiffByQuery(diffs []parser.DiffResource, q string) []parser.DiffResource {
	var out []parser.DiffResource
	for _, d := range diffs {
		if strings.Contains(strings.ToLower(d.Resource.ID), q) ||
			strings.Contains(strings.ToLower(d.Resource.Type), q) {
			out = append(out, d)
		}
	}
	return out
}
