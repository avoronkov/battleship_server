package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

const FIELD_SIZE = 10

type Field struct {
	field [FIELD_SIZE][FIELD_SIZE]cell
}

func ReadField(r io.Reader) *Field {
	f := new(Field)
	log.Printf("readField from %s", r)
	y := 0
	sc := bufio.NewScanner(r)
	for y < FIELD_SIZE {
		if !sc.Scan() {
			panic("No more lines")
		}
		line := strings.Trim(sc.Text(), " \n\t")
		if line == "" {
			continue
		}
		if len(line) < FIELD_SIZE {
			panic("too short line")
		}
		for x := 0; x < FIELD_SIZE; x++ {
			f.field[y][x] = RuneToCell(line[x])
		}
		y++
	}
	return f
}

func (f *Field) Shoot(s Shot) (cell, error) {
	state := f.field[s.Y][s.X]
	switch state {
	case EMPTY:
		f.field[s.Y][s.X] = MISS
		return MISS, nil
	case SHIP:
		f.field[s.Y][s.X] = HIT
		if f.AliveAround(s.X, s.Y) {
			return HIT, nil
		} else {
			return KILL, nil
		}
		return HIT, nil
	default:
		return EMPTY, fmt.Errorf("Already shoot at %s", s)
	}
}

func (f Field) AliveAround(x, y int) bool {
	if x > 1 && f.field[y][x-1] == SHIP {
		return true
	}
	if x < FIELD_SIZE-1 && f.field[y][x+1] == SHIP {
		return true
	}
	if y > 1 && f.field[y-1][x] == SHIP {
		return true
	}
	if y < FIELD_SIZE-1 && f.field[y+1][x] == SHIP {
		return true
	}
	return false
}

func (f Field) StillAlive() int {
	count := 0
	for _, line := range f.field {
		for _, v := range line {
			if v == SHIP {
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
