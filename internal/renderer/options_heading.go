package renderer

// WithHeading sets a custom heading text displayed at the top of the report
// body, beneath the navigation bar. An empty value is silently ignored.
func WithHeading(text string) Option {
	return func(o *Options) {
		if text == "" {
			return
		}
		o.HeadingText = text
	}
}

// WithHeadingLevel sets the HTML heading level (1–6) used for the report
// heading. Values outside the valid range are ignored.
func WithHeadingLevel(level int) Option {
	return func(o *Options) {
		if level < 1 || level > 6 {
			return
		}
		o.HeadingLevel = level
	}
}

// WithHeadingVisible controls whether the heading element is rendered.
// Defaults to true when a heading text has been set.
func WithHeadingVisible(visible bool) Option {
	return func(o *Options) {
		o.HeadingVisible = visible
	}
}
