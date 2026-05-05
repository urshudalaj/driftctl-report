package renderer_test

import (
	"strings"
	"testing"

	"github.com/your-org/driftctl-report/internal/parser"
	"github.com/your-org/driftctl-report/internal/renderer"
)

func sampleReport() *parser.DriftReport {
	return &parser.DriftReport{
		Summary: parser.Summary{
			TotalFound:   5,
			TotalManaged: 3,
			Coverage:     60.0,
		},
		ManagedResources: []parser.Resource{
			{ID: "bucket-1", Type: "aws_s3_bucket"},
		},
		UnmanagedResources: []parser.Resource{
			{ID: "sg-abc123", Type: "aws_security_group"},
		},
		MissingResources: []parser.Resource{
			{ID: "role-xyz", Type: "aws_iam_role"},
		},
		DifferentResources: []parser.Resource{
			{ID: "lambda-fn", Type: "aws_lambda_function"},
		},
	}
}

func TestNew_ReturnsRenderer(t *testing.T) {
	r, err := renderer.New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if r == nil {
		t.Fatal("expected non-nil renderer")
	}
}

func TestRender_ContainsSummaryCards(t *testing.T) {
	r, _ := renderer.New()
	html, err := r.RenderToString(sampleReport())
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	for _, want := range []string{"60.0%", "Managed", "Unmanaged", "Missing", "Drifted"} {
		if !strings.Contains(html, want) {
			t.Errorf("expected HTML to contain %q", want)
		}
	}
}

func TestRender_ContainsResourceIDs(t *testing.T) {
	r, _ := renderer.New()
	html, err := r.RenderToString(sampleReport())
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	for _, id := range []string{"sg-abc123", "role-xyz", "lambda-fn"} {
		if !strings.Contains(html, id) {
			t.Errorf("expected HTML to contain resource ID %q", id)
		}
	}
}

func TestRender_EmptyReport(t *testing.T) {
	r, _ := renderer.New()
	empty := &parser.DriftReport{}
	html, err := r.RenderToString(empty)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	if !strings.Contains(html, "driftctl Drift Report") {
		t.Error("expected HTML to contain report title")
	}
}

func TestRender_WritesToWriter(t *testing.T) {
	r, _ := renderer.New()
	var buf strings.Builder
	if err := r.Render(&buf, sampleReport()); err != nil {
		t.Fatalf("render error: %v", err)
	}
	if buf.Len() == 0 {
		t.Error("expected non-empty output written to writer")
	}
}
