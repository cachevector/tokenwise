package tokenizer

import (
	"github.com/pkoukk/tiktoken-go"
)

func TokenCount(text string, model string) (int, error) {
	tk, err := tiktoken.EncodingForModel(model)
	if err != nil {
		return 0, err
	}
	// Encode returns []int (token IDs)
	tokens := tk.Encode(text, nil, nil) // allowedSpecial, disallowedSpecial = nil (default)
	return len(tokens), nil
}
