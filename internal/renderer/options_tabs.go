package renderer

// WithTabs enables or disables the tabbed navigation view in the report.
// When enabled, resource sections are organized into clickable tabs.
func WithTabs(enabled bool) Option {
	return func(o *Options) {
		o.TabsEnabled = enabled
	}
}

// WithTabsDefaultActive sets the label of the tab that should be active
// (visible) when the report first loads. If the label does not match any
// generated tab the first tab is shown instead.
func WithTabsDefaultActive(label string) Option {
	return func(o *Options) {
		if label == "" {
			return
		}
		o.TabsDefaultActive = label
	}
}

// WithTabsPosition sets the visual position of the tab bar.
// Accepted values are "top" (default) and "left".
// Unknown values are silently ignored.
func WithTabsPosition(position string) Option {
	return func(o *Options) {
		switch position {
		case "top", "left":
			o.TabsPosition = position
		}
	}
}
