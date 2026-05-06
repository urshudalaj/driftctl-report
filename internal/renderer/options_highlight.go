package renderer

// WithHighlight enables or disables inline search-term highlighting in the
// rendered HTML. When enabled, occurrences of the active search query are
// wrapped in <mark> tags inside resource ID and type cells.
//
// Highlighting is enabled by default when a non-empty search query is set via
// [WithSearch]. Set enabled to false to suppress <mark> wrapping even when a
// query is active.
func WithHighlight(enabled bool) Option {
	return func(o *Options) {
		o.Highlight = enabled
	}
}

// WithHighlightMaxLength sets the maximum character length of a resource ID
// that will be highlighted. IDs longer than max are rendered as plain text to
// avoid expensive regex operations on very long strings. A value of 0 means
// no limit is applied.
func WithHighlightMaxLength(max int) Option {
	return func(o *Options) {
		if max < 0 {
			max = 0
		}
		o.HighlightMaxLength = max
	}
}
