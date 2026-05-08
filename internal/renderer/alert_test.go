package renderer

import (
	"testing"

	"github.com/snyk/driftctl-report/internal/parser"
)

func makeAlertAnalysis(unmanaged, deleted int) parser.Analysis {
	a := parser.Analysis{}
	for i := 0; i < unmanaged; i++ {
		a.Summary.TotalUnmanaged = append(a.Summary.TotalUnmanaged, parser.Resource{
			ResourceID:   "u-" + itoa(i),
			ResourceType: "aws_instance",
		})
	}
	for i := 0; i < deleted; i++ {
		a.Summary.TotalDeleted = append(a.Summary.TotalDeleted, parser.Resource{
			ResourceID:   "d-" + itoa(i),
			ResourceType: "aws_s3_bucket",
		})
	}
	return a
}

func TestBuildAlerts_Disabled(t *testing.T) {
	opts := DefaultOptions()
	// AlertsEnabled defaults to false
	result := buildAlerts(makeAlertAnalysis(5, 5), opts)
	if result.Enabled {
		t.Fatal("expected alerts to be disabled")
	}
	if len(result.Alerts) != 0 {
		t.Fatalf("expected 0 alerts, got %d", len(result.Alerts))
	}
}

func TestBuildAlerts_NoIssues(t *testing.T) {
	opts := DefaultOptions()
	WithAlerts(true)(&opts)
	result := buildAlerts(makeAlertAnalysis(0, 0), opts)
	if !result.Enabled {
		t.Fatal("expected alerts to be enabled")
	}
	if len(result.Alerts) != 1 {
		t.Fatalf("expected 1 info alert, got %d", len(result.Alerts))
	}
	if result.Alerts[0].Level != AlertLevelInfo {
		t.Errorf("expected info level, got %s", result.Alerts[0].Level)
	}
}

func TestBuildAlerts_UnmanagedWarning(t *testing.T) {
	opts := DefaultOptions()
	WithAlerts(true)(&opts)
	WithAlertThresholdUnmanaged(5)(&opts)
	result := buildAlerts(makeAlertAnalysis(5, 0), opts)
	if len(result.Alerts) != 1 {
		t.Fatalf("expected 1 alert, got %d", len(result.Alerts))
	}
	if result.Alerts[0].Level != AlertLevelWarning {
		t.Errorf("expected warning, got %s", result.Alerts[0].Level)
	}
}

func TestBuildAlerts_UnmanagedDanger(t *testing.T) {
	opts := DefaultOptions()
	WithAlerts(true)(&opts)
	WithAlertThresholdUnmanaged(5)(&opts)
	result := buildAlerts(makeAlertAnalysis(10, 0), opts)
	if result.Alerts[0].Level != AlertLevelDanger {
		t.Errorf("expected danger, got %s", result.Alerts[0].Level)
	}
}

func TestBuildAlerts_DeletedAndUnmanaged(t *testing.T) {
	opts := DefaultOptions()
	WithAlerts(true)(&opts)
	result := buildAlerts(makeAlertAnalysis(3, 2), opts)
	if len(result.Alerts) != 2 {
		t.Fatalf("expected 2 alerts, got %d", len(result.Alerts))
	}
}

func TestWithAlertThresholdUnmanaged_ZeroClampedToOne(t *testing.T) {
	opts := DefaultOptions()
	WithAlertThresholdUnmanaged(0)(&opts)
	if opts.AlertThresholdUnmanaged != 1 {
		t.Errorf("expected 1, got %d", opts.AlertThresholdUnmanaged)
	}
}
