package bf

import (
	"fmt"

	"github.com/mybearworld/bfgo/internal/tokenizer"
)

func Run(code []byte) error {
	tokens := tokenizer.Tokenize(code)
	fmt.Println(tokens)
	return nil
}
