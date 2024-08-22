package bf

import (
	"fmt"

	"github.com/mybearworld/bfgo/internal/asts"
	"github.com/mybearworld/bfgo/internal/tokenizer"
)

func Run(code []byte, newLine bool) error {
	tokens := tokenizer.Tokenize(code)
	program, err := asts.FromTokens(tokens)
	if err != nil {
		return err
	}
	program.Start()
	if newLine {
		fmt.Println()
	}
	return nil
}
