package renderer

import (
	"testing"
)

func TestWithExcludeType_SingleType(t *testing.T) {
	o := DefaultOptions()
	WithExcludeType("aws_s3_bucket")(o)
	if len(o.ExcludeTypes) != 1 || o.ExcludeTypes[0] != "aws_s3_bucket" {
		t.Fatalf("expected [aws_s3_bucket], got %v", o.ExcludeTypes)
	}
}

func TestWithExcludeType_MultipleTypes(t *testing.T) {
	o := DefaultOptions()
	WithExcludeType("aws_s3_bucket")(o)
	WithExcludeType("aws_iam_role")(o)
	if len(o.ExcludeTypes) != 2 {
		t.Fatalf("expected 2 excluded types, got %d", len(o.ExcludeTypes))
	}
}

func TestWithExcludeType_Deduplicates(t *testing.T) {
	o := DefaultOptions()
	WithExcludeType("aws_s3_bucket")(o)
	WithExcludeType("aws_s3_bucket")(o)
	if len(o.ExcludeTypes) != 1 {
		t.Fatalf("expected deduplication, got %d entries", len(o.ExcludeTypes))
	}
}

func TestWithExcludeType_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithExcludeType("")(o)
	if len(o.ExcludeTypes) != 0 {
		t.Fatalf("expected no excluded types for empty string, got %d", len(o.ExcludeTypes))
	}
}

func TestWithIncludeManaged_True(t *testing.T) {
	o := DefaultOptions()
	WithIncludeManaged(true)(o)
	if !o.IncludeManaged {
		t.Fatal("expected IncludeManaged to be true")
	}
}

func TestWithIncludeManaged_False(t *testing.T) {
	o := DefaultOptions()
	WithIncludeManaged(false)(o)
	if o.IncludeManaged {
		t.Fatal("expected IncludeManaged to be false")
	}
}

func TestWithExcludeIDs_AddsIDs(t *testing.T) {
	o := DefaultOptions()
	WithExcludeIDs("id-1", "id-2")(o)
	if len(o.ExcludeIDs) != 2 {
		t.Fatalf("expected 2 excluded IDs, got %d", len(o.ExcludeIDs))
	}
}

func TestWithExcludeIDs_Deduplicates(t *testing.T) {
	o := DefaultOptions()
	WithExcludeIDs("id-1")(o)
	WithExcludeIDs("id-1")(o)
	if len(o.ExcludeIDs) != 1 {
		t.Fatalf("expected deduplication, got %d entries", len(o.ExcludeIDs))
	}
}

func TestWithExcludeIDs_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithExcludeIDs("", "")(o)
	if len(o.ExcludeIDs) != 0 {
		t.Fatalf("expected no excluded IDs for empty strings, got %d", len(o.ExcludeIDs))
	}
}
