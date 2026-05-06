package renderer

import "sort"

// ResourceTypeStats holds per-type resource counts for the stats section.
type ResourceTypeStats struct {
	Type     string
	Managed  int
	Unmanaged int
	Missing  int
	Total    int
}

// ReportStats aggregates resource statistics across all types.
type ReportStats struct {
	ByType       []ResourceTypeStats
	TotalManaged int
	TotalUnmanaged int
	TotalMissing int
	TotalResources int
}

// buildStats computes per-type and overall resource statistics from the analysis.
func buildStats(a Analysis) ReportStats {
	counts := map[string]*ResourceTypeStats{}

	for _, r := range a.Managed {
		s := statsEntry(counts, r.Type)
		s.Managed++
		s.Total++
	}
	for _, r := range a.Unmanaged {
		s := statsEntry(counts, r.Type)
		s.Unmanaged++
		s.Total++
	}
	for _, r := range a.Missing {
		s := statsEntry(counts, r.Type)
		s.Missing++
		s.Total++
	}

	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	rs := ReportStats{}
	for _, k := range keys {
		e := counts[k]
		rs.ByType = append(rs.ByType, *e)
		rs.TotalManaged += e.Managed
		rs.TotalUnmanaged += e.Unmanaged
		rs.TotalMissing += e.Missing
		rs.TotalResources += e.Total
	}
	return rs
}

func statsEntry(m map[string]*ResourceTypeStats, typ string) *ResourceTypeStats {
	if _, ok := m[typ]; !ok {
		m[typ] = &ResourceTypeStats{Type: typ}
	}
	return m[typ]
}
