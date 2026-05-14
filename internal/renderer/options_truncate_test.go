package renderer

import (
	"testing"
)

func TestWithTruncateIDs_Enables(t *testing.T) {
	o := DefaultOptions()
	WithTruncateIDs(true)(o)
	if !o.TruncateIDs {
		t.Fatal("expected TruncateIDs to be true")
	}
}

func TestWithTruncateIDs_Disables(t *testing.T) {
	o := DefaultOptions()
	o.TruncateIDs = true
	WithTruncateIDs(false)(o)
	if o.TruncateIDs {
		t.Fatal("expected TruncateIDs to be false")
	}
}

func TestWithTruncateIDs_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.TruncateIDs {
		t.Fatal("expected TruncateIDs default to be false")
	}
}

func TestWithTruncateIDsLength_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithTruncateIDsLength(32)(o)
	if o.TruncateIDsLength != 32 {
		t.Fatalf("expected 32, got %d", o.TruncateIDsLength)
	}
}

func TestWithTruncateIDsLength_ZeroIgnored(t *testing.T) {
	o := DefaultOptions()
	o.TruncateIDsLength = 20
	WithTruncateIDsLength(0)(o)
	if o.TruncateIDsLength != 20 {
		t.Fatalf("expected 20, got %d", o.TruncateIDsLength)
	}
}

func TestWithTruncateIDsLength_NegativeIgnored(t *testing.T) {
	o := DefaultOptions()
	o.TruncateIDsLength = 20
	WithTruncateIDsLength(-5)(o)
	if o.TruncateIDsLength != 20 {
		t.Fatalf("expected 20, got %d", o.TruncateIDsLength)
	}
}

func TestWithTruncateIDsSuffix_SetsSuffix(t *testing.T) {
	o := DefaultOptions()
	WithTruncateIDsSuffix("...")(o)
	if o.TruncateIDsSuffix != "..." {
		t.Fatalf("expected '...', got %q", o.TruncateIDsSuffix)
	}
}

func TestWithTruncateIDsSuffix_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	o.TruncateIDsSuffix = "…"
	WithTruncateIDsSuffix("")(o)
	if o.TruncateIDsSuffix != "…" {
		t.Fatalf("expected '…', got %q", o.TruncateIDsSuffix)
	}
}
