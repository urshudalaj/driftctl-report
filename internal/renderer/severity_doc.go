// Package renderer provides HTML rendering capabilities for driftctl JSON reports.
//
// # Severity
//
// The severity module classifies infrastructure drift into five levels based on
// a weighted score derived from the number of missing, unmanaged, and differing
// resources:
//
//	- none     – score == 0 (infrastructure fully in sync)
//	- low      – score 1–5
//	- medium   – score 6–15
//	- high     – score 16–30
//	- critical – score > 30
//
// Weights applied per resource count:
//
//	- missing   × 3  (highest impact: resource gone from state)
//	- unmanaged × 2  (medium impact: resource not tracked)
//	- different × 1  (low impact: attribute drift)
//
// Each SeverityResult includes a human-readable label, a Bootstrap badge CSS
// class suitable for embedding in HTML templates, and a short description
// intended for display in the report summary section.
package renderer
