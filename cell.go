package main

import (
	"fmt"
	"io"
)

type cell int

const (
	EMPTY cell = 1
	SHIP  cell = 2
	MISS  cell = 3
	HIT   cell = 4
	KILL  cell = 5
)

func RuneToCell(c byte) cell {
	switch c {
	case '.', '0':
		return EMPTY
	case '#', '1':
		return SHIP
	case '-':
		return MISS
	case 'X':
		return HIT
	default:
		panic(fmt.Errorf("Unknown char: %c", c))
	}

}

func ReadCell(r io.Reader) cell {
	buf := make([]byte, 1)
	if _, err := r.Read(buf); err != nil {
		panic(err)
	}
	switch buf[0] {
	case '.', '0':
		return EMPTY
	case '#', '1':
		return SHIP
	case '-':
		return MISS
	case 'X':
		return HIT
	default:
		panic(fmt.Errorf("Unknown char: %c", buf[0]))
	}
}
