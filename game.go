package main

import (
	"battleship_server/cell"
	"battleship_server/field"
	"bytes"
	"fmt"
	"log"
	"path"
	"time"
)

type Game struct {
	p1         *Player
	p2         *Player
	Sleep      time.Duration
	StatusPipe chan StatusMessage
	Winner     PlayerName
	Error      error
}

func NewGame(exe1, exe2 string) *Game {
	g := new(Game)
	pn1 := PlayerName{path.Base(exe1), 1}
	g.p1 = NewPlayer(pn1, exe1)
	time.Sleep(2 * time.Second)
	pn2 := PlayerName{path.Base(exe2), 2}
	g.p2 = NewPlayer(pn2, exe2)
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
	f1, e1 := g.p1.InitField()
	if e1 == nil {
		e1 = f1.Check()
	}
	if e1 != nil {
		g.finish(g.p2, g.p1, e1)
		return
	}
	f2, e2 := g.p2.InitField()
	if e2 == nil {
		e2 = f2.Check()
	}
	if e2 != nil {
		g.finish(g.p1, g.p2, e2)
		return
	}

	currentPlayer := g.p1
	currentField := f1
	anotherPlayer := g.p2
	anotherField := f2

	currentPlayer.ShootCmd()

	fmt.Println(DisplayFields(g.p1.Name().Name, g.p2.Name().Name, f1, f2))

	wait := WaitP1

L:
	for {
		time.Sleep(g.Sleep)
		g.StatusPipe <- wait
		shot, e := currentPlayer.GetShot()
		if e != nil {
			g.finish(anotherPlayer, currentPlayer, e)
			return
		}
		result, e := anotherField.Shoot(*shot)
		if e != nil {
			g.finish(anotherPlayer, currentPlayer, e)
			return
		}
		log.Printf("shoot result: %s", result)
		fmt.Println(DisplayFields(g.p1.Name().Name, g.p2.Name().Name, f1, f2))
		if anotherField.StillAlive() == 0 {
			// Win!
			g.finish(currentPlayer, anotherPlayer, nil)
			return
		}
		currentPlayer.SendResult(result)
		switch result {
		case cell.HIT, cell.KILL:
			continue L
		case cell.MISS:
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

func (g Game) Player1() PlayerName {
	return g.p1.Name()
}

func (g Game) Player2() PlayerName {
	return g.p2.Name()
}

func (g *Game) finish(winner, loser *Player, err error) {
	winner.Win()
	loser.Lose()
	g.Winner = winner.Name()
	g.Error = err
	g.StatusPipe <- Finish
}

func DisplayFields(name1, name2 string, f1, f2 *field.Field) string {
	var buff bytes.Buffer
	// TODO
	fmt.Fprintf(&buff, "\033[H\033[2J")
	fmt.Fprintf(&buff, "\n%s: %s:\n", name1, name2)
	for l := 0; l < field.FIELD_SIZE; l++ {
		fmt.Fprintf(&buff, "%s    %s\n", f1.LineString(l), f2.LineString(l))
	}
	fmt.Fprintf(&buff, "        %2d    %2d\n", f1.StillAlive(), f2.StillAlive())
	return buff.String()
}

func Play(player1, player2 string) (winner PlayerName, err error) {
	game := NewGame(player1, player2)
	game.Sleep = *sleep

	defer game.Close()
	go game.Run()

	status := Unknown

L:
	for {
		select {
		case status = <-game.StatusPipe:
			log.Printf("Status = %d", status)
			if status == Finish {
				winner = game.Winner
				err = game.Error
				break L
			}
		case <-time.After(3 * time.Second):
			if status == WaitP1 {
				winner = game.Player2()
				err = fmt.Errorf("%s timeout", game.Player1())
			} else if status == WaitP2 {
				winner = game.Player1()
				err = fmt.Errorf("%s timeout", game.Player2())
			} else {
				err = fmt.Errorf("Unknown timeout")
			}
			break L
		}
	}
	return
}
