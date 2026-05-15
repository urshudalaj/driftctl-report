package renderer

import (
	"testing"

	"github.com/snyk/driftctl-report/internal/parser"
)

func TestBuildBreadcrumb_Disabled(t *testing.T) {
	o := DefaultOptions()
	got := buildBreadcrumb(&parser.Analysis{}, o)
	if got.Enabled {
		t.Fatal("expected Enabled=false when breadcrumb is disabled")
	}
	if len(got.Items) != 0 {
		t.Fatalf("expected no items, got %d", len(got.Items))
	}
}

func TestBuildBreadcrumb_DefaultSeparator(t *testing.T) {
	o := DefaultOptions()
	WithBreadcrumb(true)(&o)
	got := buildBreadcrumb(&parser.Analysis{}, o)
	if got.Separator != "/" {
		t.Fatalf("expected default separator '/', got %q", got.Separator)
	}
}

func TestBuildBreadcrumb_CustomSeparator(t *testing.T) {
	o := DefaultOptions()
	WithBreadcrumb(true)(&o)
	WithBreadcrumbSeparator(">")(&o)
	got := buildBreadcrumb(&parser.Analysis{}, o)
	if got.Separator != ">" {
		t.Fatalf("expected separator '>', got %q", got.Separator)
	}
}

func TestBuildBreadcrumb_ItemsPreserveOrder(t *testing.T) {
	o := DefaultOptions()
	WithBreadcrumb(true)(&o)
	WithBreadcrumbItems(
		BreadcrumbItem{Label: "Home", URL: "/"},
		BreadcrumbItem{Label: "Reports", URL: "/reports"},
		BreadcrumbItem{Label: "Drift"},
	)(&o)
	got := buildBreadcrumb(&parser.Analysis{}, o)
	if len(got.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(got.Items))
	}
	if got.Items[0].Label != "Home" || got.Items[1].Label != "Reports" || got.Items[2].Label != "Drift" {
		t.Fatalf("unexpected item order: %+v", got.Items)
	}
}

func TestBuildBreadcrumb_EmptyLabelItemsIgnored(t *testing.T) {
	o := DefaultOptions()
	WithBreadcrumb(true)(&o)
	WithBreadcrumbItems(
		BreadcrumbItem{Label: "", URL: "/ignored"},
		BreadcrumbItem{Label: "Valid"},
	)(&o)
	got := buildBreadcrumb(&parser.Analysis{}, o)
	if len(got.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(got.Items))
	}
}

func TestBuildBreadcrumb_ItemURLOptional(t *testing.T) {
	o := DefaultOptions()
	WithBreadcrumb(true)(&o)
	WithBreadcrumbItems(BreadcrumbItem{Label: "Current"})(&o)
	got := buildBreadcrumb(&parser.Analysis{}, o)
	if got.Items[0].URL != "" {
		t.Fatalf("expected empty URL, got %q", got.Items[0].URL)
	}
}

func TestWithBreadcrumbSeparator_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithBreadcrumb(true)(&o)
	WithBreadcrumbSeparator("")(&o)
	got := buildBreadcrumb(&parser.Analysis{}, o)
	if got.Separator != "/" {
		t.Fatalf("empty separator should fall back to '/', got %q", got.Separator)
	}
}
