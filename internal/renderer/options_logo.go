package renderer

// WithLogo sets a custom logo URL or data URI to display in the report header.
// An empty value clears any previously set logo.
func WithLogo(url string) Option {
	return func(o *Options) {
		if url == "" {
			o.LogoURL = ""
			return
		}
		o.LogoURL = url
	}
}

// WithLogoAlt sets the alternative text for the logo image.
// An empty value is silently ignored.
func WithLogoAlt(alt string) Option {
	return func(o *Options) {
		if alt == "" {
			return
		}
		o.LogoAlt = alt
	}
}

// WithLogoHeight sets the display height (in pixels) of the logo.
// Values less than or equal to zero are ignored.
func WithLogoHeight(px int) Option {
	return func(o *Options) {
		if px <= 0 {
			return
		}
		o.LogoHeight = px
	}
}
