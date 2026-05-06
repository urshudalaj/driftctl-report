package renderer

import (
	"testing"

	"github.com/snyk/driftctl-report/internal/parser"
)

func searchReport() parser.DriftctlReport {
	return parser.DriftctlReport{
		Summary: parser.Summary{
			Managed: []parser.Resource{
				{ID: "aws_s3_bucket.logs", Type: "aws_s3_bucket"},
				{ID: "aws_iam_role.deployer", Type: "aws_iam_role"},
			},
			Unmanaged: []parser.Resource{
				{ID: "aws_lambda_function.handler", Type: "aws_lambda_function"},
			},
			Missing: []parser.Resource{
				{ID: "aws_s3_bucket.backup", Type: "aws_s3_bucket"},
			},
		},
		Differences: []parser.DiffResource{
			{Resource: parser.Resource{ID: "aws_iam_role.deployer", Type: "aws_iam_role"}},
		},
	}
}

func TestSearchResources_EmptyQuery(t *testing.T) {
	r := searchReport()
	result := searchResources(r, SearchOptions{Query: ""})
	if len(result.Summary.Managed) != 2 {
		t.Errorf("expected 2 managed, got %d", len(result.Summary.Managed))
	}
}

func TestSearchResources_MatchByID(t *testing.T) {
	r := searchReport()
	result := searchResources(r, SearchOptions{Query: "logs"})
	if len(result.Summary.Managed) != 1 {
		t.Fatalf("expected 1 managed, got %d", len(result.Summary.Managed))
	}
	if result.Summary.Managed[0].ID != "aws_s3_bucket.logs" {
		t.Errorf("unexpected resource: %s", result.Summary.Managed[0].ID)
	}
}

func TestSearchResources_MatchByType(t *testing.T) {
	r := searchReport()
	result := searchResources(r, SearchOptions{Query: "aws_s3_bucket"})
	if len(result.Summary.Managed) != 1 {
		t.Errorf("expected 1 managed, got %d", len(result.Summary.Managed))
	}
	if len(result.Summary.Missing) != 1 {
		t.Errorf("expected 1 missing, got %d", len(result.Summary.Missing))
	}
}

func TestSearchResources_CaseInsensitive(t *testing.T) {
	r := searchReport()
	result := searchResources(r, SearchOptions{Query: "IAM"})
	if len(result.Summary.Managed) != 1 {
		t.Errorf("expected 1 managed, got %d", len(result.Summary.Managed))
	}
	if len(result.Differences) != 1 {
		t.Errorf("expected 1 diff, got %d", len(result.Differences))
	}
}

func TestSearchResources_NoMatch(t *testing.T) {
	r := searchReport()
	result := searchResources(r, SearchOptions{Query: "nonexistent"})
	if len(result.Summary.Managed) != 0 {
		t.Errorf("expected 0 managed, got %d", len(result.Summary.Managed))
	}
	if len(result.Summary.Unmanaged) != 0 {
		t.Errorf("expected 0 unmanaged, got %d", len(result.Summary.Unmanaged))
	}
}
