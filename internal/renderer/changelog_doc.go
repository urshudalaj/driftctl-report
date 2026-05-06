// Package renderer provides HTML rendering capabilities for driftctl JSON reports.
//
// # Changelog
//
// The changelog module builds a structured, human-readable summary of all
// resource changes detected by driftctl. It groups changes into three
// categories:
//
//   - missing:   resources declared in IaC but absent in the cloud provider.
//   - unmanaged: resources found in the cloud provider but not tracked by IaC.
//   - changed:   resources that exist in both but whose attributes differ.
//
// Within each group, entries are sorted lexicographically by resource type
// and then by resource ID to ensure deterministic, diff-friendly output.
//
// Usage:
//
//	sections := buildChangelog(analysis)
//	for _, s := range sections {
//		fmt.Println(s.Kind, len(s.Entries))
//	}
package renderer
