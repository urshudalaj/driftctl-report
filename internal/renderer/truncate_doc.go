// Package renderer — truncate module
//
// # ID Truncation
//
// Long resource IDs (e.g. AWS ARNs or GCP resource paths) can make the HTML
// report hard to read. The truncation module shortens IDs that exceed a
// configurable character limit and appends a suffix to indicate the value was
// cut.
//
// # Configuration
//
// Use the following option functions to control truncation:
//
//	WithTruncateIDs(true)            — enable truncation (default: false)
//	WithTruncateIDsLength(40)        — max rune length before cutting (default: 40)
//	WithTruncateIDsSuffix("\u2026")  — suffix appended to cut IDs (default: "…")
//
// # Output
//
// buildTruncated returns a TruncateResult containing a TruncateEntry for every
// unique resource ID across managed, unmanaged, deleted, and drifted buckets.
// Each entry exposes the original ID alongside its display form so templates
// can render tooltips or title attributes with the full value.
package renderer
