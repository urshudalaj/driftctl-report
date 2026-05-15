package renderer

import "fmt"

// HeadingData holds the resolved heading configuration used by the template.
type HeadingData struct {
	// Enabled reports whether the heading block should be rendered.
	Enabled bool
	// Text is the heading string shown to the user.
	Text string
	// Tag is the HTML element tag, e.g. "h1" or "h2".
	Tag string
	// Level is the numeric heading level (1–6).
	Level int
}

// buildHeading constructs a HeadingData value from the current Options.
// If no heading text has been configured, or visibility is explicitly disabled,
// Enabled will be false and the template should skip the element entirely.
func buildHeading(o Options) HeadingData {
	if o.HeadingText == "" || !o.HeadingVisible {
		return HeadingData{Enabled: false}
	}

	level := o.HeadingLevel
	if level < 1 || level > 6 {
		level = 1
	}

	return HeadingData{
		Enabled: true,
		Text:    o.HeadingText,
		Tag:     fmt.Sprintf("h%d", level),
		Level:   level,
	}
}

// IsValid reports whether the HeadingData is in a renderable state.
// A heading is considered valid when it is enabled and contains non-empty text
// with a level in the accepted HTML range (1–6).
func (h HeadingData) IsValid() bool {
	return h.Enabled && h.Text != "" && h.Level >= 1 && h.Level <= 6
}
