package renderer

import "testing"

func TestWithWatermark_SetsText(t *testing.T) {
	o := DefaultOptions()
	WithWatermark("Confidential")(o)
	if o.WatermarkText != "Confidential" {
		t.Errorf("expected WatermarkText=Confidential, got %q", o.WatermarkText)
	}
}

func TestWithWatermark_EmptyClears(t *testing.T) {
	o := DefaultOptions()
	WithWatermark("Draft")(o)
	WithWatermark("")(o)
	if o.WatermarkText != "" {
		t.Errorf("expected WatermarkText to be cleared, got %q", o.WatermarkText)
	}
}

func TestWithWatermark_DefaultIsEmpty(t *testing.T) {
	o := DefaultOptions()
	if o.WatermarkText != "" {
		t.Errorf("expected default WatermarkText to be empty, got %q", o.WatermarkText)
	}
}

func TestWithWatermarkURL_SetsURL(t *testing.T) {
	o := DefaultOptions()
	WithWatermarkURL("https://example.com")(o)
	if o.WatermarkURL != "https://example.com" {
		t.Errorf("expected WatermarkURL=https://example.com, got %q", o.WatermarkURL)
	}
}

func TestWithWatermarkURL_DefaultIsEmpty(t *testing.T) {
	o := DefaultOptions()
	if o.WatermarkURL != "" {
		t.Errorf("expected default WatermarkURL to be empty, got %q", o.WatermarkURL)
	}
}

func TestWithWatermarkPosition_Valid(t *testing.T) {
	positions := []string{"top-left", "top-right", "bottom-left", "bottom-right"}
	for _, pos := range positions {
		o := DefaultOptions()
		WithWatermarkPosition(pos)(o)
		if o.WatermarkPosition != pos {
			t.Errorf("expected WatermarkPosition=%q, got %q", pos, o.WatermarkPosition)
		}
	}
}

func TestWithWatermarkPosition_UnknownIgnored(t *testing.T) {
	o := DefaultOptions()
	defaultPos := o.WatermarkPosition
	WithWatermarkPosition("center")(o)
	if o.WatermarkPosition != defaultPos {
		t.Errorf("expected WatermarkPosition to remain %q, got %q", defaultPos, o.WatermarkPosition)
	}
}

func TestWithWatermarkPosition_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	defaultPos := o.WatermarkPosition
	WithWatermarkPosition("")(o)
	if o.WatermarkPosition != defaultPos {
		t.Errorf("expected WatermarkPosition to remain %q, got %q", defaultPos, o.WatermarkPosition)
	}
}

func TestWithWatermarkPosition_DefaultIsBottomRight(t *testing.T) {
	o := DefaultOptions()
	if o.WatermarkPosition != "bottom-right" {
		t.Errorf("expected default WatermarkPosition=bottom-right, got %q", o.WatermarkPosition)
	}
}
