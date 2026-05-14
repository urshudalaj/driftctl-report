package renderer

// WithTruncateIDs enables or disables truncation of long resource IDs in the
// rendered output. When enabled, IDs longer than the configured length are
// shortened with an ellipsis suffix.
func WithTruncateIDs(enabled bool) Option {
	return func(o *Options) {
		o.TruncateIDs = enabled
	}
}

// WithTruncateIDsLength sets the maximum character length for resource IDs
// before they are truncated. Values less than 1 are ignored; the existing
// value is preserved.
func WithTruncateIDsLength(n int) Option {
	return func(o *Options) {
		if n < 1 {
			return
		}
		o.TruncateIDsLength = n
	}
}

// WithTruncateIDsSuffix sets the suffix appended to a truncated ID, e.g.
// "…" or "...". An empty value is ignored.
func WithTruncateIDsSuffix(suffix string) Option {
	return func(o *Options) {
		if suffix == "" {
			return
		}
		o.TruncateIDsSuffix = suffix
	}
}
