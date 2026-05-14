package renderer

import "testing"

func TestWithTabs_Enables(t *testing.T) {
	o := DefaultOptions()
	WithTabs(true)(o)
	if !o.TabsEnabled {
		t.Fatal("expected TabsEnabled to be true")
	}
}

func TestWithTabs_Disables(t *testing.T) {
	o := DefaultOptions()
	o.TabsEnabled = true
	WithTabs(false)(o)
	if o.TabsEnabled {
		t.Fatal("expected TabsEnabled to be false")
	}
}

func TestWithTabs_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.TabsEnabled {
		t.Fatal("expected TabsEnabled default to be false")
	}
}

func TestWithTabsDefaultActive_SetsLabel(t *testing.T) {
	o := DefaultOptions()
	WithTabsDefaultActive("Unmanaged")(o)
	if o.TabsDefaultActive != "Unmanaged" {
		t.Fatalf("expected 'Unmanaged', got %q", o.TabsDefaultActive)
	}
}

func TestWithTabsDefaultActive_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	o.TabsDefaultActive = "Managed"
	WithTabsDefaultActive("")(o)
	if o.TabsDefaultActive != "Managed" {
		t.Fatalf("expected label to remain 'Managed', got %q", o.TabsDefaultActive)
	}
}

func TestWithTabsPosition_Top(t *testing.T) {
	o := DefaultOptions()
	WithTabsPosition("top")(o)
	if o.TabsPosition != "top" {
		t.Fatalf("expected 'top', got %q", o.TabsPosition)
	}
}

func TestWithTabsPosition_Left(t *testing.T) {
	o := DefaultOptions()
	WithTabsPosition("left")(o)
	if o.TabsPosition != "left" {
		t.Fatalf("expected 'left', got %q", o.TabsPosition)
	}
}

func TestWithTabsPosition_Unknown_Ignored(t *testing.T) {
	o := DefaultOptions()
	o.TabsPosition = "top"
	WithTabsPosition("diagonal")(o)
	if o.TabsPosition != "top" {
		t.Fatalf("expected position to remain 'top', got %q", o.TabsPosition)
	}
}
