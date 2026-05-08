package renderer

// WithSnapshot enables or disables the snapshot/history section in the report.
// When enabled, a trend indicator is shown comparing the current run to
// previous entries supplied via WithSnapshotHistory.
func WithSnapshot(enabled bool) Option {
	return func(o *Options) {
		o.SnapshotEnabled = enabled
	}
}

// WithSnapshotHistory provides historical SnapshotEntry values that are
// combined with the current analysis to display a coverage trend over time.
// Passing a nil or empty slice is valid and results in no trend indicator.
func WithSnapshotHistory(entries []SnapshotEntry) Option {
	return func(o *Options) {
		if len(entries) == 0 {
			o.SnapshotHistory = nil
			return
		}
		copy := make([]SnapshotEntry, len(entries))
		for i, e := range entries {
			copy[i] = e
		}
		o.SnapshotHistory = copy
	}
}
