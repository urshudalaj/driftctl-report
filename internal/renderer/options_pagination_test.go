package renderer

import "testing"

func TestWithPageSize_Valid(t *testing.T) {
	opts := DefaultOptions()
	if err := WithPageSize(50)(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.PageSize != 50 {
		t.Fatalf("expected PageSize 50, got %d", opts.PageSize)
	}
}

func TestWithPageSize_Zero(t *testing.T) {
	opts := DefaultOptions()
	original := opts.PageSize
	if err := WithPageSize(0)(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.PageSize != original {
		t.Fatalf("expected PageSize unchanged (%d), got %d", original, opts.PageSize)
	}
}

func TestWithPageSize_Negative(t *testing.T) {
	opts := DefaultOptions()
	original := opts.PageSize
	if err := WithPageSize(-5)(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.PageSize != original {
		t.Fatalf("expected PageSize unchanged (%d), got %d", original, opts.PageSize)
	}
}

func TestWithPage_Valid(t *testing.T) {
	opts := DefaultOptions()
	if err := WithPage(3)(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.Page != 3 {
		t.Fatalf("expected Page 3, got %d", opts.Page)
	}
}

func TestWithPage_Zero(t *testing.T) {
	opts := DefaultOptions()
	if err := WithPage(0)(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.Page != 1 {
		t.Fatalf("expected Page clamped to 1, got %d", opts.Page)
	}
}

func TestWithPage_Negative(t *testing.T) {
	opts := DefaultOptions()
	if err := WithPage(-10)(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.Page != 1 {
		t.Fatalf("expected Page clamped to 1, got %d", opts.Page)
	}
}
