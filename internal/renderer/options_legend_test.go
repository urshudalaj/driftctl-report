package renderer

import (
	"testing"
)

func TestWithLegend_Enables(t *testing.T) {
	o := DefaultOptions()
	WithLegend(true)(o)
	if !o.LegendEnabled {
		t.Fatal("expected LegendEnabled to be true")
	}
}

func TestWithLegend_Disables(t *testing.T) {
	o := DefaultOptions()
	o.LegendEnabled = true
	WithLegend(false)(o)
	if o.LegendEnabled {
		t.Fatal("expected LegendEnabled to be false")
	}
}

func TestWithLegend_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.LegendEnabled {
		t.Fatal("expected LegendEnabled default to be false")
	}
}

func TestWithLegendPosition_Top(t *testing.T) {
	o := DefaultOptions()
	WithLegendPosition("top")(o)
	if o.LegendPosition != "top" {
		t.Fatalf("expected LegendPosition 'top', got %q", o.LegendPosition)
	}
}

func TestWithLegendPosition_Bottom(t *testing.T) {
	o := DefaultOptions()
	WithLegendPosition("bottom")(o)
	if o.LegendPosition != "bottom" {
		t.Fatalf("expected LegendPosition 'bottom', got %q", o.LegendPosition)
	}
}

func TestWithLegendPosition_Unknown_Ignored(t *testing.T) {
	o := DefaultOptions()
	o.LegendPosition = "bottom"
	WithLegendPosition("left")(o)
	if o.LegendPosition != "bottom" {
		t.Fatalf("expected LegendPosition to remain 'bottom', got %q", o.LegendPosition)
	}
}

func TestWithLegendPosition_Empty_Ignored(t *testing.T) {
	o := DefaultOptions()
	o.LegendPosition = "top"
	WithLegendPosition("")(o)
	if o.LegendPosition != "top" {
		t.Fatalf("expected LegendPosition to remain 'top', got %q", o.LegendPosition)
	}
}

func TestWithLegendCompact_Enables(t *testing.T) {
	o := DefaultOptions()
	WithLegendCompact(true)(o)
	if !o.LegendCompact {
		t.Fatal("expected LegendCompact to be true")
	}
}

func TestWithLegendCompact_Disables(t *testing.T) {
	o := DefaultOptions()
	o.LegendCompact = true
	WithLegendCompact(false)(o)
	if o.LegendCompact {
		t.Fatal("expected LegendCompact to be false")
	}
}

func TestWithLegendCompact_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.LegendCompact {
		t.Fatal("expected LegendCompact default to be false")
	}
}
