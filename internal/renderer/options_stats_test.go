package renderer

import "testing"

func TestWithStats_Enables(t *testing.T) {
	o := DefaultOptions()
	if err := WithStats(true)(o); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !o.ShowStats {
		t.Error("expected ShowStats to be true")
	}
}

func TestWithStats_Disables(t *testing.T) {
	o := DefaultOptions()
	o.ShowStats = true
	if err := WithStats(false)(o); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if o.ShowStats {
		t.Error("expected ShowStats to be false")
	}
}

func TestWithStatsTopN_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	if err := WithStatsTopN(5)(o); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if o.StatsTopN != 5 {
		t.Errorf("expected StatsTopN=5, got %d", o.StatsTopN)
	}
}

func TestWithStatsTopN_ZeroMeansUnlimited(t *testing.T) {
	o := DefaultOptions()
	o.StatsTopN = 10
	if err := WithStatsTopN(0)(o); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if o.StatsTopN != 0 {
		t.Errorf("expected StatsTopN=0, got %d", o.StatsTopN)
	}
}

func TestWithStatsTopN_NegativeClampedToZero(t *testing.T) {
	o := DefaultOptions()
	if err := WithStatsTopN(-3)(o); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if o.StatsTopN != 0 {
		t.Errorf("expected StatsTopN clamped to 0, got %d", o.StatsTopN)
	}
}

func TestDefaultOptions_StatsDefaults(t *testing.T) {
	o := DefaultOptions()
	if !o.ShowStats {
		t.Error("expected ShowStats default to be true")
	}
	if o.StatsTopN != 0 {
		t.Errorf("expected StatsTopN default to be 0, got %d", o.StatsTopN)
	}
}
