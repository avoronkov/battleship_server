package cell

func (c Cell) String() string {
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
