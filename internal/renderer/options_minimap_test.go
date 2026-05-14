package renderer

import "testing"

func TestWithMinimap_Enables(t *testing.T) {
	opts := DefaultOptions()
	WithMinimap(true)(&opts)
	if !opts.Minimap {
		t.Fatal("expected Minimap to be true")
	}
}

func TestWithMinimap_Disables(t *testing.T) {
	opts := DefaultOptions()
	opts.Minimap = true
	WithMinimap(false)(&opts)
	if opts.Minimap {
		t.Fatal("expected Minimap to be false")
	}
}

func TestWithMinimap_DefaultIsFalse(t *testing.T) {
	opts := DefaultOptions()
	if opts.Minimap {
		t.Fatal("expected Minimap default to be false")
	}
}

func TestWithMinimapTopN_PositiveValue(t *testing.T) {
	opts := DefaultOptions()
	WithMinimapTopN(5)(&opts)
	if opts.MinimapTopN != 5 {
		t.Fatalf("expected MinimapTopN=5, got %d", opts.MinimapTopN)
	}
}

func TestWithMinimapTopN_ZeroMeansUnlimited(t *testing.T) {
	opts := DefaultOptions()
	WithMinimapTopN(0)(&opts)
	if opts.MinimapTopN != 0 {
		t.Fatalf("expected MinimapTopN=0, got %d", opts.MinimapTopN)
	}
}

func TestWithMinimapTopN_NegativeClampedToZero(t *testing.T) {
	opts := DefaultOptions()
	WithMinimapTopN(-3)(&opts)
	if opts.MinimapTopN != 0 {
		t.Fatalf("expected MinimapTopN clamped to 0, got %d", opts.MinimapTopN)
	}
}
