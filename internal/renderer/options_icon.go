package renderer

// WithIcons enables or disables icon rendering next to resource status labels.
// When enabled, small SVG/Unicode icons are prepended to managed, unmanaged,
// and deleted resource entries in the HTML report.
func WithIcons(enabled bool) Option {
	return func(o *Options) {
		o.Icons = enabled
	}
}

// WithIconSet selects the named icon set to use when icons are enabled.
// Supported values: "unicode" (default), "emoji", "svg".
// Unknown values are silently ignored.
func WithIconSet(name string) Option {
	return func(o *Options) {
		switch name {
		case "unicode", "emoji", "svg":
			o.IconSet = name
		}
	}
}

// WithIconSize sets the display size (in pixels) for SVG icons.
// Values <= 0 are ignored; the existing size is preserved.
func WithIconSize(px int) Option {
	return func(o *Options) {
		if px > 0 {
			o.IconSize = px
		}
	}
}
