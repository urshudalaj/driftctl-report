package renderer

import (
	"testing"

	"github.com/snyk/driftctl/pkg/resource"
)

func makeResource(id, rtype string) resource.Resource {
	return resource.Resource{ID: id, Type: rtype}
}

func TestSortResources_ByTypeAsc(t *testing.T) {
	resources := []resource.Resource{
		makeResource("c", "aws_s3_bucket"),
		makeResource("a", "aws_iam_role"),
		makeResource("b", "aws_iam_role"),
	}
	sortResources(resources, SortByTypeAsc)
	if resources[0].ResourceType() != "aws_iam_role" {
		t.Errorf("expected aws_iam_role first, got %s", resources[0].ResourceType())
	}
	if resources[1].ResourceId() != "a" {
		t.Errorf("expected id 'a' second, got %s", resources[1].ResourceId())
	}
}

func TestSortResources_ByTypeDesc(t *testing.T) {
	resources := []resource.Resource{
		makeResource("a", "aws_iam_role"),
		makeResource("c", "aws_s3_bucket"),
	}
	sortResources(resources, SortByTypeDesc)
	if resources[0].ResourceType() != "aws_s3_bucket" {
		t.Errorf("expected aws_s3_bucket first, got %s", resources[0].ResourceType())
	}
}

func TestSortResources_ByIDDesc(t *testing.T) {
	resources := []resource.Resource{
		makeResource("alpha", "aws_iam_role"),
		makeResource("gamma", "aws_iam_role"),
		makeResource("beta", "aws_iam_role"),
	}
	sortResources(resources, SortByIDDesc)
	if resources[0].ResourceId() != "gamma" {
		t.Errorf("expected gamma first, got %s", resources[0].ResourceId())
	}
}

func TestSortResources_ByIDAsc(t *testing.T) {
	resources := []resource.Resource{
		makeResource("gamma", "aws_iam_role"),
		makeResource("alpha", "aws_iam_role"),
		makeResource("beta", "aws_iam_role"),
	}
	sortResources(resources, SortByIDAsc)
	if resources[0].ResourceId() != "alpha" {
		t.Errorf("expected alpha first, got %s", resources[0].ResourceId())
	}
	if resources[2].ResourceId() != "gamma" {
		t.Errorf("expected gamma last, got %s", resources[2].ResourceId())
	}
}

func TestSortResources_Empty(t *testing.T) {
	var resources []resource.Resource
	sortResources(resources, SortByTypeAsc)
	if len(resources) != 0 {
		t.Error("expected empty slice to remain empty")
	}
}

func TestSortDiffResources_ByTypeAsc(t *testing.T) {
	r1 := makeResource("z", "aws_s3_bucket")
	r2 := makeResource("a", "aws_iam_role")
	diffs := []resource.ResourceDiff{
		{Res: &r1},
		{Res: &r2},
	}
	sortDiffResources(diffs, SortByTypeAsc)
	if diffs[0].Res.ResourceType() != "aws_iam_role" {
		t.Errorf("expected aws_iam_role first, got %s", diffs[0].Res.ResourceType())
	}
}
