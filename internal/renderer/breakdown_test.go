package renderer

import (
	"testing"

	"github.com/snyk/driftctl/pkg/resource"
)

func makeBreakdownAnalysis() Analysis {
	return Analysis{
		Managed: []resource.Resource{
			{ResourceType: "aws_s3_bucket", ResourceID: "bucket-1"},
			{ResourceType: "aws_s3_bucket", ResourceID: "bucket-2"},
			{ResourceType: "aws_iam_role", ResourceID: "role-1"},
		},
		Unmanaged: []resource.Resource{
			{ResourceType: "aws_s3_bucket", ResourceID: "bucket-3"},
			{ResourceType: "aws_lambda_function", ResourceID: "fn-1"},
		},
		Deleted: []resource.Resource{
			{ResourceType: "aws_iam_role", ResourceID: "role-2"},
		},
	}
}

func TestBuildBreakdown_Disabled(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowBreakdown = false
	result := buildBreakdown(makeBreakdownAnalysis(), opts)
	if result.Enabled {
		t.Fatal("expected breakdown to be disabled")
	}
	if len(result.Entries) != 0 {
		t.Fatalf("expected no entries, got %d", len(result.Entries))
	}
}

func TestBuildBreakdown_AllTypes(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowBreakdown = true
	result := buildBreakdown(makeBreakdownAnalysis(), opts)
	if !result.Enabled {
		t.Fatal("expected breakdown to be enabled")
	}
	if len(result.Entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(result.Entries))
	}
}

func TestBuildBreakdown_SortedByTotalDesc(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowBreakdown = true
	result := buildBreakdown(makeBreakdownAnalysis(), opts)
	if result.Entries[0].Type != "aws_s3_bucket" {
		t.Fatalf("expected aws_s3_bucket first, got %s", result.Entries[0].Type)
	}
}

func TestBuildBreakdown_TopN(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowBreakdown = true
	opts.BreakdownTopN = 2
	result := buildBreakdown(makeBreakdownAnalysis(), opts)
	if len(result.Entries) != 2 {
		t.Fatalf("expected 2 entries with TopN=2, got %d", len(result.Entries))
	}
}

func TestBuildBreakdown_CountsAreCorrect(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowBreakdown = true
	result := buildBreakdown(makeBreakdownAnalysis(), opts)
	var s3 BreakdownEntry
	for _, e := range result.Entries {
		if e.Type == "aws_s3_bucket" {
			s3 = e
		}
	}
	if s3.Managed != 2 || s3.Unmanaged != 1 || s3.Deleted != 0 || s3.Total != 3 {
		t.Fatalf("unexpected s3 counts: %+v", s3)
	}
}

func TestBuildBreakdown_Empty(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowBreakdown = true
	result := buildBreakdown(Analysis{}, opts)
	if !result.Enabled {
		t.Fatal("expected enabled even with empty analysis")
	}
	if len(result.Entries) != 0 {
		t.Fatalf("expected 0 entries, got %d", len(result.Entries))
	}
}
