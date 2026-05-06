// Package renderer provides HTML report rendering for driftctl JSON output.
//
// # Search
//
// The search feature allows callers to narrow rendered output to resources
// whose ID or type contains a given substring. Matching is case-insensitive.
//
// Usage:
//
//	renderer, err := renderer.New(
//		report,
//		renderer.WithSearch("aws_s3"),
//	)
//
// The search is applied after type-filtering and before sorting and
// pagination, so the result set reflects only resources that match the
// query string.
//
// An empty query string disables search and returns all resources.
package renderer
