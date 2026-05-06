// Package renderer provides functionality for rendering driftctl scan results
// into human-readable HTML reports.
//
// # Sorting
//
// The sorter module exposes sortResources and sortDiffResources helpers that
// reorder resource slices in-place before they are handed to the HTML template.
// Four sort orders are available via the SortOrder type:
//
//   - SortByTypeAsc  – alphabetical by resource type, then by ID (default)
//   - SortByTypeDesc – reverse alphabetical by resource type, then by ID
//   - SortByIDAsc    – alphabetical by resource ID
//   - SortByIDDesc   – reverse alphabetical by resource ID
//
// The desired order is selected through the WithSortOrder Option when
// constructing a Renderer:
//
//	r, err := renderer.New(report,
//		renderer.WithSortOrder(renderer.SortByIDAsc),
//	)
package renderer
