package renderer

import "sort"

// BookmarkEntry represents a single bookmarked resource shown in the panel.
type BookmarkEntry struct {
	ID   string
	Type string
}

// BookmarkData holds all data required to render the bookmarks panel.
type BookmarkData struct {
	Enabled bool
	Label   string
	Entries []BookmarkEntry
}

// buildBookmarks constructs the BookmarkData from the analysis and options.
// Resources whose IDs appear in BookmarkIDs are included in the panel.
// If BookmarksEnabled is false, an empty disabled struct is returned.
func buildBookmarks(a *Analysis, o Options) BookmarkData {
	if !o.BookmarksEnabled {
		return BookmarkData{}
	}

	label := o.BookmarkLabel
	if label == "" {
		label = "Bookmarks"
	}

	if len(o.BookmarkIDs) == 0 {
		return BookmarkData{Enabled: true, Label: label}
	}

	index := make(map[string]struct{}, len(o.BookmarkIDs))
	for _, id := range o.BookmarkIDs {
		index[id] = struct{}{}
	}

	var entries []BookmarkEntry
	for _, r := range a.Managed {
		if _, ok := index[r.ResourceID]; ok {
			entries = append(entries, BookmarkEntry{ID: r.ResourceID, Type: r.ResourceType})
		}
	}
	for _, r := range a.Unmanaged {
		if _, ok := index[r.ResourceID]; ok {
			entries = append(entries, BookmarkEntry{ID: r.ResourceID, Type: r.ResourceType})
		}
	}
	for _, r := range a.Deleted {
		if _, ok := index[r.ResourceID]; ok {
			entries = append(entries, BookmarkEntry{ID: r.ResourceID, Type: r.ResourceType})
		}
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Type != entries[j].Type {
			return entries[i].Type < entries[j].Type
		}
		return entries[i].ID < entries[j].ID
	})

	return BookmarkData{Enabled: true, Label: label, Entries: entries}
}
