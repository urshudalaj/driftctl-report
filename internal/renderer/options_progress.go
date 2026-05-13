package renderer

// WithProgress enables or disables the per-type managed-ratio progress bar
// section in the rendered report.
func WithProgress(enabled bool) Option {
	return func(o *Options) {
		o.ShowProgress = enabled
	}
}

// WithProgressTopN limits the progress bar section to the top N resource
// types (sorted alphabetically). Zero means no limit.
func WithProgressTopN(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.ProgressTopN = n
	}
}
