package renderer

// WithVersion sets the tool version string embedded in the report metadata.
// An empty value is ignored; the default "dev" label is preserved.
func WithVersion(v string) Option {
	return func(o *Options) {
		if v != "" {
			o.Version = v
		}
	}
}

// WithHostname sets the hostname embedded in the report metadata.
// An empty value is ignored; the default "unknown" label is preserved.
func WithHostname(h string) Option {
	return func(o *Options) {
		if h != "" {
			o.Hostname = h
		}
	}
}

// WithInputFile records the source file path in the report metadata so readers
// can trace the report back to the original driftctl JSON output.
func WithInputFile(path string) Option {
	return func(o *Options) {
		o.InputFile = path
	}
}
