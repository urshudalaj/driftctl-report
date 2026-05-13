package renderer

import "sort"

// TrendDirection indicates whether drift is improving, degrading, or stable.
type TrendDirection string

const (
	TrendImproving  TrendDirection = "improving"
	TrendDegrading  TrendDirection = "degrading"
	TrendStable     TrendDirection = "stable"
	TrendInsufficient TrendDirection = "insufficient"
)

// TrendPoint is a single observation in the trend series.
type TrendPoint struct {
	Label    string
	Managed  int
	Unmanaged int
	Deleted  int
	Total    int
	DriftPct float64
}

// TrendData holds the computed trend series and direction.
type TrendData struct {
	Enabled   bool
	Points    []TrendPoint
	Direction TrendDirection
	Delta     float64 // change in drift % from first to last point
}

// buildTrend computes a trend summary from a slice of historical snapshots.
// Each snapshot is a map with keys: label, managed, unmanaged, deleted.
func buildTrend(history []map[string]interface{}, enabled bool) TrendData {
	if !enabled {
		return TrendData{}
	}

	points := make([]TrendPoint, 0, len(history))
	for _, h := range history {
		p := snapshotToPoint(h)
		points = append(points, p)
	}

	// Preserve insertion order (caller is responsible for ordering).
	sort.SliceStable(points, func(i, j int) bool {
		return points[i].Label < points[j].Label
	})

	direction, delta := computeDirection(points)

	return TrendData{
		Enabled:   true,
		Points:    points,
		Direction: direction,
		Delta:     delta,
	}
}

func snapshotToPoint(h map[string]interface{}) TrendPoint {
	get := func(key string) int {
		if v, ok := h[key]; ok {
			if n, ok := v.(int); ok {
				return n
			}
		}
		return 0
	}
	label, _ := h["label"].(string)
	m := get("managed")
	u := get("unmanaged")
	d := get("deleted")
	total := m + u + d
	var pct float64
	if total > 0 {
		pct = roundTwo(float64(u+d) / float64(total) * 100)
	}
	return TrendPoint{Label: label, Managed: m, Unmanaged: u, Deleted: d, Total: total, DriftPct: pct}
}

func computeDirection(pts []TrendPoint) (TrendDirection, float64) {
	if len(pts) < 2 {
		return TrendInsufficient, 0
	}
	first := pts[0].DriftPct
	last := pts[len(pts)-1].DriftPct
	delta := roundTwo(last - first)
	switch {
	case delta < 0:
		return TrendImproving, delta
	case delta > 0:
		return TrendDegrading, delta
	default:
		return TrendStable, 0
	}
}
