package cell

import (
	"fmt"
)

type Cell int

const (
	EMPTY Cell = iota
	SHIP
	MISS
	HIT
	KILL
)

func RuneToCell(c byte) Cell {
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
