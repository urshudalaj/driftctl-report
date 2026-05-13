// Package renderer — table module
//
// # Resource Table
//
// The table module renders a paginated, sortable table of all resources found
// in the driftctl analysis. Each row shows the resource ID, type, and drift
// status (managed / unmanaged / deleted) with a Bootstrap contextual class.
//
// # Usage
//
// Enable the table section via the WithTable option:
//
//	 renderer.New(analysis,
//	     renderer.WithTable(true),
//	     renderer.WithTablePageSize(20),
//	     renderer.WithPage(1),
//	     renderer.WithSortField("type"),
//	     renderer.WithSortOrder("asc"),
//	 )
//
// # Columns
//
// The default columns are: type, id, status.  Override with WithTableColumns.
//
// # Pagination
//
// Page size defaults to 25 when not set or set to zero.  Pages beyond the
// total are clamped to the last valid page.
package renderer
