package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/toon-format/toon-go"

	"tokenwise/assets"
	"tokenwise/internal/converter"
	"tokenwise/internal/tokenizer"
	"tokenwise/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.CustomTheme{})
	w := a.NewWindow("TokenWise")
	w.Resize(fyne.NewSize(900, 600))
	w.SetIcon(assets.GetAppIcon())

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

	convertBtn := widget.NewButton("Convert", func() {
		runConversion(
			inputEntry.Text, inputFormat.Selected, outputFormat.Selected,
			outputLabel, inputTokens, outputTokens, tokensSaved,
		)
	})

	convertBtn.Importance = widget.HighImportance

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
	case inFmt == "JSON" && outFmt == "JSON":
		var tmp map[string]any
		if err = json.Unmarshal([]byte(input), &tmp); err != nil {
			err = fmt.Errorf("invalid JSON input")
			break
		}
		result = input
	case inFmt == "TOON" && outFmt == "TOON":
		var tmp map[string]any
		if err = toon.Unmarshal([]byte(input), &tmp); err != nil {
			err = fmt.Errorf("invalid TOON input")
			break
		}
		result = input

	default:
		err = fmt.Errorf("unsupported format combination")
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
	if inTokens >= outTokens {
		tokensSaved.SetText(fmt.Sprintf("Tokens Saved: %d", inTokens-outTokens))
	} else {
		tokensSaved.SetText(fmt.Sprintf("Extra Tokens: %d", outTokens-inTokens))
	}
}
