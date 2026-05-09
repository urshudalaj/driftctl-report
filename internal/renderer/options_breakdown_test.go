package renderer

import "testing"

func TestWithBreakdown_Enables(t *testing.T) {
	opts := DefaultOptions()
	WithBreakdown(true)(&opts)
	if !opts.ShowBreakdown {
		t.Fatal("expected ShowBreakdown to be true")
	}
}

func TestWithBreakdown_Disables(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowBreakdown = true
	WithBreakdown(false)(&opts)
	if opts.ShowBreakdown {
		t.Fatal("expected ShowBreakdown to be false")
	}
}

func TestWithBreakdown_DefaultIsFalse(t *testing.T) {
	opts := DefaultOptions()
	if opts.ShowBreakdown {
		t.Fatal("expected ShowBreakdown default to be false")
	}
}

func TestWithBreakdownTopN_PositiveValue(t *testing.T) {
	opts := DefaultOptions()
	WithBreakdownTopN(5)(&opts)
	if opts.BreakdownTopN != 5 {
		t.Fatalf("expected BreakdownTopN=5, got %d", opts.BreakdownTopN)
	}
}

func TestWithBreakdownTopN_ZeroMeansUnlimited(t *testing.T) {
	opts := DefaultOptions()
	WithBreakdownTopN(0)(&opts)
	if opts.BreakdownTopN != 0 {
		t.Fatalf("expected BreakdownTopN=0, got %d", opts.BreakdownTopN)
	}
}

func TestWithBreakdownTopN_NegativeClampedToZero(t *testing.T) {
	opts := DefaultOptions()
	WithBreakdownTopN(-3)(&opts)
	if opts.BreakdownTopN != 0 {
		t.Fatalf("expected BreakdownTopN clamped to 0, got %d", opts.BreakdownTopN)
	}
}

func TestWithBreakdownTopN_DefaultIsZero(t *testing.T) {
	opts := DefaultOptions()
	if opts.BreakdownTopN != 0 {
		t.Fatalf("expected default BreakdownTopN=0, got %d", opts.BreakdownTopN)
	}
}
