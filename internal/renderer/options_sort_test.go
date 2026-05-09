package renderer

import "testing"

func TestWithSortField_Type(t *testing.T) {
	o := DefaultOptions()
	WithSortField(SortByType)(o)
	if o.SortField != SortByType {
		t.Fatalf("expected SortByType, got %q", o.SortField)
	}
}

func TestWithSortField_ID(t *testing.T) {
	o := DefaultOptions()
	WithSortField(SortByID)(o)
	if o.SortField != SortByID {
		t.Fatalf("expected SortByID, got %q", o.SortField)
	}
}

func TestWithSortField_Unknown_Ignored(t *testing.T) {
	o := DefaultOptions()
	original := o.SortField
	WithSortField(SortField("unknown"))(o)
	if o.SortField != original {
		t.Fatalf("expected field to remain %q, got %q", original, o.SortField)
	}
}

func TestWithSortOrder_Asc(t *testing.T) {
	o := DefaultOptions()
	WithSortOrder(SortAsc)(o)
	if o.SortOrder != SortAsc {
		t.Fatalf("expected SortAsc, got %q", o.SortOrder)
	}
}

func TestWithSortOrder_Desc(t *testing.T) {
	o := DefaultOptions()
	WithSortOrder(SortDesc)(o)
	if o.SortOrder != SortDesc {
		t.Fatalf("expected SortDesc, got %q", o.SortOrder)
	}
}

func TestWithSortOrder_Unknown_Ignored(t *testing.T) {
	o := DefaultOptions()
	original := o.SortOrder
	WithSortOrder(SortOrder("random"))(o)
	if o.SortOrder != original {
		t.Fatalf("expected order to remain %q, got %q", original, o.SortOrder)
	}
}

func TestDefaultOptions_SortDefaults(t *testing.T) {
	o := DefaultOptions()
	if o.SortField != SortByType {
		t.Fatalf("expected default SortField to be SortByType, got %q", o.SortField)
	}
	if o.SortOrder != SortAsc {
		t.Fatalf("expected default SortOrder to be SortAsc, got %q", o.SortOrder)
	}
}
