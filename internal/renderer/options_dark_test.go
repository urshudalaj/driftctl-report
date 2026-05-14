package renderer

import "testing"

func TestWithDarkMode_Enables(t *testing.T) {
	o := DefaultOptions()
	WithDarkMode(true)(&o)
	if !o.DarkMode {
		t.Fatal("expected DarkMode=true")
	}
}

func TestWithDarkMode_Disables(t *testing.T) {
	o := DefaultOptions()
	WithDarkMode(true)(&o)
	WithDarkMode(false)(&o)
	if o.DarkMode {
		t.Fatal("expected DarkMode=false after disabling")
	}
}

func TestWithDarkMode_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.DarkMode {
		t.Fatal("expected DarkMode to default to false")
	}
}

func TestWithDarkModeToggle_Enables(t *testing.T) {
	o := DefaultOptions()
	WithDarkModeToggle(true)(&o)
	if !o.DarkModeToggle {
		t.Fatal("expected DarkModeToggle=true")
	}
}

func TestWithDarkModeToggle_Disables(t *testing.T) {
	o := DefaultOptions()
	WithDarkModeToggle(true)(&o)
	WithDarkModeToggle(false)(&o)
	if o.DarkModeToggle {
		t.Fatal("expected DarkModeToggle=false")
	}
}

func TestWithDarkModeDefault_Dark(t *testing.T) {
	o := DefaultOptions()
	WithDarkModeDefault("dark")(&o)
	if o.DarkModeDefault != "dark" {
		t.Errorf("expected 'dark', got %q", o.DarkModeDefault)
	}
}

func TestWithDarkModeDefault_Light(t *testing.T) {
	o := DefaultOptions()
	WithDarkModeDefault("light")(&o)
	if o.DarkModeDefault != "light" {
		t.Errorf("expected 'light', got %q", o.DarkModeDefault)
	}
}

func TestWithDarkModeDefault_UnknownIgnored(t *testing.T) {
	o := DefaultOptions()
	WithDarkModeDefault("solarized")(&o)
	if o.DarkModeDefault != "" {
		t.Errorf("expected empty string for unknown mode, got %q", o.DarkModeDefault)
	}
}
