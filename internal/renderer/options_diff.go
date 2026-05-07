package renderer

// WithDiff enables or disables the attribute-level diff section in the report.
// When enabled, each drifted resource will list its changed attributes
// alongside the IaC ("from") and real-state ("to") values.
func WithDiff(enabled bool) Option {
	return func(o *Options) {
		o.ShowDiff = enabled
	}
}

// WithDiffLimit sets the maximum number of diff entries rendered.
// A value of 0 means unlimited. Negative values are clamped to 0.
func WithDiffLimit(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.DiffLimit = n
	}
}
