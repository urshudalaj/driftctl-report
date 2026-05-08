package renderer

// WithRecommendations enables or disables the recommendations section in the report.
// When enabled, the renderer will produce actionable suggestions based on the
// drift analysis results (unmanaged, deleted, and drifted resources).
func WithRecommendations(enabled bool) Option {
	return func(o *Options) {
		o.Recommendations = enabled
	}
}

// WithRecommendationsMaxItems limits how many recommendations are shown.
// A value of 0 means unlimited. Negative values are clamped to 0.
func WithRecommendationsMaxItems(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.RecommendationsMaxItems = n
	}
}
