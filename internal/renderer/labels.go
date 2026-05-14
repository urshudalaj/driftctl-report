package renderer

import "sort"

// LabelEntry holds the resolved labels for a single resource.
type LabelEntry struct {
	ResourceID string
	Pairs      []LabelPair
}

// LabelPair is a single key/value label.
type LabelPair struct {
	Key   string
	Value string
}

// LabelsData is the top-level structure passed to the template.
type LabelsData struct {
	Enabled bool
	Entries []LabelEntry
}

// buildLabels constructs LabelsData from the renderer options and the
// set of resource IDs present in the current analysis.
func buildLabels(ids []string, o Options) LabelsData {
	if !o.LabelsEnabled || len(o.LabelMap) == 0 {
		return LabelsData{}
	}

	keyFilter := buildKeySet(o.LabelKeys)

	var entries []LabelEntry
	for _, id := range ids {
		labels, ok := o.LabelMap[id]
		if !ok || len(labels) == 0 {
			continue
		}
		pairs := filterPairs(labels, keyFilter)
		if len(pairs) == 0 {
			continue
		}
		entries = append(entries, LabelEntry{
			ResourceID: id,
			Pairs:      pairs,
		})
	}

	return LabelsData{
		Enabled: true,
		Entries: entries,
	}
}

// buildKeySet converts a slice of allowed keys into a lookup map.
// A nil/empty slice means all keys are allowed.
func buildKeySet(keys []string) map[string]struct{} {
	if len(keys) == 0 {
		return nil
	}
	s := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

// filterPairs returns sorted LabelPairs from m, restricted to keyFilter when
// non-nil.
func filterPairs(m map[string]string, keyFilter map[string]struct{}) []LabelPair {
	var pairs []LabelPair
	for k, v := range m {
		if keyFilter != nil {
			if _, ok := keyFilter[k]; !ok {
				continue
			}
		}
		pairs = append(pairs, LabelPair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Key != pairs[j].Key {
			return pairs[i].Key < pairs[j].Key
		}
		return pairs[i].Value < pairs[j].Value
	})
	return pairs
}
