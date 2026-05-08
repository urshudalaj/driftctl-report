package renderer

import (
	"fmt"
	"sort"
)

// Recommendation holds a single actionable suggestion derived from the drift analysis.
type Recommendation struct {
	Severity    string
	BadgeClass  string
	Title       string
	Description string
	Count       int
}

// RecommendationsData is passed to the template for the recommendations section.
type RecommendationsData struct {
	Enabled         bool
	Recommendations []Recommendation
}

// buildRecommendations produces actionable recommendations based on the analysis.
func buildRecommendations(a Analysis, opts Options) RecommendationsData {
	if !opts.Recommendations {
		return RecommendationsData{}
	}

	var recs []Recommendation

	unmanaged := len(a.Summary.TotalUnmanaged)
	if unmanaged > 0 {
		recs = append(recs, Recommendation{
			Severity:    severityLabel(badgeLevelForCount(unmanaged, 5, 20)),
			BadgeClass:  severityBadgeClass(badgeLevelForCount(unmanaged, 5, 20)),
			Title:       "Import unmanaged resources",
			Description: fmt.Sprintf("%d resource(s) exist in your infrastructure but are not tracked by IaC. Run `terraform import` or add them to your configuration.", unmanaged),
			Count:       unmanaged,
		})
	}

	deleted := len(a.Summary.TotalDeleted)
	if deleted > 0 {
		recs = append(recs, Recommendation{
			Severity:    severityLabel(badgeLevelForCount(deleted, 1, 10)),
			BadgeClass:  severityBadgeClass(badgeLevelForCount(deleted, 1, 10)),
			Title:       "Reconcile deleted resources",
			Description: fmt.Sprintf("%d resource(s) are defined in IaC but missing from the live infrastructure. Review and either re-apply or remove the definitions.", deleted),
			Count:       deleted,
		})
	}

	drifted := len(a.Summary.TotalDrifted)
	if drifted > 0 {
		recs = append(recs, Recommendation{
			Severity:    severityLabel(badgeLevelForCount(drifted, 3, 15)),
			BadgeClass:  severityBadgeClass(badgeLevelForCount(drifted, 3, 15)),
			Title:       "Fix configuration drift",
			Description: fmt.Sprintf("%d resource(s) have drifted from their desired state. Run `terraform plan` to review and `terraform apply` to remediate.", drifted),
			Count:       drifted,
		})
	}

	sort.SliceStable(recs, func(i, j int) bool {
		return recs[i].Count > recs[j].Count
	})

	return RecommendationsData{
		Enabled:         true,
		Recommendations: recs,
	}
}
