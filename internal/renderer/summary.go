package renderer

import "github.com/snyk/driftctl/pkg/analyser"

// SummarySection holds aggregated counts and coverage percentage
// derived from a driftctl analysis result, used to populate the
// top-level summary cards in the HTML report.
type SummarySection struct {
	Managed    int
	Unmanaged  int
	Missing    int
	Drifted    int
	Total      int
	Coverage   float64
}

// buildSummary computes a SummarySection from the given analysis.
// Coverage is expressed as a percentage (0–100) rounded to two
// decimal places. When Total is zero, Coverage is reported as 0.
func buildSummary(a analyser.Analysis) SummarySection {
	managed := len(a.Managed())
	unmanaged := len(a.Unmanaged())
	missing := len(a.Missing())
	drifted := len(a.Differences())
	total := managed + unmanaged + missing

	var coverage float64
	if total > 0 {
		coverage = roundTwo(float64(managed) / float64(total) * 100)
	}

	return SummarySection{
		Managed:   managed,
		Unmanaged: unmanaged,
		Missing:   missing,
		Drifted:   drifted,
		Total:     total,
		Coverage:  coverage,
	}
}

// roundTwo rounds f to two decimal places.
func roundTwo(f float64) float64 {
	const factor = 100.0
	return float64(int(f*factor+0.5)) / factor
}
