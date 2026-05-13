package renderer

import (
	"sort"

	"github.com/snyk/driftctl/pkg/analyser"
)

// TableRow represents a single row in the resource table view.
type TableRow struct {
	ID     string
	Type   string
	Status string
	Class  string
}

// TableData holds all rows and pagination metadata for the table view.
type TableData struct {
	Enabled     bool
	Rows        []TableRow
	TotalRows   int
	Page        int
	PageSize    int
	TotalPages  int
	SortField   string
	SortOrder   string
	SearchQuery string
}

// buildTable assembles TableData from the analysis using current options.
func buildTable(a analyser.Analysis, opts Options) TableData {
	if !opts.TableEnabled {
		return TableData{}
	}

	rows := collectRows(a)

	if opts.SortField == "id" {
		sort.SliceStable(rows, func(i, j int) bool {
			return resourceLess(rows[i].ID, rows[j].ID, rows[i].Type, rows[j].Type, opts.SortOrder == "asc")
		})
	} else {
		sort.SliceStable(rows, func(i, j int) bool {
			return resourceLess(rows[i].Type, rows[j].Type, rows[i].ID, rows[j].ID, opts.SortOrder == "asc")
		})
	}

	total := len(rows)
	pageSize := opts.PageSize
	if pageSize <= 0 {
		pageSize = 25
	}
	totalPages := (total + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}
	page := opts.Page
	if page < 1 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > total {
		end = total
	}
	paged := rows[start:end]

	return TableData{
		Enabled:     true,
		Rows:        paged,
		TotalRows:   total,
		Page:        page,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		SortField:   opts.SortField,
		SortOrder:   opts.SortOrder,
		SearchQuery: opts.SearchQuery,
	}
}

func collectRows(a analyser.Analysis) []TableRow {
	var rows []TableRow
	for _, r := range a.Managed() {
		rows = append(rows, TableRow{ID: r.ResourceId(), Type: r.ResourceType(), Status: "managed", Class: "success"})
	}
	for _, r := range a.Unmanaged() {
		rows = append(rows, TableRow{ID: r.ResourceId(), Type: r.ResourceType(), Status: "unmanaged", Class: "warning"})
	}
	for _, r := range a.Deleted() {
		rows = append(rows, TableRow{ID: r.ResourceId(), Type: r.ResourceType(), Status: "deleted", Class: "danger"})
	}
	return rows
}
