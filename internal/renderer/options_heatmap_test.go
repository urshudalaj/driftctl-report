package renderer

import "testing"

func TestWithHeatmap_Enables(t *testing.T) {
	opts := DefaultOptions()
	WithHeatmap(true)(&opts)
	if !opts.Heatmap {
		t.Fatal("expected Heatmap to be true")
	}
}

func TestWithHeatmap_Disables(t *testing.T) {
	opts := DefaultOptions()
	opts.Heatmap = true
	WithHeatmap(false)(&opts)
	if opts.Heatmap {
		t.Fatal("expected Heatmap to be false")
	}
}

func TestWithHeatmap_DefaultIsFalse(t *testing.T) {
	opts := DefaultOptions()
	if opts.Heatmap {
		t.Fatal("expected default Heatmap to be false")
	}
}

func TestWithHeatmapTopN_PositiveValue(t *testing.T) {
	opts := DefaultOptions()
	WithHeatmapTopN(5)(&opts)
	if opts.HeatmapTopN != 5 {
		t.Fatalf("expected HeatmapTopN=5, got %d", opts.HeatmapTopN)
	}
}

func TestWithHeatmapTopN_ZeroMeansUnlimited(t *testing.T) {
	opts := DefaultOptions()
	WithHeatmapTopN(10)(&opts)
	WithHeatmapTopN(0)(&opts)
	if opts.HeatmapTopN != 0 {
		t.Fatalf("expected HeatmapTopN=0, got %d", opts.HeatmapTopN)
	}
}

func TestWithHeatmapTopN_NegativeClampedToZero(t *testing.T) {
	opts := DefaultOptions()
	WithHeatmapTopN(-3)(&opts)
	if opts.HeatmapTopN != 0 {
		t.Fatalf("expected HeatmapTopN=0, got %d", opts.HeatmapTopN)
	}
}

func TestWithHeatmapTopN_DefaultIsZero(t *testing.T) {
	opts := DefaultOptions()
	if opts.HeatmapTopN != 0 {
		t.Fatalf("expected default HeatmapTopN=0, got %d", opts.HeatmapTopN)
	}
}
