package renderer

// WithWatermark sets a custom watermark text displayed in the report footer.
// An empty string disables the watermark.
func WithWatermark(text string) Option {
	return func(o *Options) {
		o.WatermarkText = text
	}
}

// WithWatermarkURL sets a URL that the watermark text links to.
// Has no effect if WatermarkText is empty.
func WithWatermarkURL(url string) Option {
	return func(o *Options) {
		o.WatermarkURL = url
	}
}

// WithWatermarkPosition sets the position of the watermark in the report.
// Accepted values: "footer" (default), "header", "both".
// Unknown values are silently ignored.
func WithWatermarkPosition(pos string) Option {
	return func(o *Options) {
		switch pos {
		case "footer", "header", "both":
			o.WatermarkPosition = pos
		}
	}
}
