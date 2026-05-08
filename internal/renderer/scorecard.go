package renderer

import "math"

// ScoreCard holds a numeric drift health score and its breakdown.
type ScoreCard struct {
	// Score is a value from 0 (fully drifted) to 100 (fully managed, no drift).
	Score       float64
	Grade       string
	Managed     int
	Unmanaged   int
	Missing     int
	Drifted     int
	Total       int
	Description string
}

// buildScorecard computes a drift health scorecard from the analysis.
// The score penalises unmanaged, missing and drifted resources relative
// to the total number of resources seen.
func buildScorecard(a Analysis) ScoreCard {
	total := a.Summary.TotalResources
	if total == 0 {
		return ScoreCard{
			Score:       100,
			Grade:       "A",
			Description: "No resources found.",
		}
	}

	penalty := float64(a.Summary.TotalUnmanaged+a.Summary.TotalMissing) +
		float64(a.Summary.TotalDrifted)*0.5

	raw := math.Max(0, 100-penalty/float64(total)*100)
	score := math.Round(raw*100) / 100

	return ScoreCard{
		Score:       score,
		Grade:       scorecardGrade(score),
		Managed:     a.Summary.TotalManaged,
		Unmanaged:   a.Summary.TotalUnmanaged,
		Missing:     a.Summary.TotalMissing,
		Drifted:     a.Summary.TotalDrifted,
		Total:       total,
		Description: scorecardDescription(score),
	}
}

func scorecardGrade(score float64) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 75:
		return "B"
	case score >= 60:
		return "C"
	case score >= 40:
		return "D"
	default:
		return "F"
	}
}

func scorecardDescription(score float64) string {
	switch {
	case score >= 90:
		return "Infrastructure is well managed with minimal drift."
	case score >= 75:
		return "Minor drift detected; review unmanaged resources."
	case score >= 60:
		return "Moderate drift; action recommended."
	case score >= 40:
		return "Significant drift detected; immediate review advised."
	default:
		return "Critical drift level; infrastructure is largely unmanaged."
	}
}
