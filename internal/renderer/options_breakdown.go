package renderer

// WithBreakdown enables or disables the per-type resource breakdown section.
func WithBreakdown(enabled bool) Option {
	return func(o *Options) {
		o.ShowBreakdown = enabled
	}
}

// WithBreakdownTopN limits the breakdown table to the top N resource types
// sorted by total resource count descending. Zero means show all types.
func WithBreakdownTopN(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.BreakdownTopN = n
	}
}
