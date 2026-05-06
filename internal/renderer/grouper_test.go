package renderer

import (
	"testing"

	"github.com/snyk/driftctl/pkg/analyser"
)

func makeRes(typ, id string) analyser.Resource {
	return analyser.Resource{Type: typ, ResourceID: id}
}

func makeDiffRes(typ, id string) analyser.DiffResource {
	return analyser.DiffResource{Type: typ, ResourceID: id}
}

func TestGroupByType_Empty(t *testing.T) {
	groups := groupByType(nil)
	if len(groups) != 0 {
		t.Fatalf("expected 0 groups, got %d", len(groups))
	}
}

func TestGroupByType_SingleType(t *testing.T) {
	resources := []analyser.Resource{
		makeRes("aws_s3_bucket", "bucket-1"),
		makeRes("aws_s3_bucket", "bucket-2"),
	}
	groups := groupByType(resources)
	if len(groups) != 1 {
		t.Fatalf("expected 1 group, got %d", len(groups))
	}
	if groups[0].Type != "aws_s3_bucket" {
		t.Errorf("unexpected type: %s", groups[0].Type)
	}
	if groups[0].Count != 2 {
		t.Errorf("expected count 2, got %d", groups[0].Count)
	}
}

func TestGroupByType_MultipleTypes_PreservesOrder(t *testing.T) {
	resources := []analyser.Resource{
		makeRes("aws_lambda_function", "fn-a"),
		makeRes("aws_s3_bucket", "bucket-1"),
		makeRes("aws_lambda_function", "fn-b"),
		makeRes("aws_iam_role", "role-1"),
	}
	groups := groupByType(resources)
	if len(groups) != 3 {
		t.Fatalf("expected 3 groups, got %d", len(groups))
	}
	expected := []string{"aws_lambda_function", "aws_s3_bucket", "aws_iam_role"}
	for i, g := range groups {
		if g.Type != expected[i] {
			t.Errorf("group[%d]: expected %s, got %s", i, expected[i], g.Type)
		}
	}
	if groups[0].Count != 2 {
		t.Errorf("lambda group count: expected 2, got %d", groups[0].Count)
	}
}

func TestGroupDiffByType_Empty(t *testing.T) {
	groups := groupDiffByType(nil)
	if len(groups) != 0 {
		t.Fatalf("expected 0 groups, got %d", len(groups))
	}
}

func TestGroupDiffByType_MixedTypes(t *testing.T) {
	resources := []analyser.DiffResource{
		makeDiffRes("aws_instance", "i-001"),
		makeDiffRes("aws_instance", "i-002"),
		makeDiffRes("aws_vpc", "vpc-001"),
	}
	groups := groupDiffByType(resources)
	if len(groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(groups))
	}
	if groups[0].Type != "aws_instance" || groups[0].Count != 2 {
		t.Errorf("unexpected first group: %+v", groups[0])
	}
	if groups[1].Type != "aws_vpc" || groups[1].Count != 1 {
		t.Errorf("unexpected second group: %+v", groups[1])
	}
}
