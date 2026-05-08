// Package renderer provides HTML rendering capabilities for driftctl reports.
//
// # Highlight
//
// The highlight module wraps occurrences of a search query inside resource IDs
// and type strings with an HTML <mark> element so that matching text is visually
// emphasised in the rendered report.
//
// highlightQuery is the primary entry-point: given a raw string and a query it
// returns the string with every case-insensitive occurrence of the query wrapped
// in <mark>…</mark> tags.  If the query is empty or the string does not contain
// the query the original string is returned unchanged.
//
// highlightHTML handles the low-level wrapping and is responsible for preserving
// the original casing of the matched text while surrounding it with the tag.
//
// HighlightMaxLength controls the maximum length of a string that will be
// processed; strings longer than this threshold are returned as-is to avoid
// performance issues with very large resource IDs or diff values.
//
// matchesSearch (re-exported from search_doc.go) is used to determine whether
// a resource should be included in the filtered view before highlighting is
// applied.
package renderer
