package renderer

// WithTable enables the resource table section in the report.
func WithTable(enabled bool) Option {
	return func(o *Options) {
		o.TableEnabled = enabled
	}
}

// WithTablePageSize sets the number of rows per page in the resource table.
// Values <= 0 are ignored; the renderer defaults to 25.
func WithTablePageSize(size int) Option {
	return func(o *Options) {
		if size > 0 {
			o.PageSize = size
		}
	}
}

// WithTableColumns sets the ordered list of columns to display in the table.
// An empty slice is ignored.
func WithTableColumns(cols []string) Option {
	return func(o *Options) {
		if len(cols) == 0 {
			return
		}
		o.TableColumns = make([]string, len(cols))
		copy(o.TableColumns, cols)
	}
}
