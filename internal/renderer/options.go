package renderer

// Options controls how the HTML report is rendered.
type Options struct {
	// Title is the report page title.
	Title string
	// OnlyDrifted restricts output to resources with detected drift.
	OnlyDrifted bool
	// ResourceType filters output to a specific resource type (empty = all).
	ResourceType string
	// MaxResources caps the number of resources shown per bucket (0 = unlimited).
	MaxResources int
	// SortOrder controls the ordering of resources in the report.
	SortOrder SortOrder
}

// Option is a functional option for Options.
type Option func(*Options)

// DefaultOptions returns an Options struct populated with sensible defaults.
func DefaultOptions() Options {
	return Options{
		Title:        "Drift Report",
		OnlyDrifted:  false,
		ResourceType: "",
		MaxResources: 0,
		SortOrder:    SortByTypeAsc,
	}
}

// WithTitle sets the HTML report title.
func WithTitle(title string) Option {
	return func(o *Options) {
		if title != "" {
			o.Title = title
		}
	}
}

// WithOnlyDrifted restricts the report to drifted resources only.
func WithOnlyDrifted(v bool) Option {
	return func(o *Options) {
		o.OnlyDrifted = v
	}
}

// WithResourceType filters the report to a single resource type.
func WithResourceType(rt string) Option {
	return func(o *Options) {
		o.ResourceType = rt
	}
}

// WithMaxResources caps the number of resources rendered per section.
func WithMaxResources(n int) Option {
	return func(o *Options) {
		if n >= 0 {
			o.MaxResources = n
		}
	}
}

// WithSortOrder sets the sort order applied to all resource lists.
func WithSortOrder(order SortOrder) Option {
	return func(o *Options) {
		o.SortOrder = order
	}
}

// applyOptions merges a set of Option functions onto a base Options value.
func applyOptions(base Options, opts []Option) Options {
	for _, o := range opts {
		o(&base)
	}
	return base
}

// Validate checks that the Options fields contain acceptable values.
// It returns an error describing the first invalid field encountered.
func (o Options) Validate() error {
	if o.Title == "" {
		return fmt.Errorf("renderer: Title must not be empty")
	}
	if o.MaxResources < 0 {
		return fmt.Errorf("renderer: MaxResources must be >= 0, got %d", o.MaxResources)
	}
	return nil
}
