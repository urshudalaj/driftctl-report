package renderer

import "strings"

// PrintStyles holds the CSS injected for print-friendly rendering.
type PrintStyles struct {
	Enabled    bool
	HideNav    bool
	PageBreaks bool
	CSS        string
}

// buildPrintStyles constructs the PrintStyles value from the current
// options. It returns a zero-value struct when print mode is disabled.
func buildPrintStyles(o *Options) PrintStyles {
	if !o.PrintFriendly {
		return PrintStyles{}
	}

	var b strings.Builder
	b.WriteString("@media print {\n")

	if o.PrintHideNav {
		b.WriteString("  nav, .navbar { display: none !important; }\n")
	}

	if o.PrintPageBreaks {
		b.WriteString("  .report-section { page-break-before: always; }\n")
		b.WriteString("  .report-section:first-of-type { page-break-before: avoid; }\n")
	}

	b.WriteString("  body { font-size: 11pt; color: #000; background: #fff; }\n")
	b.WriteString("  a[href]::after { content: none; }\n")
	b.WriteString("}\n")

	return PrintStyles{
		Enabled:    true,
		HideNav:    o.PrintHideNav,
		PageBreaks: o.PrintPageBreaks,
		CSS:        b.String(),
	}
}
