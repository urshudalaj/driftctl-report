package renderer

import (
	"testing"
)

func TestWithColumns_KnownColumnsAccepted(t *testing.T) {
	o := DefaultOptions()
	WithColumns("id", "type", "status")(&o)
	if len(o.Columns) != 3 {
		t.Fatalf("expected 3 columns, got %d", len(o.Columns))
	}
}

func TestWithColumns_UnknownColumnsIgnored(t *testing.T) {
	o := DefaultOptions()
	WithColumns("id", "unknown", "bogus")(&o)
	if len(o.Columns) != 1 {
		t.Fatalf("expected 1 column, got %d", len(o.Columns))
	}
	if o.Columns[0] != "id" {
		t.Errorf("expected \"id\", got %q", o.Columns[0])
	}
}

func TestWithColumns_EmptyCallAddsNothing(t *testing.T) {
	o := DefaultOptions()
	WithColumns()(&o)
	if len(o.Columns) != 0 {
		t.Fatalf("expected 0 columns, got %d", len(o.Columns))
	}
}

func TestWithColumnsAll_SetsAllFive(t *testing.T) {
	o := DefaultOptions()
	WithColumnsAll()(&o)
	if len(o.Columns) != 5 {
		t.Fatalf("expected 5 columns, got %d", len(o.Columns))
	}
}

func TestWithColumnsReset_ClearsColumns(t *testing.T) {
	o := DefaultOptions()
	WithColumnsAll()(&o)
	WithColumnsReset()(&o)
	if len(o.Columns) != 0 {
		t.Fatalf("expected 0 columns after reset, got %d", len(o.Columns))
	}
}

func TestWithColumns_AccumulatesAcrossCalls(t *testing.T) {
	o := DefaultOptions()
	WithColumns("id")(&o)
	WithColumns("type")(&o)
	if len(o.Columns) != 2 {
		t.Fatalf("expected 2 columns, got %d", len(o.Columns))
	}
}

func TestWithColumns_DefaultIsEmpty(t *testing.T) {
	o := DefaultOptions()
	if len(o.Columns) != 0 {
		t.Fatalf("expected default columns to be empty, got %d", len(o.Columns))
	}
}
