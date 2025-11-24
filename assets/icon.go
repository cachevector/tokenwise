package assets

import (
	"os"

	"fyne.io/fyne/v2"
)

func GetAppIcon() fyne.Resource {
	data, err := os.ReadFile("assets/icons/icon.png")
	if err != nil {
		panic(err)
	}
	return fyne.NewStaticResource("icon.png", data)
}
