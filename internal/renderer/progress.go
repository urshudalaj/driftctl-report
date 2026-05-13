package renderer

// ProgressBar represents a single progress bar entry used in the report.
type ProgressBar struct {
	Label      string
	Value      int
	Max        int
	Percentage float64
	BarClass   string // CSS class: success, warning, danger, info
}

// buildProgress constructs a list of ProgressBar entries from the analysis,
// one per resource type showing the managed vs total ratio.
func buildProgress(a Analysis, opts Options) []ProgressBar {
	if !opts.ShowProgress {
		return nil
	}

	typeTotals := map[string][2]int{} // [managed, total]

	for _, r := range a.Managed {
		e := typeTotals[r.Type]
		e[0]++
		e[1]++
		typeTotals[r.Type] = e
	}
	for _, r := range a.Unmanaged {
		e := typeTotals[r.Type]
		e[1]++
		typeTotals[r.Type] = e
	}
	for _, r := range a.Deleted {
		e := typeTotals[r.Type]
		e[1]++
		typeTotals[r.Type] = e
	}

	types := sortedKeys(typeTotals)

	bars := make([]ProgressBar, 0, len(types))
	for _, t := range types {
		counts := typeTotals[t]
		managed, total := counts[0], counts[1]
		if total == 0 {
			continue
		}
		pct := float64(managed) / float64(total) * 100
		bars = append(bars, ProgressBar{
			Label:      t,
			Value:      managed,
			Max:        total,
			Percentage: roundTwo(pct),
			BarClass:   progressBarClass(pct),
		})
	}

	if opts.ProgressTopN > 0 && len(bars) > opts.ProgressTopN {
		bars = bars[:opts.ProgressTopN]
	}

	return bars
}

func progressBarClass(pct float64) string {
	switch {
	case pct >= 90:
		return "success"
	case pct >= 60:
		return "info"
	case pct >= 30:
		return "warning"
	default:
		return "danger"
	}
}

func sortedKeys(m map[string][2]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sortStrings(keys)
	return keys
}

func sortStrings(s []string) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}
