package renderer

// CollapsibleSection represents a single collapsible group in the report.
type CollapsibleSection struct {
	// ID is a unique HTML-safe identifier for the section.
	ID string
	// Label is the visible heading text for the toggle control.
	Label string
	// Open indicates whether the section should be rendered in the open state.
	Open bool
	// Animated indicates whether the toggle should use a CSS transition.
	Animated bool
	// ItemCount is the number of items inside the section.
	ItemCount int
}

// CollapsibleData holds all collapsible sections derived from the report.
type CollapsibleData struct {
	Enabled  bool
	Sections []CollapsibleSection
}

// buildCollapsible constructs CollapsibleData from the rendered resource groups.
// Each unique resource type becomes one collapsible section.
func buildCollapsible(o Options, types []string, counts map[string]int) CollapsibleData {
	if !o.Collapsible {
		return CollapsibleData{}
	}

	sections := make([]CollapsibleSection, 0, len(types))
	for _, t := range types {
		sections = append(sections, CollapsibleSection{
			ID:        tocAnchor(t),
			Label:     t,
			Open:      o.CollapsibleDefaultOpen,
			Animated:  o.CollapsibleAnimated,
			ItemCount: counts[t],
		})
	}

	return CollapsibleData{
		Enabled:  true,
		Sections: sections,
	}
}
