package renderer

import "testing"

func TestWithProgress_Enables(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowProgress = false
	WithProgress(true)(&opts)
	if !opts.ShowProgress {
		t.Error("expected ShowProgress to be true")
	}
}

func TestWithProgress_Disables(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowProgress = true
	WithProgress(false)(&opts)
	if opts.ShowProgress {
		t.Error("expected ShowProgress to be false")
	}
}

func TestWithProgress_DefaultIsFalse(t *testing.T) {
	opts := DefaultOptions()
	if opts.ShowProgress {
		t.Error("expected ShowProgress default to be false")
	}
}

func TestWithProgressTopN_PositiveValue(t *testing.T) {
	opts := DefaultOptions()
	WithProgressTopN(5)(&opts)
	if opts.ProgressTopN != 5 {
		t.Errorf("expected ProgressTopN=5, got %d", opts.ProgressTopN)
	}
}

func TestWithProgressTopN_ZeroMeansUnlimited(t *testing.T) {
	opts := DefaultOptions()
	WithProgressTopN(0)(&opts)
	if opts.ProgressTopN != 0 {
		t.Errorf("expected ProgressTopN=0, got %d", opts.ProgressTopN)
	}
}

func TestWithProgressTopN_NegativeClampedToZero(t *testing.T) {
	opts := DefaultOptions()
	WithProgressTopN(-3)(&opts)
	if opts.ProgressTopN != 0 {
		t.Errorf("expected ProgressTopN clamped to 0, got %d", opts.ProgressTopN)
	}
}
