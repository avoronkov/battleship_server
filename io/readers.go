package io

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"battleship_server/cell"
	"battleship_server/field"
	"battleship_server/shot"
)

func ReadCell(r io.Reader) cell.Cell {
	buf := make([]byte, 1)
	if _, err := r.Read(buf); err != nil {
		panic(err)
	}
	switch buf[0] {
	case '.', '0':
		return cell.EMPTY
	case '#', '1':
		return cell.SHIP
	case '-':
		return cell.MISS
	case 'X':
		return cell.HIT
	default:
		panic(fmt.Errorf("Unknown char: %c", buf[0]))
	}
}

func ReadField(r io.Reader) *field.Field {
	f := new(field.Field)
	log.Printf("readField from %s", r)
	y := 0
	sc := bufio.NewScanner(r)
	for y < field.FIELD_SIZE {
		if !sc.Scan() {
			panic("No more lines")
		}
		line := strings.Trim(sc.Text(), " \n\t")
		if line == "" {
			continue
		}
		if len(line) < field.FIELD_SIZE {
			panic("too short line")
		}
		for x := 0; x < field.FIELD_SIZE; x++ {
			f.Set(x, y, cell.RuneToCell(line[x]))
		}
		y++
	}
	return f
}

func ReadShot(r io.Reader) (*shot.Shot, error) {
	sc := bufio.NewScanner(r)
	if !sc.Scan() {
		return nil, fmt.Errorf("Shot: no more lines")
	}
	fields := strings.Fields(sc.Text())
	if len(fields) < 2 {
		return nil, fmt.Errorf("Incorrect shot: %s", fields)
	}
	var x int = int(fields[0][0] - 'A')
	if len(fields[0]) != 1 || x < 0 || x >= field.FIELD_SIZE {
		return nil, fmt.Errorf("invalid x coordinate: %s", fields[0])
	}
	var y int = int(fields[1][0] - '0')
	if len(fields[1]) != 1 || y < 0 || y >= field.FIELD_SIZE {
		return nil, fmt.Errorf("invalid y coordinate: %s", fields[1])
	}
	return &shot.Shot{x, y}, nil
}

