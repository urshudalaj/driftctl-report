package renderer

// WithPrint enables or disables the print-friendly CSS media styles
// injected into the HTML report. When enabled, a <style> block with
// @media print rules is appended so the report renders cleanly on paper
// or when saved as PDF via the browser.
func WithPrint(enabled bool) Option {
	return func(o *Options) {
		o.PrintFriendly = enabled
	}
}

// WithPrintHideNav controls whether the navigation bar is hidden in
// print mode. Defaults to true when print-friendly mode is active.
func WithPrintHideNav(hide bool) Option {
	return func(o *Options) {
		o.PrintHideNav = hide
	}
}

// WithPrintPageBreaks enables automatic page-break hints between major
// report sections (summary, changelog, timeline, etc.) in print output.
func WithPrintPageBreaks(enabled bool) Option {
	return func(o *Options) {
		o.PrintPageBreaks = enabled
	}
}
