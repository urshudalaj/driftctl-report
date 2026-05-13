package renderer

import "testing"

func TestWithTooltip_Enables(t *testing.T) {
	o := DefaultOptions()
	WithTooltip(true)(o)
	if !o.Tooltip {
		t.Fatal("expected Tooltip to be true")
	}
}

func TestWithTooltip_Disables(t *testing.T) {
	o := DefaultOptions()
	o.Tooltip = true
	WithTooltip(false)(o)
	if o.Tooltip {
		t.Fatal("expected Tooltip to be false")
	}
}

func TestWithTooltip_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.Tooltip {
		t.Fatal("expected Tooltip default to be false")
	}
}

func TestWithTooltipPlacement_Top(t *testing.T) {
	o := DefaultOptions()
	WithTooltipPlacement("top")(o)
	if o.TooltipPlacement != "top" {
		t.Fatalf("expected 'top', got %q", o.TooltipPlacement)
	}
}

func TestWithTooltipPlacement_Bottom(t *testing.T) {
	o := DefaultOptions()
	WithTooltipPlacement("bottom")(o)
	if o.TooltipPlacement != "bottom" {
		t.Fatalf("expected 'bottom', got %q", o.TooltipPlacement)
	}
}

func TestWithTooltipPlacement_Unknown_Ignored(t *testing.T) {
	o := DefaultOptions()
	o.TooltipPlacement = "top"
	WithTooltipPlacement("diagonal")(o)
	if o.TooltipPlacement != "top" {
		t.Fatalf("expected placement unchanged, got %q", o.TooltipPlacement)
	}
}

func TestWithTooltipMaxWidth_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithTooltipMaxWidth(320)(o)
	if o.TooltipMaxWidth != 320 {
		t.Fatalf("expected 320, got %d", o.TooltipMaxWidth)
	}
}

func TestWithTooltipMaxWidth_ZeroIgnored(t *testing.T) {
	o := DefaultOptions()
	o.TooltipMaxWidth = 200
	WithTooltipMaxWidth(0)(o)
	if o.TooltipMaxWidth != 200 {
		t.Fatalf("expected 200 unchanged, got %d", o.TooltipMaxWidth)
	}
}

func TestWithTooltipMaxWidth_NegativeIgnored(t *testing.T) {
	o := DefaultOptions()
	o.TooltipMaxWidth = 200
	WithTooltipMaxWidth(-50)(o)
	if o.TooltipMaxWidth != 200 {
		t.Fatalf("expected 200 unchanged, got %d", o.TooltipMaxWidth)
	}
}
