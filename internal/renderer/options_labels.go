package renderer

// WithLabels enables the resource labels section in the report.
// When enabled, each resource entry may display user-defined key/value
// label pairs sourced from the WithLabelMap option.
func WithLabels(enabled bool) Option {
	return func(o *Options) {
		o.LabelsEnabled = enabled
	}
}

// WithLabelMap sets a map of resource ID to label key/value pairs.
// Only IDs present in the analysis are rendered; unknown IDs are silently
// ignored. An empty map is accepted and disables label display even when
// WithLabels(true) is set.
func WithLabelMap(m map[string]map[string]string) Option {
	return func(o *Options) {
		if len(m) == 0 {
			return
		}
		o.LabelMap = m
	}
}

// WithLabelKeys restricts which label keys are shown in the report.
// Keys are matched case-sensitively. Passing an empty slice clears any
// previous restriction and shows all keys.
func WithLabelKeys(keys []string) Option {
	return func(o *Options) {
		if len(keys) == 0 {
			o.LabelKeys = nil
			return
		}
		o.LabelKeys = keys
	}
}
