package asts

import (
	"errors"
	"fmt"
	"os"

	"github.com/mybearworld/bfgo/internal/tokenizer"
)

type Node interface {
	Run(tape *Tape, ptr *int)
}

const maxTapeSize = 300000

type Tape [maxTapeSize]byte

type Program struct {
	Nodes []Node
}

func (program Program) Start() {
	program.Run(&Tape{0}, new(int))
}

func (program Program) Run(tape *Tape, ptr *int) {
	for _, node := range program.Nodes {
		node.Run(tape, ptr)
	}
}

type IncrementCell struct{}

func (IncrementCell) Run(tape *Tape, ptr *int) {
	tape[*ptr]++
}

type DecrementCell struct{}

func (DecrementCell) Run(tape *Tape, ptr *int) {
	tape[*ptr]--
}

type IncrementPointer struct{}

func (IncrementPointer) Run(tape *Tape, ptr *int) {
	*ptr++
	if *ptr > maxTapeSize {
		*ptr = 0
	}
}

type DecrementPointer struct{}

func (DecrementPointer) Run(tape *Tape, ptr *int) {
	*ptr--
	if *ptr < 0 {
		*ptr = maxTapeSize
	}
}

type Input struct{}

func (Input) Run(tape *Tape, ptr *int) {
	buf := make([]byte, 1)
	os.Stdin.Read(buf)
	tape[*ptr] = buf[0]
}

type Output struct{}

func (Output) Run(tape *Tape, ptr *int) {
	fmt.Printf("%c", tape[*ptr])
}

type Loop struct {
	Nodes []Node
}

func (loop Loop) Run(tape *Tape, ptr *int) {
	for tape[*ptr] != 0 {
		for _, node := range loop.Nodes {
			node.Run(tape, ptr)
		}
	}
}

func tokensToNodes(tokens []tokenizer.Token) ([]Node, error) {
	nodes := []Node{}
	for i := 0; i < len(tokens); {
		token := tokens[i]
		node, ok := tokenToNode(token)
		if ok {
			nodes = append(nodes, node)
			i++
			continue
		}
		switch token {
		case tokenizer.Unrecognized:
			i++
		case tokenizer.LoopBegin:
			loopNodes, err := tokensToNodes(tokens[i+1:])
			if err != nil {
				switch err := err.(type) {
				case LoopEndError:
					nodes = append(nodes, Loop{Nodes: loopNodes})
					// the two loop characters need to be skipped, so add two
					i += err.TokenIndex + 2
					continue
				default:
					return nodes, err
				}
			}
			return nodes, errors.New("unclosed loop")
		case tokenizer.LoopEnd:
			return nodes, LoopEndError{i}
		default:
			return nodes, fmt.Errorf("unexpected token type %d", token)
		}
	}
	return nodes, nil
}

func FromTokens(tokens []tokenizer.Token) (Program, error) {
	nodes, err := tokensToNodes(tokens)
	if err != nil {
		return Program{}, err
	}
	return Program{nodes}, nil
}

type LoopEndError struct {
	TokenIndex int
}

func (LoopEndError) Error() string {
	return "unmatched loop ending character"
}

func tokenToNode(token tokenizer.Token) (Node, bool) {
	switch token {
	case tokenizer.IncrementCell:
		return IncrementCell{}, true
	case tokenizer.DecrementCell:
		return DecrementCell{}, true
	case tokenizer.IncrementPointer:
		return IncrementPointer{}, true
	case tokenizer.DecrementPointer:
		return DecrementPointer{}, true
	case tokenizer.Input:
		return Input{}, true
	case tokenizer.Output:
		return Output{}, true
	}
	return nil, false
}
