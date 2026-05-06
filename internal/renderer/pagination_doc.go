// Package renderer — pagination support.
//
// # Pagination
//
// Large driftctl reports can contain hundreds of resources. The pagination
// sub-system breaks resource lists into fixed-size pages so that the HTML
// report remains fast to load and easy to navigate.
//
// ## Types
//
//   - [Paginator] — stateless helper that computes [Page] descriptors.
//   - [Page]      — value type carrying page number, total pages, offset and
//     navigation flags (HasNext / HasPrev).
//
// ## Options
//
// Two [Option] functions control pagination at render time:
//
//   - [WithPageSize] — number of items per page (default 25).
//   - [WithPage]     — 1-based page to render (default 1).
//
// ## Helper
//
// [SliceStrings] applies a [Page] descriptor to a string slice, returning
// only the items that belong to that page. It is used internally by the
// renderer but is exported for use in tests and custom templates.
package renderer
