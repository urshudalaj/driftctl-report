package renderer

import "testing"

func TestWithBadgeLevelThresholds_ValidValues(t *testing.T) {
	o := DefaultOptions()
	WithBadgeLevelThresholds(5, 20)(o)
	if o.BadgeLevelLow != 5 {
		t.Errorf("expected BadgeLevelLow=5, got %d", o.BadgeLevelLow)
	}
	if o.BadgeLevelHigh != 20 {
		t.Errorf("expected BadgeLevelHigh=20, got %d", o.BadgeLevelHigh)
	}
}

func TestWithBadgeLevelThresholds_ZeroLowIgnored(t *testing.T) {
	o := DefaultOptions()
	prev := o.BadgeLevelLow
	WithBadgeLevelThresholds(0, 20)(o)
	if o.BadgeLevelLow != prev {
		t.Error("zero low should be ignored")
	}
}

func TestWithBadgeLevelThresholds_NegativeHighIgnored(t *testing.T) {
	o := DefaultOptions()
	prev := o.BadgeLevelHigh
	WithBadgeLevelThresholds(5, -1)(o)
	if o.BadgeLevelHigh != prev {
		t.Error("negative high should be ignored")
	}
}

func TestWithBadgeLevelThresholds_LowEqualHighIgnored(t *testing.T) {
	o := DefaultOptions()
	prev := o.BadgeLevelLow
	WithBadgeLevelThresholds(10, 10)(o)
	if o.BadgeLevelLow != prev {
		t.Error("low == high should be ignored")
	}
}

func TestWithBadgeLevelThresholds_LowGreaterThanHighIgnored(t *testing.T) {
	o := DefaultOptions()
	prev := o.BadgeLevelLow
	WithBadgeLevelThresholds(15, 5)(o)
	if o.BadgeLevelLow != prev {
		t.Error("low > high should be ignored")
	}
}

func TestWithBadgeSuccessLabel_SetsLabel(t *testing.T) {
	o := DefaultOptions()
	WithBadgeSuccessLabel("ok")(o)
	if o.BadgeSuccessLabel != "ok" {
		t.Errorf("expected 'ok', got %q", o.BadgeSuccessLabel)
	}
}

func TestWithBadgeSuccessLabel_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithBadgeSuccessLabel("healthy")(o)
	WithBadgeSuccessLabel("")(o)
	if o.BadgeSuccessLabel != "healthy" {
		t.Error("empty label should not overwrite existing value")
	}
}

func TestWithBadgeDangerLabel_SetsLabel(t *testing.T) {
	o := DefaultOptions()
	WithBadgeDangerLabel("critical")(o)
	if o.BadgeDangerLabel != "critical" {
		t.Errorf("expected 'critical', got %q", o.BadgeDangerLabel)
	}
}

func TestWithBadgeDangerLabel_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithBadgeDangerLabel("error")(o)
	WithBadgeDangerLabel("")(o)
	if o.BadgeDangerLabel != "error" {
		t.Error("empty label should not overwrite existing value")
	}
}
