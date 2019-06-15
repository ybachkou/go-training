package main

import (
	"match/match"
	"match/player"
)

func main() {
	pl1 := player.Player{Name: "Alexsandr", Skill: 70}
	pl2 := player.Player{Name: "Stanislav", Skill: 70}
	m := match.Match{}
	m.Start(&pl1, &pl2)
}
