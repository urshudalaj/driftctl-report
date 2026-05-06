// Package renderer provides HTML rendering capabilities for driftctl JSON
// reports.
//
// # Highlight
//
// The highlight sub-module provides utilities for visually marking search
// query matches within resource IDs and type strings in the rendered HTML
// output.
//
// highlightQuery splits a text string into a slice of [HighlightResult]
// fragments. Each fragment carries the original text segment and a boolean
// indicating whether it matched the search query. The comparison is always
// case-insensitive.
//
// highlightHTML converts those fragments into an HTML string where matched
// portions are wrapped in <mark> tags, making them visually distinct in the
// browser without requiring any JavaScript.
//
// Both functions treat an empty query as a no-op, returning the original
// string as a single non-highlighted fragment. This keeps template logic
// simple — callers do not need to guard against empty queries.
package renderer
