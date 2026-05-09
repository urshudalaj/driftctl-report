package renderer

import (
	"testing"

	"github.com/snyk/driftctl/pkg/resource"
)

func makeComplianceAnalysis(managed, unmanaged, deleted int, t string) Analysis {
	a := Analysis{}
	for i := 0; i < managed; i++ {
		a.Managed = append(a.Managed, resource.Resource{ResourceType: t, ResourceId: itoa(i)})
	}
	for i := 0; i < unmanaged; i++ {
		a.Unmanaged = append(a.Unmanaged, resource.Resource{ResourceType: t, ResourceId: itoa(i)})
	}
	for i := 0; i < deleted; i++ {
		a.Deleted = append(a.Deleted, resource.Resource{ResourceType: t, ResourceId: itoa(i)})
	}
	return a
}

func TestBuildCompliance_Disabled(t *testing.T) {
	a := makeComplianceAnalysis(5, 3, 1, "aws_s3_bucket")
	opts := DefaultOptions()
	opts.Compliance = false
	r := buildCompliance(a, opts)
	if r.Enabled {
		t.Fatal("expected compliance to be disabled")
	}
}

func TestBuildCompliance_FullyManaged(t *testing.T) {
	a := makeComplianceAnalysis(10, 0, 0, "aws_s3_bucket")
	opts := DefaultOptions()
	opts.Compliance = true
	r := buildCompliance(a, opts)
	if !r.Enabled {
		t.Fatal("expected compliance to be enabled")
	}
	if r.Overall != 100.0 {
		t.Fatalf("expected 100%% compliance, got %.2f", r.Overall)
	}
	if r.Grade != "A" {
		t.Fatalf("expected grade A, got %s", r.Grade)
	}
}

func TestBuildCompliance_AllUnmanaged(t *testing.T) {
	a := makeComplianceAnalysis(0, 10, 0, "aws_instance")
	opts := DefaultOptions()
	opts.Compliance = true
	r := buildCompliance(a, opts)
	if r.Overall != 0.0 {
		t.Fatalf("expected 0%% compliance, got %.2f", r.Overall)
	}
	if r.Grade != "F" {
		t.Fatalf("expected grade F, got %s", r.Grade)
	}
}

func TestBuildCompliance_PartialCompliance(t *testing.T) {
	a := makeComplianceAnalysis(8, 2, 0, "aws_iam_role")
	opts := DefaultOptions()
	opts.Compliance = true
	r := buildCompliance(a, opts)
	if r.Overall != 80.0 {
		t.Fatalf("expected 80%% compliance, got %.2f", r.Overall)
	}
	if r.Grade != "B" {
		t.Fatalf("expected grade B, got %s", r.Grade)
	}
}

func TestBuildCompliance_TopN(t *testing.T) {
	a := makeComplianceAnalysis(3, 1, 0, "aws_s3_bucket")
	a.Unmanaged = append(a.Unmanaged, resource.Resource{ResourceType: "aws_instance", ResourceId: "x"})
	a.Unmanaged = append(a.Unmanaged, resource.Resource{ResourceType: "aws_iam_role", ResourceId: "y"})
	opts := DefaultOptions()
	opts.Compliance = true
	opts.ComplianceTopN = 2
	r := buildCompliance(a, opts)
	if len(r.Entries) > 2 {
		t.Fatalf("expected at most 2 entries, got %d", len(r.Entries))
	}
}

func TestComplianceGrade_Boundaries(t *testing.T) {
	cases := []struct {
		pct   float64
		want  string
	}{
		{100, "A"}, {95, "A"}, {94, "B"}, {80, "B"}, {79, "C"},
		{60, "C"}, {59, "D"}, {40, "D"}, {39, "F"}, {0, "F"},
	}
	for _, tc := range cases {
		got := complianceGrade(tc.pct)
		if got != tc.want {
			t.Errorf("complianceGrade(%.0f) = %s, want %s", tc.pct, got, tc.want)
		}
	}
}
