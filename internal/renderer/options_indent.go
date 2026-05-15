package renderer

// WithIndent enables or disables indented HTML output.
// When enabled, the rendered HTML will be formatted with consistent
// indentation to aid readability and debugging.
func WithIndent(enabled bool) Option {
	return func(o *Options) {
		o.IndentOutput = enabled
	}
}

// WithIndentSize sets the number of spaces used per indentation level.
// Values less than 1 are ignored; the existing value is preserved.
func WithIndentSize(size int) Option {
	return func(o *Options) {
		if size < 1 {
			return
		}
		o.IndentSize = size
	}
}

// WithIndentChar sets the character used for indentation.
// Only "\t" (tab) or " " (space) are accepted; other values are ignored.
func WithIndentChar(ch string) Option {
	return func(o *Options) {
		if ch != "\t" && ch != " " {
			return
		}
		o.IndentChar = ch
	}
}
