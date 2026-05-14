package renderer

// WithBadges enables or disables the badge row in the report header.
// Badges provide at-a-glance severity indicators for managed, unmanaged,
// and deleted resource counts.
func WithBadges(enabled bool) Option {
	return func(o *Options) {
		o.BadgesEnabled = enabled
	}
}

// WithBadgeLabels overrides the default labels used for each badge category.
// Keys must be "managed", "unmanaged", or "deleted"; unknown keys are ignored.
func WithBadgeLabels(labels map[string]string) Option {
	return func(o *Options) {
		if o.BadgeLabels == nil {
			o.BadgeLabels = make(map[string]string)
		}
		for k, v := range labels {
			switch k {
			case "managed", "unmanaged", "deleted":
				if v != "" {
					o.BadgeLabels[k] = v
				}
			}
		}
	}
}

// WithBadgeShowZero controls whether badges with a zero count are still
// rendered. When false (the default), zero-count badges are hidden.
func WithBadgeShowZero(show bool) Option {
	return func(o *Options) {
		o.BadgeShowZero = show
	}
}
