package renderer

import (
	"strings"

	"github.com/snyk/driftctl/pkg/analyser"
)

// DiffEntry represents a single attribute change between IaC and real state.
type DiffEntry struct {
	ResourceID   string
	ResourceType string
	Attribute    string
	From         string
	To           string
	ChangeType   string // "updated", "added", "deleted"
}

// DiffSection holds all diff entries for template rendering.
type DiffSection struct {
	Enabled bool
	Entries []DiffEntry
	Total   int
	Trunc   bool
}

// buildDiff extracts attribute-level diffs from the analysis result.
// It returns a DiffSection ready for template consumption.
func buildDiff(analysis analyser.Analysis, opts Options) DiffSection {
	if !opts.ShowDiff {
		return DiffSection{}
	}

	var entries []DiffEntry

	for _, res := range analysis.Differences() {
		for attr, change := range res.Changelog {
			ct := classifyChange(change.Type)
			entries = append(entries, DiffEntry{
				ResourceID:   res.Res.ResourceId(),
				ResourceType: res.Res.ResourceType(),
				Attribute:    attr,
				From:         formatValue(change.From),
				To:           formatValue(change.To),
				ChangeType:   ct,
			})
		}
	}

	sortDiffEntries(entries)

	total := len(entries)
	trunc := false
	if opts.DiffLimit > 0 && total > opts.DiffLimit {
		entries = entries[:opts.DiffLimit]
		trunc = true
	}

	return DiffSection{
		Enabled: true,
		Entries: entries,
		Total:   total,
		Trunc:   trunc,
	}
}

func classifyChange(t string) string {
	switch strings.ToLower(t) {
	case "create":
		return "added"
	case "delete":
		return "deleted"
	default:
		return "updated"
	}
}

func formatValue(v interface{}) string {
	if v == nil {
		return "<nil>"
	}
	return strings.TrimSpace(strings.ReplaceAll(fmt.Sprintf("%v", v), "\n", " "))
}

func sortDiffEntries(entries []DiffEntry) {
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].ResourceType != entries[j].ResourceType {
			return entries[i].ResourceType < entries[j].ResourceType
		}
		if entries[i].ResourceID != entries[j].ResourceID {
			return entries[i].ResourceID < entries[j].ResourceID
		}
		return entries[i].Attribute < entries[j].Attribute
	})
}
