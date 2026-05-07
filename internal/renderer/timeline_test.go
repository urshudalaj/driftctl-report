package renderer

import (
	"testing"
	"time"
)

func makeTimelineAnalysis() Analysis {
	now := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	return Analysis{
		Date: now,
		Options: RenderOptions{ShowTimeline: true},
		Summary: Summary{
			Unmanaged: []Resource{
				{Type: "aws_s3_bucket", ID: "my-bucket"},
			},
			Missing: []Resource{
				{Type: "aws_iam_role", ID: "my-role"},
			},
			Drifted: []DiffResource{
				{Resource: Resource{Type: "aws_lambda_function", ID: "my-fn"}},
			},
		},
	}
}

func TestBuildTimeline_Disabled(t *testing.T) {
	a := makeTimelineAnalysis()
	a.Options.ShowTimeline = false
	td := buildTimeline(a, 0)
	if td.Enabled {
		t.Fatal("expected timeline to be disabled")
	}
	if len(td.Events) != 0 {
		t.Fatalf("expected 0 events, got %d", len(td.Events))
	}
}

func TestBuildTimeline_AllKinds(t *testing.T) {
	a := makeTimelineAnalysis()
	td := buildTimeline(a, 0)
	if !td.Enabled {
		t.Fatal("expected timeline to be enabled")
	}
	if len(td.Events) != 3 {
		t.Fatalf("expected 3 events, got %d", len(td.Events))
	}
}

func TestBuildTimeline_MaxEvents(t *testing.T) {
	a := makeTimelineAnalysis()
	td := buildTimeline(a, 2)
	if len(td.Events) != 2 {
		t.Fatalf("expected 2 events after limit, got %d", len(td.Events))
	}
}

func TestBuildTimeline_SortedByTypeAndID(t *testing.T) {
	a := makeTimelineAnalysis()
	td := buildTimeline(a, 0)
	for i := 1; i < len(td.Events); i++ {
		prev := td.Events[i-1]
		curr := td.Events[i]
		if prev.Type > curr.Type {
			t.Errorf("events not sorted by type: %s > %s", prev.Type, curr.Type)
		}
	}
}

func TestBuildTimeline_LabelFormat(t *testing.T) {
	a := makeTimelineAnalysis()
	td := buildTimeline(a, 0)
	for _, ev := range td.Events {
		expected := ev.Type + " / " + ev.ID
		if ev.Label != expected {
			t.Errorf("expected label %q, got %q", expected, ev.Label)
		}
	}
}

func TestBuildTimeline_ZeroMaxMeansUnlimited(t *testing.T) {
	a := makeTimelineAnalysis()
	td := buildTimeline(a, 0)
	if len(td.Events) != 3 {
		t.Fatalf("expected all 3 events, got %d", len(td.Events))
	}
}
