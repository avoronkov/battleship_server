package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Not enough args. Usage: <prog1> <prog2> [--slow]")
	}
	game := NewGame(os.Args[1], os.Args[2])
	if len(os.Args) > 3 && os.Args[3] == "--slow" {
		game.Sleep = 600 * time.Millisecond
	}
	defer game.Close()
	go game.Run()

	status := Unknown
	winner := ""
	var err error = nil
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
	if winner != "" {
		fmt.Printf("%s is winner!\n", winner)
		if err != nil {
			fmt.Printf("(%s)\n", err)
		}
	} else {
		fmt.Printf("Tech failure: %s\n", err)
	}
}
