package renderer

import "testing"

func TestWithRecommendations_Enables(t *testing.T) {
	opts := DefaultOptions()
	WithRecommendations(true)(&opts)
	if !opts.Recommendations {
		t.Error("expected Recommendations to be true")
	}
}

func TestWithRecommendations_Disables(t *testing.T) {
	opts := DefaultOptions()
	opts.Recommendations = true
	WithRecommendations(false)(&opts)
	if opts.Recommendations {
		t.Error("expected Recommendations to be false")
	}
}

func TestWithRecommendations_DefaultIsFalse(t *testing.T) {
	opts := DefaultOptions()
	if opts.Recommendations {
		t.Error("expected Recommendations default to be false")
	}
}

func TestWithRecommendationsMaxItems_PositiveValue(t *testing.T) {
	opts := DefaultOptions()
	WithRecommendationsMaxItems(5)(&opts)
	if opts.RecommendationsMaxItems != 5 {
		t.Errorf("expected 5, got %d", opts.RecommendationsMaxItems)
	}
}

func TestWithRecommendationsMaxItems_ZeroMeansUnlimited(t *testing.T) {
	opts := DefaultOptions()
	WithRecommendationsMaxItems(0)(&opts)
	if opts.RecommendationsMaxItems != 0 {
		t.Errorf("expected 0, got %d", opts.RecommendationsMaxItems)
	}
}

func TestWithRecommendationsMaxItems_NegativeClampedToZero(t *testing.T) {
	opts := DefaultOptions()
	WithRecommendationsMaxItems(-3)(&opts)
	if opts.RecommendationsMaxItems != 0 {
		t.Errorf("expected 0, got %d", opts.RecommendationsMaxItems)
	}
}

func TestWithRecommendationsMaxItems_DefaultIsZero(t *testing.T) {
	opts := DefaultOptions()
	if opts.RecommendationsMaxItems != 0 {
		t.Errorf("expected default 0, got %d", opts.RecommendationsMaxItems)
	}
}
