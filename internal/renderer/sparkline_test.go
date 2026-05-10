package renderer

import "testing"

func makeSparklineAnalysis() Analysis {
	return Analysis{
		UnmanagedResources: []Resource{
			{ID: "bucket-1", Type: "aws_s3_bucket"},
			{ID: "bucket-2", Type: "aws_s3_bucket"},
			{ID: "fn-1", Type: "aws_lambda_function"},
		},
		DeletedResources: []Resource{
			{ID: "sg-1", Type: "aws_security_group"},
		},
	}
}

func TestBuildSparkline_Disabled(t *testing.T) {
	opts := DefaultOptions()
	opts.Sparkline = false
	sd := buildSparkline(makeSparklineAnalysis(), opts)
	if sd.Enabled {
		t.Fatal("expected sparkline disabled")
	}
	if len(sd.Points) != 0 {
		t.Fatalf("expected no points, got %d", len(sd.Points))
	}
}

func TestBuildSparkline_PointCount(t *testing.T) {
	opts := DefaultOptions()
	opts.Sparkline = true
	sd := buildSparkline(makeSparklineAnalysis(), opts)
	if !sd.Enabled {
		t.Fatal("expected sparkline enabled")
	}
	// 3 distinct types: aws_lambda_function, aws_s3_bucket, aws_security_group
	if len(sd.Points) != 3 {
		t.Fatalf("expected 3 points, got %d", len(sd.Points))
	}
}

func TestBuildSparkline_SortedByType(t *testing.T) {
	opts := DefaultOptions()
	opts.Sparkline = true
	sd := buildSparkline(makeSparklineAnalysis(), opts)
	for i := 1; i < len(sd.Points); i++ {
		if sd.Points[i].Label < sd.Points[i-1].Label {
			t.Fatalf("points not sorted: %s before %s", sd.Points[i-1].Label, sd.Points[i].Label)
		}
	}
}

func TestBuildSparkline_MaxAndMin(t *testing.T) {
	opts := DefaultOptions()
	opts.Sparkline = true
	sd := buildSparkline(makeSparklineAnalysis(), opts)
	if sd.Max != 2 {
		t.Fatalf("expected max=2, got %d", sd.Max)
	}
	if sd.Min != 1 {
		t.Fatalf("expected min=1, got %d", sd.Min)
	}
}

func TestBuildSparkline_TrendUp(t *testing.T) {
	opts := DefaultOptions()
	opts.Sparkline = true
	a := Analysis{
		UnmanagedResources: []Resource{
			{ID: "a-1", Type: "aws_a"},
			{ID: "b-1", Type: "aws_b"},
			{ID: "b-2", Type: "aws_b"},
			{ID: "b-3", Type: "aws_b"},
		},
	}
	sd := buildSparkline(a, opts)
	if sd.Trend != "up" {
		t.Fatalf("expected trend=up, got %s", sd.Trend)
	}
}

func TestBuildSparkline_EmptyAnalysis(t *testing.T) {
	opts := DefaultOptions()
	opts.Sparkline = true
	sd := buildSparkline(Analysis{}, opts)
	if !sd.Enabled {
		t.Fatal("expected enabled")
	}
	if sd.Trend != "flat" {
		t.Fatalf("expected flat trend for empty, got %s", sd.Trend)
	}
}
