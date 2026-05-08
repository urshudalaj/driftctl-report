package renderer

import "github.com/snyk/driftctl-report/internal/parser"

// AlertLevel represents the severity of an alert.
type AlertLevel string

const (
	AlertLevelInfo    AlertLevel = "info"
	AlertLevelWarning AlertLevel = "warning"
	AlertLevelDanger  AlertLevel = "danger"
)

// Alert represents a single actionable alert derived from drift analysis.
type Alert struct {
	Level   AlertLevel
	Title   string
	Message string
}

// AlertData holds the full set of alerts for template rendering.
type AlertData struct {
	Enabled bool
	Alerts  []Alert
}

// buildAlerts inspects the analysis and produces alerts based on configured thresholds.
func buildAlerts(analysis parser.Analysis, opts Options) AlertData {
	if !opts.AlertsEnabled {
		return AlertData{Enabled: false}
	}

	var alerts []Alert

	unmanaged := len(analysis.Summary.TotalUnmanaged)
	if unmanaged > 0 {
		level := AlertLevelWarning
		if unmanaged >= opts.AlertThresholdUnmanaged*2 {
			level = AlertLevelDanger
		}
		alerts = append(alerts, Alert{
			Level:   level,
			Title:   "Unmanaged resources detected",
			Message: itoa(unmanaged) + " resource(s) exist in your infrastructure but are not tracked by IaC.",
		})
	}

	deleted := len(analysis.Summary.TotalDeleted)
	if deleted > 0 {
		level := AlertLevelWarning
		if deleted >= opts.AlertThresholdDeleted*2 {
			level = AlertLevelDanger
		}
		alerts = append(alerts, Alert{
			Level:   level,
			Title:   "Deleted resources detected",
			Message: itoa(deleted) + " resource(s) are defined in IaC but missing from your infrastructure.",
		})
	}

	if len(alerts) == 0 {
		alerts = append(alerts, Alert{
			Level:   AlertLevelInfo,
			Title:   "No issues detected",
			Message: "All managed resources are in sync with your infrastructure.",
		})
	}

	return AlertData{Enabled: true, Alerts: alerts}
}
