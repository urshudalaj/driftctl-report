package renderer

import (
	"regexp"
	"strings"
)

// HighlightResult holds a text fragment with an optional highlight flag.
type HighlightResult struct {
	Text        string
	Highlighted bool
}

// highlightQuery splits text into fragments, marking substrings that match
// query (case-insensitive) as highlighted. Returns a single non-highlighted
// fragment when query is empty.
func highlightQuery(text, query string) []HighlightResult {
	if query == "" {
		return []HighlightResult{{Text: text, Highlighted: false}}
	}

	pattern := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(query))
	indices := pattern.FindAllStringIndex(text, -1)
	if len(indices) == 0 {
		return []HighlightResult{{Text: text, Highlighted: false}}
	}

	var results []HighlightResult
	prev := 0
	for _, loc := range indices {
		if loc[0] > prev {
			results = append(results, HighlightResult{Text: text[prev:loc[0]], Highlighted: false})
		}
		results = append(results, HighlightResult{Text: text[loc[0]:loc[1]], Highlighted: true})
		prev = loc[1]
	}
	if prev < len(text) {
		results = append(results, HighlightResult{Text: text[prev:], Highlighted: false})
	}
	return results
}

// highlightHTML renders highlight results as an HTML string, wrapping
// highlighted fragments in <mark> tags.
func highlightHTML(text, query string) string {
	frags := highlightQuery(text, query)
	var sb strings.Builder
	for _, f := range frags {
		if f.Highlighted {
			sb.WriteString("<mark>")
			sb.WriteString(f.Text)
			sb.WriteString("</mark>")
		} else {
			sb.WriteString(f.Text)
		}
	}
	return sb.String()
}
