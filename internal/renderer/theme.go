package renderer

// Theme defines the visual styling options for the HTML report.
type Theme struct {
	// PrimaryColor is the main accent color used in headers and highlights.
	PrimaryColor string
	// DriftedColor is the color used to indicate drifted resources.
	DriftedColor string
	// MissingColor is the color used to indicate missing resources.
	MissingColor string
	// UnmanagedColor is the color used to indicate unmanaged resources.
	UnmanagedColor string
	// FontFamily is the CSS font-family string applied globally.
	FontFamily string
}

// defaultTheme returns the built-in theme used when no theme is specified.
func defaultTheme() Theme {
	return Theme{
		PrimaryColor:   "#4F46E5",
		DriftedColor:   "#F59E0B",
		MissingColor:   "#EF4444",
		UnmanagedColor: "#6B7280",
		FontFamily:     "'Inter', system-ui, sans-serif",
	}
}

// predefinedThemes holds named themes available via WithTheme.
var predefinedThemes = map[string]Theme{
	"default": defaultTheme(),
	"dark": {
		PrimaryColor:   "#818CF8",
		DriftedColor:   "#FCD34D",
		MissingColor:   "#F87171",
		UnmanagedColor: "#9CA3AF",
		FontFamily:     "'Inter', system-ui, sans-serif",
	},
	"high-contrast": {
		PrimaryColor:   "#1D4ED8",
		DriftedColor:   "#B45309",
		MissingColor:   "#B91C1C",
		UnmanagedColor: "#374151",
		FontFamily:     "'Courier New', monospace",
	},
}

// ThemeByName returns a predefined theme by name.
// If the name is not found, the default theme is returned along with false.
func ThemeByName(name string) (Theme, bool) {
	t, ok := predefinedThemes[name]
	if !ok {
		return defaultTheme(), false
	}
	return t, true
}
