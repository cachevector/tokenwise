package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"tokenwise/internal/converter"
	"tokenwise/internal/tokenizer"
)

func main() {
	a := app.New()
	w := a.NewWindow("TokenWise")
	w.Resize(fyne.NewSize(900, 600))

	inputEntry := widget.NewMultiLineEntry()
	inputEntry.SetPlaceHolder("Paste your input here...")
	outputEntry := widget.NewMultiLineEntry()
	outputEntry.SetPlaceHolder("Output will appear here...")
	outputEntry.Disable()

	var inputFormat, outputFormat *widget.Select
	formatOptions := []string{"JSON", "TOON"}
	inputFormat = widget.NewSelect(formatOptions, func(string) { updateConversion(inputEntry, outputEntry, inputFormat.Selected, outputFormat.Selected) })
	inputFormat.Selected = "JSON"
	outputFormat = widget.NewSelect(formatOptions, func(string) { updateConversion(inputEntry, outputEntry, inputFormat.Selected, outputFormat.Selected) })
	outputFormat.Selected = "TOON"

	inputTokens := widget.NewLabel("Input Tokens: 0")
	outputTokens := widget.NewLabel("Output Tokens: 0")
	tokensSaved := widget.NewLabel("Tokens Saved: 0")

	inputEntry.OnChanged = func(_ string) {
		updateConversion(inputEntry, outputEntry, inputFormat.Selected, outputFormat.Selected)

		inTokens, err := tokenizer.TokenCount(inputEntry.Text, "gpt-4")
		if err != nil {
			log.Fatal(err)
		}
		outTokens, err := tokenizer.TokenCount(outputEntry.Text, "gpt-4")
		if err != nil {
			log.Fatal(err)
		}

		inputTokens.SetText(fmt.Sprintf("Input Tokens: %d", inTokens))
		outputTokens.SetText(fmt.Sprintf("Output Tokens: %d", outTokens))
		tokensSaved.SetText(fmt.Sprintf("Tokens Saved: %d", inTokens-outTokens))
	}

	formatContainer := container.NewHBox(
		widget.NewLabel("Input Format:"), inputFormat,
		widget.NewLabel("Output Format:"), outputFormat,
	)

	tokensContainer := container.NewHBox(inputTokens, outputTokens, tokensSaved)
	panes := container.NewHSplit(inputEntry, outputEntry)
	panes.Offset = 0.5

	content := container.NewBorder(formatContainer, tokensContainer, nil, nil, panes)
	w.SetContent(content)

	w.ShowAndRun()
}

func updateConversion(inputEntry, outputEntry *widget.Entry, inFmt, outFmt string) {
	input := inputEntry.Text
	if input == "" {
		outputEntry.SetText("")
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
		outputEntry.SetText(fmt.Sprintf("Conversion error: %v", err))
	} else {
		outputEntry.SetText(result)
	}
}
