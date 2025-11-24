package theme

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:embed fonts/JetBrainsMono-Regular.ttf
var jetbrainsRegular []byte

//go:embed fonts/JetBrainsMono-Bold.ttf
var jetbrainsBold []byte

func RegularFont() fyne.Resource {
	return fyne.NewStaticResource("JetBrainsMono-Regular", jetbrainsRegular)
}

func BoldFont() fyne.Resource {
	return fyne.NewStaticResource("JetBrainsMono-Bold", jetbrainsBold)
}

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
		return BoldFont()
	}
	return RegularFont()
}

func (CustomTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (CustomTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNameText:
		return 12 // normal text
	case theme.SizeNameHeadingText:
		return 16
	case theme.SizeNameSubHeadingText:
		return 14
	case theme.SizeNameCaptionText:
		return 10
	}
	return theme.DefaultTheme().Size(n)
}
