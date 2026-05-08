package renderer

// WithAnnotations attaches a slice of Annotation values to the render options.
// Annotations are matched against resource IDs present in the report and
// surfaced in the HTML output as labels or notes alongside each resource.
func WithAnnotations(annotations []Annotation) Option {
	return func(o *Options) {
		if len(annotations) == 0 {
			return
		}
		// Deduplicate by ResourceID; last one wins.
		seen := make(map[string]int, len(annotations))
		merged := make([]Annotation, 0, len(annotations))
		for _, a := range annotations {
			if a.ResourceID == "" {
				continue
			}
			if idx, exists := seen[a.ResourceID]; exists {
				merged[idx] = a
			} else {
				seen[a.ResourceID] = len(merged)
				merged = append(merged, a)
			}
		}
		o.Annotations = merged
	}
}
