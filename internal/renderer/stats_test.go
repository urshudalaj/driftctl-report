package renderer

import (
	"testing"

	"github.com/snyk/driftctl/pkg/analyser"
)

func makeStatsAnalysis() Analysis {
	return Analysis{
		Managed: []analyser.Resource{
			{ResourceType: "aws_s3_bucket", ResourceID: "bucket-1"},
			{ResourceType: "aws_s3_bucket", ResourceID: "bucket-2"},
			{ResourceType: "aws_iam_role", ResourceID: "role-1"},
		},
		Unmanaged: []analyser.Resource{
			{ResourceType: "aws_s3_bucket", ResourceID: "bucket-3"},
		},
		Missing: []analyser.Resource{
			{ResourceType: "aws_iam_role", ResourceID: "role-2"},
			{ResourceType: "aws_lambda_function", ResourceID: "fn-1"},
		},
	}
}

func TestBuildStats_Empty(t *testing.T) {
	stats := buildStats(Analysis{})
	if len(stats.ByType) != 0 {
		t.Fatalf("expected no types, got %d", len(stats.ByType))
	}
	if stats.TotalResources != 0 {
		t.Errorf("expected TotalResources=0, got %d", stats.TotalResources)
	}
}

func TestBuildStats_Totals(t *testing.T) {
	stats := buildStats(makeStatsAnalysis())
	if stats.TotalManaged != 3 {
		t.Errorf("TotalManaged: want 3, got %d", stats.TotalManaged)
	}
	if stats.TotalUnmanaged != 1 {
		t.Errorf("TotalUnmanaged: want 1, got %d", stats.TotalUnmanaged)
	}
	if stats.TotalMissing != 2 {
		t.Errorf("TotalMissing: want 2, got %d", stats.TotalMissing)
	}
	if stats.TotalResources != 6 {
		t.Errorf("TotalResources: want 6, got %d", stats.TotalResources)
	}
}

func TestBuildStats_ByType_Sorted(t *testing.T) {
	stats := buildStats(makeStatsAnalysis())
	if len(stats.ByType) != 3 {
		t.Fatalf("expected 3 types, got %d", len(stats.ByType))
	}
	if stats.ByType[0].Type != "aws_iam_role" {
		t.Errorf("first type: want aws_iam_role, got %s", stats.ByType[0].Type)
	}
	if stats.ByType[2].Type != "aws_s3_bucket" {
		t.Errorf("last type: want aws_s3_bucket, got %s", stats.ByType[2].Type)
	}
}

func TestBuildStats_PerTypeCounts(t *testing.T) {
	stats := buildStats(makeStatsAnalysis())
	var s3 ResourceTypeStats
	for _, e := range stats.ByType {
		if e.Type == "aws_s3_bucket" {
			s3 = e
		}
	}
	if s3.Managed != 2 {
		t.Errorf("s3 managed: want 2, got %d", s3.Managed)
	}
	if s3.Unmanaged != 1 {
		t.Errorf("s3 unmanaged: want 1, got %d", s3.Unmanaged)
	}
	if s3.Missing != 0 {
		t.Errorf("s3 missing: want 0, got %d", s3.Missing)
	}
}
