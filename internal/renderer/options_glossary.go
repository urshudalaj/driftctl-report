package renderer

// WithGlossary enables or disables the glossary section in the report.
// The glossary provides human-readable explanations for drift statuses and
// resource types found in the report.
func WithGlossary(enabled bool) Option {
	return func(o *Options) {
		o.GlossaryEnabled = enabled
	}
}

// WithGlossaryTerms appends custom term/definition pairs to the glossary.
// Keys are terms (e.g. "unmanaged") and values are their definitions.
// Empty keys or values are silently ignored.
func WithGlossaryTerms(terms map[string]string) Option {
	return func(o *Options) {
		if o.GlossaryTerms == nil {
			o.GlossaryTerms = make(map[string]string)
		}
		for k, v := range terms {
			if k != "" && v != "" {
				o.GlossaryTerms[k] = v
			}
		}
	}
}

// WithGlossaryMaxItems sets the maximum number of terms shown in the glossary.
// A value of zero means unlimited. Negative values are clamped to zero.
func WithGlossaryMaxItems(n int) Option {
	return func(o *Options) {
		if n < 0 {
			n = 0
		}
		o.GlossaryMaxItems = n
	}
}
