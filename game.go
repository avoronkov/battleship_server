package main

import (
	"bytes"
	"fmt"
	"path"
	"time"
)

type Game struct {
	p1         *Player
	p2         *Player
	Sleep      time.Duration
	StatusPipe chan StatusMessage
	Winner     string
	Error      error
}

func NewGame(exe1, exe2 string) *Game {
	g := new(Game)
	g.p1 = NewPlayer(path.Base(exe1), exe1)
	g.p2 = NewPlayer(path.Base(exe2), exe2)
	g.StatusPipe = make(chan StatusMessage)
	g.Sleep = 300 * time.Millisecond
	return g
}

func (g *Game) Close() error {
	g.p1.Close()
	g.p2.Close()
	return nil
}

func (g *Game) Run() {
	f1 := g.p1.InitField()
	f2 := g.p2.InitField()
	if err := checkField(f1); err != nil {
		g.Error = err
		g.Winner = g.Player2()
		g.StatusPipe <- Finish
		return
	}
	if err := checkField(f2); err != nil {
		g.Error = err
		g.Winner = g.Player1()
		g.StatusPipe <- Finish
		return
	}

	currentPlayer := g.p1
	currentField := f1
	anotherPlayer := g.p2
	anotherField := f2

	currentPlayer.ShootCmd()

	fmt.Println(DisplayFields(g.p1.Name(), g.p2.Name(), f1, f2))

	wait := WaitP1

L:
	for {
		time.Sleep(g.Sleep)
		g.StatusPipe <- wait
		shot := currentPlayer.GetShot()
		result, e := anotherField.Shoot(*shot)
		if e != nil {
			g.Winner = anotherPlayer.Name()
			g.Error = e
			g.StatusPipe <- Finish
			return
		}
		fmt.Println(DisplayFields(g.p1.Name(), g.p2.Name(), f1, f2))
		if anotherField.StillAlive() == 0 {
			// Win!
			currentPlayer.Win()
			anotherPlayer.Lose()
			g.Winner = currentPlayer.Name()
			g.StatusPipe <- Finish
			return
		}
		currentPlayer.SendResult(result)
		switch result {
		case HIT, KILL:
			continue L
		case MISS:
			// switch players
			currentPlayer, anotherPlayer = anotherPlayer, currentPlayer
			currentField, anotherField = anotherField, currentField
			currentPlayer.ShootCmd()
			if wait == WaitP1 {
				wait = WaitP2
			} else {
				wait = WaitP1
			}

		}
	}
}

func (g Game) Player1() string {
	return fmt.Sprintf("%s (1)", g.p1.Name())
}

func (g Game) Player2() string {
	return fmt.Sprintf("%s (2)", g.p2.Name())
}

func checkField(f *Field) error {
	return nil
}

func DisplayFields(name1, name2 string, f1, f2 *Field) string {
	var buff bytes.Buffer
	// TODO
	fmt.Fprintf(&buff, "\033[H\033[2J")
	fmt.Fprintf(&buff, "\n%s: %s:\n", name1, name2)
	for l := 0; l < FIELD_SIZE; l++ {
		fmt.Fprintf(&buff, "%s    %s\n", f1.LineString(l), f2.LineString(l))
	}
	fmt.Fprintf(&buff, "        %2d    %2d\n", f1.StillAlive(), f2.StillAlive())
	return buff.String()
}
