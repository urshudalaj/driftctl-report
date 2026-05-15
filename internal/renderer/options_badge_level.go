package renderer

// WithBadgeLevelThresholds sets custom numeric thresholds that control when a
// badge transitions from success → warning → danger.
//
// low  – counts strictly below this value render as success (green).
// high – counts at or above this value render as danger (red).
// Values between low and high render as warning (yellow).
//
// Both values must be positive and low must be less than high; invalid
// combinations are silently ignored and the defaults are preserved.
func WithBadgeLevelThresholds(low, high int) Option {
	return func(o *Options) {
		if low <= 0 || high <= 0 || low >= high {
			return
		}
		o.BadgeLevelLow = low
		o.BadgeLevelHigh = high
	}
}

// WithBadgeSuccessLabel overrides the accessible label text shown on success
// (green) badges. An empty value is ignored.
func WithBadgeSuccessLabel(label string) Option {
	return func(o *Options) {
		if label == "" {
			return
		}
		o.BadgeSuccessLabel = label
	}
}

// WithBadgeDangerLabel overrides the accessible label text shown on danger
// (red) badges. An empty value is ignored.
func WithBadgeDangerLabel(label string) Option {
	return func(o *Options) {
		if label == "" {
			return
		}
		o.BadgeDangerLabel = label
	}
}
