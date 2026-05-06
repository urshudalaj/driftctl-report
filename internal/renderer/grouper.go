package renderer

import "github.com/snyk/driftctl/pkg/analyser"

// ResourceGroup holds resources grouped under a single resource type label.
type ResourceGroup struct {
	Type      string
	Resources []analyser.Resource
	Count     int
}

// DiffGroup holds diff resources grouped under a single resource type label.
type DiffGroup struct {
	Type      string
	Resources []analyser.DiffResource
	Count     int
}

// groupByType partitions a flat slice of resources into ResourceGroups,
// preserving the order in which each type is first encountered.
func groupByType(resources []analyser.Resource) []ResourceGroup {
	order := make([]string, 0)
	index := make(map[string]int)

	for _, r := range resources {
		if _, seen := index[r.Type]; !seen {
			index[r.Type] = len(order)
			order = append(order, r.Type)
		}
	}

	groups := make([]ResourceGroup, len(order))
	for i, t := range order {
		groups[i] = ResourceGroup{Type: t}
	}

	for _, r := range resources {
		i := index[r.Type]
		groups[i].Resources = append(groups[i].Resources, r)
	}

	for i := range groups {
		groups[i].Count = len(groups[i].Resources)
	}

	return groups
}

// groupDiffByType partitions a flat slice of diff resources into DiffGroups,
// preserving the order in which each type is first encountered.
func groupDiffByType(resources []analyser.DiffResource) []DiffGroup {
	order := make([]string, 0)
	index := make(map[string]int)

	for _, r := range resources {
		if _, seen := index[r.Type]; !seen {
			index[r.Type] = len(order)
			order = append(order, r.Type)
		}
	}

	groups := make([]DiffGroup, len(order))
	for i, t := range order {
		groups[i] = DiffGroup{Type: t}
	}

	for _, r := range resources {
		i := index[r.Type]
		groups[i].Resources = append(groups[i].Resources, r)
	}

	for i := range groups {
		groups[i].Count = len(groups[i].Resources)
	}

	return groups
}
