// Package renderer — badge module
//
// # Badge
//
// The badge module renders small pill-shaped indicators in the report header
// that summarise the overall health of the scanned infrastructure at a glance.
//
// Each badge corresponds to one resource category:
//
//   - managed   — resources tracked and in sync with IaC
//   - unmanaged — resources that exist in the cloud but are not in state
//   - deleted   — resources present in state but missing from the cloud
//
// # Severity levels
//
// Badge colour is determined by [badgeLevelForCount]:
//
//	0          → success  (green)
//	1–9        → warning  (amber)
//	10+        → danger   (red)
//
// Managed resources always receive the success level regardless of count.
//
// # Options
//
// Use [WithBadges] to enable or disable the section, [WithBadgeLabels] to
// override default category labels, and [WithBadgeShowZero] to control
// whether zero-count badges are rendered.
package renderer
