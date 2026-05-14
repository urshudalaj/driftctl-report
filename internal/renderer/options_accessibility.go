package renderer

// WithAccessibility enables ARIA labels and role attributes in the rendered
// HTML output, improving compatibility with screen readers and assistive
// technologies.
func WithAccessibility(enabled bool) Option {
	return func(o *Options) {
		o.Accessibility = enabled
	}
}

// WithAccessibilitySkipNav adds a "skip to main content" link at the top of
// the page, allowing keyboard users to bypass the navigation bar.
func WithAccessibilitySkipNav(enabled bool) Option {
	return func(o *Options) {
		o.AccessibilitySkipNav = enabled
	}
}

// WithAccessibilityAnnounce enables a live region that announces dynamic
// content changes (e.g. search results) to screen readers.
func WithAccessibilityAnnounce(enabled bool) Option {
	return func(o *Options) {
		o.AccessibilityAnnounce = enabled
	}
}
