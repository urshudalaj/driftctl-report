package renderer

// WithTruncateIDs enables or disables ID truncation in the rendered output.
func WithTruncateIDs(enabled bool) Option {
	return func(o *Options) {
		o.TruncateIDs = enabled
	}
}

// WithTruncateIDsLength sets the maximum character length for displayed resource IDs.
// Values <= 0 are ignored; the existing setting is preserved.
func WithTruncateIDsLength(n int) Option {
	return func(o *Options) {
		if n <= 0 {
			return
		}
		o.TruncateIDsLen = n
	}
}

// WithTruncateIDsSuffix sets the suffix appended to truncated IDs (e.g. "…").
// An empty string is ignored; the existing setting is preserved.
func WithTruncateIDsSuffix(suffix string) Option {
	return func(o *Options) {
		if suffix == "" {
			return
		}
		o.TruncateIDsSuffix = suffix
	}
}
