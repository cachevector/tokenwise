package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"tokenwise/internal/converter"
	"tokenwise/internal/tokenizer"
)

// -------------- GREEN THEME -----------------

type greenTheme struct{}

func (greenTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	if n == theme.ColorNamePrimary {
		return color.NRGBA{0, 150, 0, 255} // GREEN
	}
	return theme.DefaultTheme().Color(n, v)
}
func (greenTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}
func (greenTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}
func (greenTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

// -------------------------------------------

func main() {
	a := app.New()
	a.Settings().SetTheme(&greenTheme{})
	w := a.NewWindow("TokenWise")
	w.Resize(fyne.NewSize(900, 600))

	inputEntry := widget.NewMultiLineEntry()
	inputEntry.SetPlaceHolder("Paste your input here...")

	outputLabel := widget.NewLabel("")
	outputLabel.Wrapping = fyne.TextWrapWord
	outputPane := container.NewVScroll(outputLabel)

	formatOptions := []string{"JSON", "TOON"}
	inputFormat := widget.NewSelect(formatOptions, nil)
	inputFormat.Selected = "JSON"
	outputFormat := widget.NewSelect(formatOptions, nil)
	outputFormat.Selected = "TOON"

	inputTokens := widget.NewLabel("Input Tokens: 0")
	outputTokens := widget.NewLabel("Output Tokens: 0")
	tokensSaved := widget.NewLabel("Tokens Saved: 0")

	convertBtn := widget.NewButtonWithIcon("Convert", theme.ConfirmIcon(), func() {
		runConversion(
			inputEntry.Text, inputFormat.Selected, outputFormat.Selected,
			outputLabel, inputTokens, outputTokens, tokensSaved,
		)
	})

	formatContainer := container.NewHBox(
		widget.NewLabel("Input Format:"), inputFormat,
		widget.NewLabel("Output Format:"), outputFormat,
		convertBtn,
	)

	tokensContainer := container.NewHBox(inputTokens, outputTokens, tokensSaved)
	panes := container.NewHSplit(inputEntry, outputPane)
	panes.Offset = 0.5

	content := container.NewBorder(formatContainer, tokensContainer, nil, nil, panes)
	w.SetContent(content)

	w.ShowAndRun()
}

func runConversion(input, inFmt, outFmt string,
	outputLabel *widget.Label,
	inputTokens, outputTokens, tokensSaved *widget.Label) {

	if input == "" {
		outputLabel.SetText("")
		return
	}

	var result string
	var err error

	switch {
	case inFmt == "JSON" && outFmt == "TOON":
		result, err = converter.JSONToTOON(input)
	case inFmt == "TOON" && outFmt == "JSON":
		result, err = converter.TOONToJSON(input)
	default:
		result = input
	}

	if err != nil {
		outputLabel.SetText(fmt.Sprintf("Conversion error: %v", err))
		return
	}

	outputLabel.SetText(result)

	inTokens, err := tokenizer.TokenCount(input, "gpt-4")
	if err != nil {
		log.Println("Token count error:", err)
		return
	}
	outTokens, err := tokenizer.TokenCount(result, "gpt-4")
	if err != nil {
		log.Println("Token count error:", err)
		return
	}

	inputTokens.SetText(fmt.Sprintf("Input Tokens: %d", inTokens))
	outputTokens.SetText(fmt.Sprintf("Output Tokens: %d", outTokens))
	tokensSaved.SetText(fmt.Sprintf("Tokens Saved: %d", inTokens-outTokens))
}
