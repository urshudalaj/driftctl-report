package renderer

// WithTooltip enables or disables inline tooltips on resource entries.
// When enabled, hovering over a resource ID in the HTML report shows a
// small popover with the resource type and status.
func WithTooltip(enabled bool) Option {
	return func(o *Options) {
		o.Tooltip = enabled
	}
}

// WithTooltipPlacement sets the preferred placement of tooltips relative
// to the target element. Accepted values: "top", "bottom", "left", "right".
// Unknown values are silently ignored and the previous setting is kept.
func WithTooltipPlacement(placement string) Option {
	switch placement {
	case "top", "bottom", "left", "right":
		return func(o *Options) {
			o.TooltipPlacement = placement
		}
	}
	return func(o *Options) {}
}

// WithTooltipMaxWidth sets the maximum width (in pixels) of the tooltip
// popover. Values <= 0 are ignored.
func WithTooltipMaxWidth(px int) Option {
	if px <= 0 {
		return func(o *Options) {}
	}
	return func(o *Options) {
		o.TooltipMaxWidth = px
	}
}
