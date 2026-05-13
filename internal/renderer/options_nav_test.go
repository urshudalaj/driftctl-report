package renderer

import (
	"testing"
)

func TestWithNav_Enables(t *testing.T) {
	o := DefaultOptions()
	WithNav(true)(o)
	if !o.NavEnabled {
		t.Fatal("expected NavEnabled to be true")
	}
}

func TestWithNav_Disables(t *testing.T) {
	o := DefaultOptions()
	o.NavEnabled = true
	WithNav(false)(o)
	if o.NavEnabled {
		t.Fatal("expected NavEnabled to be false")
	}
}

func TestWithNav_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.NavEnabled {
		t.Fatal("expected NavEnabled default to be false")
	}
}

func TestWithNavBrand_SetsText(t *testing.T) {
	o := DefaultOptions()
	WithNavBrand("My Report")(o)
	if o.NavBrand != "My Report" {
		t.Fatalf("expected NavBrand 'My Report', got %q", o.NavBrand)
	}
}

func TestWithNavBrand_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	o.NavBrand = "Existing"
	WithNavBrand("")(o)
	if o.NavBrand != "Existing" {
		t.Fatalf("expected NavBrand to remain 'Existing', got %q", o.NavBrand)
	}
}

func TestWithNavBrand_DefaultIsEmpty(t *testing.T) {
	o := DefaultOptions()
	if o.NavBrand != "" {
		t.Fatalf("expected NavBrand default to be empty, got %q", o.NavBrand)
	}
}

func TestWithNavSticky_True(t *testing.T) {
	o := DefaultOptions()
	WithNavSticky(true)(o)
	if !o.NavSticky {
		t.Fatal("expected NavSticky to be true")
	}
}

func TestWithNavSticky_False(t *testing.T) {
	o := DefaultOptions()
	o.NavSticky = true
	WithNavSticky(false)(o)
	if o.NavSticky {
		t.Fatal("expected NavSticky to be false")
	}
}
