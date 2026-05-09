package renderer

import "sort"

// ComplianceEntry represents the compliance posture for a single resource type.
type ComplianceEntry struct {
	Type       string
	Managed    int
	Unmanaged  int
	Deleted    int
	Total      int
	Compliance float64 // percentage of managed resources
	Grade      string
}

// ComplianceReport holds the overall compliance data passed to the template.
type ComplianceReport struct {
	Enabled  bool
	Overall  float64
	Grade    string
	Entries  []ComplianceEntry
	TopN     int
}

// buildCompliance constructs a ComplianceReport from the analysis.
func buildCompliance(a Analysis, opts Options) ComplianceReport {
	if !opts.Compliance {
		return ComplianceReport{}
	}

	typeMap := map[string]*ComplianceEntry{}

	for _, r := range a.Managed {
		e := entryForType(typeMap, r.Type)
		e.Managed++
		e.Total++
	}
	for _, r := range a.Unmanaged {
		e := entryForType(typeMap, r.Type)
		e.Unmanaged++
		e.Total++
	}
	for _, r := range a.Deleted {
		e := entryForType(typeMap, r.Type)
		e.Deleted++
		e.Total++
	}

	entries := make([]ComplianceEntry, 0, len(typeMap))
	totalManaged, totalAll := 0, 0
	for _, e := range typeMap {
		if e.Total > 0 {
			e.Compliance = float64(e.Managed) / float64(e.Total) * 100
		}
		e.Grade = complianceGrade(e.Compliance)
		totalManaged += e.Managed
		totalAll += e.Total
		entries = append(entries, *e)
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Compliance != entries[j].Compliance {
			return entries[i].Compliance < entries[j].Compliance
		}
		return entries[i].Type < entries[j].Type
	})

	if opts.ComplianceTopN > 0 && len(entries) > opts.ComplianceTopN {
		entries = entries[:opts.ComplianceTopN]
	}

	overall := 0.0
	if totalAll > 0 {
		overall = float64(totalManaged) / float64(totalAll) * 100
	}

	return ComplianceReport{
		Enabled: true,
		Overall: roundTwo(overall),
		Grade:   complianceGrade(overall),
		Entries: entries,
		TopN:    opts.ComplianceTopN,
	}
}

func entryForType(m map[string]*ComplianceEntry, t string) *ComplianceEntry {
	if _, ok := m[t]; !ok {
		m[t] = &ComplianceEntry{Type: t}
	}
	return m[t]
}

func complianceGrade(pct float64) string {
	switch {
	case pct >= 95:
		return "A"
	case pct >= 80:
		return "B"
	case pct >= 60:
		return "C"
	case pct >= 40:
		return "D"
	default:
		return "F"
	}
}
