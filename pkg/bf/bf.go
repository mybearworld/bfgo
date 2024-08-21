package bf

import (
	"fmt"

	"github.com/mybearworld/bfgo/internal/asts"
	"github.com/mybearworld/bfgo/internal/tokenizer"
)

func Run(code []byte) error {
	tokens := tokenizer.Tokenize(code)
	ast, err := asts.FromTokens(tokens)
	if err != nil {
		return err
	}
	fmt.Println(ast)
	return nil
}
