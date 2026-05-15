package renderer

import (
	"strings"
	"testing"
)

func TestBuildEmbed_AllDisabled(t *testing.T) {
	o := DefaultOptions()
	ed := buildEmbed(o)
	if ed.Enabled {
		t.Fatal("expected Enabled=false when all embed options are off")
	}
	if ed.CSS != "" || ed.JS != "" || ed.FontFaces != "" {
		t.Fatal("expected all fields empty when disabled")
	}
}

func TestBuildEmbed_CSSOnly(t *testing.T) {
	o := DefaultOptions()
	WithEmbedCSS(true)(&o)
	ed := buildEmbed(o)
	if !ed.Enabled {
		t.Fatal("expected Enabled=true")
	}
	if ed.CSS == "" {
		t.Fatal("expected non-empty CSS")
	}
	if ed.JS != "" {
		t.Fatal("expected empty JS")
	}
}

func TestBuildEmbed_JSOnly(t *testing.T) {
	o := DefaultOptions()
	WithEmbedJS(true)(&o)
	ed := buildEmbed(o)
	if !ed.Enabled {
		t.Fatal("expected Enabled=true")
	}
	if ed.JS == "" {
		t.Fatal("expected non-empty JS")
	}
	if ed.CSS != "" {
		t.Fatal("expected empty CSS")
	}
}

func TestBuildEmbed_FontsOnly(t *testing.T) {
	o := DefaultOptions()
	WithEmbedFonts(true)(&o)
	ed := buildEmbed(o)
	if !ed.Enabled {
		t.Fatal("expected Enabled=true")
	}
	// font faces may be empty in default implementation; just check no panic
	_ = ed.FontFaces
}

func TestBuildEmbed_CSSContainsColorVars(t *testing.T) {
	o := DefaultOptions()
	WithEmbedCSS(true)(&o)
	WithColorManaged("#00ff00")(&o)
	ed := buildEmbed(o)
	if !strings.Contains(ed.CSS, "#00ff00") {
		t.Errorf("expected CSS to contain managed color, got: %s", ed.CSS)
	}
}

func TestBuildEmbed_JSContainsDOMListener(t *testing.T) {
	o := DefaultOptions()
	WithEmbedJS(true)(&o)
	ed := buildEmbed(o)
	if !strings.Contains(ed.JS, "DOMContentLoaded") {
		t.Errorf("expected JS to contain DOMContentLoaded, got: %s", ed.JS)
	}
}

func TestWithEmbedCSS_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.EmbedCSS {
		t.Fatal("expected EmbedCSS default to be false")
	}
}

func TestWithEmbedJS_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.EmbedJS {
		t.Fatal("expected EmbedJS default to be false")
	}
}

func TestWithEmbedFonts_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.EmbedFonts {
		t.Fatal("expected EmbedFonts default to be false")
	}
}
