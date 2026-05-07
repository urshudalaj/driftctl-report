// Package renderer provides utilities for rendering driftctl analysis
// results as human-readable HTML reports.
//
// # Grouper
//
// The grouper module organises flat resource slices into ordered buckets
// keyed by resource type. This makes it straightforward for templates to
// render a collapsible section per type rather than a single flat table.
//
// groupByType accepts a slice of managed/unmanaged resources and returns
// a slice of Bucket values whose order is determined by the first
// occurrence of each type in the input slice.
//
// groupDiffByType performs the same operation for drifted (changed)
// resources, which carry additional diff metadata.
//
// Both functions are pure and produce no side-effects; they are safe to
// call concurrently.
package renderer
