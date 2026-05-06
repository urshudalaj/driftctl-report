package renderer

import (
	"testing"
)

func TestDefaultTheme_HasExpectedColors(t *testing.T) {
	theme := defaultTheme()

	if theme.PrimaryColor == "" {
		t.Error("expected PrimaryColor to be set")
	}
	if theme.DriftedColor == "" {
		t.Error("expected DriftedColor to be set")
	}
	if theme.MissingColor == "" {
		t.Error("expected MissingColor to be set")
	}
	if theme.UnmanagedColor == "" {
		t.Error("expected UnmanagedColor to be set")
	}
	if theme.FontFamily == "" {
		t.Error("expected FontFamily to be set")
	}
}

func TestThemeByName_KnownTheme(t *testing.T) {
	for _, name := range []string{"default", "dark", "high-contrast"} {
		t.Run(name, func(t *testing.T) {
			theme, ok := ThemeByName(name)
			if !ok {
				t.Fatalf("expected theme %q to be found", name)
			}
			if theme.PrimaryColor == "" {
				t.Errorf("theme %q has empty PrimaryColor", name)
			}
		})
	}
}

func TestThemeByName_UnknownTheme(t *testing.T) {
	theme, ok := ThemeByName("nonexistent")
	if ok {
		t.Error("expected ok=false for unknown theme")
	}
	def := defaultTheme()
	if theme.PrimaryColor != def.PrimaryColor {
		t.Errorf("expected fallback to default theme, got PrimaryColor=%q", theme.PrimaryColor)
	}
}

func TestWithTheme_AppliesTheme(t *testing.T) {
	customTheme := Theme{
		PrimaryColor:   "#FF0000",
		DriftedColor:   "#00FF00",
		MissingColor:   "#0000FF",
		UnmanagedColor: "#FFFFFF",
		FontFamily:     "monospace",
	}

	opts := DefaultOptions()
	WithTheme(customTheme)(&opts)

	if opts.Theme.PrimaryColor != "#FF0000" {
		t.Errorf("expected PrimaryColor #FF0000, got %q", opts.Theme.PrimaryColor)
	}
	if opts.Theme.FontFamily != "monospace" {
		t.Errorf("expected FontFamily monospace, got %q", opts.Theme.FontFamily)
	}
}

func TestWithThemeName_KnownName(t *testing.T) {
	opts := DefaultOptions()
	err := WithThemeName("dark")(&opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	dark, _ := ThemeByName("dark")
	if opts.Theme.PrimaryColor != dark.PrimaryColor {
		t.Errorf("expected dark theme PrimaryColor %q, got %q", dark.PrimaryColor, opts.Theme.PrimaryColor)
	}
}

func TestWithThemeName_UnknownName(t *testing.T) {
	opts := DefaultOptions()
	err := WithThemeName("bogus")(&opts)
	if err == nil {
		t.Error("expected error for unknown theme name")
	}
}
