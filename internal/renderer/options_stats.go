package renderer

// WithStats enables or disables the per-type statistics section in the report.
//
// When enabled (default: true), the rendered HTML will include a breakdown
// table showing managed, unmanaged, and missing counts per resource type.
func WithStats(enabled bool) Option {
	return func(o *Options) error {
		o.ShowStats = enabled
		return nil
	}
}

// WithStatsTopN limits the stats table to the top N resource types by total
// resource count. A value of 0 means no limit (all types are shown).
//
// Negative values are clamped to 0.
func WithStatsTopN(n int) Option {
	return func(o *Options) error {
		if n < 0 {
			n = 0
		}
		o.StatsTopN = n
		return nil
	}
}
