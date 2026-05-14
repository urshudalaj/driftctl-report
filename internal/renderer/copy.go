package renderer

import "html"

// CopyEntry represents a single resource ID that can be copied to the
// clipboard, together with the display label and feedback text.
type CopyEntry struct {
	// ID is the raw resource identifier (unescaped).
	ID string
	// Label is the accessible aria-label for the copy button.
	Label string
	// Feedback is the transient confirmation text shown after a copy.
	Feedback string
	// EscapedID is the HTML-escaped form of ID, safe for use in attributes.
	EscapedID string
}

// CopyData holds all copy-button entries produced for the report.
type CopyData struct {
	Enabled  bool
	Entries  []CopyEntry
}

// buildCopyData constructs a CopyData value from the analysis and options.
// When the copy button feature is disabled it returns a zero-value struct
// with Enabled set to false so templates can short-circuit rendering.
func buildCopyData(a Analysis, o Options) CopyData {
	if !o.CopyButton {
		return CopyData{}
	}

	label := o.CopyButtonLabel
	if label == "" {
		label = "Copy ID"
	}
	feedback := o.CopyButtonFeedback
	if feedback == "" {
		feedback = "Copied!"
	}

	var entries []CopyEntry
	for _, r := range a.Managed {
		entries = append(entries, copyEntry(r.ResourceID, label, feedback))
	}
	for _, r := range a.Unmanaged {
		entries = append(entries, copyEntry(r.ResourceID, label, feedback))
	}
	for _, r := range a.Deleted {
		entries = append(entries, copyEntry(r.ResourceID, label, feedback))
	}

	return CopyData{
		Enabled: true,
		Entries: entries,
	}
}

func copyEntry(id, label, feedback string) CopyEntry {
	return CopyEntry{
		ID:        id,
		Label:     label,
		Feedback:  feedback,
		EscapedID: html.EscapeString(id),
	}
}
