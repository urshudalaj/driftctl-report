package renderer

import (
	"testing"
)

func TestWithSearch_SetsQuery(t *testing.T) {
	opts := DefaultOptions()
	if err := WithSearch("bucket")(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.SearchQuery != "bucket" {
		t.Errorf("expected SearchQuery 'bucket', got %q", opts.SearchQuery)
	}
}

func TestWithSearch_EmptyQuery(t *testing.T) {
	opts := DefaultOptions()
	if err := WithSearch("")(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.SearchQuery != "" {
		t.Errorf("expected empty SearchQuery, got %q", opts.SearchQuery)
	}
}

func TestWithSearch_OverwritesPreviousQuery(t *testing.T) {
	opts := DefaultOptions()
	_ = WithSearch("first")(opts)
	_ = WithSearch("second")(opts)
	if opts.SearchQuery != "second" {
		t.Errorf("expected 'second', got %q", opts.SearchQuery)
	}
}

func TestWithSearch_DefaultIsEmpty(t *testing.T) {
	opts := DefaultOptions()
	if opts.SearchQuery != "" {
		t.Errorf("expected empty default SearchQuery, got %q", opts.SearchQuery)
	}
}
