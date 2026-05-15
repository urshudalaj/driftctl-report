package renderer

import "testing"

func TestBuildHeading_Disabled_NoText(t *testing.T) {
	o := DefaultOptions()
	h := buildHeading(o)
	if h.Enabled {
		t.Fatal("expected heading to be disabled when no text is set")
	}
}

func TestBuildHeading_Disabled_ExplicitlyHidden(t *testing.T) {
	o := DefaultOptions()
	WithHeading("My Report")(&o)
	WithHeadingVisible(false)(&o)
	h := buildHeading(o)
	if h.Enabled {
		t.Fatal("expected heading to be disabled when visible=false")
	}
}

func TestBuildHeading_DefaultLevel(t *testing.T) {
	o := DefaultOptions()
	WithHeading("Drift Report")(&o)
	WithHeadingVisible(true)(&o)
	h := buildHeading(o)
	if !h.Enabled {
		t.Fatal("expected heading to be enabled")
	}
	if h.Level != 1 {
		t.Fatalf("expected default level 1, got %d", h.Level)
	}
	if h.Tag != "h1" {
		t.Fatalf("expected tag h1, got %s", h.Tag)
	}
}

func TestBuildHeading_CustomLevel(t *testing.T) {
	o := DefaultOptions()
	WithHeading("Section")(&o)
	WithHeadingVisible(true)(&o)
	WithHeadingLevel(3)(&o)
	h := buildHeading(o)
	if h.Level != 3 {
		t.Fatalf("expected level 3, got %d", h.Level)
	}
	if h.Tag != "h3" {
		t.Fatalf("expected tag h3, got %s", h.Tag)
	}
}

func TestBuildHeading_InvalidLevelFallsBackToOne(t *testing.T) {
	o := DefaultOptions()
	WithHeading("Title")(&o)
	WithHeadingVisible(true)(&o)
	// level 0 is invalid and should be ignored by the option
	WithHeadingLevel(0)(&o)
	h := buildHeading(o)
	if h.Level != 1 {
		t.Fatalf("expected fallback level 1, got %d", h.Level)
	}
}

func TestBuildHeading_TextPreserved(t *testing.T) {
	o := DefaultOptions()
	WithHeading("Infrastructure Drift")(&o)
	WithHeadingVisible(true)(&o)
	h := buildHeading(o)
	if h.Text != "Infrastructure Drift" {
		t.Fatalf("expected text 'Infrastructure Drift', got %q", h.Text)
	}
}

func TestWithHeading_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithHeading("")(&o)
	if o.HeadingText != "" {
		t.Fatal("empty heading should not overwrite existing value")
	}
}

func TestWithHeadingLevel_OutOfRangeIgnored(t *testing.T) {
	o := DefaultOptions()
	WithHeadingLevel(7)(&o)
	if o.HeadingLevel != 0 {
		t.Fatalf("out-of-range level should be ignored, got %d", o.HeadingLevel)
	}
}
