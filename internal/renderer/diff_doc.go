// Package renderer — diff module
//
// # Attribute-Level Diff
//
// The diff module extracts per-attribute change details from driftctl's
// analysis result and surfaces them as a structured DiffSection for
// HTML rendering.
//
// # How It Works
//
//  1. buildDiff iterates over all Differences() in the analysis.
//  2. For each drifted resource, it walks the Changelog map and creates
//     a DiffEntry per attribute, capturing the "from" (IaC) and "to"
//     (real-state) values along with the change type.
//  3. Entries are sorted by resource type → resource ID → attribute name
//     for deterministic, human-readable output.
//  4. When DiffLimit > 0, the slice is truncated and the Trunc flag is
//     set so the template can display a "showing N of M" notice.
//
// # Options
//
// Use WithDiff(true) to enable the section (disabled by default).
// Use WithDiffLimit(n) to cap the number of rendered entries.
package renderer
