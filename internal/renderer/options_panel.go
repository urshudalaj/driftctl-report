package renderer

// WithPanel enables or disables the collapsible detail panel section
// in the rendered report. When enabled, each resource group is
// rendered inside an expandable panel component.
func WithPanel(enabled bool) Option {
	return func(o *Options) {
		o.PanelEnabled = enabled
	}
}

// WithPanelDefaultOpen controls whether panels are expanded by default
// when the report is first loaded. Has no effect if PanelEnabled is false.
func WithPanelDefaultOpen(open bool) Option {
	return func(o *Options) {
		o.PanelDefaultOpen = open
	}
}

// WithPanelMaxItems sets the maximum number of items shown inside a
// single panel before a "show more" affordance is rendered.
// A value of 0 means unlimited.
func WithPanelMaxItems(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.PanelMaxItems = n
	}
}
