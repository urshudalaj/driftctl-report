package renderer

import (
	"testing"

	"github.com/snyk/driftctl/pkg/analyser"
)

func makeTableAnalysis(managed, unmanaged, deleted int) analyser.Analysis {
	return makeBreakdownAnalysis(managed, unmanaged, deleted)
}

func TestBuildTable_Disabled(t *testing.T) {
	a := makeTableAnalysis(2, 1, 1)
	opts := DefaultOptions()
	// TableEnabled defaults to false
	result := buildTable(a, opts)
	if result.Enabled {
		t.Fatal("expected table to be disabled")
	}
	if len(result.Rows) != 0 {
		t.Fatalf("expected no rows, got %d", len(result.Rows))
	}
}

func TestBuildTable_RowCount(t *testing.T) {
	a := makeTableAnalysis(3, 2, 1)
	opts := DefaultOptions()
	opts.TableEnabled = true
	opts.PageSize = 100
	result := buildTable(a, opts)
	if !result.Enabled {
		t.Fatal("expected table to be enabled")
	}
	if result.TotalRows != 6 {
		t.Fatalf("expected 6 total rows, got %d", result.TotalRows)
	}
	if len(result.Rows) != 6 {
		t.Fatalf("expected 6 rows on page, got %d", len(result.Rows))
	}
}

func TestBuildTable_Pagination(t *testing.T) {
	a := makeTableAnalysis(5, 5, 5)
	opts := DefaultOptions()
	opts.TableEnabled = true
	opts.PageSize = 4
	opts.Page = 2
	result := buildTable(a, opts)
	if result.TotalRows != 15 {
		t.Fatalf("expected 15 total rows, got %d", result.TotalRows)
	}
	if result.TotalPages != 4 {
		t.Fatalf("expected 4 total pages, got %d", result.TotalPages)
	}
	if len(result.Rows) != 4 {
		t.Fatalf("expected 4 rows on page 2, got %d", len(result.Rows))
	}
	if result.Page != 2 {
		t.Fatalf("expected page 2, got %d", result.Page)
	}
}

func TestBuildTable_StatusClasses(t *testing.T) {
	a := makeTableAnalysis(1, 1, 1)
	opts := DefaultOptions()
	opts.TableEnabled = true
	opts.PageSize = 100
	result := buildTable(a, opts)
	classes := map[string]bool{}
	for _, row := range result.Rows {
		classes[row.Class] = true
	}
	for _, want := range []string{"success", "warning", "danger"} {
		if !classes[want] {
			t.Errorf("expected class %q in rows", want)
		}
	}
}

func TestBuildTable_DefaultPageSize(t *testing.T) {
	a := makeTableAnalysis(10, 10, 10)
	opts := DefaultOptions()
	opts.TableEnabled = true
	// PageSize 0 should default to 25
	opts.PageSize = 0
	opts.Page = 1
	result := buildTable(a, opts)
	if result.PageSize != 25 {
		t.Fatalf("expected default page size 25, got %d", result.PageSize)
	}
}

func TestBuildTable_PageBeyondTotal_Clamps(t *testing.T) {
	a := makeTableAnalysis(2, 0, 0)
	opts := DefaultOptions()
	opts.TableEnabled = true
	opts.PageSize = 10
	opts.Page = 99
	result := buildTable(a, opts)
	if result.Page != 1 {
		t.Fatalf("expected page clamped to 1, got %d", result.Page)
	}
}
