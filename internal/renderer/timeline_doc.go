// Package renderer provides HTML rendering capabilities for driftctl JSON reports.
//
// # Timeline
//
// The timeline module builds a chronological sequence of drift events derived
// from the analysis results. Each event captures a resource kind (missing,
// unmanaged, or changed), its type, and its ID so that auditors can trace
// infrastructure changes over time.
//
// Use [WithTimeline] to enable or disable the timeline section in the rendered
// report, and [WithTimelineLimit] to cap the number of events displayed.
//
// When the limit is 0 (the default), all events are included. Negative values
// are silently clamped to 0.
//
// Events are sorted first by resource type, then by resource ID, providing a
// stable and predictable ordering across repeated report generations.
package renderer
