package renderer

import "sort"

// HeatmapCell represents a single cell in the resource type heatmap.
type HeatmapCell struct {
	Type       string
	Managed    int
	Unmanaged  int
	Deleted    int
	Total      int
	DriftRatio float64 // 0.0–1.0
	HeatLevel  string  // "none", "low", "medium", "high"
}

// HeatmapData holds the full heatmap output.
type HeatmapData struct {
	Enabled bool
	Cells   []HeatmapCell
}

// buildHeatmap constructs a HeatmapData from the analysis.
func buildHeatmap(a Analysis, opts Options) HeatmapData {
	if !opts.Heatmap {
		return HeatmapData{}
	}

	typeMap := make(map[string]*HeatmapCell)

	for _, r := range a.Managed {
		c := cellFor(typeMap, r.Type)
		c.Managed++
		c.Total++
	}
	for _, r := range a.Unmanaged {
		c := cellFor(typeMap, r.Type)
		c.Unmanaged++
		c.Total++
	}
	for _, r := range a.Deleted {
		c := cellFor(typeMap, r.Type)
		c.Deleted++
		c.Total++
	}

	cells := make([]HeatmapCell, 0, len(typeMap))
	for _, c := range typeMap {
		if c.Total > 0 {
			drifted := c.Unmanaged + c.Deleted
			c.DriftRatio = float64(drifted) / float64(c.Total)
			c.HeatLevel = heatLevel(c.DriftRatio)
		}
		cells = append(cells, *c)
	}

	sort.Slice(cells, func(i, j int) bool {
		if cells[i].DriftRatio != cells[j].DriftRatio {
			return cells[i].DriftRatio > cells[j].DriftRatio
		}
		return cells[i].Type < cells[j].Type
	})

	return HeatmapData{Enabled: true, Cells: cells}
}

func cellFor(m map[string]*HeatmapCell, typ string) *HeatmapCell {
	if _, ok := m[typ]; !ok {
		m[typ] = &HeatmapCell{Type: typ}
	}
	return m[typ]
}

func heatLevel(ratio float64) string {
	switch {
	case ratio == 0:
		return "none"
	case ratio < 0.25:
		return "low"
	case ratio < 0.60:
		return "medium"
	default:
		return "high"
	}
}
