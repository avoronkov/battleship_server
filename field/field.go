package field

import (
	"battleship_server/cell"
	"battleship_server/shot"
	"bytes"
	"fmt"
)

const FIELD_SIZE = 10

type Field struct {
	field [FIELD_SIZE][FIELD_SIZE]cell.Cell
}

func (f *Field) Set(x, y int, c cell.Cell) {
	f.field[y][x] = c
}

func (f *Field) Shoot(s shot.Shot) (cell.Cell, error) {
	state := f.field[s.Y][s.X]
	switch state {
	case cell.EMPTY:
		f.field[s.Y][s.X] = cell.MISS
		return cell.MISS, nil
	case cell.SHIP:
		f.field[s.Y][s.X] = cell.HIT
		if f.killed(s.X, s.Y) {
			return cell.HIT, nil
		} else {
			return cell.KILL, nil
		}
		return cell.HIT, nil
	default:
		return cell.EMPTY, fmt.Errorf("Already shoot at %s", s)
	}
}

func (f Field) killed(x, y int) bool {
	cells := []shot.Shot{{x, y}}
	checked_cells := []shot.Shot{}
L:
	for len(cells) > 0 {
		c := cells[0]
		cells = cells[1:]
		for _, cc := range checked_cells {
			if c == cc {
				continue L
			}
		}
		checked_cells = append(checked_cells, c)
		if c.X > 0 {
			st := f.field[c.Y][c.X-1]
			if st == cell.SHIP {
				return false
			}
			if st == cell.HIT {
				cells = append(cells, shot.Shot{c.X - 1, c.Y})
			}
		}
		if c.X < FIELD_SIZE-1 {
			st := f.field[c.Y][c.X+1]
			if st == cell.SHIP {
				return false
			}
			if st == cell.HIT {
				cells = append(cells, shot.Shot{c.X + 1, c.Y})
			}
		}
		if c.Y > 0 {
			st := f.field[c.Y-1][c.X]
			if st == cell.SHIP {
				return false
			}
			if st == cell.HIT {
				cells = append(cells, shot.Shot{c.X, c.Y - 1})
			}
		}
		if c.Y < FIELD_SIZE-1 {
			st := f.field[c.Y+1][c.X]
			if st == cell.SHIP {
				return false
			}
			if st == cell.HIT {
				cells = append(cells, shot.Shot{c.X, c.Y - 1})
			}
		}
	}
	return true
}

func (f Field) Check() error {
	return nil
}

func (f Field) AliveAround(x, y int) bool {
	if x > 1 && f.field[y][x-1] == cell.SHIP {
		return true
	}
	if x < FIELD_SIZE-1 && f.field[y][x+1] == cell.SHIP {
		return true
	}
	if y > 1 && f.field[y-1][x] == cell.SHIP {
		return true
	}
	if y < FIELD_SIZE-1 && f.field[y+1][x] == cell.SHIP {
		return true
	}
	return false
}

func (f Field) StillAlive() int {
	count := 0
	for _, line := range f.field {
		for _, v := range line {
			if v == cell.SHIP {
				count++
			}
		}
	}
	return count
}

func (f Field) String() string {
	var buffer bytes.Buffer
	for _, line := range f.field {
		for _, v := range line {
			buffer.WriteString(v.String())
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (f Field) LineString(l int) string {
	var buffer bytes.Buffer
	for _, v := range f.field[l] {
		buffer.WriteString(v.String())
	}
	return buffer.String()
}
