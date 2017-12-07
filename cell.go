package main

import "io"

type cell int

const (
	EMPTY cell = 1
	SHIP  cell = 2
	MISS  cell = 3
	HIT   cell = 4
	KILL  cell = 5
)

func (c cell) String() string {
	switch c {
	case EMPTY:
		return "."
	case SHIP:
		return "#"
	case MISS:
		return "O"
	case HIT, KILL:
		return "X"
	default:
		return "?"
	}
}

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
		panic("Unknown char")
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
		panic("Unknown char")
	}
}
