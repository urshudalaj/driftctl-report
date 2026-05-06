package renderer

import (
	"testing"
)

func TestBuildTOC_Empty(t *testing.T) {
	toc := buildTOC(nil, nil, nil)
	if len(toc.Entries) != 0 {
		t.Fatalf("expected 0 entries, got %d", len(toc.Entries))
	}
}

func TestBuildTOC_SingleType(t *testing.T) {
	types := []string{"aws_s3_bucket"}
	counts := map[string]int{"aws_s3_bucket": 3}
	drifted := map[string]bool{"aws_s3_bucket": true}

	toc := buildTOC(types, counts, drifted)

	if len(toc.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(toc.Entries))
	}
	e := toc.Entries[0]
	if e.Label != "aws_s3_bucket" {
		t.Errorf("expected label aws_s3_bucket, got %s", e.Label)
	}
	if e.Count != 3 {
		t.Errorf("expected count 3, got %d", e.Count)
	}
	if !e.Drifted {
		t.Error("expected drifted=true")
	}
}

func TestBuildTOC_MultipleTypes_PreservesOrder(t *testing.T) {
	types := []string{"aws_iam_role", "aws_s3_bucket", "aws_vpc"}
	counts := map[string]int{"aws_iam_role": 1, "aws_s3_bucket": 5, "aws_vpc": 2}
	drifted := map[string]bool{"aws_vpc": true}

	toc := buildTOC(types, counts, drifted)

	if len(toc.Entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(toc.Entries))
	}
	if toc.Entries[0].Label != "aws_iam_role" {
		t.Errorf("expected first entry aws_iam_role, got %s", toc.Entries[0].Label)
	}
	if toc.Entries[2].Drifted != true {
		t.Error("expected aws_vpc to be marked drifted")
	}
}

func TestTOCAnchor_AlphanumericUnchanged(t *testing.T) {
	result := tocAnchor("awsS3Bucket123")
	if result != "awsS3Bucket123" {
		t.Errorf("expected awsS3Bucket123, got %s", result)
	}
}

func TestTOCAnchor_SpecialCharsReplaced(t *testing.T) {
	result := tocAnchor("aws_s3.bucket/type")
	expected := "aws-s3-bucket-type"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestTOCAnchor_EmptyString(t *testing.T) {
	result := tocAnchor("")
	if result != "" {
		t.Errorf("expected empty string, got %s", result)
	}
}
