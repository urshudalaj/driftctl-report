package renderer

// WithCSV enables CSV export as an additional output alongside HTML.
// When enabled, a CSV file containing resource rows will be produced.
func WithCSV(enabled bool) Option {
	return func(o *Options) {
		o.ExportCSV = enabled
	}
}

// WithCSVFilename sets the filename used when writing CSV output.
// If empty the option is silently ignored and the default is kept.
func WithCSVFilename(name string) Option {
	return func(o *Options) {
		if name == "" {
			return
		}
		o.CSVFilename = name
	}
}

// WithCSVDelimiter sets the field delimiter for CSV output.
// Defaults to comma. If the rune is zero the option is ignored.
func WithCSVDelimiter(r rune) Option {
	return func(o *Options) {
		if r == 0 {
			return
		}
		o.CSVDelimiter = r
	}
}
