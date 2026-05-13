package renderer

import (
	"testing"
)

func TestWithPanel_Enables(t *testing.T) {
	o := DefaultOptions()
	WithPanel(true)(o)
	if !o.PanelEnabled {
		t.Fatal("expected PanelEnabled to be true")
	}
}

func TestWithPanel_Disables(t *testing.T) {
	o := DefaultOptions()
	o.PanelEnabled = true
	WithPanel(false)(o)
	if o.PanelEnabled {
		t.Fatal("expected PanelEnabled to be false")
	}
}

func TestWithPanel_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.PanelEnabled {
		t.Fatal("expected PanelEnabled default to be false")
	}
}

func TestWithPanelDefaultOpen_Enables(t *testing.T) {
	o := DefaultOptions()
	WithPanelDefaultOpen(true)(o)
	if !o.PanelDefaultOpen {
		t.Fatal("expected PanelDefaultOpen to be true")
	}
}

func TestWithPanelDefaultOpen_Disables(t *testing.T) {
	o := DefaultOptions()
	o.PanelDefaultOpen = true
	WithPanelDefaultOpen(false)(o)
	if o.PanelDefaultOpen {
		t.Fatal("expected PanelDefaultOpen to be false")
	}
}

func TestWithPanelMaxItems_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithPanelMaxItems(20)(o)
	if o.PanelMaxItems != 20 {
		t.Fatalf("expected PanelMaxItems=20, got %d", o.PanelMaxItems)
	}
}

func TestWithPanelMaxItems_ZeroMeansUnlimited(t *testing.T) {
	o := DefaultOptions()
	WithPanelMaxItems(0)(o)
	if o.PanelMaxItems != 0 {
		t.Fatalf("expected PanelMaxItems=0, got %d", o.PanelMaxItems)
	}
}

func TestWithPanelMaxItems_NegativeClampedToZero(t *testing.T) {
	o := DefaultOptions()
	WithPanelMaxItems(-5)(o)
	if o.PanelMaxItems != 0 {
		t.Fatalf("expected PanelMaxItems clamped to 0, got %d", o.PanelMaxItems)
	}
}
