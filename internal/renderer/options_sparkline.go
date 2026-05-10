package renderer

// WithSparkline enables or disables the per-type drift sparkline section.
// When enabled, the report renders a small bar-chart style overview of
// unmanaged and deleted resource counts grouped by resource type.
func WithSparkline(enabled bool) Option {
	return func(o *Options) {
		o.Sparkline = enabled
	}
}

// WithSparklineTopN limits the sparkline to the top-N resource types by
// drift count. Zero means no limit (all types are shown).
func WithSparklineTopN(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.SparklineTopN = n
	}
}
