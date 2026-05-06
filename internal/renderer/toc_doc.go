// Package renderer — toc.go
//
// The table-of-contents (TOC) feature generates a navigable index of resource
// type sections that appear in the rendered HTML report.
//
// # Overview
//
// buildTOC accepts the ordered list of resource types present in the filtered
// report together with per-type resource counts and a set of types that contain
// at least one drifted resource. It returns a [TableOfContents] value whose
// [TOCEntry] slice mirrors the input order so callers control sorting.
//
// # Anchor generation
//
// tocAnchor converts an arbitrary resource-type string (e.g. "aws_s3_bucket")
// into a safe HTML id attribute value by replacing every character that is not
// ASCII alphanumeric with a hyphen. This keeps anchors valid across all
// browsers without requiring an external dependency.
//
// # Template integration
//
// The [TableOfContents] value is injected into the template data map under the
// key "TOC" so that the HTML template can render a sticky sidebar or in-page
// navigation list with jump links to each resource-type section.
package renderer
