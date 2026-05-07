package renderer

import "testing"

func TestWithDiff_Enables(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowDiff = false
	WithDiff(true)(&opts)
	if !opts.ShowDiff {
		t.Error("expected ShowDiff to be true")
	}
}

func TestWithDiff_Disables(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowDiff = true
	WithDiff(false)(&opts)
	if opts.ShowDiff {
		t.Error("expected ShowDiff to be false")
	}
}

func TestWithDiffLimit_PositiveValue(t *testing.T) {
	opts := DefaultOptions()
	WithDiffLimit(50)(&opts)
	if opts.DiffLimit != 50 {
		t.Errorf("expected DiffLimit=50, got %d", opts.DiffLimit)
	}
}

func TestWithDiffLimit_ZeroMeansUnlimited(t *testing.T) {
	opts := DefaultOptions()
	opts.DiffLimit = 10
	WithDiffLimit(0)(&opts)
	if opts.DiffLimit != 0 {
		t.Errorf("expected DiffLimit=0, got %d", opts.DiffLimit)
	}
}

func TestWithDiffLimit_NegativeClampedToZero(t *testing.T) {
	opts := DefaultOptions()
	WithDiffLimit(-5)(&opts)
	if opts.DiffLimit != 0 {
		t.Errorf("expected DiffLimit clamped to 0, got %d", opts.DiffLimit)
	}
}

func TestWithDiff_DefaultIsFalse(t *testing.T) {
	opts := DefaultOptions()
	if opts.ShowDiff {
		t.Error("expected ShowDiff default to be false")
	}
}

func TestWithDiffLimit_DefaultIsZero(t *testing.T) {
	opts := DefaultOptions()
	if opts.DiffLimit != 0 {
		t.Errorf("expected DiffLimit default to be 0, got %d", opts.DiffLimit)
	}
}
