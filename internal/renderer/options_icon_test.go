package renderer

import (
	"testing"
)

func TestWithIcons_Enables(t *testing.T) {
	o := DefaultOptions()
	WithIcons(true)(o)
	if !o.Icons {
		t.Fatal("expected Icons to be true")
	}
}

func TestWithIcons_Disables(t *testing.T) {
	o := DefaultOptions()
	o.Icons = true
	WithIcons(false)(o)
	if o.Icons {
		t.Fatal("expected Icons to be false")
	}
}

func TestWithIcons_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.Icons {
		t.Fatal("expected Icons default to be false")
	}
}

func TestWithIconSet_Unicode(t *testing.T) {
	o := DefaultOptions()
	WithIconSet("unicode")(o)
	if o.IconSet != "unicode" {
		t.Fatalf("expected IconSet=unicode, got %q", o.IconSet)
	}
}

func TestWithIconSet_Emoji(t *testing.T) {
	o := DefaultOptions()
	WithIconSet("emoji")(o)
	if o.IconSet != "emoji" {
		t.Fatalf("expected IconSet=emoji, got %q", o.IconSet)
	}
}

func TestWithIconSet_SVG(t *testing.T) {
	o := DefaultOptions()
	WithIconSet("svg")(o)
	if o.IconSet != "svg" {
		t.Fatalf("expected IconSet=svg, got %q", o.IconSet)
	}
}

func TestWithIconSet_Unknown_Ignored(t *testing.T) {
	o := DefaultOptions()
	o.IconSet = "unicode"
	WithIconSet("neon")(o)
	if o.IconSet != "unicode" {
		t.Fatalf("expected IconSet to remain unicode, got %q", o.IconSet)
	}
}

func TestWithIconSize_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithIconSize(24)(o)
	if o.IconSize != 24 {
		t.Fatalf("expected IconSize=24, got %d", o.IconSize)
	}
}

func TestWithIconSize_ZeroIgnored(t *testing.T) {
	o := DefaultOptions()
	o.IconSize = 16
	WithIconSize(0)(o)
	if o.IconSize != 16 {
		t.Fatalf("expected IconSize to remain 16, got %d", o.IconSize)
	}
}

func TestWithIconSize_NegativeIgnored(t *testing.T) {
	o := DefaultOptions()
	o.IconSize = 16
	WithIconSize(-5)(o)
	if o.IconSize != 16 {
		t.Fatalf("expected IconSize to remain 16, got %d", o.IconSize)
	}
}
