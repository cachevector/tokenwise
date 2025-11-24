package theme

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"

	"tokenwise/assets"
)

type CustomTheme struct{}

func (CustomTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNamePrimary, theme.ColorNameButton:
		return color.NRGBA{R: 0, G: 85, B: 119, A: 255}
	}
	return theme.DefaultTheme().Color(n, v)
}

func (CustomTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Bold {
		return assets.BoldFont()
	}
	return assets.RegularFont()
}

func (CustomTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (CustomTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNameText:
		return 12
	case theme.SizeNameHeadingText:
		return 16
	case theme.SizeNameSubHeadingText:
		return 14
	case theme.SizeNameCaptionText:
		return 10
	}
	return theme.DefaultTheme().Size(n)
}
