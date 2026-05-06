package renderer

import (
	"testing"
)

func TestHighlightQuery_EmptyQuery(t *testing.T) {
	results := highlightQuery("aws_s3_bucket", "")
	if len(results) != 1 {
		t.Fatalf("expected 1 fragment, got %d", len(results))
	}
	if results[0].Highlighted {
		t.Error("expected fragment to not be highlighted")
	}
	if results[0].Text != "aws_s3_bucket" {
		t.Errorf("unexpected text: %q", results[0].Text)
	}
}

func TestHighlightQuery_NoMatch(t *testing.T) {
	results := highlightQuery("aws_s3_bucket", "lambda")
	if len(results) != 1 || results[0].Highlighted {
		t.Error("expected single non-highlighted fragment")
	}
}

func TestHighlightQuery_SingleMatch(t *testing.T) {
	results := highlightQuery("aws_s3_bucket", "s3")
	if len(results) != 3 {
		t.Fatalf("expected 3 fragments, got %d", len(results))
	}
	if results[0].Text != "aws_" || results[0].Highlighted {
		t.Errorf("unexpected prefix: %+v", results[0])
	}
	if results[1].Text != "s3" || !results[1].Highlighted {
		t.Errorf("unexpected match fragment: %+v", results[1])
	}
	if results[2].Text != "_bucket" || results[2].Highlighted {
		t.Errorf("unexpected suffix: %+v", results[2])
	}
}

func TestHighlightQuery_CaseInsensitive(t *testing.T) {
	results := highlightQuery("AWS_S3_Bucket", "s3")
	var found bool
	for _, r := range results {
		if r.Highlighted {
			found = true
			if r.Text != "S3" {
				t.Errorf("expected original casing S3, got %q", r.Text)
			}
		}
	}
	if !found {
		t.Error("expected a highlighted fragment")
	}
}

func TestHighlightQuery_MultipleMatches(t *testing.T) {
	results := highlightQuery("aXaXa", "X")
	highlighted := 0
	for _, r := range results {
		if r.Highlighted {
			highlighted++
		}
	}
	if highlighted != 2 {
		t.Errorf("expected 2 highlighted fragments, got %d", highlighted)
	}
}

func TestHighlightHTML_WrapsWithMark(t *testing.T) {
	out := highlightHTML("aws_s3_bucket", "s3")
	expected := "aws_<mark>s3</mark>_bucket"
	if out != expected {
		t.Errorf("expected %q, got %q", expected, out)
	}
}

func TestHighlightHTML_EmptyQuery(t *testing.T) {
	out := highlightHTML("aws_s3_bucket", "")
	if out != "aws_s3_bucket" {
		t.Errorf("unexpected output: %q", out)
	}
}

func TestHighlightHTML_NoMatch(t *testing.T) {
	out := highlightHTML("aws_s3_bucket", "lambda")
	if out != "aws_s3_bucket" {
		t.Errorf("unexpected output: %q", out)
	}
}
