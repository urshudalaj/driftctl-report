package renderer

import "testing"

func TestBuildRTL_Disabled(t *testing.T) {
	o := DefaultOptions()
	cfg := buildRTL(o)

	if cfg.Enabled {
		t.Fatal("expected Enabled=false when RTL option not set")
	}
	if cfg.Dir != "ltr" {
		t.Fatalf("expected Dir=ltr, got %q", cfg.Dir)
	}
	if cfg.Lang != "en" {
		t.Fatalf("expected Lang=en, got %q", cfg.Lang)
	}
}

func TestBuildRTL_EnabledDefaults(t *testing.T) {
	o := DefaultOptions()
	WithRTL(true)(&o)
	cfg := buildRTL(o)

	if !cfg.Enabled {
		t.Fatal("expected Enabled=true")
	}
	if cfg.Dir != "rtl" {
		t.Fatalf("expected default Dir=rtl, got %q", cfg.Dir)
	}
	if cfg.Lang != "ar" {
		t.Fatalf("expected default Lang=ar, got %q", cfg.Lang)
	}
}

func TestBuildRTL_CustomLang(t *testing.T) {
	o := DefaultOptions()
	WithRTL(true)(&o)
	WithRTLLang("he")(&o)
	cfg := buildRTL(o)

	if cfg.Lang != "he" {
		t.Fatalf("expected Lang=he, got %q", cfg.Lang)
	}
}

func TestBuildRTL_CustomDir(t *testing.T) {
	o := DefaultOptions()
	WithRTL(true)(&o)
	WithRTLDir("ltr")(&o)
	cfg := buildRTL(o)

	if cfg.Dir != "ltr" {
		t.Fatalf("expected Dir=ltr, got %q", cfg.Dir)
	}
}

func TestWithRTLDir_InvalidIgnored(t *testing.T) {
	o := DefaultOptions()
	WithRTL(true)(&o)
	WithRTLDir("auto")(&o) // invalid value
	cfg := buildRTL(o)

	// Should fall back to default "rtl" since the invalid value was ignored
	if cfg.Dir != "rtl" {
		t.Fatalf("expected Dir=rtl after invalid override, got %q", cfg.Dir)
	}
}

func TestWithRTLLang_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithRTL(true)(&o)
	WithRTLLang("fa")(&o)
	WithRTLLang("")(&o) // empty should be ignored

	if o.RTLLang != "fa" {
		t.Fatalf("expected RTLLang=fa after empty override, got %q", o.RTLLang)
	}
}

func TestWithRTL_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.RTL {
		t.Fatal("expected RTL default to be false")
	}
}
