package renderer

// WithOutputFormat sets the desired output format for the exported report.
// Supported values are "html" (default) and "json".
// Unknown formats are silently ignored and the default is preserved.
func WithOutputFormat(format string) Option {
	return func(o *Options) {
		switch format {
		case "html", "json":
			o.OutputFormat = format
		}
	}
}

// WithFilename sets the suggested filename (without extension) used when
// exporting the report to disk. An empty value is ignored.
func WithFilename(name string) Option {
	return func(o *Options) {
		if name != "" {
			o.Filename = name
		}
	}
}

// WithEmbedAssets controls whether CSS/JS assets are inlined into the HTML
// output (true) or referenced as external files (false).
func WithEmbedAssets(embed bool) Option {
	return func(o *Options) {
		o.EmbedAssets = embed
	}
}
