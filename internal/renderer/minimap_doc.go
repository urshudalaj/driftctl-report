// Package renderer — minimap section.
//
// The minimap provides a compact grid overview of infrastructure drift
// organised by resource type. Each cell in the grid represents one resource
// type and is colour-coded according to its drift percentage:
//
//   - minimap-ok     (0 % drifted)   — fully managed, no issues
//   - minimap-low    (< 25 % drifted) — minor drift, low urgency
//   - minimap-medium (< 60 % drifted) — moderate drift, review recommended
//   - minimap-high   (≥ 60 % drifted) — severe drift, immediate attention
//
// Usage:
//
//	renderer.New(report,
//	    renderer.WithMinimap(true),
//	    renderer.WithMinimapTopN(10),
//	)
//
// WithMinimapTopN(0) disables the limit and renders all resource types.
package renderer
