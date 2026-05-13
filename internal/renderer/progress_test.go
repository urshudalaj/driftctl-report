package renderer

import (
	"testing"
)

func makeProgressAnalysis() Analysis {
	return Analysis{
		Managed: []Resource{
			{ID: "a", Type: "aws_s3_bucket"},
			{ID: "b", Type: "aws_s3_bucket"},
			{ID: "c", Type: "aws_iam_role"},
		},
		Unmanaged: []Resource{
			{ID: "d", Type: "aws_s3_bucket"},
			{ID: "e", Type: "aws_lambda_function"},
		},
		Deleted: []Resource{
			{ID: "f", Type: "aws_iam_role"},
		},
	}
}

func TestBuildProgress_Disabled(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowProgress = false
	bars := buildProgress(makeProgressAnalysis(), opts)
	if bars != nil {
		t.Errorf("expected nil when disabled, got %v", bars)
	}
}

func TestBuildProgress_BarCount(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowProgress = true
	bars := buildProgress(makeProgressAnalysis(), opts)
	// aws_iam_role, aws_lambda_function, aws_s3_bucket
	if len(bars) != 3 {
		t.Errorf("expected 3 bars, got %d", len(bars))
	}
}

func TestBuildProgress_SortedByType(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowProgress = true
	bars := buildProgress(makeProgressAnalysis(), opts)
	if bars[0].Label != "aws_iam_role" {
		t.Errorf("expected first bar aws_iam_role, got %s", bars[0].Label)
	}
	if bars[2].Label != "aws_s3_bucket" {
		t.Errorf("expected last bar aws_s3_bucket, got %s", bars[2].Label)
	}
}

func TestBuildProgress_PercentageCalculation(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowProgress = true
	bars := buildProgress(makeProgressAnalysis(), opts)
	// aws_s3_bucket: 2 managed / 3 total = 66.67
	var s3bar ProgressBar
	for _, b := range bars {
		if b.Label == "aws_s3_bucket" {
			s3bar = b
			break
		}
	}
	if s3bar.Value != 2 || s3bar.Max != 3 {
		t.Errorf("unexpected s3 counts: value=%d max=%d", s3bar.Value, s3bar.Max)
	}
	if s3bar.Percentage != 66.67 {
		t.Errorf("expected 66.67%%, got %f", s3bar.Percentage)
	}
}

func TestBuildProgress_TopN(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowProgress = true
	opts.ProgressTopN = 2
	bars := buildProgress(makeProgressAnalysis(), opts)
	if len(bars) != 2 {
		t.Errorf("expected 2 bars with TopN=2, got %d", len(bars))
	}
}

func TestProgressBarClass_Levels(t *testing.T) {
	cases := []struct {
		pct   float64
		want  string
	}{
		{100, "success"},
		{90, "success"},
		{75, "info"},
		{60, "info"},
		{45, "warning"},
		{10, "danger"},
	}
	for _, tc := range cases {
		got := progressBarClass(tc.pct)
		if got != tc.want {
			t.Errorf("pct=%.0f: expected %s, got %s", tc.pct, tc.want, got)
		}
	}
}
