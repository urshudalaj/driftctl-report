package renderer

import (
	"testing"
)

func makeMinimapAnalysis() analysisInput {
	return analysisInput{
		Managed: []resource{
			{Type: "aws_s3_bucket", ID: "bucket-1"},
			{Type: "aws_s3_bucket", ID: "bucket-2"},
			{Type: "aws_iam_role", ID: "role-1"},
		},
		Unmanaged: []resource{
			{Type: "aws_s3_bucket", ID: "bucket-3"},
			{Type: "aws_lambda_function", ID: "fn-1"},
		},
		Deleted: []resource{
			{Type: "aws_iam_role", ID: "role-2"},
		},
	}
}

func TestBuildMinimap_Disabled(t *testing.T) {
	opts := DefaultOptions()
	opts.Minimap = false
	result := buildMinimap(makeMinimapAnalysis(), opts)
	if result.Enabled {
		t.Fatal("expected minimap to be disabled")
	}
	if len(result.Entries) != 0 {
		t.Fatalf("expected no entries, got %d", len(result.Entries))
	}
}

func TestBuildMinimap_EntryCount(t *testing.T) {
	opts := DefaultOptions()
	opts.Minimap = true
	result := buildMinimap(makeMinimapAnalysis(), opts)
	if !result.Enabled {
		t.Fatal("expected minimap to be enabled")
	}
	// 3 distinct types: aws_iam_role, aws_lambda_function, aws_s3_bucket
	if len(result.Entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(result.Entries))
	}
}

func TestBuildMinimap_SortedByType(t *testing.T) {
	opts := DefaultOptions()
	opts.Minimap = true
	result := buildMinimap(makeMinimapAnalysis(), opts)
	for i := 1; i < len(result.Entries); i++ {
		if result.Entries[i].Type < result.Entries[i-1].Type {
			t.Errorf("entries not sorted: %s before %s", result.Entries[i-1].Type, result.Entries[i].Type)
		}
	}
}

func TestBuildMinimap_DriftPctCalculation(t *testing.T) {
	opts := DefaultOptions()
	opts.Minimap = true
	result := buildMinimap(makeMinimapAnalysis(), opts)

	var s3 *MinimapEntry
	for i := range result.Entries {
		if result.Entries[i].Type == "aws_s3_bucket" {
			s3 = &result.Entries[i]
			break
		}
	}
	if s3 == nil {
		t.Fatal("aws_s3_bucket entry not found")
	}
	// 2 managed + 1 unmanaged = 3 total, 1 drifted => 33.33%
	if s3.Total != 3 {
		t.Errorf("expected total 3, got %d", s3.Total)
	}
	if s3.DriftPct < 33 || s3.DriftPct > 34 {
		t.Errorf("unexpected drift pct: %f", s3.DriftPct)
	}
}

func TestBuildMinimap_TopN(t *testing.T) {
	opts := DefaultOptions()
	opts.Minimap = true
	opts.MinimapTopN = 2
	result := buildMinimap(makeMinimapAnalysis(), opts)
	if len(result.Entries) != 2 {
		t.Fatalf("expected 2 entries with TopN=2, got %d", len(result.Entries))
	}
}

func TestBuildMinimap_ColorClasses(t *testing.T) {
	cases := []struct {
		pct      float64
		wantClass string
	}{
		{0, "minimap-ok"},
		{10, "minimap-low"},
		{40, "minimap-medium"},
		{80, "minimap-high"},
	}
	for _, tc := range cases {
		got := minimapColorClass(tc.pct)
		if got != tc.wantClass {
			t.Errorf("pct %.0f: want %s, got %s", tc.pct, tc.wantClass, got)
		}
	}
}
