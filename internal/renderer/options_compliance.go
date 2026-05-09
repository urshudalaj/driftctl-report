package renderer

// WithCompliance enables or disables the compliance section in the report.
func WithCompliance(enabled bool) Option {
	return func(o *Options) {
		o.Compliance = enabled
	}
}

// WithComplianceTopN limits the compliance table to the top N worst-performing
// resource types. A value of zero means no limit.
func WithComplianceTopN(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.ComplianceTopN = n
	}
}
