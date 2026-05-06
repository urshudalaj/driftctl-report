package renderer

// WithChangelog enables or disables the changelog section in the rendered report.
// When enabled (the default), a full per-resource change list is appended after
// the summary cards. Disabling it produces a more compact report.
func WithChangelog(enabled bool) Option {
	return func(o *Options) error {
		o.ShowChangelog = enabled
		return nil
	}
}

// WithChangelogLimit caps the maximum number of entries shown per changelog
// section (missing / unmanaged / changed). Zero or negative values mean
// unlimited. This is independent of WithMaxResources which limits table rows.
func WithChangelogLimit(n int) Option {
	return func(o *Options) error {
		if n < 0 {
			n = 0
		}
		o.ChangelogLimit = n
		return nil
	}
}
