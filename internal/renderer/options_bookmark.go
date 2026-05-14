package renderer

// WithBookmarks enables or disables the bookmarks panel in the report.
// When enabled, users can mark individual resources for follow-up.
func WithBookmarks(enabled bool) Option {
	return func(o *Options) {
		o.BookmarksEnabled = enabled
	}
}

// WithBookmarkIDs pre-populates the bookmarks panel with the given resource IDs.
// Empty strings are silently ignored. Duplicate IDs are deduplicated.
func WithBookmarkIDs(ids ...string) Option {
	return func(o *Options) {
		seen := make(map[string]struct{}, len(ids))
		for _, id := range ids {
			if id == "" {
				continue
			}
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = struct{}{}
			o.BookmarkIDs = append(o.BookmarkIDs, id)
		}
	}
}

// WithBookmarkLabel sets the display label shown in the bookmarks panel header.
// An empty value is ignored and the default label is retained.
func WithBookmarkLabel(label string) Option {
	return func(o *Options) {
		if label != "" {
			o.BookmarkLabel = label
		}
	}
}
