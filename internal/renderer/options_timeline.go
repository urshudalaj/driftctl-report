package renderer

// WithTimeline enables or disables the timeline section in the report.
// The timeline lists drift events (missing, unmanaged, changed) in a
// chronological, scannable view. Disabled by default.
func WithTimeline(enabled bool) Option {
	return func(o *RenderOptions) {
		o.ShowTimeline = enabled
	}
}

// WithTimelineLimit sets the maximum number of events shown in the timeline.
// A value of 0 means unlimited. Negative values are clamped to 0.
func WithTimelineLimit(n int) Option {
	return func(o *RenderOptions) {
		if n < 0 {
			n = 0
		}
		o.TimelineLimit = n
	}
}
