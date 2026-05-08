package renderer

// WithScorecard enables or disables the drift health scorecard section.
// The scorecard is disabled by default.
func WithScorecard(enabled bool) Option {
	return func(o *Options) {
		o.ShowScorecard = enabled
	}
}

// WithScorecardThreshold sets the minimum acceptable score (0–100).
// When the computed score falls below this threshold the scorecard grade
// is rendered with a warning indicator.
// Values outside [0, 100] are clamped.
func WithScorecardThreshold(threshold float64) Option {
	return func(o *Options) {
		if threshold < 0 {
			threshold = 0
		}
		if threshold > 100 {
			threshold = 100
		}
		o.ScorecardThreshold = threshold
	}
}
