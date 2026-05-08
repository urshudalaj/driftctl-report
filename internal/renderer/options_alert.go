package renderer

// WithAlerts enables or disables the alerts section in the report.
// When enabled, resources exceeding configured thresholds are highlighted.
func WithAlerts(enabled bool) Option {
	return func(o *Options) {
		o.AlertsEnabled = enabled
	}
}

// WithAlertThresholdUnmanaged sets the minimum number of unmanaged resources
// that triggers a warning-level alert. Values <= 0 are clamped to 1.
func WithAlertThresholdUnmanaged(n int) Option {
	return func(o *Options) {
		if n <= 0 {
			n = 1
		}
		o.AlertThresholdUnmanaged = n
	}
}

// WithAlertThresholdDeleted sets the minimum number of deleted resources
// that triggers a warning-level alert. Values <= 0 are clamped to 1.
func WithAlertThresholdDeleted(n int) Option {
	return func(o *Options) {
		if n <= 0 {
			n = 1
		}
		o.AlertThresholdDeleted = n
	}
}
