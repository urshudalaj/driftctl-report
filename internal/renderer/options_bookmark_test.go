package renderer

import "testing"

func TestWithBookmarks_Enables(t *testing.T) {
	o := DefaultOptions()
	WithBookmarks(true)(&o)
	if !o.BookmarksEnabled {
		t.Fatal("expected BookmarksEnabled to be true")
	}
}

func TestWithBookmarks_Disables(t *testing.T) {
	o := DefaultOptions()
	WithBookmarks(true)(&o)
	WithBookmarks(false)(&o)
	if o.BookmarksEnabled {
		t.Fatal("expected BookmarksEnabled to be false")
	}
}

func TestWithBookmarks_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.BookmarksEnabled {
		t.Fatal("expected BookmarksEnabled default to be false")
	}
}

func TestWithBookmarkIDs_SetsIDs(t *testing.T) {
	o := DefaultOptions()
	WithBookmarkIDs("id-1", "id-2")(&o)
	if len(o.BookmarkIDs) != 2 {
		t.Fatalf("expected 2 bookmark IDs, got %d", len(o.BookmarkIDs))
	}
}

func TestWithBookmarkIDs_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithBookmarkIDs("", "id-1", "")(&o)
	if len(o.BookmarkIDs) != 1 {
		t.Fatalf("expected 1 bookmark ID, got %d", len(o.BookmarkIDs))
	}
}

func TestWithBookmarkIDs_Deduplicates(t *testing.T) {
	o := DefaultOptions()
	WithBookmarkIDs("id-1", "id-1", "id-2")(&o)
	if len(o.BookmarkIDs) != 2 {
		t.Fatalf("expected 2 unique IDs, got %d", len(o.BookmarkIDs))
	}
}

func TestWithBookmarkLabel_SetsLabel(t *testing.T) {
	o := DefaultOptions()
	WithBookmarkLabel("My Bookmarks")(&o)
	if o.BookmarkLabel != "My Bookmarks" {
		t.Fatalf("expected label 'My Bookmarks', got %q", o.BookmarkLabel)
	}
}

func TestWithBookmarkLabel_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithBookmarkLabel("Initial")(&o)
	WithBookmarkLabel("")(&o)
	if o.BookmarkLabel != "Initial" {
		t.Fatalf("expected label to remain 'Initial', got %q", o.BookmarkLabel)
	}
}
