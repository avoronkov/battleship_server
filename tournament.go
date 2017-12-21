package main

type Statistics struct {
	Stats []Result
}

type Result struct {
	Player1, Player2 string
	Winner           string
	Error            error
}

func (s *Statistics) AddResult(player1, player2, winner string, err error) {
	s.Stats = append(s.Stats, Result{
		Player1: player1,
		Player2: player2,
		Winner:  winner,
		Error:   err,
	})
}

func (s Statistics) TotalStats() map[string]PlayerStat {
	total := make(map[string]PlayerStat)
	for _, res := range s.Stats {
		if res.Winner != "" {
			pw := total[res.Winner]
			pw.TotalWin += 1
			if res.Error == nil {
				pw.CleanWin += 1
			}
			total[res.Winner] = pw

			loser := res.Player1
			if loser == res.Winner {
				loser = res.Player2
			}
			pl := total[loser]
			pl.TotalLoose += 1
			if res.Error != nil {
				pl.TechLoose += 1
			}
			total[loser] = pl
		}
	}
	return total
}

type PlayerStat struct {
	TotalWin   int
	CleanWin   int
	TotalLoose int
	TechLoose  int
}
