package renderer

// TableOfContents represents a navigable section index for the HTML report.
// Each entry corresponds to a resource type section rendered on the page.
type TableOfContents struct {
	Entries []TOCEntry
}

// TOCEntry is a single item in the table of contents.
type TOCEntry struct {
	Label   string
	Anchor  string
	Count   int
	Drifted bool
}

// buildTOC constructs a TableOfContents from the resource type buckets
// present in the report data passed to the renderer.
func buildTOC(types []string, counts map[string]int, driftedTypes map[string]bool) TableOfContents {
	entries := make([]TOCEntry, 0, len(types))
	for _, t := range types {
		entries = append(entries, TOCEntry{
			Label:   t,
			Anchor:  tocAnchor(t),
			Count:   counts[t],
			Drifted: driftedTypes[t],
		})
	}
	return TableOfContents{Entries: entries}
}

// tocAnchor converts a resource type string into a safe HTML anchor id
// by replacing non-alphanumeric characters with hyphens.
func tocAnchor(resourceType string) string {
	out := make([]byte, len(resourceType))
	for i := 0; i < len(resourceType); i++ {
		c := resourceType[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			out[i] = c
		} else {
			out[i] = '-'
		}
	}
	return string(out)
}
