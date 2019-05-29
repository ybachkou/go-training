package match

import (
	"fmt"
	"match/player"
	"sync"
)

type Match struct {
	player.Player
	player.MatchPlayer
}

func createRoutinesAndChannel(pl1, pl2 *player.Player) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		pl1.Play(*pl1, ch)
		wg.Done()
	}()

	go func() {
		pl2.Play(*pl2, ch)
		wg.Done()
	}()

	ch <- 1
	wg.Wait()
}

func scoreGamePoints(i int, pl1, pl2 *player.Player) {
	for {
		createRoutinesAndChannel(pl1, pl2)

		if pl1.GamePoints == 3 {
			fmt.Printf("\n%s WON GAME!\n", pl1.Name)
			pl1.MatchPlayer.Score[i]++
			pl1.GamePoints = 0
			break
		}
		if pl2.GamePoints == 3 {
			fmt.Printf("\n%s WON GAME!\n", pl2.Name)
			pl2.MatchPlayer.Score[i]++
			pl2.GamePoints = 0
			break
		}
	}
}

func scoreSets(i int, pl1, pl2 *player.Player) {
	for {
		fmt.Println("\n\tNEW GAME")
		scoreGamePoints(i, pl1, pl2)
		if pl1.MatchPlayer.Score[i] == 6 {
			fmt.Printf("\n%s WON SET!\n\n", pl1.Name)
			break
		}
		if pl2.MatchPlayer.Score[i] == 6 {
			fmt.Printf("\n%s WON SET!\n\n", pl2.Name)
			break
		}
	}
}

func printScore(pl1, pl2 *player.Player) {
	if (pl1.MatchPlayer.Score[0] > pl2.MatchPlayer.Score[0] && pl1.MatchPlayer.Score[1] > pl2.MatchPlayer.Score[1]) || (pl1.MatchPlayer.Score[0] < pl2.MatchPlayer.Score[0] && pl1.MatchPlayer.Score[1] < pl2.MatchPlayer.Score[1]) {
		fmt.Printf("\tMATCH RESULTS(2 sets):\n")
		fmt.Printf("%s: %d-%d \n%s: %d-%d", pl1.Name, pl1.MatchPlayer.Score[0], pl1.MatchPlayer.Score[1], pl2.Name, pl2.MatchPlayer.Score[0], pl2.MatchPlayer.Score[1])
	} else {
		scoreSets(2, pl1, pl2)
		fmt.Printf("\tMATCH RESULTS(3 sets):\n")
		fmt.Printf("%s: %d-%d-%d \n%s: %d-%d-%d", pl1.Name, pl1.MatchPlayer.Score[0], pl1.MatchPlayer.Score[1], pl1.MatchPlayer.Score[2], pl2.Name, pl2.MatchPlayer.Score[0], pl2.MatchPlayer.Score[1], pl2.MatchPlayer.Score[2])
	}
}

func (m *Match) Start(pl1, pl2 *player.Player) {
	for i := 0; i < 2; i++ {
		scoreSets(i, pl1, pl2)
	}
	printScore(pl1, pl2)
}
