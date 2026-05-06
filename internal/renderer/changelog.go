package renderer

import (
	"fmt"
	"sort"

	"github.com/snyk/driftctl/pkg/analyser"
)

// ChangelogEntry represents a single resource change for the changelog section.
type ChangelogEntry struct {
	ResourceID   string
	ResourceType string
	ChangeKind   string // "missing", "unmanaged", "changed"
	DiffCount    int
}

// ChangelogSection groups entries by change kind.
type ChangelogSection struct {
	Kind    string
	Entries []ChangelogEntry
}

// buildChangelog constructs an ordered changelog from the analyser result.
// Entries are grouped by kind and sorted by resource type then ID within each group.
func buildChangelog(result *analyser.Analysis) []ChangelogSection {
	var missing, unmanaged, changed []ChangelogEntry

	for _, r := range result.Missing() {
		missing = append(missing, ChangelogEntry{
			ResourceID:   r.ResourceId(),
			ResourceType: r.ResourceType(),
			ChangeKind:   "missing",
		})
	}

	for _, r := range result.Unmanaged() {
		unmanaged = append(unmanaged, ChangelogEntry{
			ResourceID:   r.ResourceId(),
			ResourceType: r.ResourceType(),
			ChangeKind:   "unmanaged",
		})
	}

	for _, r := range result.Differences() {
		changed = append(changed, ChangelogEntry{
			ResourceID:   r.Res.ResourceId(),
			ResourceType: r.Res.ResourceType(),
			ChangeKind:   "changed",
			DiffCount:    len(r.Changelog),
		})
	}

	sortEntries(missing)
	sortEntries(unmanaged)
	sortEntries(changed)

	var sections []ChangelogSection
	if len(missing) > 0 {
		sections = append(sections, ChangelogSection{Kind: "missing", Entries: missing})
	}
	if len(unmanaged) > 0 {
		sections = append(sections, ChangelogSection{Kind: "unmanaged", Entries: unmanaged})
	}
	if len(changed) > 0 {
		sections = append(sections, ChangelogSection{Kind: "changed", Entries: changed})
	}
	return sections
}

func sortEntries(entries []ChangelogEntry) {
	sort.Slice(entries, func(i, j int) bool {
		ki := fmt.Sprintf("%s/%s", entries[i].ResourceType, entries[i].ResourceID)
		kj := fmt.Sprintf("%s/%s", entries[j].ResourceType, entries[j].ResourceID)
		return ki < kj
	})
}
