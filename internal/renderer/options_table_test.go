package renderer

import "testing"

func TestWithTable_Enables(t *testing.T) {
	opts := DefaultOptions()
	WithTable(true)(&opts)
	if !opts.TableEnabled {
		t.Fatal("expected TableEnabled to be true")
	}
}

func TestWithTable_Disables(t *testing.T) {
	opts := DefaultOptions()
	opts.TableEnabled = true
	WithTable(false)(&opts)
	if opts.TableEnabled {
		t.Fatal("expected TableEnabled to be false")
	}
}

func TestWithTable_DefaultIsFalse(t *testing.T) {
	opts := DefaultOptions()
	if opts.TableEnabled {
		t.Fatal("expected TableEnabled default to be false")
	}
}

func TestWithTablePageSize_PositiveValue(t *testing.T) {
	opts := DefaultOptions()
	WithTablePageSize(50)(&opts)
	if opts.PageSize != 50 {
		t.Fatalf("expected PageSize 50, got %d", opts.PageSize)
	}
}

func TestWithTablePageSize_ZeroIgnored(t *testing.T) {
	opts := DefaultOptions()
	opts.PageSize = 20
	WithTablePageSize(0)(&opts)
	if opts.PageSize != 20 {
		t.Fatalf("expected PageSize unchanged at 20, got %d", opts.PageSize)
	}
}

func TestWithTablePageSize_NegativeIgnored(t *testing.T) {
	opts := DefaultOptions()
	opts.PageSize = 30
	WithTablePageSize(-5)(&opts)
	if opts.PageSize != 30 {
		t.Fatalf("expected PageSize unchanged at 30, got %d", opts.PageSize)
	}
}

func TestWithTableColumns_SetsColumns(t *testing.T) {
	opts := DefaultOptions()
	cols := []string{"type", "id", "status"}
	WithTableColumns(cols)(&opts)
	if len(opts.TableColumns) != 3 {
		t.Fatalf("expected 3 columns, got %d", len(opts.TableColumns))
	}
	for i, c := range cols {
		if opts.TableColumns[i] != c {
			t.Errorf("column[%d]: expected %q, got %q", i, c, opts.TableColumns[i])
		}
	}
}

func TestWithTableColumns_EmptyIgnored(t *testing.T) {
	opts := DefaultOptions()
	opts.TableColumns = []string{"id"}
	WithTableColumns([]string{})(&opts)
	if len(opts.TableColumns) != 1 {
		t.Fatalf("expected columns unchanged, got %d", len(opts.TableColumns))
	}
}

func TestWithTableColumns_CopiesSlice(t *testing.T) {
	opts := DefaultOptions()
	cols := []string{"type", "id"}
	WithTableColumns(cols)(&opts)
	cols[0] = "mutated"
	if opts.TableColumns[0] == "mutated" {
		t.Fatal("expected TableColumns to be an independent copy")
	}
}
