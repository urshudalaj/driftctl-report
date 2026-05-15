package renderer

// WithColumns sets which resource columns are visible in the report output.
// Valid column names are: "id", "type", "status", "region", "tags".
// Unknown column names are silently ignored.
func WithColumns(columns ...string) Option {
	known := map[string]bool{
		"id":     true,
		"type":   true,
		"status": true,
		"region": true,
		"tags":   true,
	}
	return func(o *Options) {
		for _, c := range columns {
			if known[c] {
				o.Columns = append(o.Columns, c)
			}
		}
	}
}

// WithColumnsAll enables all available columns in the report output.
func WithColumnsAll() Option {
	return func(o *Options) {
		o.Columns = []string{"id", "type", "status", "region", "tags"}
	}
}

// WithColumnsReset clears any previously configured column list,
// reverting to the default set of visible columns.
func WithColumnsReset() Option {
	return func(o *Options) {
		o.Columns = nil
	}
}
