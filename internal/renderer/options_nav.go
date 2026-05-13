package renderer

// WithNav enables or disables the sticky navigation bar in the HTML report.
// When enabled, a top navigation bar is rendered with links to each major
// section of the report (Summary, Resources, Diff, etc.).
func WithNav(enabled bool) Option {
	return func(o *Options) {
		o.NavEnabled = enabled
	}
}

// WithNavBrand sets the brand/title text displayed in the navigation bar.
// If text is empty the call is ignored and the existing value is preserved.
func WithNavBrand(text string) Option {
	return func(o *Options) {
		if text == "" {
			return
		}
		o.NavBrand = text
	}
}

// WithNavSticky controls whether the navigation bar stays fixed at the top
// of the viewport while the user scrolls. Defaults to true when nav is enabled.
func WithNavSticky(sticky bool) Option {
	return func(o *Options) {
		o.NavSticky = sticky
	}
}
