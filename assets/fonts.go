package assets

import (
	_ "embed"

	"fyne.io/fyne/v2"
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
