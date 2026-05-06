package renderer

// WithPageSize sets the number of resources displayed per page in the HTML
// report. Values less than 1 are ignored and the default (25) is retained.
func WithPageSize(size int) Option {
	return func(o *Options) error {
		if size < 1 {
			return nil
		}
		o.PageSize = size
		return nil
	}
}

// WithPage sets the current page number for the rendered report.
// Values less than 1 are clamped to 1.
func WithPage(page int) Option {
	return func(o *Options) error {
		if page < 1 {
			page = 1
		}
		o.Page = page
		return nil
	}
}
