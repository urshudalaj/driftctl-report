package renderer

import "fmt"

// WithTheme sets a custom Theme on the Options.
// It returns an OptionFunc that applies the provided theme directly.
func WithTheme(theme Theme) OptionFunc {
	return func(o *Options) error {
		o.Theme = theme
		return nil
	}
}

// WithThemeName looks up a predefined theme by name and applies it to Options.
// It returns an error if the theme name is not recognised.
func WithThemeName(name string) OptionFunc {
	return func(o *Options) error {
		theme, ok := ThemeByName(name)
		if !ok {
			return fmt.Errorf("renderer: unknown theme %q; available themes are: default, dark, high-contrast", name)
		}
		o.Theme = theme
		return nil
	}
}
