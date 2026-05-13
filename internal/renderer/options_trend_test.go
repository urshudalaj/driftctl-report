package renderer

import "testing"

func TestWithTrend_Enables(t *testing.T) {
	o := DefaultOptions()
	WithTrend(true)(o)
	if !o.TrendEnabled {
		t.Fatal("expected TrendEnabled to be true")
	}
}

func TestWithTrend_Disables(t *testing.T) {
	o := DefaultOptions()
	o.TrendEnabled = true
	WithTrend(false)(o)
	if o.TrendEnabled {
		t.Fatal("expected TrendEnabled to be false")
	}
}

func TestWithTrend_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.TrendEnabled {
		t.Fatal("expected TrendEnabled default to be false")
	}
}

func TestWithTrendHistory_SetsHistory(t *testing.T) {
	o := DefaultOptions()
	h := []map[string]interface{}{
		{"label": "2024-01", "managed": 5, "unmanaged": 2, "deleted": 0},
	}
	WithTrendHistory(h)(o)
	if len(o.TrendHistory) != 1 {
		t.Fatalf("expected 1 history entry, got %d", len(o.TrendHistory))
	}
}

func TestWithTrendHistory_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	prev := []map[string]interface{}{{"label": "2024-01"}}
	o.TrendHistory = prev
	WithTrendHistory([]map[string]interface{}{})(o)
	if len(o.TrendHistory) != 1 {
		t.Fatal("expected history to be unchanged when empty slice provided")
	}
}

func TestWithTrendMaxPoints_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithTrendMaxPoints(6)(o)
	if o.TrendMaxPoints != 6 {
		t.Fatalf("expected 6, got %d", o.TrendMaxPoints)
	}
}

func TestWithTrendMaxPoints_ZeroMeansUnlimited(t *testing.T) {
	o := DefaultOptions()
	WithTrendMaxPoints(0)(o)
	if o.TrendMaxPoints != 0 {
		t.Fatalf("expected 0, got %d", o.TrendMaxPoints)
	}
}

func TestWithTrendMaxPoints_NegativeClampedToZero(t *testing.T) {
	o := DefaultOptions()
	WithTrendMaxPoints(-3)(o)
	if o.TrendMaxPoints != 0 {
		t.Fatalf("expected 0 after clamp, got %d", o.TrendMaxPoints)
	}
}
