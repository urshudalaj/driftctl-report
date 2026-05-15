package renderer

// RTLConfig holds the resolved right-to-left rendering configuration
// that is passed to the HTML template.
type RTLConfig struct {
	// Enabled indicates whether RTL mode is active.
	Enabled bool
	// Dir is the HTML dir attribute value, either "rtl" or "ltr".
	Dir string
	// Lang is the HTML lang attribute value.
	Lang string
}

// buildRTL resolves the RTL configuration from the given options.
// When RTL is disabled it returns a default LTR config so the template
// always has a valid, non-nil value to render.
func buildRTL(o Options) RTLConfig {
	if !o.RTL {
		return RTLConfig{
			Enabled: false,
			Dir:     "ltr",
			Lang:    "en",
		}
	}

	dir := o.RTLDir
	if dir == "" {
		dir = "rtl"
	}

	lang := o.RTLLang
	if lang == "" {
		lang = "ar"
	}

	return RTLConfig{
		Enabled: true,
		Dir:     dir,
		Lang:    lang,
	}
}
