package renderer

import (
	"sort"

	"github.com/snyk/driftctl-report/internal/parser"
)

// FilterOptions controls which resources appear in the rendered report.
type FilterOptions struct {
	// OnlyDrifted, when true, excludes resources that are in sync.
	OnlyDrifted bool
	// ResourceTypes limits output to the given resource types.
	// An empty slice means all types are included.
	ResourceTypes []string
}

// applyFilters returns a copy of the report with resources filtered
// according to opts. The original report is not modified.
func applyFilters(report parser.DriftctlReport, opts FilterOptions) parser.DriftctlReport {
	typeSet := buildTypeSet(opts.ResourceTypes)

	filtered := parser.DriftctlReport{
		Summary:       report.Summary,
		ManagedResources:   filterResources(report.ManagedResources, false, typeSet),
		UnmanagedResources: filterResources(report.UnmanagedResources, false, typeSet),
		DeletedResources:   filterResources(report.DeletedResources, false, typeSet),
		DifferentResources: filterDiffResources(report.DifferentResources, typeSet),
	}

	if opts.OnlyDrifted {
		filtered.ManagedResources = nil
	}

	return filtered
}

func buildTypeSet(types []string) map[string]struct{} {
	if len(types) == 0 {
		return nil
	}
	s := make(map[string]struct{}, len(types))
	for _, t := range types {
		s[t] = struct{}{}
	}
	return s
}

func filterResources(resources []parser.Resource, _ bool, typeSet map[string]struct{}) []parser.Resource {
	if typeSet == nil {
		return resources
	}
	var out []parser.Resource
	for _, r := range resources {
		if _, ok := typeSet[r.Type]; ok {
			out = append(out, r)
		}
	}
	return out
}

func filterDiffResources(resources []parser.DifferentResource, typeSet map[string]struct{}) []parser.DifferentResource {
	if typeSet == nil {
		return resources
	}
	var out []parser.DifferentResource
	for _, r := range resources {
		if _, ok := typeSet[r.Type]; ok {
			out = append(out, r)
		}
	}
	return out
}

// uniqueResourceTypes returns a sorted list of all resource types present
// across every category in the report.
func uniqueResourceTypes(report parser.DriftctlReport) []string {
	seen := make(map[string]struct{})
	for _, r := range report.ManagedResources {
		seen[r.Type] = struct{}{}
	}
	for _, r := range report.UnmanagedResources {
		seen[r.Type] = struct{}{}
	}
	for _, r := range report.DeletedResources {
		seen[r.Type] = struct{}{}
	}
	for _, r := range report.DifferentResources {
		seen[r.Type] = struct{}{}
	}
	types := make([]string, 0, len(seen))
	for t := range seen {
		types = append(types, t)
	}
	sort.Strings(types)
	return types
}
