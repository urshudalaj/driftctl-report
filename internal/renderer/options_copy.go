package renderer

// WithCopyButton enables or disables the copy-to-clipboard button that appears
// next to resource IDs in the rendered report.
func WithCopyButton(enabled bool) Option {
	return func(o *Options) {
		o.CopyButton = enabled
	}
}

// WithCopyButtonLabel sets the accessible label used for the copy button.
// If label is empty, the option is ignored and the existing label is kept.
func WithCopyButtonLabel(label string) Option {
	return func(o *Options) {
		if label == "" {
			return
		}
		o.CopyButtonLabel = label
	}
}

// WithCopyButtonFeedback sets the short text shown briefly after a successful
// copy action (e.g. "Copied!"). If text is empty, the option is ignored.
func WithCopyButtonFeedback(text string) Option {
	return func(o *Options) {
		if text == "" {
			return
		}
		o.CopyButtonFeedback = text
	}
}
