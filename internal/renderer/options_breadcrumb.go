package renderer

// WithBreadcrumb enables or disables the breadcrumb navigation trail
// rendered at the top of the report.
func WithBreadcrumb(enabled bool) Option {
	return func(o *Options) {
		o.BreadcrumbEnabled = enabled
	}
}

// WithBreadcrumbSeparator sets the separator string rendered between
// breadcrumb segments. Defaults to "/" when empty.
func WithBreadcrumbSeparator(sep string) Option {
	return func(o *Options) {
		if sep == "" {
			return
		}
		o.BreadcrumbSeparator = sep
	}
}

// WithBreadcrumbItems appends one or more BreadcrumbItem entries to the
// breadcrumb trail. Items are rendered in the order they are added.
func WithBreadcrumbItems(items ...BreadcrumbItem) Option {
	return func(o *Options) {
		for _, item := range items {
			if item.Label == "" {
				continue
			}
			o.BreadcrumbItems = append(o.BreadcrumbItems, item)
		}
	}
}

// BreadcrumbItem represents a single segment in the breadcrumb trail.
type BreadcrumbItem struct {
	Label string
	URL   string // optional; if empty the item is rendered as plain text
}
