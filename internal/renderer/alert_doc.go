// Package renderer provides alert generation for drift reports.
//
// # Alerts
//
// The alert subsystem inspects a parsed driftctl analysis and emits
// structured Alert values that the HTML template renders as banners.
//
// Alerts are produced in three levels:
//
//   - info    – no issues found; all resources are in sync.
//   - warning – the number of problematic resources meets the configured
//     threshold (AlertThresholdUnmanaged / AlertThresholdDeleted).
//   - danger  – the count exceeds twice the configured threshold.
//
// # Configuration
//
// Use the functional options to control alert behaviour:
//
//	renderer.New(analysis,
//	    renderer.WithAlerts(true),
//	    renderer.WithAlertThresholdUnmanaged(10),
//	    renderer.WithAlertThresholdDeleted(5),
//	)
//
// Alerts are disabled by default and must be explicitly enabled via
// WithAlerts(true).
package renderer
