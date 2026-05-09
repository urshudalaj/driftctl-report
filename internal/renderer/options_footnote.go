package renderer

// WithFootnote appends a custom footnote string to the report.
// Multiple calls accumulate footnotes in order.
func WithFootnote(text string) Option {
	return func(o *Options) {
		if text == "" {
			return
		}
		o.Footnotes = append(o.Footnotes, text)
	}
}

// WithFootnoteLimit sets the maximum number of footnotes rendered.
// Zero means unlimited. Negative values are clamped to zero.
func WithFootnoteLimit(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.FootnoteLimit = n
	}
}
