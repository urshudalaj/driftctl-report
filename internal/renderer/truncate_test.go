package renderer

import (
	"strings"
	"testing"
)

func makeTruncateAnalysis() Analysis {
	return Analysis{
		Managed:  []Resource{{ResourceID: "aws_s3_bucket.my-very-long-bucket-name-that-exceeds-limit", ResourceType: "aws_s3_bucket"}},
		Unmanaged: []Resource{{ResourceID: "short", ResourceType: "aws_iam_role"}},
		Deleted:  []Resource{{ResourceID: "aws_ec2_instance.i-0123456789abcdef0", ResourceType: "aws_instance"}},
		Drifted:  []DriftedResource{{ResourceID: "aws_lambda.fn", ResourceType: "aws_lambda_function"}},
	}
}

func TestBuildTruncated_Disabled(t *testing.T) {
	opts := DefaultOptions()
	opts.TruncateIDs = false
	res := buildTruncated(makeTruncateAnalysis(), opts)
	if res.Enabled {
		t.Fatal("expected Enabled=false")
	}
	if len(res.Entries) != 0 {
		t.Fatalf("expected 0 entries, got %d", len(res.Entries))
	}
}

func TestBuildTruncated_EntryCount(t *testing.T) {
	opts := DefaultOptions()
	WithTruncateIDs(true)(&opts)
	res := buildTruncated(makeTruncateAnalysis(), opts)
	if !res.Enabled {
		t.Fatal("expected Enabled=true")
	}
	// 4 unique IDs across all buckets
	if len(res.Entries) != 4 {
		t.Fatalf("expected 4 entries, got %d", len(res.Entries))
	}
}

func TestBuildTruncated_LongIDIsTruncated(t *testing.T) {
	opts := DefaultOptions()
	WithTruncateIDs(true)(&opts)
	WithTruncateIDsLength(20)(&opts)
	WithTruncateIDsSuffix("...")(&opts)
	res := buildTruncated(makeTruncateAnalysis(), opts)

	var found *TruncateEntry
	for i := range res.Entries {
		if strings.HasPrefix(res.Entries[i].Original, "aws_s3_bucket") {
			found = &res.Entries[i]
			break
		}
	}
	if found == nil {
		t.Fatal("long entry not found")
	}
	if !found.Truncated {
		t.Error("expected Truncated=true for long ID")
	}
	if !strings.HasSuffix(found.Display, "...") {
		t.Errorf("expected display to end with '...', got %q", found.Display)
	}
}

func TestBuildTruncated_ShortIDNotTruncated(t *testing.T) {
	opts := DefaultOptions()
	WithTruncateIDs(true)(&opts)
	WithTruncateIDsLength(40)(&opts)
	res := buildTruncated(makeTruncateAnalysis(), opts)

	for _, e := range res.Entries {
		if e.Original == "short" {
			if e.Truncated {
				t.Error("short ID should not be truncated")
			}
			if e.Display != "short" {
				t.Errorf("expected display=short, got %q", e.Display)
			}
			return
		}
	}
	t.Fatal("short entry not found")
}

func TestBuildTruncated_DeduplicatesIDs(t *testing.T) {
	a := Analysis{
		Managed:  []Resource{{ResourceID: "dup", ResourceType: "t"}},
		Unmanaged: []Resource{{ResourceID: "dup", ResourceType: "t"}},
	}
	opts := DefaultOptions()
	WithTruncateIDs(true)(&opts)
	res := buildTruncated(a, opts)
	if len(res.Entries) != 1 {
		t.Fatalf("expected 1 deduplicated entry, got %d", len(res.Entries))
	}
}

func TestTruncateID_DefaultSuffix(t *testing.T) {
	long := strings.Repeat("a", 50)
	e := truncateID(long, 10, "\u2026")
	if !e.Truncated {
		t.Error("expected truncated")
	}
	if !strings.HasSuffix(e.Display, "\u2026") {
		t.Errorf("expected ellipsis suffix, got %q", e.Display)
	}
}
