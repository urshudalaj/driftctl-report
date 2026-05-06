package renderer

// WithGroupByType enables or disables the grouping of resources by their
// resource type in the rendered HTML output. When enabled (the default),
// resources are presented under collapsible type headings rather than in a
// single flat table.
func WithGroupByType(enabled bool) Option {
	return func(o *Options) error {
		o.GroupByType = enabled
		return nil
	}
}

// WithGroupCollapsed controls whether grouped type sections are rendered in a
// collapsed state by default. Has no effect when GroupByType is false.
func WithGroupCollapsed(collapsed bool) Option {
	return func(o *Options) error {
		o.GroupCollapsed = collapsed
		return nil
	}
}
