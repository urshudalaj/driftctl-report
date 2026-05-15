package renderer

import "unicode/utf8"

// TruncateResult holds the output of a truncation operation.
type TruncateResult struct {
	Enabled  bool
	Entries  []TruncateEntry
}

// TruncateEntry maps an original resource ID to its (possibly truncated) display form.
type TruncateEntry struct {
	Original  string
	Display   string
	Truncated bool
}

// buildTruncated produces a TruncateResult for all resource IDs found in the
// analysis, applying the truncation settings from opts.
func buildTruncated(a Analysis, opts Options) TruncateResult {
	if !opts.TruncateIDs {
		return TruncateResult{Enabled: false}
	}

	maxLen := opts.TruncateIDsLen
	if maxLen <= 0 {
		maxLen = 40
	}
	suffix := opts.TruncateIDsSuffix
	if suffix == "" {
		suffix = "\u2026"
	}

	seen := make(map[string]struct{})
	var entries []TruncateEntry

	add := func(id string) {
		if _, ok := seen[id]; ok {
			return
		}
		seen[id] = struct{}{}
		entry := truncateID(id, maxLen, suffix)
		entries = append(entries, entry)
	}

	for _, r := range a.Managed {
		add(r.ResourceID)
	}
	for _, r := range a.Unmanaged {
		add(r.ResourceID)
	}
	for _, r := range a.Deleted {
		add(r.ResourceID)
	}
	for _, r := range a.Drifted {
		add(r.ResourceID)
	}

	return TruncateResult{Enabled: true, Entries: entries}
}

// truncateID shortens id to at most maxLen runes, appending suffix if cut.
func truncateID(id string, maxLen int, suffix string) TruncateEntry {
	if utf8.RuneCountInString(id) <= maxLen {
		return TruncateEntry{Original: id, Display: id, Truncated: false}
	}
	runes := []rune(id)
	cut := string(runes[:maxLen]) + suffix
	return TruncateEntry{Original: id, Display: cut, Truncated: true}
}
