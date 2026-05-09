// Package renderer provides HTML rendering for driftctl JSON reports.
//
// # Breakdown
//
// The breakdown feature produces a per-resource-type table that shows how many
// resources of each type are managed, unmanaged, or deleted. Rows are sorted
// by total resource count descending so the most impactful types appear first.
//
// # Usage
//
//	r := renderer.New(analysis,
//	    renderer.WithBreakdown(true),
//	    renderer.WithBreakdownTopN(10),
//	)
//
// Setting TopN to 0 (the default) disables truncation and shows all types.
// Negative values are silently clamped to 0.
package renderer
