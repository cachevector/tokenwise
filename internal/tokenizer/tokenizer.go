package main

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

/*
func main() {
	text := `{"users":[{"id":1,"name":"Alice"}]}` // TODO
	count, err := TokenCount(text, "gpt-4o")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token count: %d\n", count)
}
*/
