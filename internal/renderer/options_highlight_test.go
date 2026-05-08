package renderer

import (
	"testing"
)

func TestWithHighlight_Enables(t *testing.T) {
	o := DefaultOptions()
	WithHighlight(true)(o)
	if !o.Highlight {
		t.Error("expected Highlight to be true")
	}
}

func TestWithHighlight_Disables(t *testing.T) {
	o := DefaultOptions()
	o.Highlight = true
	WithHighlight(false)(o)
	if o.Highlight {
		t.Error("expected Highlight to be false")
	}
}

func TestWithHighlight_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.Highlight {
		t.Error("expected default Highlight to be false")
	}
}

func TestWithHighlightMaxLength_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithHighlightMaxLength(120)(o)
	if o.HighlightMaxLength != 120 {
		t.Errorf("expected HighlightMaxLength 120, got %d", o.HighlightMaxLength)
	}
}

func TestWithHighlightMaxLength_ZeroMeansUnlimited(t *testing.T) {
	o := DefaultOptions()
	WithHighlightMaxLength(0)(o)
	if o.HighlightMaxLength != 0 {
		t.Errorf("expected HighlightMaxLength 0, got %d", o.HighlightMaxLength)
	}
}

func TestWithHighlightMaxLength_NegativeClampedToZero(t *testing.T) {
	o := DefaultOptions()
	WithHighlightMaxLength(-5)(o)
	if o.HighlightMaxLength != 0 {
		t.Errorf("expected HighlightMaxLength clamped to 0, got %d", o.HighlightMaxLength)
	}
}

func TestWithHighlightMaxLength_DefaultValue(t *testing.T) {
	o := DefaultOptions()
	if o.HighlightMaxLength != 200 {
		t.Errorf("expected default HighlightMaxLength 200, got %d", o.HighlightMaxLength)
	}
}
