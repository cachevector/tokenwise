package converter

import (
	"encoding/json"
	"fmt"

	"github.com/toon-format/toon-go"
)

func Convert(input string, from, to Format) (string, error) {
	switch from {
	case FormatJSON:
		if to == FormatTOON {
			return JSONToTOON(input)
		}
	case FormatTOON:
		if to == FormatJSON {
			return TOONToJSON(input)
		}
	}
	return "", fmt.Errorf("unsupported conversion")
}

// JSON -> TOON
func JSONToTOON(input string) (string, error) {
	var doc map[string]any
	if err := json.Unmarshal([]byte(input), &doc); err != nil {
		return "", err
	}

	// Marshal to TOON with length markers
	b, err := toon.Marshal(doc, toon.WithLengthMarkers(true))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// TOON -> JSON
func TOONToJSON(input string) (string, error) {
	var doc map[string]any
	if err := toon.Unmarshal([]byte(input), &doc); err != nil {
		return "", err
	}

	j, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return "", err
	}
	return string(j), nil
}
