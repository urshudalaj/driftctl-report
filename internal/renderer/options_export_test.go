package renderer

import "testing"

func TestWithOutputFormat_HTML(t *testing.T) {
	o := DefaultOptions()
	WithOutputFormat("html")(o)
	if o.OutputFormat != "html" {
		t.Fatalf("expected html, got %s", o.OutputFormat)
	}
}

func TestWithOutputFormat_JSON(t *testing.T) {
	o := DefaultOptions()
	WithOutputFormat("json")(o)
	if o.OutputFormat != "json" {
		t.Fatalf("expected json, got %s", o.OutputFormat)
	}
}

func TestWithOutputFormat_Unknown_Ignored(t *testing.T) {
	o := DefaultOptions()
	defaultFmt := o.OutputFormat
	WithOutputFormat("xml")(o)
	if o.OutputFormat != defaultFmt {
		t.Fatalf("expected format to remain %q, got %q", defaultFmt, o.OutputFormat)
	}
}

func TestWithOutputFormat_Empty_Ignored(t *testing.T) {
	o := DefaultOptions()
	defaultFmt := o.OutputFormat
	WithOutputFormat("")(o)
	if o.OutputFormat != defaultFmt {
		t.Fatalf("expected format to remain %q, got %q", defaultFmt, o.OutputFormat)
	}
}

func TestWithFilename_SetsName(t *testing.T) {
	o := DefaultOptions()
	WithFilename("drift-report")(o)
	if o.Filename != "drift-report" {
		t.Fatalf("expected drift-report, got %s", o.Filename)
	}
}

func TestWithFilename_Empty_Ignored(t *testing.T) {
	o := DefaultOptions()
	original := o.Filename
	WithFilename("")(o)
	if o.Filename != original {
		t.Fatalf("expected filename to remain %q, got %q", original, o.Filename)
	}
}

func TestWithEmbedAssets_True(t *testing.T) {
	o := DefaultOptions()
	WithEmbedAssets(true)(o)
	if !o.EmbedAssets {
		t.Fatal("expected EmbedAssets to be true")
	}
}

func TestWithEmbedAssets_False(t *testing.T) {
	o := DefaultOptions()
	WithEmbedAssets(true)(o)
	WithEmbedAssets(false)(o)
	if o.EmbedAssets {
		t.Fatal("expected EmbedAssets to be false")
	}
}

func TestWithEmbedAssets_Default(t *testing.T) {
	o := DefaultOptions()
	// Default should be true so reports are self-contained.
	if !o.EmbedAssets {
		t.Fatal("expected EmbedAssets default to be true")
	}
}
