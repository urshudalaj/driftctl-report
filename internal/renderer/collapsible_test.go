package renderer

import "testing"

func TestBuildCollapsible_Disabled(t *testing.T) {
	o := DefaultOptions()
	data := buildCollapsible(o, []string{"aws_s3_bucket"}, map[string]int{"aws_s3_bucket": 3})
	if data.Enabled {
		t.Fatal("expected Enabled=false when WithCollapsible not set")
	}
	if len(data.Sections) != 0 {
		t.Fatalf("expected 0 sections, got %d", len(data.Sections))
	}
}

func TestBuildCollapsible_SectionCount(t *testing.T) {
	o := DefaultOptions(WithCollapsible(true))
	types := []string{"aws_s3_bucket", "aws_iam_role", "aws_lambda_function"}
	counts := map[string]int{"aws_s3_bucket": 2, "aws_iam_role": 5, "aws_lambda_function": 1}
	data := buildCollapsible(o, types, counts)
	if !data.Enabled {
		t.Fatal("expected Enabled=true")
	}
	if len(data.Sections) != 3 {
		t.Fatalf("expected 3 sections, got %d", len(data.Sections))
	}
}

func TestBuildCollapsible_DefaultOpenFalse(t *testing.T) {
	o := DefaultOptions(WithCollapsible(true))
	data := buildCollapsible(o, []string{"aws_s3_bucket"}, map[string]int{"aws_s3_bucket": 1})
	if data.Sections[0].Open {
		t.Error("expected Open=false by default")
	}
}

func TestBuildCollapsible_DefaultOpenTrue(t *testing.T) {
	o := DefaultOptions(WithCollapsible(true), WithCollapsibleDefaultOpen(true))
	data := buildCollapsible(o, []string{"aws_s3_bucket"}, map[string]int{"aws_s3_bucket": 1})
	if !data.Sections[0].Open {
		t.Error("expected Open=true when WithCollapsibleDefaultOpen(true)")
	}
}

func TestBuildCollapsible_AnimatedPropagates(t *testing.T) {
	o := DefaultOptions(WithCollapsible(true), WithCollapsibleAnimated(true))
	data := buildCollapsible(o, []string{"aws_s3_bucket"}, map[string]int{"aws_s3_bucket": 4})
	if !data.Sections[0].Animated {
		t.Error("expected Animated=true when WithCollapsibleAnimated(true)")
	}
}

func TestBuildCollapsible_IDIsAnchorSafe(t *testing.T) {
	o := DefaultOptions(WithCollapsible(true))
	data := buildCollapsible(o, []string{"aws_s3_bucket"}, map[string]int{"aws_s3_bucket": 1})
	got := data.Sections[0].ID
	want := tocAnchor("aws_s3_bucket")
	if got != want {
		t.Errorf("expected ID %q, got %q", want, got)
	}
}

func TestBuildCollapsible_ItemCountMatches(t *testing.T) {
	o := DefaultOptions(WithCollapsible(true))
	counts := map[string]int{"aws_iam_role": 7}
	data := buildCollapsible(o, []string{"aws_iam_role"}, counts)
	if data.Sections[0].ItemCount != 7 {
		t.Errorf("expected ItemCount=7, got %d", data.Sections[0].ItemCount)
	}
}

func TestBuildCollapsible_PreservesTypeOrder(t *testing.T) {
	o := DefaultOptions(WithCollapsible(true))
	types := []string{"zzz_type", "aaa_type", "mmm_type"}
	counts := map[string]int{"zzz_type": 1, "aaa_type": 2, "mmm_type": 3}
	data := buildCollapsible(o, types, counts)
	for i, want := range types {
		if data.Sections[i].Label != want {
			t.Errorf("position %d: expected %q, got %q", i, want, data.Sections[i].Label)
		}
	}
}
