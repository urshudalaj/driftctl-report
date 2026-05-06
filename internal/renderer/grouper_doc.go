// Package renderer provides HTML rendering capabilities for driftctl analysis
// results.
//
// # Grouper
//
// The grouper module provides utilities for organising flat resource slices
// into typed groups, making it straightforward for templates and other
// rendering stages to iterate over resources by their resource type.
//
// groupByType and groupDiffByType both preserve the order in which each
// distinct type is first encountered in the input slice, so the output is
// deterministic given a consistently ordered input (e.g. after sorting).
//
// Typical usage inside the renderer pipeline:
//
//	groups := groupByType(filteredResources)
//	for _, g := range groups {
//		// render g.Type heading and g.Resources rows
//	}
package renderer
