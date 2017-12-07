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

func ReadShot(r io.Reader) *Shot {
	sc := bufio.NewScanner(r)
	if !sc.Scan() {
		panic("Shot: no more lines")
	}
	fields := strings.Fields(sc.Text())
	var x int = int(fields[0][0] - 'A')
	if len(fields[0]) != 1 || x < 0 || x >= FIELD_SIZE {
		panic("invalid x coordinate")
	}
	var y int = int(fields[1][0] - '0')
	if len(fields[1]) != 1 || y < 0 || y >= FIELD_SIZE {
		panic("invalid x coordinate")
	}
	return &Shot{x, y}
}

func (s Shot) String() string {
	return fmt.Sprintf("%c %d", byte(s.X)+'A', s.Y)
}
