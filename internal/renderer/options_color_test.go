package renderer

import "testing"

func TestWithColorScheme_Default(t *testing.T) {
	o := DefaultOptions()
	WithColorScheme("default")(o)
	if o.ColorScheme != "default" {
		t.Fatalf("expected 'default', got %q", o.ColorScheme)
	}
}

func TestWithColorScheme_Accessible(t *testing.T) {
	o := DefaultOptions()
	WithColorScheme("accessible")(o)
	if o.ColorScheme != "accessible" {
		t.Fatalf("expected 'accessible', got %q", o.ColorScheme)
	}
}

func TestWithColorScheme_Monochrome(t *testing.T) {
	o := DefaultOptions()
	WithColorScheme("monochrome")(o)
	if o.ColorScheme != "monochrome" {
		t.Fatalf("expected 'monochrome', got %q", o.ColorScheme)
	}
}

func TestWithColorScheme_Unknown_Ignored(t *testing.T) {
	o := DefaultOptions()
	prev := o.ColorScheme
	WithColorScheme("rainbow")(o)
	if o.ColorScheme != prev {
		t.Fatalf("expected scheme to remain %q, got %q", prev, o.ColorScheme)
	}
}

func TestWithColorScheme_Empty_Ignored(t *testing.T) {
	o := DefaultOptions()
	prev := o.ColorScheme
	WithColorScheme("")(o)
	if o.ColorScheme != prev {
		t.Fatalf("expected scheme to remain %q, got %q", prev, o.ColorScheme)
	}
}

func TestWithColorManaged_SetsClass(t *testing.T) {
	o := DefaultOptions()
	WithColorManaged("text-green-600")(o)
	if o.ColorManaged != "text-green-600" {
		t.Fatalf("expected 'text-green-600', got %q", o.ColorManaged)
	}
}

func TestWithColorManaged_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithColorManaged("text-green-600")(o)
	WithColorManaged("")(o)
	if o.ColorManaged != "text-green-600" {
		t.Fatalf("expected class to remain 'text-green-600', got %q", o.ColorManaged)
	}
}

func TestWithColorUnmanaged_SetsClass(t *testing.T) {
	o := DefaultOptions()
	WithColorUnmanaged("text-yellow-500")(o)
	if o.ColorUnmanaged != "text-yellow-500" {
		t.Fatalf("expected 'text-yellow-500', got %q", o.ColorUnmanaged)
	}
}

func TestWithColorDeleted_SetsClass(t *testing.T) {
	o := DefaultOptions()
	WithColorDeleted("text-red-600")(o)
	if o.ColorDeleted != "text-red-600" {
		t.Fatalf("expected 'text-red-600', got %q", o.ColorDeleted)
	}
}

func TestWithColorDeleted_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithColorDeleted("text-red-600")(o)
	WithColorDeleted("")(o)
	if o.ColorDeleted != "text-red-600" {
		t.Fatalf("expected class to remain 'text-red-600', got %q", o.ColorDeleted)
	}
}
