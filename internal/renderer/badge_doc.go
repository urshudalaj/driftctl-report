// Package renderer — badge.go
//
// Badge generation for the HTML drift report.
//
// # Overview
//
// A Badge is a small labelled chip rendered in the report header that
// communicates a single drift metric at a glance.  Four badges are always
// produced:
//
//   - Managed   — resources fully under Terraform control (always green)
//   - Unmanaged — resources that exist in the cloud but not in state
//   - Missing   — resources declared in state but absent from the cloud
//   - Drifted   — resources whose configuration has changed since last apply
//
// # Severity Colouring
//
// Counts of zero render as "success" (green).  Counts between 1 and 4
// inclusive render as "warning" (amber).  Counts of 5 or more render as
// "danger" (red).  The managed count is always success regardless of value
// because a high managed count is a positive signal.
package renderer
