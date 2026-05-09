package renderer

// WithHeatmap enables or disables the resource type heatmap section.
// The heatmap visualises drift intensity per resource type.
func WithHeatmap(enabled bool) Option {
	return func(o *Options) {
		o.Heatmap = enabled
	}
}

// WithHeatmapTopN limits the heatmap to the top N resource types by drift
// ratio. A value of 0 means no limit (all types are shown).
func WithHeatmapTopN(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.HeatmapTopN = n
	}
}
