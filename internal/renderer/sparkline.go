package renderer

import "sort"

// SparklinePoint represents a single data point in a sparkline series.
type SparklinePoint struct {
	Label string
	Value int
}

// SparklineData holds the series data and display metadata for a sparkline.
type SparklineData struct {
	Enabled bool
	Points  []SparklinePoint
	Max     int
	Min     int
	Last    int
	Trend   string // "up", "down", "flat"
}

// buildSparkline constructs a SparklineData from the analysis, bucketing
// unmanaged resource counts by resource type for a quick visual overview.
func buildSparkline(a Analysis, opts Options) SparklineData {
	if !opts.Sparkline {
		return SparklineData{}
	}

	typeCounts := make(map[string]int)
	for _, r := range a.UnmanagedResources {
		typeCounts[r.Type]++
	}
	for _, r := range a.DeletedResources {
		typeCounts[r.Type]++
	}

	types := make([]string, 0, len(typeCounts))
	for t := range typeCounts {
		types = append(types, t)
	}
	sort.Strings(types)

	points := make([]SparklinePoint, 0, len(types))
	for _, t := range types {
		points = append(points, SparklinePoint{Label: t, Value: typeCounts[t]})
	}

	if len(points) == 0 {
		return SparklineData{Enabled: true, Points: points, Trend: "flat"}
	}

	min, max := points[0].Value, points[0].Value
	for _, p := range points[1:] {
		if p.Value < min {
			min = p.Value
		}
		if p.Value > max {
			max = p.Value
		}
	}

	last := points[len(points)-1].Value
	trend := "flat"
	if len(points) >= 2 {
		first := points[0].Value
		if last > first {
			trend = "up"
		} else if last < first {
			trend = "down"
		}
	}

	return SparklineData{
		Enabled: true,
		Points:  points,
		Max:     max,
		Min:     min,
		Last:    last,
		Trend:   trend,
	}
}
