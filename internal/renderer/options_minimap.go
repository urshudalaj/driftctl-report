package renderer

// WithMinimap enables or disables the minimap grid section in the report.
// The minimap provides a compact per-type drift overview.
func WithMinimap(enabled bool) Option {
	return func(o *Options) {
		o.Minimap = enabled
	}
}

// WithMinimapTopN limits the minimap to the N first resource types (sorted
// alphabetically). A value of 0 means no limit.
func WithMinimapTopN(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.MinimapTopN = n
	}
}
