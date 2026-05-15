package renderer

import (
	"strings"
)

// EmbedData holds all inlined asset strings for a self-contained HTML report.
type EmbedData struct {
	Enabled    bool
	CSS        string
	JS         string
	FontFaces  string
}

// buildEmbed constructs an EmbedData value from the renderer options.
// When EmbedCSS, EmbedJS, or EmbedFonts are disabled their respective fields
// are left empty so the template can skip the inline blocks.
func buildEmbed(o Options) EmbedData {
	if !o.EmbedCSS && !o.EmbedJS && !o.EmbedFonts {
		return EmbedData{}
	}

	ed := EmbedData{Enabled: true}

	if o.EmbedCSS {
		ed.CSS = defaultEmbedCSS(o)
	}

	if o.EmbedJS {
		ed.JS = defaultEmbedJS()
	}

	if o.EmbedFonts {
		ed.FontFaces = defaultEmbedFontFaces()
	}

	return ed
}

// defaultEmbedCSS returns a minimal set of CSS rules derived from the active
// theme. In a real implementation this would read from embedded FS assets.
func defaultEmbedCSS(o Options) string {
	var b strings.Builder
	b.WriteString(":root{")
	b.WriteString("--color-managed:" + o.ColorManaged + ";")
	b.WriteString("--color-unmanaged:" + o.ColorUnmanaged + ";")
	b.WriteString("--color-deleted:" + o.ColorDeleted + ";")
	b.WriteString("}")
	return b.String()
}

// defaultEmbedJS returns a minimal JS snippet for interactive features.
func defaultEmbedJS() string {
	return "document.addEventListener('DOMContentLoaded',function(){" +
		"var t=document.getElementById('theme-toggle');" +
		"if(t){t.addEventListener('click',function(){" +
		"document.body.classList.toggle('dark');});}});"
}

// defaultEmbedFontFaces returns an empty string; real usage would return
// base64-encoded @font-face declarations.
func defaultEmbedFontFaces() string {
	return ""
}
