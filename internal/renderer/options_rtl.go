package renderer

// WithRTL enables right-to-left text direction in the rendered HTML report.
// This is useful for languages such as Arabic, Hebrew, or Persian.
func WithRTL(enabled bool) Option {
	return func(o *Options) {
		o.RTL = enabled
	}
}

// WithRTLLang sets the HTML lang attribute used when RTL mode is active.
// If lang is empty, the option is ignored.
func WithRTLLang(lang string) Option {
	return func(o *Options) {
		if lang == "" {
			return
		}
		o.RTLLang = lang
	}
}

// WithRTLDir overrides the dir attribute value (e.g. "rtl" or "ltr").
// If dir is neither "rtl" nor "ltr" the option is ignored.
func WithRTLDir(dir string) Option {
	return func(o *Options) {
		if dir != "rtl" && dir != "ltr" {
			return
		}
		o.RTLDir = dir
	}
}
