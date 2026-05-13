package renderer

// WithColorScheme sets a named color scheme for resource status indicators.
// Valid values are "default", "accessible", and "monochrome".
// Unknown values are silently ignored.
func WithColorScheme(scheme string) Option {
	return func(o *Options) {
		switch scheme {
		case "default", "accessible", "monochrome":
			o.ColorScheme = scheme
		}
	}
}

// WithColorManaged sets the CSS color class used for managed resources.
// An empty value is ignored.
func WithColorManaged(class string) Option {
	return func(o *Options) {
		if class != "" {
			o.ColorManaged = class
		}
	}
}

// WithColorUnmanaged sets the CSS color class used for unmanaged resources.
// An empty value is ignored.
func WithColorUnmanaged(class string) Option {
	return func(o *Options) {
		if class != "" {
			o.ColorUnmanaged = class
		}
	}
}

// WithColorDeleted sets the CSS color class used for deleted resources.
// An empty value is ignored.
func WithColorDeleted(class string) Option {
	return func(o *Options) {
		if class != "" {
			o.ColorDeleted = class
		}
	}
}
