// Package renderer — tabs module
//
// # Tabbed Navigation
//
// The tabs feature organises the four resource buckets (Managed, Unmanaged,
// Deleted, Drifted) into a tabbed interface so that readers can focus on one
// category at a time without scrolling through the entire report.
//
// # Options
//
//   - [WithTabs] — enable or disable the tab bar (default: false).
//   - [WithTabsDefaultActive] — choose which tab label is shown on first load.
//     Falls back to the first non-empty tab when the label is not found.
//   - [WithTabsPosition] — "top" (default) renders a horizontal tab bar above
//     the content; "left" renders a vertical sidebar.
//
// # Behaviour
//
// Tabs whose resource bucket is empty are silently omitted from the rendered
// output.  If all buckets are empty the tab bar is rendered but contains no
// tab items.  Tabs are sorted alphabetically by label to keep the order
// deterministic across runs.
package renderer
