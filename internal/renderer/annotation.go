package renderer

import (
	"sort"
	"strings"
)

// Annotation holds a user-defined label attached to a resource ID.
type Annotation struct {
	ResourceID string
	Label      string
	Note       string
}

// AnnotationMap is a lookup from resource ID to its annotation.
type AnnotationMap map[string]Annotation

// buildAnnotations resolves annotations for the resources present in the
// report, returning only those whose resource ID appears in knownIDs.
func buildAnnotations(annotations []Annotation, knownIDs []string) AnnotationMap {
	if len(annotations) == 0 || len(knownIDs) == 0 {
		return AnnotationMap{}
	}

	idSet := make(map[string]struct{}, len(knownIDs))
	for _, id := range knownIDs {
		idSet[strings.ToLower(id)] = struct{}{}
	}

	result := make(AnnotationMap, len(annotations))
	for _, a := range annotations {
		if _, ok := idSet[strings.ToLower(a.ResourceID)]; ok {
			result[a.ResourceID] = a
		}
	}
	return result
}

// sortedAnnotations returns annotations sorted by ResourceID for deterministic
// template rendering.
func sortedAnnotations(am AnnotationMap) []Annotation {
	out := make([]Annotation, 0, len(am))
	for _, a := range am {
		out = append(out, a)
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].ResourceID < out[j].ResourceID
	})
	return out
}
