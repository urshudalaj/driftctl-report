package renderer

import (
	"testing"
)

func TestBuildAnnotations_Empty(t *testing.T) {
	am := buildAnnotations(nil, []string{"aws_s3_bucket.foo"})
	if len(am) != 0 {
		t.Fatalf("expected empty map, got %d entries", len(am))
	}
}

func TestBuildAnnotations_NoKnownIDs(t *testing.T) {
	anns := []Annotation{{ResourceID: "aws_s3_bucket.foo", Label: "prod", Note: "critical"}}
	am := buildAnnotations(anns, nil)
	if len(am) != 0 {
		t.Fatalf("expected empty map, got %d entries", len(am))
	}
}

func TestBuildAnnotations_FiltersUnknown(t *testing.T) {
	anns := []Annotation{
		{ResourceID: "aws_s3_bucket.foo", Label: "prod"},
		{ResourceID: "aws_lambda_function.bar", Label: "test"},
	}
	am := buildAnnotations(anns, []string{"aws_s3_bucket.foo"})
	if len(am) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(am))
	}
	if _, ok := am["aws_s3_bucket.foo"]; !ok {
		t.Error("expected aws_s3_bucket.foo in annotation map")
	}
}

func TestBuildAnnotations_CaseInsensitiveMatch(t *testing.T) {
	anns := []Annotation{
		{ResourceID: "AWS_S3_Bucket.Foo", Label: "prod"},
	}
	am := buildAnnotations(anns, []string{"aws_s3_bucket.foo"})
	if len(am) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(am))
	}
}

func TestBuildAnnotations_PreservesLabelAndNote(t *testing.T) {
	anns := []Annotation{
		{ResourceID: "res.a", Label: "important", Note: "do not delete"},
	}
	am := buildAnnotations(anns, []string{"res.a"})
	a := am["res.a"]
	if a.Label != "important" {
		t.Errorf("expected label 'important', got %q", a.Label)
	}
	if a.Note != "do not delete" {
		t.Errorf("expected note 'do not delete', got %q", a.Note)
	}
}

func TestSortedAnnotations_OrderedByID(t *testing.T) {
	am := AnnotationMap{
		"res.z": {ResourceID: "res.z", Label: "last"},
		"res.a": {ResourceID: "res.a", Label: "first"},
		"res.m": {ResourceID: "res.m", Label: "middle"},
	}
	sorted := sortedAnnotations(am)
	if len(sorted) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(sorted))
	}
	if sorted[0].ResourceID != "res.a" || sorted[1].ResourceID != "res.m" || sorted[2].ResourceID != "res.z" {
		t.Errorf("unexpected order: %v", sorted)
	}
}

func TestSortedAnnotations_Empty(t *testing.T) {
	sorted := sortedAnnotations(AnnotationMap{})
	if len(sorted) != 0 {
		t.Fatalf("expected empty slice, got %d", len(sorted))
	}
}
