package renderer

import (
	"testing"
)

func TestWithTimeline_Enables(t *testing.T) {
	o := DefaultOptions()
	WithTimeline(true)(o)
	if !o.ShowTimeline {
		t.Error("expected ShowTimeline to be true")
	}
}

func TestWithTimeline_Disables(t *testing.T) {
	o := DefaultOptions()
	o.ShowTimeline = true
	WithTimeline(false)(o)
	if o.ShowTimeline {
		t.Error("expected ShowTimeline to be false")
	}
}

func TestWithTimelineLimit_PositiveValue(t *testing.T) {
	o := DefaultOptions()
	WithTimelineLimit(25)(o)
	if o.TimelineLimit != 25 {
		t.Errorf("expected TimelineLimit 25, got %d", o.TimelineLimit)
	}
}

func TestWithTimelineLimit_ZeroMeansUnlimited(t *testing.T) {
	o := DefaultOptions()
	o.TimelineLimit = 10
	WithTimelineLimit(0)(o)
	if o.TimelineLimit != 0 {
		t.Errorf("expected TimelineLimit 0, got %d", o.TimelineLimit)
	}
}

func TestWithTimelineLimit_NegativeClampedToZero(t *testing.T) {
	o := DefaultOptions()
	WithTimelineLimit(-5)(o)
	if o.TimelineLimit != 0 {
		t.Errorf("expected TimelineLimit clamped to 0, got %d", o.TimelineLimit)
	}
}

func TestWithTimeline_DefaultIsFalse(t *testing.T) {
	o := DefaultOptions()
	if o.ShowTimeline {
		t.Error("expected ShowTimeline default to be false")
	}
}

func TestWithTimelineLimit_DefaultIsZero(t *testing.T) {
	o := DefaultOptions()
	if o.TimelineLimit != 0 {
		t.Errorf("expected TimelineLimit default 0, got %d", o.TimelineLimit)
	}
}
