package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var tournament = flag.Bool("t", false, "Let's fight!")
var sleep = flag.Duration("s", 50*time.Millisecond, "sleeping interval")

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		log.Fatal("Not enough args. Usage: <prog1> <prog2> [--slow]")
	}
	if !*tournament {
		winner, err := Play(flag.Arg(0), flag.Arg(1))
		if winner.Name != "" {
			fmt.Printf("%v is winner!\n", winner)
			if err != nil {
				fmt.Printf("(%s)\n", err)
			}
		} else {
			fmt.Printf("Tech failure: %s\n", err)
		}
	} else {
		stat := Statistics{}
		for x, p1 := range flag.Args() {
			for y, p2 := range flag.Args() {
				if x != y {
					winner, err := Play(p1, p2)
					stat.AddResult(path.Base(p1), path.Base(p2), winner.Name, err)
				}
			}
		}
		file, err := os.OpenFile("stats.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Print(err)
		}
		defer file.Close()
		total := stat.TotalStats()
		for player, st := range total {
			line := fmt.Sprintf("%s: wins %d (clean: %d) times, looses %d (tech: %d) times.\n",
				player, st.TotalWin, st.CleanWin, st.TotalLoose, st.TechLoose)
			fmt.Fprint(file, line)
			fmt.Print(line)
		}
		fmt.Fprintln(file, "")
		for _, res := range stat.Stats {
			fmt.Fprintf(file, "%s vs %s : winner is %s", res.Player1, res.Player2, res.Winner)
			if res.Error != nil {
				fmt.Fprintf(file, " (error: %v)", res.Error)
			}
			fmt.Fprintf(file, "\n")
		}
	}
}
