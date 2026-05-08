package renderer

import (
	"sort"
	"time"
)

// SnapshotEntry represents a point-in-time summary of drift state.
type SnapshotEntry struct {
	Timestamp   time.Time
	Managed     int
	Unmanaged   int
	Missing     int
	Drifted     int
	CoveragePC  float64
}

// SnapshotData holds the rendered snapshot section.
type SnapshotData struct {
	Enabled  bool
	Entries  []SnapshotEntry
	Latest   SnapshotEntry
	HasTrend bool
	Trend    string // "improving", "degrading", "stable"
}

// buildSnapshot constructs a SnapshotData from a slice of historical entries
// plus the current analysis. Entries are sorted oldest-first.
func buildSnapshot(current analysisInput, history []SnapshotEntry, enabled bool) SnapshotData {
	if !enabled {
		return SnapshotData{}
	}

	latest := SnapshotEntry{
		Timestamp:  time.Now().UTC(),
		Managed:    current.Managed,
		Unmanaged:  current.Unmanaged,
		Missing:    current.Missing,
		Drifted:    current.Drifted,
		CoveragePC: current.CoveragePC,
	}

	all := make([]SnapshotEntry, len(history))
	copy(all, history)
	all = append(all, latest)

	sort.Slice(all, func(i, j int) bool {
		return all[i].Timestamp.Before(all[j].Timestamp)
	})

	trend := "stable"
	hasTrend := false
	if len(all) >= 2 {
		hasTrend = true
		prev := all[len(all)-2]
		switch {
		case latest.CoveragePC > prev.CoveragePC:
			trend = "improving"
		case latest.CoveragePC < prev.CoveragePC:
			trend = "degrading"
		}
	}

	return SnapshotData{
		Enabled:  true,
		Entries:  all,
		Latest:   latest,
		HasTrend: hasTrend,
		Trend:    trend,
	}
}

// analysisInput is a small value type used by buildSnapshot and other builders.
type analysisInput struct {
	Managed    int
	Unmanaged  int
	Missing    int
	Drifted    int
	CoveragePC float64
}
