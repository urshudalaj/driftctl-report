package renderer

import "github.com/snyk/driftctl/pkg/resource"

// SeverityLevel represents the drift severity classification.
type SeverityLevel string

const (
	SeverityNone     SeverityLevel = "none"
	SeverityLow      SeverityLevel = "low"
	SeverityMedium   SeverityLevel = "medium"
	SeverityHigh     SeverityLevel = "high"
	SeverityCritical SeverityLevel = "critical"
)

// SeverityResult holds the computed severity for a report.
type SeverityResult struct {
	Level       SeverityLevel
	Label       string
	BadgeClass  string
	Description string
	Score       int
}

// computeSeverity derives a SeverityResult from drift counts.
// Score is calculated as: missing*3 + unmanaged*2 + different*1.
func computeSeverity(missing, unmanaged, different int) SeverityResult {
	score := missing*3 + unmanaged*2 + different*1

	var level SeverityLevel
	switch {
	case score == 0:
		level = SeverityNone
	case score <= 5:
		level = SeverityLow
	case score <= 15:
		level = SeverityMedium
	case score <= 30:
		level = SeverityHigh
	default:
		level = SeverityCritical
	}

	return SeverityResult{
		Level:       level,
		Label:       severityLabel(level),
		BadgeClass:  severityBadgeClass(level),
		Description: severityDescription(level),
		Score:       score,
	}
}

func severityLabel(l SeverityLevel) string {
	switch l {
	case SeverityNone:
		return "No Drift"
	case SeverityLow:
		return "Low"
	case SeverityMedium:
		return "Medium"
	case SeverityHigh:
		return "High"
	default:
		return "Critical"
	}
}

func severityBadgeClass(l SeverityLevel) string {
	switch l {
	case SeverityNone:
		return "badge-success"
	case SeverityLow:
		return "badge-info"
	case SeverityMedium:
		return "badge-warning"
	case SeverityHigh, SeverityCritical:
		return "badge-danger"
	default:
		return "badge-secondary"
	}
}

func severityDescription(l SeverityLevel) string {
	switch l {
	case SeverityNone:
		return "Infrastructure is fully in sync."
	case SeverityLow:
		return "Minor drift detected; review when convenient."
	case SeverityMedium:
		return "Moderate drift detected; schedule remediation."
	case SeverityHigh:
		return "Significant drift detected; remediate soon."
	default:
		return "Critical drift detected; immediate action required."
	}
}

// Ensure the resource import is used only when needed by other files.
var _ = resource.Resource{}
