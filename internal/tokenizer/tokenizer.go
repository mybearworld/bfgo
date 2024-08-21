package tokenizer

type Token int

const (
	IncrementCell Token = iota
	DecrementCell
	IncrementPointer
	DecrementPointer
	Input
	Output
	LoopBegin
	LoopEnd
	Unrecognized
)

func charToToken(char byte) Token {
	switch char {
	case '+':
		return IncrementCell
	case '-':
		return DecrementCell
	case '>':
		return IncrementPointer
	case '<':
		return DecrementPointer
	case ',':
		return Input
	case '.':
		return Output
	case '[':
		return LoopBegin
	case ']':
		return LoopEnd
	default:
		return Unrecognized
	}
}

func Tokenize(code []byte) []Token {
	tokens := []Token{}
	for _, char := range code {
		tokens = append(tokens, charToToken(char))
	}
	return tokens
}
