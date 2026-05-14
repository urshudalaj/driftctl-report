package renderer

import "testing"

func makeCopyAnalysis() Analysis {
	return Analysis{
		Managed:   []Resource{{ResourceID: "aws_s3_bucket.logs", ResourceType: "aws_s3_bucket"}},
		Unmanaged: []Resource{{ResourceID: "aws_iam_role.admin", ResourceType: "aws_iam_role"}},
		Deleted:   []Resource{{ResourceID: "aws_vpc.old", ResourceType: "aws_vpc"}},
	}
}

func TestBuildCopyData_Disabled(t *testing.T) {
	o := DefaultOptions()
	// CopyButton defaults to false
	data := buildCopyData(makeCopyAnalysis(), o)
	if data.Enabled {
		t.Fatal("expected Enabled=false when CopyButton is off")
	}
	if len(data.Entries) != 0 {
		t.Fatalf("expected 0 entries, got %d", len(data.Entries))
	}
}

func TestBuildCopyData_Enabled_EntryCount(t *testing.T) {
	o := DefaultOptions()
	WithCopyButton(true)(&o)
	data := buildCopyData(makeCopyAnalysis(), o)
	if !data.Enabled {
		t.Fatal("expected Enabled=true")
	}
	// 1 managed + 1 unmanaged + 1 deleted = 3
	if len(data.Entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(data.Entries))
	}
}

func TestBuildCopyData_DefaultLabelAndFeedback(t *testing.T) {
	o := DefaultOptions()
	WithCopyButton(true)(&o)
	data := buildCopyData(makeCopyAnalysis(), o)
	for _, e := range data.Entries {
		if e.Label == "" {
			t.Error("expected non-empty default Label")
		}
		if e.Feedback == "" {
			t.Error("expected non-empty default Feedback")
		}
	}
}

func TestBuildCopyData_CustomLabelAndFeedback(t *testing.T) {
	o := DefaultOptions()
	WithCopyButton(true)(&o)
	WithCopyButtonLabel("Copy")(&o)
	WithCopyButtonFeedback("Done")(&o)
	data := buildCopyData(makeCopyAnalysis(), o)
	for _, e := range data.Entries {
		if e.Label != "Copy" {
			t.Errorf("expected Label='Copy', got %q", e.Label)
		}
		if e.Feedback != "Done" {
			t.Errorf("expected Feedback='Done', got %q", e.Feedback)
		}
	}
}

func TestBuildCopyData_EscapedID(t *testing.T) {
	a := Analysis{
		Managed: []Resource{{ResourceID: "<script>alert(1)</script>", ResourceType: "aws_s3_bucket"}},
	}
	o := DefaultOptions()
	WithCopyButton(true)(&o)
	data := buildCopyData(a, o)
	if len(data.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(data.Entries))
	}
	got := data.Entries[0].EscapedID
	if got == data.Entries[0].ID {
		t.Error("expected EscapedID to differ from raw ID for HTML-unsafe input")
	}
}

func TestWithCopyButtonLabel_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithCopyButtonLabel("Original")(&o)
	WithCopyButtonLabel("")(&o)
	if o.CopyButtonLabel != "Original" {
		t.Errorf("expected label to remain 'Original', got %q", o.CopyButtonLabel)
	}
}

func TestWithCopyButtonFeedback_EmptyIgnored(t *testing.T) {
	o := DefaultOptions()
	WithCopyButtonFeedback("Yes!")(&o)
	WithCopyButtonFeedback("")(&o)
	if o.CopyButtonFeedback != "Yes!" {
		t.Errorf("expected feedback to remain 'Yes!', got %q", o.CopyButtonFeedback)
	}
}
