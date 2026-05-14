package renderer

// WithDarkMode enables or disables dark mode rendering for the HTML report.
// When enabled, the report uses a dark colour palette suitable for low-light
// environments. Defaults to false.
func WithDarkMode(enabled bool) Option {
	return func(o *Options) {
		o.DarkMode = enabled
	}
}

// WithDarkModeToggle controls whether a dark/light mode toggle button is
// rendered in the report navigation bar. Requires WithNav to be enabled.
// Defaults to false.
func WithDarkModeToggle(show bool) Option {
	return func(o *Options) {
		o.DarkModeToggle = show
	}
}

// WithDarkModeDefault sets the preferred default mode when the toggle is
// present. Accepted values are "dark" and "light". Any other value is
// silently ignored.
func WithDarkModeDefault(mode string) Option {
	return func(o *Options) {
		if mode == "dark" || mode == "light" {
			o.DarkModeDefault = mode
		}
	}
}
