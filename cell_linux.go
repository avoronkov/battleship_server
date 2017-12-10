package main

func (c cell) String() string {
	switch c {
	case EMPTY:
		return "."
	case SHIP:
		return "\x1b[32m#\x1b[0m"
	case MISS:
		return "\x1b[34mO\x1b[0m"
	case HIT, KILL:
		return "\x1b[31mX\x1b[0m"
	default:
		return "?"
	}
}
