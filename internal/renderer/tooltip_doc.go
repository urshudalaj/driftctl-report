// Package renderer — tooltip module.
//
// # Tooltip
//
// The tooltip feature adds lightweight hover popovers to resource entries
// rendered in the HTML report. When a user hovers over a resource ID, a
// small popover appears showing the resource type and its drift status
// (managed, unmanaged, or deleted).
//
// # Options
//
//   - [WithTooltip]         — enable/disable the feature (default: false).
//   - [WithTooltipPlacement] — preferred popover direction: "top" (default),
//     "bottom", "left", or "right". Unknown values are ignored.
//   - [WithTooltipMaxWidth]  — maximum popover width in pixels. Values <= 0
//     are ignored and the previous value is preserved.
//
// # Defaults
//
// Tooltips are disabled by default to keep the report lightweight.
// Enable them explicitly via [WithTooltip](true).
package renderer
