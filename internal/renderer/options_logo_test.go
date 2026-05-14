package renderer

import "testing"

func TestWithLogo_SetsURL(t *testing.T) {
	o := DefaultOptions()
	WithLogo("https://example.com/logo.png")(o)
	if o.LogoURL != "https://example.com/logo.png" {
		t.Fatalf("expected logo URL to be set, got %q", o.LogoURL)
	}
}

func TestWithLogo_EmptyClears(t *testing.T) {
	o := DefaultOptions()
	WithLogo("https://example.com/logo.png")(o)
	WithLogo("")(o)
	if o.LogoURL != "" {
		t.Fatalf("expected logo URL to be cleared, got %q", o.LogoURL)
	}
}

func TestWithLogo_DefaultIsEmpty(t *testing.T) {
	o := DefaultOptions()
	if o.LogoURL != "" {
		t.Fatalf("expected default logo URL to be empty, got %q", o.LogoURL)
	}
}

func TestWithLogoAlt_SetsAlt(t *testing.T) {
	o := DefaultOptions()
	WithLogoAlt("Company Logo")(o)
	if o.LogoAlt != "Company Logo" {
		t.Fatalf("expected logo alt to be set, got %q", o.LogoAlt)
	}
}

func TestWithLogoAlt_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithLogoAlt("Initial")(o)
	WithLogoAlt("")(o)
	if o.LogoAlt != "Initial" {
		t.Fatalf("expected logo alt to be preserved, got %q", o.LogoAlt)
	}
}

func TestWithLogoHeight_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithLogoHeight(48)(o)
	if o.LogoHeight != 48 {
		t.Fatalf("expected logo height 48, got %d", o.LogoHeight)
	}
}

func TestWithLogoHeight_ZeroIgnored(t *testing.T) {
	o := DefaultOptions()
	WithLogoHeight(32)(o)
	WithLogoHeight(0)(o)
	if o.LogoHeight != 32 {
		t.Fatalf("expected logo height to be preserved, got %d", o.LogoHeight)
	}
}

func TestWithLogoHeight_NegativeIgnored(t *testing.T) {
	o := DefaultOptions()
	WithLogoHeight(32)(o)
	WithLogoHeight(-10)(o)
	if o.LogoHeight != 32 {
		t.Fatalf("expected logo height to be preserved, got %d", o.LogoHeight)
	}
}
