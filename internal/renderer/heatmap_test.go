package renderer

import (
	"testing"

	"github.com/snyk/driftctl/enumeration/resource"
)

func makeHeatmapAnalysis() Analysis {
	return Analysis{
		Managed: []resource.Resource{
			{ResourceType: "aws_s3_bucket", ResourceId: "bucket-1"},
			{ResourceType: "aws_s3_bucket", ResourceId: "bucket-2"},
			{ResourceType: "aws_iam_role", ResourceId: "role-1"},
		},
		Unmanaged: []resource.Resource{
			{ResourceType: "aws_s3_bucket", ResourceId: "bucket-3"},
			{ResourceType: "aws_lambda_function", ResourceId: "fn-1"},
		},
		Deleted: []resource.Resource{
			{ResourceType: "aws_iam_role", ResourceId: "role-2"},
		},
	}
}

func TestBuildHeatmap_Disabled(t *testing.T) {
	opts := DefaultOptions()
	opts.Heatmap = false
	h := buildHeatmap(makeHeatmapAnalysis(), opts)
	if h.Enabled {
		t.Fatal("expected heatmap to be disabled")
	}
	if len(h.Cells) != 0 {
		t.Fatalf("expected no cells, got %d", len(h.Cells))
	}
}

func TestBuildHeatmap_CellCount(t *testing.T) {
	opts := DefaultOptions()
	opts.Heatmap = true
	h := buildHeatmap(makeHeatmapAnalysis(), opts)
	if !h.Enabled {
		t.Fatal("expected heatmap to be enabled")
	}
	if len(h.Cells) != 3 {
		t.Fatalf("expected 3 cells, got %d", len(h.Cells))
	}
}

func TestBuildHeatmap_SortedByDriftRatioDesc(t *testing.T) {
	opts := DefaultOptions()
	opts.Heatmap = true
	h := buildHeatmap(makeHeatmapAnalysis(), opts)
	for i := 1; i < len(h.Cells); i++ {
		if h.Cells[i].DriftRatio > h.Cells[i-1].DriftRatio {
			t.Errorf("cells not sorted by drift ratio desc at index %d", i)
		}
	}
}

func TestBuildHeatmap_HeatLevels(t *testing.T) {
	opts := DefaultOptions()
	opts.Heatmap = true
	h := buildHeatmap(makeHeatmapAnalysis(), opts)

	levels := map[string]string{}
	for _, c := range h.Cells {
		levels[c.Type] = c.HeatLevel
	}

	// aws_lambda_function: 1 unmanaged / 1 total = 1.0 → high
	if levels["aws_lambda_function"] != "high" {
		t.Errorf("expected high for aws_lambda_function, got %s", levels["aws_lambda_function"])
	}
	// aws_s3_bucket: 1 unmanaged / 3 total ≈ 0.33 → medium
	if levels["aws_s3_bucket"] != "medium" {
		t.Errorf("expected medium for aws_s3_bucket, got %s", levels["aws_s3_bucket"])
	}
	// aws_iam_role: 0 unmanaged, 1 deleted / 2 total = 0.5 → medium
	if levels["aws_iam_role"] != "medium" {
		t.Errorf("expected medium for aws_iam_role, got %s", levels["aws_iam_role"])
	}
}

func TestHeatLevel_Boundaries(t *testing.T) {
	cases := []struct {
		ratio float64
		want  string
	}{
		{0.0, "none"},
		{0.1, "low"},
		{0.25, "medium"},
		{0.59, "medium"},
		{0.60, "high"},
		{1.0, "high"},
	}
	for _, tc := range cases {
		got := heatLevel(tc.ratio)
		if got != tc.want {
			t.Errorf("heatLevel(%v) = %q, want %q", tc.ratio, got, tc.want)
		}
	}
}
