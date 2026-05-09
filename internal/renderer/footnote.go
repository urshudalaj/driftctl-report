package renderer

// FootnoteEntry represents a single footnote line in the report.
type FootnoteEntry struct {
	Index int
	Text  string
}

// FootnoteData holds the rendered footnote section.
type FootnoteData struct {
	Enabled   bool
	Footnotes []FootnoteEntry
	Truncated bool
	Total     int
}

// buildFootnotes constructs the footnote section from the current options.
// If no footnotes are configured, the section is disabled.
func buildFootnotes(opts Options) FootnoteData {
	if len(opts.Footnotes) == 0 {
		return FootnoteData{}
	}

	src := opts.Footnotes
	total := len(src)
	truncated := false

	if opts.FootnoteLimit > 0 && len(src) > opts.FootnoteLimit {
		src = src[:opts.FootnoteLimit]
		truncated = true
	}

	entries := make([]FootnoteEntry, len(src))
	for i, text := range src {
		entries[i] = FootnoteEntry{Index: i + 1, Text: text}
	}

	return FootnoteData{
		Enabled:   true,
		Footnotes: entries,
		Truncated: truncated,
		Total:     total,
	}
}
