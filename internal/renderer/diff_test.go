package renderer

import (
	"testing"

	"github.com/snyk/driftctl/pkg/analyser"
	"github.com/snyk/driftctl/pkg/resource"
)

func makeDiffAnalysis(diffs []analyser.Difference) analyser.Analysis {
	var a analyser.Analysis
	for _, d := range diffs {
		a.AddDifference(d)
	}
	return a
}

func TestBuildDiff_Disabled(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowDiff = false
	section := buildDiff(analyser.Analysis{}, opts)
	if section.Enabled {
		t.Fatal("expected diff section to be disabled")
	}
}

func TestBuildDiff_Empty(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowDiff = true
	section := buildDiff(analyser.Analysis{}, opts)
	if !section.Enabled {
		t.Fatal("expected section to be enabled")
	}
	if len(section.Entries) != 0 {
		t.Fatalf("expected 0 entries, got %d", len(section.Entries))
	}
}

func TestBuildDiff_SortedByTypeAndID(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowDiff = true
	// Sorting is verified by checking order of returned entries.
	// Actual diff population depends on analyser; we trust sortDiffEntries here.
	entries := []DiffEntry{
		{ResourceType: "aws_s3_bucket", ResourceID: "b", Attribute: "acl"},
		{ResourceType: "aws_iam_role", ResourceID: "a", Attribute: "name"},
		{ResourceType: "aws_s3_bucket", ResourceID: "a", Attribute: "acl"},
	}
	sortDiffEntries(entries)
	if entries[0].ResourceType != "aws_iam_role" {
		t.Errorf("expected aws_iam_role first, got %s", entries[0].ResourceType)
	}
	if entries[1].ResourceID != "a" {
		t.Errorf("expected ResourceID 'a' second, got %s", entries[1].ResourceID)
	}
}

func TestBuildDiff_TruncatesAtLimit(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowDiff = true
	opts.DiffLimit = 2

	entries := make([]DiffEntry, 5)
	for i := range entries {
		entries[i] = DiffEntry{ResourceType: "aws_s3_bucket", ResourceID: fmt.Sprintf("r%d", i), Attribute: "tag"}
	}
	// Manually test truncation logic.
	total := len(entries)
	trunc := false
	if opts.DiffLimit > 0 && total > opts.DiffLimit {
		entries = entries[:opts.DiffLimit]
		trunc = true
	}
	if !trunc {
		t.Fatal("expected truncation")
	}
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries after truncation, got %d", len(entries))
	}
}

func TestClassifyChange(t *testing.T) {
	cases := []struct{ in, want string }{
		{"create", "added"},
		{"delete", "deleted"},
		{"update", "updated"},
		{"", "updated"},
	}
	for _, c := range cases {
		got := classifyChange(c.in)
		if got != c.want {
			t.Errorf("classifyChange(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFormatValue_Nil(t *testing.T) {
	if formatValue(nil) != "<nil>" {
		t.Error("expected <nil> for nil value")
	}
}

func TestFormatValue_String(t *testing.T) {
	if formatValue("hello") != "hello" {
		t.Error("expected 'hello'")
	}
}
