package renderer

import (
	"sort"

	"github.com/snyk/driftctl/pkg/resource"
)

// SortOrder defines the order in which resources are sorted.
type SortOrder int

const (
	// SortByTypeAsc sorts resources by type ascending.
	SortByTypeAsc SortOrder = iota
	// SortByTypeDesc sorts resources by type descending.
	SortByTypeDesc
	// SortByIDAsc sorts resources by ID ascending.
	SortByIDAsc
	// SortByIDDesc sorts resources by ID descending.
	SortByIDDesc
)

// sortResources sorts a slice of resources in-place according to the given order.
func sortResources(resources []resource.Resource, order SortOrder) {
	sort.SliceStable(resources, func(i, j int) bool {
		return resourceLess(resources[i], resources[j], order)
	})
}

// sortDiffResources sorts a slice of diff resources in-place according to the given order.
func sortDiffResources(resources []resource.ResourceDiff, order SortOrder) {
	sort.SliceStable(resources, func(i, j int) bool {
		if resources[i].Res == nil || resources[j].Res == nil {
			return false
		}
		return resourceLess(*resources[i].Res, *resources[j].Res, order)
	})
}

func resourceLess(a, b resource.Resource, order SortOrder) bool {
	switch order {
	case SortByTypeAsc:
		if a.ResourceType() == b.ResourceType() {
			return a.ResourceId() < b.ResourceId()
		}
		return a.ResourceType() < b.ResourceType()
	case SortByTypeDesc:
		if a.ResourceType() == b.ResourceType() {
			return a.ResourceId() > b.ResourceId()
		}
		return a.ResourceType() > b.ResourceType()
	case SortByIDDesc:
		return a.ResourceId() > b.ResourceId()
	default: // SortByIDAsc
		return a.ResourceId() < b.ResourceId()
	}
}
