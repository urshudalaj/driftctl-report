package renderer

import (
	"testing"

	"github.com/snyk/driftctl-report/internal/parser"
)

func baseReport() parser.DriftctlReport {
	return parser.DriftctlReport{
		ManagedResources: []parser.Resource{
			{ID: "bucket-1", Type: "aws_s3_bucket"},
			{ID: "sg-1", Type: "aws_security_group"},
		},
		UnmanagedResources: []parser.Resource{
			{ID: "bucket-2", Type: "aws_s3_bucket"},
		},
		DeletedResources: []parser.Resource{
			{ID: "sg-2", Type: "aws_security_group"},
		},
		DifferentResources: []parser.DifferentResource{
			{ID: "bucket-3", Type: "aws_s3_bucket"},
		},
	}
}

func TestApplyFilters_NoOptions(t *testing.T) {
	report := baseReport()
	out := applyFilters(report, FilterOptions{})

	if len(out.ManagedResources) != 2 {
		t.Errorf("expected 2 managed resources, got %d", len(out.ManagedResources))
	}
	if len(out.UnmanagedResources) != 1 {
		t.Errorf("expected 1 unmanaged resource, got %d", len(out.UnmanagedResources))
	}
}

func TestApplyFilters_OnlyDrifted(t *testing.T) {
	report := baseReport()
	out := applyFilters(report, FilterOptions{OnlyDrifted: true})

	if len(out.ManagedResources) != 0 {
		t.Errorf("expected 0 managed resources when OnlyDrifted=true, got %d", len(out.ManagedResources))
	}
	if len(out.UnmanagedResources) != 1 {
		t.Errorf("expected 1 unmanaged resource, got %d", len(out.UnmanagedResources))
	}
	if len(out.DifferentResources) != 1 {
		t.Errorf("expected 1 different resource, got %d", len(out.DifferentResources))
	}
}

func TestApplyFilters_ResourceTypeFilter(t *testing.T) {
	report := baseReport()
	out := applyFilters(report, FilterOptions{ResourceTypes: []string{"aws_s3_bucket"}})

	if len(out.ManagedResources) != 1 {
		t.Errorf("expected 1 managed s3 resource, got %d", len(out.ManagedResources))
	}
	if len(out.DeletedResources) != 0 {
		t.Errorf("expected 0 deleted resources after type filter, got %d", len(out.DeletedResources))
	}
}

func TestUniqueResourceTypes(t *testing.T) {
	report := baseReport()
	types := uniqueResourceTypes(report)

	if len(types) != 2 {
		t.Errorf("expected 2 unique types, got %d: %v", len(types), types)
	}
	if types[0] != "aws_s3_bucket" || types[1] != "aws_security_group" {
		t.Errorf("unexpected type order: %v", types)
	}
}

func TestUniqueResourceTypes_Empty(t *testing.T) {
	types := uniqueResourceTypes(parser.DriftctlReport{})
	if len(types) != 0 {
		t.Errorf("expected 0 types for empty report, got %d", len(types))
	}
}
