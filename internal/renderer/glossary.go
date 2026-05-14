package renderer

import "sort"

// GlossaryEntry holds a single term and its definition for display.
type GlossaryEntry struct {
	Term       string
	Definition string
}

// GlossaryData is the view model passed to the template for the glossary section.
type GlossaryData struct {
	Enabled bool
	Entries []GlossaryEntry
}

// defaultGlossaryTerms provides built-in definitions for common drift concepts.
var defaultGlossaryTerms = map[string]string{
	"managed":   "Resource is tracked and managed by IaC (e.g. Terraform).",
	"unmanaged": "Resource exists in the cloud but is not tracked by any IaC definition.",
	"deleted":   "Resource is defined in IaC but no longer exists in the cloud provider.",
	"drifted":   "Resource is managed by IaC but its live state differs from the desired state.",
	"missing":   "Resource declared in IaC state but absent from the cloud provider.",
}

// buildGlossary constructs the GlossaryData view model from the given options.
// Custom terms supplied via WithGlossaryTerms are merged with the built-in
// defaults; custom entries take precedence over defaults with the same key.
func buildGlossary(opts Options) GlossaryData {
	if !opts.GlossaryEnabled {
		return GlossaryData{}
	}

	merged := make(map[string]string, len(defaultGlossaryTerms)+len(opts.GlossaryTerms))
	for k, v := range defaultGlossaryTerms {
		merged[k] = v
	}
	for k, v := range opts.GlossaryTerms {
		merged[k] = v
	}

	keys := make([]string, 0, len(merged))
	for k := range merged {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	max := opts.GlossaryMaxItems
	if max <= 0 || max > len(keys) {
		max = len(keys)
	}

	entries := make([]GlossaryEntry, 0, max)
	for _, k := range keys[:max] {
		entries = append(entries, GlossaryEntry{Term: k, Definition: merged[k]})
	}

	return GlossaryData{
		Enabled: true,
		Entries: entries,
	}
}
