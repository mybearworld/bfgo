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

func FromTokens(tokens []tokenizer.Token) (Node, error) {
	return nil, errors.New("not implemented")
}
