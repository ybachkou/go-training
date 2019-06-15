package player

import (
	"fmt"
	"math/rand"
	"time"
)

var random *rand.Rand

type Player struct {
	Name       string
	Skill      int
	GamePoints int
	MatchPlayer
}

type MatchPlayer struct {
	Score [3]int
}

func (p *Player) Play(pl Player, ch chan int) Player {
	for {
		ball, ok := <-ch
		if !ok {
			fmt.Printf("%s WON TENNIS SERVE!\n", pl.Name)
			p.GamePoints++
			return *p
		}
		if random.Intn(100) > pl.Skill {
			fmt.Printf("%s missed the ball\n", pl.Name)
			close(ch)
			return *p
		}
		fmt.Printf("%s hit the ball %d\n", pl.Name, ball)
		ball++
		ch <- ball
	}
}

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}
