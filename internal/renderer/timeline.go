package renderer

import (
	"sort"
	"time"
)

// TimelineEvent represents a single drift event in chronological order.
type TimelineEvent struct {
	Timestamp time.Time
	Kind      string // "missing", "unmanaged", "changed"
	Type      string
	ID        string
	Label     string
}

// TimelineData holds the ordered list of events for the report timeline section.
type TimelineData struct {
	Events  []TimelineEvent
	Enabled bool
}

// buildTimeline constructs a TimelineData from the analysis, limited to the
// most recent maxEvents entries (0 = unlimited).
func buildTimeline(a Analysis, maxEvents int) TimelineData {
	if !a.Options.ShowTimeline {
		return TimelineData{}
	}

	var events []TimelineEvent

	for _, r := range a.Summary.Unmanaged {
		events = append(events, TimelineEvent{
			Timestamp: a.Date,
			Kind:      "unmanaged",
			Type:      r.Type,
			ID:        r.ID,
			Label:     r.Type + " / " + r.ID,
		})
	}

	for _, r := range a.Summary.Missing {
		events = append(events, TimelineEvent{
			Timestamp: a.Date,
			Kind:      "missing",
			Type:      r.Type,
			ID:        r.ID,
			Label:     r.Type + " / " + r.ID,
		})
	}

	for _, r := range a.Summary.Drifted {
		events = append(events, TimelineEvent{
			Timestamp: a.Date,
			Kind:      "changed",
			Type:      r.Resource.Type,
			ID:        r.Resource.ID,
			Label:     r.Resource.Type + " / " + r.Resource.ID,
		})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].Type != events[j].Type {
			return events[i].Type < events[j].Type
		}
		return events[i].ID < events[j].ID
	})

	if maxEvents > 0 && len(events) > maxEvents {
		events = events[:maxEvents]
	}

	return TimelineData{
		Events:  events,
		Enabled: true,
	}
}
