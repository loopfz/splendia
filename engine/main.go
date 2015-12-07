package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var playerList arrayFlags
	flag.Var(&playerList, "player", "player program path")
	maxTurn := flag.Int("maxturn", 0, "max number of turns")
	flag.Parse()

	var pl []*Player
	for _, p := range playerList {
		pl = append(pl, NewForkPlayer(p))
	}

	g, err := NewGame(pl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err)
		os.Exit(1)
	}

	printScores(g.Run(*maxTurn))
}

func printScores(winners []*Player) {
	if len(winners) == 1 {
		fmt.Println("WINNER:")
	} else {
		fmt.Println("WINNERS (TIE):")
	}
	for _, p := range winners {
		fmt.Printf("Player %d (%s): %d points, %d cards\n", p.Number, p.Program, p.VP, p.NumberCards)
	}
}
