package renderer

// WithTrend enables or disables the trend section in the report.
func WithTrend(enabled bool) Option {
	return func(o *Options) {
		o.TrendEnabled = enabled
	}
}

// WithTrendHistory supplies historical snapshot data for trend computation.
// Each entry should be a map with keys: "label", "managed", "unmanaged", "deleted".
func WithTrendHistory(history []map[string]interface{}) Option {
	return func(o *Options) {
		if len(history) > 0 {
			o.TrendHistory = history
		}
	}
}

// WithTrendMaxPoints limits the number of data points shown in the trend chart.
// A value of 0 means unlimited.
func WithTrendMaxPoints(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.TrendMaxPoints = n
	}
}
