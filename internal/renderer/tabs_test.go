package renderer

import "testing"

func makeTabsAnalysis() *Analysis {
	return &Analysis{
		Managed:   []Resource{{ID: "r1", Type: "aws_s3_bucket"}},
		Unmanaged: []Resource{{ID: "r2", Type: "aws_instance"}},
		Deleted:   []Resource{{ID: "r3", Type: "aws_iam_role"}},
		Drifted:   []DiffResource{{Resource: Resource{ID: "r4", Type: "aws_vpc"}}},
	}
}

func TestBuildTabs_Disabled(t *testing.T) {
	o := DefaultOptions()
	d := buildTabs(makeTabsAnalysis(), o)
	if d.Enabled {
		t.Fatal("expected tabs to be disabled")
	}
}

func TestBuildTabs_TabCount(t *testing.T) {
	o := DefaultOptions()
	WithTabs(true)(o)
	d := buildTabs(makeTabsAnalysis(), o)
	if len(d.Tabs) != 4 {
		t.Fatalf("expected 4 tabs, got %d", len(d.Tabs))
	}
}

func TestBuildTabs_DefaultPosition(t *testing.T) {
	o := DefaultOptions()
	WithTabs(true)(o)
	d := buildTabs(makeTabsAnalysis(), o)
	if d.Position != "top" {
		t.Fatalf("expected position 'top', got %q", d.Position)
	}
}

func TestBuildTabs_CustomPosition(t *testing.T) {
	o := DefaultOptions()
	WithTabs(true)(o)
	WithTabsPosition("left")(o)
	d := buildTabs(makeTabsAnalysis(), o)
	if d.Position != "left" {
		t.Fatalf("expected position 'left', got %q", d.Position)
	}
}

func TestBuildTabs_FirstTabActiveByDefault(t *testing.T) {
	o := DefaultOptions()
	WithTabs(true)(o)
	d := buildTabs(makeTabsAnalysis(), o)
	if !d.Tabs[0].Active {
		t.Fatal("expected first tab to be active")
	}
	for _, tab := range d.Tabs[1:] {
		if tab.Active {
			t.Fatalf("expected only first tab active, but %q is also active", tab.Label)
		}
	}
}

func TestBuildTabs_DefaultActiveByLabel(t *testing.T) {
	o := DefaultOptions()
	WithTabs(true)(o)
	WithTabsDefaultActive("Unmanaged")(o)
	d := buildTabs(makeTabsAnalysis(), o)
	var found bool
	for _, tab := range d.Tabs {
		if tab.Label == "Unmanaged" && tab.Active {
			found = true
		}
	}
	if !found {
		t.Fatal("expected 'Unmanaged' tab to be active")
	}
}

func TestBuildTabs_EmptyAnalysis(t *testing.T) {
	o := DefaultOptions()
	WithTabs(true)(o)
	d := buildTabs(&Analysis{}, o)
	if len(d.Tabs) != 0 {
		t.Fatalf("expected 0 tabs for empty analysis, got %d", len(d.Tabs))
	}
}

func TestBuildTabs_ZeroCountTabsOmitted(t *testing.T) {
	o := DefaultOptions()
	WithTabs(true)(o)
	a := &Analysis{
		Managed: []Resource{{ID: "r1", Type: "aws_s3_bucket"}},
	}
	d := buildTabs(a, o)
	if len(d.Tabs) != 1 {
		t.Fatalf("expected 1 tab, got %d", len(d.Tabs))
	}
	if d.Tabs[0].Label != "Managed" {
		t.Fatalf("expected 'Managed' tab, got %q", d.Tabs[0].Label)
	}
}
