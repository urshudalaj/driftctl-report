package renderer

import "testing"

func TestWithTruncateIDs_Enables(t *testing.T) {
	o := DefaultOptions()
	WithTruncateIDs(true)(&o)
	if !o.TruncateIDs {
		t.Error("expected TruncateIDs=true")
	}
}

func TestWithTruncateIDs_Disables(t *testing.T) {
	o := DefaultOptions()
	WithTruncateIDs(true)(&o)
	WithTruncateIDs(false)(&o)
	if o.TruncateIDs {
		t.Error("expected TruncateIDs=false after disabling")
	}
}

func TestWithTruncateIDs_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.TruncateIDs {
		t.Error("expected TruncateIDs default to be false")
	}
}

func TestWithTruncateIDsLength_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithTruncateIDsLength(32)(&o)
	if o.TruncateIDsLen != 32 {
		t.Errorf("expected TruncateIDsLen=32, got %d", o.TruncateIDsLen)
	}
}

func TestWithTruncateIDsLength_ZeroIgnored(t *testing.T) {
	o := DefaultOptions()
	WithTruncateIDsLength(20)(&o)
	WithTruncateIDsLength(0)(&o)
	if o.TruncateIDsLen != 20 {
		t.Errorf("expected TruncateIDsLen=20 after zero ignored, got %d", o.TruncateIDsLen)
	}
}

func TestWithTruncateIDsLength_NegativeIgnored(t *testing.T) {
	o := DefaultOptions()
	WithTruncateIDsLength(15)(&o)
	WithTruncateIDsLength(-5)(&o)
	if o.TruncateIDsLen != 15 {
		t.Errorf("expected TruncateIDsLen=15 after negative ignored, got %d", o.TruncateIDsLen)
	}
}

func TestWithTruncateIDsSuffix_SetsValue(t *testing.T) {
	o := DefaultOptions()
	WithTruncateIDsSuffix("...")(&o)
	if o.TruncateIDsSuffix != "..." {
		t.Errorf("expected TruncateIDsSuffix='...', got %q", o.TruncateIDsSuffix)
	}
}

func TestWithTruncateIDsSuffix_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithTruncateIDsSuffix("…")(&o)
	WithTruncateIDsSuffix("")(&o)
	if o.TruncateIDsSuffix != "…" {
		t.Errorf("expected suffix preserved, got %q", o.TruncateIDsSuffix)
	}
}
