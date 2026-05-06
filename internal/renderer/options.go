package renderer

// Options controls report generation behaviour.
type Options struct {
	// Title is displayed in the HTML report header.
	Title string

	// OnlyDrifted limits the report to resources that have drifted
	// (missing or changed) and excludes managed/untracked resources.
	OnlyDrifted bool

	// ResourceType, when non-empty, restricts the report to resources
	// whose type matches this value (e.g. "aws_s3_bucket").
	ResourceType string

	// MaxResources caps the total number of resources shown across all
	// sections. Zero means no cap.
	MaxResources int
}

// DefaultOptions returns an Options value with sensible defaults.
func DefaultOptions() Options {
	return Options{
		Title:        "Drift Report",
		OnlyDrifted:  false,
		ResourceType: "",
		MaxResources: 0,
	}
}

// Option is a functional option that mutates an Options value.
type Option func(*Options)

// WithTitle sets the report title.
func WithTitle(t string) Option {
	return func(o *Options) { o.Title = t }
}

// WithOnlyDrifted enables the drifted-only filter.
func WithOnlyDrifted() Option {
	return func(o *Options) { o.OnlyDrifted = true }
}

// WithResourceType restricts the report to a single resource type.
func WithResourceType(rt string) Option {
	return func(o *Options) { o.ResourceType = rt }
}

// WithMaxResources caps the number of resources displayed.
func WithMaxResources(n int) Option {
	return func(o *Options) { o.MaxResources = n }
}

// apply runs all functional options against o.
func (o *Options) apply(opts []Option) {
	for _, fn := range opts {
		fn(o)
	}
}
