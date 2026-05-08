package renderer

import "testing"

func TestWithScorecard_Enables(t *testing.T) {
	o := DefaultOptions()
	WithScorecard(true)(o)
	if !o.ShowScorecard {
		t.Error("expected ShowScorecard to be true")
	}
}

func TestWithScorecard_Disables(t *testing.T) {
	o := DefaultOptions()
	o.ShowScorecard = true
	WithScorecard(false)(o)
	if o.ShowScorecard {
		t.Error("expected ShowScorecard to be false")
	}
}

func TestWithScorecard_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.ShowScorecard {
		t.Error("expected ShowScorecard default to be false")
	}
}

func TestWithScorecardThreshold_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithScorecardThreshold(70)(o)
	if o.ScorecardThreshold != 70 {
		t.Errorf("expected threshold 70, got %v", o.ScorecardThreshold)
	}
}

func TestWithScorecardThreshold_ZeroIsValid(t *testing.T) {
	o := DefaultOptions()
	WithScorecardThreshold(0)(o)
	if o.ScorecardThreshold != 0 {
		t.Errorf("expected threshold 0, got %v", o.ScorecardThreshold)
	}
}

func TestWithScorecardThreshold_NegativeClampedToZero(t *testing.T) {
	o := DefaultOptions()
	WithScorecardThreshold(-10)(o)
	if o.ScorecardThreshold != 0 {
		t.Errorf("expected threshold clamped to 0, got %v", o.ScorecardThreshold)
	}
}

func TestWithScorecardThreshold_AboveHundredClampedTo100(t *testing.T) {
	o := DefaultOptions()
	WithScorecardThreshold(150)(o)
	if o.ScorecardThreshold != 100 {
		t.Errorf("expected threshold clamped to 100, got %v", o.ScorecardThreshold)
	}
}
