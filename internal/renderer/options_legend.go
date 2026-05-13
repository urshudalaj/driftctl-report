package renderer

// WithLegend enables or disables the legend section in the rendered report.
// The legend provides a key explaining the meaning of status labels, badge
// colours, and severity indicators used throughout the report.
func WithLegend(enabled bool) Option {
	return func(o *Options) {
		o.LegendEnabled = enabled
	}
}

// WithLegendPosition sets the position of the legend within the report.
// Accepted values are "top" and "bottom"; any other value is ignored.
func WithLegendPosition(position string) Option {
	return func(o *Options) {
		if position != "top" && position != "bottom" {
			return
		}
		o.LegendPosition = position
	}
}

// WithLegendCompact renders the legend in a compact single-row layout when
// set to true. Defaults to false (expanded multi-row layout).
func WithLegendCompact(compact bool) Option {
	return func(o *Options) {
		o.LegendCompact = compact
	}
}
