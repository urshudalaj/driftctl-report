package renderer

import "sort"

// BreakdownEntry holds per-type resource counts for a single category.
type BreakdownEntry struct {
	Type      string
	Managed   int
	Unmanaged int
	Deleted   int
	Total     int
}

// BreakdownData is the top-level structure passed to the template.
type BreakdownData struct {
	Enabled bool
	Entries []BreakdownEntry
	TopN    int
}

// buildBreakdown aggregates resource counts by type across all categories.
func buildBreakdown(a Analysis, opts Options) BreakdownData {
	if !opts.ShowBreakdown {
		return BreakdownData{}
	}

	counts := map[string]*BreakdownEntry{}

	for _, r := range a.Managed {
		e := entryFor(counts, r.Type)
		e.Managed++
		e.Total++
	}
	for _, r := range a.Unmanaged {
		e := entryFor(counts, r.Type)
		e.Unmanaged++
		e.Total++
	}
	for _, r := range a.Deleted {
		e := entryFor(counts, r.Type)
		e.Deleted++
		e.Total++
	}

	entries := make([]BreakdownEntry, 0, len(counts))
	for _, e := range counts {
		entries = append(entries, *e)
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Total != entries[j].Total {
			return entries[i].Total > entries[j].Total
		}
		return entries[i].Type < entries[j].Type
	})

	if opts.BreakdownTopN > 0 && len(entries) > opts.BreakdownTopN {
		entries = entries[:opts.BreakdownTopN]
	}

	return BreakdownData{
		Enabled: true,
		Entries: entries,
		TopN:    opts.BreakdownTopN,
	}
}

func entryFor(m map[string]*BreakdownEntry, typ string) *BreakdownEntry {
	if _, ok := m[typ]; !ok {
		m[typ] = &BreakdownEntry{Type: typ}
	}
	return m[typ]
}
