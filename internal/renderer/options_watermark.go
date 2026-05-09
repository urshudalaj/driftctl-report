package renderer

// WithWatermark sets the watermark text displayed on the report.
// An empty string clears any previously set watermark.
func WithWatermark(text string) Option {
	return func(o *Options) {
		o.WatermarkText = text
	}
}

// WithWatermarkURL sets a URL that the watermark text links to.
// Has no effect if no watermark text is set.
func WithWatermarkURL(url string) Option {
	return func(o *Options) {
		o.WatermarkURL = url
	}
}

// WithWatermarkPosition sets the position of the watermark on the page.
// Accepted values: "top-left", "top-right", "bottom-left", "bottom-right".
// Unknown values are silently ignored and the default is preserved.
func WithWatermarkPosition(pos string) Option {
	valid := map[string]bool{
		"top-left":     true,
		"top-right":    true,
		"bottom-left":  true,
		"bottom-right": true,
	}
	return func(o *Options) {
		if valid[pos] {
			o.WatermarkPosition = pos
		}
	}
}
