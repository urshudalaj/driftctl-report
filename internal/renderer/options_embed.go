package renderer

// WithEmbedCSS enables or disables inlining of CSS styles directly into the
// HTML output rather than referencing an external stylesheet.
func WithEmbedCSS(enabled bool) Option {
	return func(o *Options) {
		o.EmbedCSS = enabled
	}
}

// WithEmbedJS enables or disables inlining of JavaScript directly into the
// HTML output rather than referencing an external script file.
func WithEmbedJS(enabled bool) Option {
	return func(o *Options) {
		o.EmbedJS = enabled
	}
}

// WithEmbedFonts enables or disables inlining of web font data (base64) into
// the HTML output so the report is fully self-contained.
func WithEmbedFonts(enabled bool) Option {
	return func(o *Options) {
		o.EmbedFonts = enabled
	}
}
