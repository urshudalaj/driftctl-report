package renderer

import (
	"testing"

	"github.com/snyk/driftctl/pkg/analyser"
	"github.com/snyk/driftctl/pkg/resource"
)

func makeSummaryAnalysis(managed, unmanaged, missing int) analyser.Analysis {
	a := analyser.Analysis{}
	for i := 0; i < managed; i++ {
		a.AddManaged(&resource.Resource{ResourceId: fmt.Sprintf("managed-%d", i), ResourceType: "aws_s3_bucket"})
	}
	for i := 0; i < unmanaged; i++ {
		a.AddUnmanaged(&resource.Resource{ResourceId: fmt.Sprintf("unmanaged-%d", i), ResourceType: "aws_s3_bucket"})
	}
	for i := 0; i < missing; i++ {
		a.AddMissing(&resource.Resource{ResourceId: fmt.Sprintf("missing-%d", i), ResourceType: "aws_s3_bucket"})
	}
	return a
}

func TestBuildSummary_AllZero(t *testing.T) {
	s := buildSummary(analyser.Analysis{})
	if s.Total != 0 || s.Coverage != 0 {
		t.Fatalf("expected zero totals, got %+v", s)
	}
}

func TestBuildSummary_FullCoverage(t *testing.T) {
	a := makeSummaryAnalysis(5, 0, 0)
	s := buildSummary(a)
	if s.Coverage != 100.0 {
		t.Fatalf("expected 100%% coverage, got %.2f", s.Coverage)
	}
	if s.Managed != 5 || s.Total != 5 {
		t.Fatalf("unexpected counts: %+v", s)
	}
}

func TestBuildSummary_PartialCoverage(t *testing.T) {
	a := makeSummaryAnalysis(1, 1, 2)
	s := buildSummary(a)
	// managed=1, total=4 → 25%
	if s.Coverage != 25.0 {
		t.Fatalf("expected 25%% coverage, got %.2f", s.Coverage)
	}
	if s.Unmanaged != 1 || s.Missing != 2 {
		t.Fatalf("unexpected counts: %+v", s)
	}
}

func TestBuildSummary_RoundsCorrectly(t *testing.T) {
	a := makeSummaryAnalysis(1, 0, 2) // 1/3 ≈ 33.33
	s := buildSummary(a)
	if s.Coverage != 33.33 {
		t.Fatalf("expected 33.33, got %.2f", s.Coverage)
	}
}

func TestRoundTwo(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{33.3333, 33.33},
		{99.999, 100.0},
		{0.0, 0.0},
		{50.505, 50.51},
	}
	for _, c := range cases {
		got := roundTwo(c.in)
		if got != c.want {
			t.Errorf("roundTwo(%.4f) = %.2f, want %.2f", c.in, got, c.want)
		}
	}
}
