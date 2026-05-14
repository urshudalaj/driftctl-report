package renderer

import "testing"

func makeBookmarkAnalysis() *Analysis {
	return &Analysis{
		Managed: []Resource{
			{ResourceID: "vpc-aaa", ResourceType: "aws_vpc"},
			{ResourceID: "sg-bbb", ResourceType: "aws_security_group"},
		},
		Unmanaged: []Resource{
			{ResourceID: "bucket-ccc", ResourceType: "aws_s3_bucket"},
		},
		Deleted: []Resource{
			{ResourceID: "igw-ddd", ResourceType: "aws_internet_gateway"},
		},
	}
}

func TestBuildBookmarks_Disabled(t *testing.T) {
	o := DefaultOptions()
	data := buildBookmarks(makeBookmarkAnalysis(), o)
	if data.Enabled {
		t.Fatal("expected bookmarks to be disabled by default")
	}
	if len(data.Entries) != 0 {
		t.Fatalf("expected no entries when disabled, got %d", len(data.Entries))
	}
}

func TestBuildBookmarks_EnabledNoIDs(t *testing.T) {
	o := DefaultOptions()
	WithBookmarks(true)(&o)
	data := buildBookmarks(makeBookmarkAnalysis(), o)
	if !data.Enabled {
		t.Fatal("expected bookmarks to be enabled")
	}
	if len(data.Entries) != 0 {
		t.Fatalf("expected 0 entries with no bookmark IDs, got %d", len(data.Entries))
	}
}

func TestBuildBookmarks_MatchesAcrossAllBuckets(t *testing.T) {
	o := DefaultOptions()
	WithBookmarks(true)(&o)
	WithBookmarkIDs("vpc-aaa", "bucket-ccc", "igw-ddd")(&o)
	data := buildBookmarks(makeBookmarkAnalysis(), o)
	if len(data.Entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(data.Entries))
	}
}

func TestBuildBookmarks_DefaultLabel(t *testing.T) {
	o := DefaultOptions()
	WithBookmarks(true)(&o)
	data := buildBookmarks(makeBookmarkAnalysis(), o)
	if data.Label != "Bookmarks" {
		t.Fatalf("expected default label 'Bookmarks', got %q", data.Label)
	}
}

func TestBuildBookmarks_CustomLabel(t *testing.T) {
	o := DefaultOptions()
	WithBookmarks(true)(&o)
	WithBookmarkLabel("Flagged Resources")(&o)
	data := buildBookmarks(makeBookmarkAnalysis(), o)
	if data.Label != "Flagged Resources" {
		t.Fatalf("expected custom label, got %q", data.Label)
	}
}

func TestBuildBookmarks_SortedByTypeAndID(t *testing.T) {
	o := DefaultOptions()
	WithBookmarks(true)(&o)
	WithBookmarkIDs("igw-ddd", "sg-bbb", "vpc-aaa")(&o)
	data := buildBookmarks(makeBookmarkAnalysis(), o)
	if len(data.Entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(data.Entries))
	}
	if data.Entries[0].Type > data.Entries[1].Type {
		t.Error("entries not sorted by type ascending")
	}
}

func TestBuildBookmarks_UnknownIDsIgnored(t *testing.T) {
	o := DefaultOptions()
	WithBookmarks(true)(&o)
	WithBookmarkIDs("does-not-exist")(&o)
	data := buildBookmarks(makeBookmarkAnalysis(), o)
	if len(data.Entries) != 0 {
		t.Fatalf("expected 0 entries for unknown IDs, got %d", len(data.Entries))
	}
}
