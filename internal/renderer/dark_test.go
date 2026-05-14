package renderer

import (
	"strings"
	"testing"
)

func TestBuildDarkStyles_Disabled(t *testing.T) {
	o := DefaultOptions()
	ds := buildDarkStyles(o)
	if ds.Enabled {
		t.Fatal("expected Enabled=false when dark mode is off")
	}
	if ds.CSS != "" {
		t.Fatal("expected empty CSS when disabled")
	}
}

func TestBuildDarkStyles_DarkModeOnly(t *testing.T) {
	o := DefaultOptions()
	WithDarkMode(true)(&o)
	ds := buildDarkStyles(o)
	if !ds.Enabled {
		t.Fatal("expected Enabled=true")
	}
	if !strings.Contains(ds.CSS, "--bg") {
		t.Error("expected CSS to contain --bg variable")
	}
	if ds.ToggleScript != "" {
		t.Error("expected no toggle script when toggle is disabled")
	}
}

func TestBuildDarkStyles_WithToggle(t *testing.T) {
	o := DefaultOptions()
	WithDarkModeToggle(true)(&o)
	ds := buildDarkStyles(o)
	if !ds.Enabled {
		t.Fatal("expected Enabled=true when toggle is on")
	}
	if !strings.Contains(ds.ToggleScript, "theme-toggle") {
		t.Error("expected toggle script to reference theme-toggle element")
	}
}

func TestBuildDarkStyles_DefaultModeLight(t *testing.T) {
	o := DefaultOptions()
	WithDarkModeToggle(true)(&o)
	ds := buildDarkStyles(o)
	if ds.DefaultMode != "light" {
		t.Errorf("expected default mode 'light', got %q", ds.DefaultMode)
	}
}

func TestBuildDarkStyles_DefaultModeDark(t *testing.T) {
	o := DefaultOptions()
	WithDarkModeToggle(true)(&o)
	WithDarkModeDefault("dark")(&o)
	ds := buildDarkStyles(o)
	if ds.DefaultMode != "dark" {
		t.Errorf("expected default mode 'dark', got %q", ds.DefaultMode)
	}
	if !strings.Contains(ds.ToggleScript, "'dark'") {
		t.Error("expected toggle script to embed dark as default")
	}
}

func TestBuildDarkStyles_InvalidDefaultIgnored(t *testing.T) {
	o := DefaultOptions()
	WithDarkModeDefault("invalid")(&o)
	if o.DarkModeDefault != "" {
		t.Errorf("expected empty DarkModeDefault for invalid input, got %q", o.DarkModeDefault)
	}
}
