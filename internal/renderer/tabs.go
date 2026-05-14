package renderer

import "sort"

// Tab represents a single tab in the tabbed navigation view.
type Tab struct {
	// Label is the human-readable tab heading.
	Label string
	// Anchor is the HTML id used to reference the tab panel.
	Anchor string
	// Count is the number of resources represented by this tab.
	Count int
	// Active indicates whether this tab should be shown on initial render.
	Active bool
}

// TabsData holds all data required to render the tab bar.
type TabsData struct {
	Enabled  bool
	Position string // "top" or "left"
	Tabs     []Tab
}

// buildTabs constructs the TabsData from the current analysis and options.
// Each distinct resource status (managed, unmanaged, deleted, drifted) becomes
// its own tab. Tabs with zero resources are omitted unless they are the only
// available tab.
func buildTabs(a *Analysis, o *Options) TabsData {
	if !o.TabsEnabled {
		return TabsData{}
	}

	position := o.TabsPosition
	if position == "" {
		position = "top"
	}

	candidates := []Tab{
		{Label: "Managed", Anchor: "tab-managed", Count: len(a.Managed)},
		{Label: "Unmanaged", Anchor: "tab-unmanaged", Count: len(a.Unmanaged)},
		{Label: "Deleted", Anchor: "tab-deleted", Count: len(a.Deleted)},
		{Label: "Drifted", Anchor: "tab-drifted", Count: len(a.Drifted)},
	}

	var tabs []Tab
	for _, c := range candidates {
		if c.Count > 0 {
			tabs = append(tabs, c)
		}
	}

	if len(tabs) == 0 {
		return TabsData{Enabled: true, Position: position}
	}

	// Mark the default-active tab.
	activated := false
	for i := range tabs {
		if tabs[i].Label == o.TabsDefaultActive {
			tabs[i].Active = true
			activated = true
			break
		}
	}
	if !activated {
		tabs[0].Active = true
	}

	sort.SliceStable(tabs, func(i, j int) bool {
		return tabs[i].Label < tabs[j].Label
	})

	return TabsData{Enabled: true, Position: position, Tabs: tabs}
}
