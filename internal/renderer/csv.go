package renderer

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/snyk/driftctl-report/internal/parser"
)

// CSVRow represents a single resource row in the CSV export.
type CSVRow struct {
	ResourceType string
	ResourceID   string
	Status       string
}

// buildCSVRows converts a parsed report into a flat slice of CSV rows.
func buildCSVRows(report *parser.Analysis) []CSVRow {
	var rows []CSVRow

	for _, r := range report.Summary.TotalUnmanaged {
		rows = append(rows, CSVRow{
			ResourceType: r.Type,
			ResourceID:   r.ID,
			Status:       "unmanaged",
		})
	}
	for _, r := range report.Summary.TotalDeleted {
		rows = append(rows, CSVRow{
			ResourceType: r.Type,
			ResourceID:   r.ID,
			Status:       "deleted",
		})
	}
	for _, r := range report.Summary.TotalManaged {
		rows = append(rows, CSVRow{
			ResourceType: r.Type,
			ResourceID:   r.ID,
			Status:       "managed",
		})
	}

	return rows
}

// WriteCSV writes resource rows to w using the provided delimiter.
// The first row is always a header.
func WriteCSV(w io.Writer, rows []CSVRow, delimiter rune) error {
	cw := csv.NewWriter(w)
	cw.Comma = delimiter

	if err := cw.Write([]string{"resource_type", "resource_id", "status"}); err != nil {
		return fmt.Errorf("csv: write header: %w", err)
	}

	for _, row := range rows {
		if err := cw.Write([]string{row.ResourceType, row.ResourceID, row.Status}); err != nil {
			return fmt.Errorf("csv: write row: %w", err)
		}
	}

	cw.Flush()
	return cw.Error()
}
