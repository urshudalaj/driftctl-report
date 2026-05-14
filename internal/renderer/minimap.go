package renderer

// MinimapEntry represents a single resource type slot in the minimap overview.
type MinimapEntry struct {
	Type      string
	Managed   int
	Unmanaged int
	Deleted   int
	Total     int
	DriftPct  float64
	ColorClass string
}

// MinimapData holds all entries rendered in the minimap section.
type MinimapData struct {
	Enabled bool
	Entries []MinimapEntry
	TopN    int
}

// buildMinimap constructs a compact grid overview of drift per resource type.
func buildMinimap(a analysisInput, opts Options) MinimapData {
	if !opts.Minimap {
		return MinimapData{}
	}

	typeTotals := map[string]*MinimapEntry{}

	for _, r := range a.Managed {
		e := entryForMinimap(typeTotals, r.Type)
		e.Managed++
		e.Total++
	}
	for _, r := range a.Unmanaged {
		e := entryForMinimap(typeTotals, r.Type)
		e.Unmanaged++
		e.Total++
	}
	for _, r := range a.Deleted {
		e := entryForMinimap(typeTotals, r.Type)
		e.Deleted++
		e.Total++
	}

	entries := make([]MinimapEntry, 0, len(typeTotals))
	for _, e := range typeTotals {
		if e.Total > 0 {
			drifted := e.Unmanaged + e.Deleted
			e.DriftPct = float64(drifted) / float64(e.Total) * 100
			e.ColorClass = minimapColorClass(e.DriftPct)
		}
		entries = append(entries, *e)
	}

	sortStrings2(entries)

	topN := opts.MinimapTopN
	if topN > 0 && len(entries) > topN {
		entries = entries[:topN]
	}

	return MinimapData{
		Enabled: true,
		Entries: entries,
		TopN:    topN,
	}
}

func entryForMinimap(m map[string]*MinimapEntry, t string) *MinimapEntry {
	if _, ok := m[t]; !ok {
		m[t] = &MinimapEntry{Type: t}
	}
	return m[t]
}

func minimapColorClass(pct float64) string {
	switch {
	case pct == 0:
		return "minimap-ok"
	case pct < 25:
		return "minimap-low"
	case pct < 60:
		return "minimap-medium"
	default:
		return "minimap-high"
	}
}

func sortStrings2(entries []MinimapEntry) {
	for i := 1; i < len(entries); i++ {
		for j := i; j > 0 && entries[j].Type < entries[j-1].Type; j-- {
			entries[j], entries[j-1] = entries[j-1], entries[j]
		}
	}
}
