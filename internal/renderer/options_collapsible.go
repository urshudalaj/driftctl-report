package renderer

// WithCollapsible enables or disables collapsible sections in the report.
// When enabled, each resource-type group can be expanded or collapsed by
// the user via a toggle control in the rendered HTML.
func WithCollapsible(enabled bool) Option {
	return func(o *Options) {
		o.Collapsible = enabled
	}
}

// WithCollapsibleDefaultOpen controls whether collapsible sections start
// in the open (expanded) state. Has no effect if WithCollapsible is false.
func WithCollapsibleDefaultOpen(open bool) Option {
	return func(o *Options) {
		o.CollapsibleDefaultOpen = open
	}
}

// WithCollapsibleAnimated enables a CSS transition animation when sections
// are toggled. Has no effect if WithCollapsible is false.
func WithCollapsibleAnimated(animated bool) Option {
	return func(o *Options) {
		o.CollapsibleAnimated = animated
	}
}
