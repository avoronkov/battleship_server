package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Shot struct {
	X int
	Y int
}

func ReadShot(r io.Reader) (*Shot, error) {
	sc := bufio.NewScanner(r)
	if !sc.Scan() {
		return nil, fmt.Errorf("Shot: no more lines")
	}
	fields := strings.Fields(sc.Text())
	if len(fields) < 2 {
		return nil, fmt.Errorf("Incorrect shot: %s", fields)
	}
	var x int = int(fields[0][0] - 'A')
	if len(fields[0]) != 1 || x < 0 || x >= FIELD_SIZE {
		return nil, fmt.Errorf("invalid x coordinate: %s", fields[0])
	}
	var y int = int(fields[1][0] - '0')
	if len(fields[1]) != 1 || y < 0 || y >= FIELD_SIZE {
		return nil, fmt.Errorf("invalid y coordinate: %s", fields[1])
	}
	return &Shot{x, y}, nil
}

func (s Shot) String() string {
	return fmt.Sprintf("%c %d", byte(s.X)+'A', s.Y)
}
