package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	WHITE Color = "WHITE"
	RED   Color = "RED"
	GREEN Color = "GREEN"
	BLUE  Color = "BLUE"
	BLACK Color = "BLACK"
	GOLD  Color = "GOLD"

	MAX_VP = 15

	TIER1 = "TIER1"
	TIER2 = "TIER2"
	TIER3 = "TIER3"

	BUY_CARD_ACT     = "BUY_CARD"
	RESERVE_CARD_ACT = "RESERVE_CARD"
	TOKENS_ACT       = "TOKENS"

	MAX_TOKENS     = 10
	MAX_RESERVED   = 3
	PLAYER_TIMEOUT = 5 * time.Second
)

type Color string

type ColorCount map[Color]uint

type Card struct {
	Color Color      `json:"color"`
	Cost  ColorCount `json:"cost"`
	VP    uint       `json:"vp"`
}

type Noble struct {
	Name         string     `json:"name"`
	Requirements ColorCount `json:"requirements"`
	VP           uint       `json:"vp"`
}

var Colors = []Color{WHITE, RED, GREEN, BLUE, BLACK}

type Game struct {
	Players      []*Player           `json:"players"`
	ActivePlayer uint                `json:"active_player"`
	Decks        map[string]CardDeck `json:"-"`
	Cards        map[string]*Card    `json:"cards"`
	DeckSizes    map[string]uint     `json:"deck_sizes"`
	Nobles       []*Noble            `json:"nobles"`
	Tokens       ColorCount          `json:"tokens"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame(players []*Player) (*Game, error) {
	g := &Game{}

	g.AddPlayers(players)

	if len(g.Players) <= 1 {
		return nil, errors.New("Not enough players!")
	}

	g.InitTokens()
	g.InitNobles()
	g.InitDecks()
	g.InitFirstPlayer()

	return g, nil
}

func (g *Game) Run(maxTurn int) []*Player {

	fmt.Printf("Player %d (%s) is first player!\n", g.Players[0].Number, g.Players[0].Program)

	lastTurn := false
	for curTurn := 0; !lastTurn && (maxTurn == 0 || curTurn < maxTurn); curTurn++ {
		fmt.Println("------------------------------")
		fmt.Printf("          NEW TURN (%d)\n", curTurn+1)
		fmt.Println("------------------------------")
		for _, p := range g.Players {
			g.ActivePlayer = p.Number
			move, err := p.GetMove(g)
			if err == nil {
				err = move.Apply(p, g)
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "[ERROR] Player %d (%s): %v\n", p.Number, p.Program, err)
				DefaultMove(p, g)
			}
			if p.VP >= MAX_VP {
				lastTurn = true
			}
		}
	}

	fmt.Println("------------------------------")
	fmt.Println("          END GAME")
	fmt.Println("------------------------------")

	return g.ComputeWinners()
}

func (g *Game) AddPlayers(pl []*Player) {
	for _, p := range pl {
		p.Number = uint(len(g.Players) + 1)
		g.Players = append(g.Players, p)
	}
}

func (g *Game) InitTokens() {
	var tokenCount uint
	switch len(g.Players) {
	case 2:
		tokenCount = 4
	case 3:
		tokenCount = 5
	case 4:
		tokenCount = 7
	}

	g.Tokens = make(ColorCount)
	for _, c := range Colors {
		g.Tokens[c] = tokenCount
	}
	g.Tokens[GOLD] = 5
}

func (g *Game) InitNobles() {

	nobles := Nobles
	numNobles := len(g.Players) + 1

	for numNobles > 0 {
		i := rand.Intn(len(nobles))
		n := nobles[i]
		nobles[i] = nobles[len(nobles)-1]
		nobles = nobles[:len(nobles)-1]
		g.Nobles = append(g.Nobles, n)
		numNobles--
	}
}

func (g *Game) InitDecks() {
	g.Decks = make(map[string]CardDeck)
	g.Decks[TIER1] = DeckTier1.Shuffle()
	g.Decks[TIER2] = DeckTier2.Shuffle()
	g.Decks[TIER3] = DeckTier3.Shuffle()
	g.DeckSizes = make(map[string]uint)
	g.DeckSizes[TIER1] = uint(len(g.Decks[TIER1]))
	g.DeckSizes[TIER2] = uint(len(g.Decks[TIER2]))
	g.DeckSizes[TIER3] = uint(len(g.Decks[TIER3]))

	slots := []string{
		TIER1 + "_1", TIER1 + "_2", TIER1 + "_3", TIER1 + "_4",
		TIER2 + "_1", TIER2 + "_2", TIER2 + "_3", TIER2 + "_4",
		TIER3 + "_1", TIER3 + "_2", TIER3 + "_3", TIER3 + "_4",
	}

	for _, slot := range slots {
		err := g.RevealCard(slot)
		if err != nil {
			panic(err)
		}
	}
}

func (g *Game) InitFirstPlayer() {
	i := rand.Intn(len(g.Players))
	pl := g.Players[i:]
	for j := 0; j < i; j++ {
		pl = append(pl, g.Players[j])
	}
	g.Players = pl
}

func (g *Game) ComputeWinners() []*Player {
	MaxVP := uint(0)
	for _, p := range g.Players {
		if p.VP > MaxVP {
			MaxVP = p.VP
		}
		fmt.Printf("Player %d (%s) has %d vp, %d cards\n", p.Number, p.Program, p.VP, p.NumberCards)
	}
	MinCards := uint(1000)
	for _, p := range g.Players {
		if p.VP == MaxVP && p.NumberCards < MinCards {
			MinCards = p.NumberCards
		}
	}
	winners := []*Player{}
	for _, p := range g.Players {
		if p.VP == MaxVP && p.NumberCards == MinCards {
			winners = append(winners, p)
		}
	}

	return winners
}

func (g *Game) RevealCard(slot string) error {
	if g.Cards == nil {
		g.Cards = make(map[string]*Card)
	}

	var prefix string
	switch {
	case strings.HasPrefix(slot, TIER1):
		prefix = TIER1
	case strings.HasPrefix(slot, TIER2):
		prefix = TIER2
	case strings.HasPrefix(slot, TIER3):
		prefix = TIER3
	}

	suffix := slot[len(prefix):]

	ok := false
	acceptedSuffixes := []string{"_1", "_2", "_3", "_4"}
	for _, s := range acceptedSuffixes {
		if suffix == s {
			ok = true
		}
	}

	if !ok {
		return fmt.Errorf("Bad card slot: %s", slot)
	}

	delete(g.Cards, slot)

	c := g.DrawCard(prefix)
	if c != nil {
		g.Cards[slot] = c
	}
	return nil
}

func (g *Game) DrawCard(deck string) *Card {
	if len(g.Decks[deck]) == 0 {
		return nil
	}

	i := rand.Intn(len(g.Decks[deck]))
	c := g.Decks[deck][i]
	g.Decks[deck][i] = g.Decks[deck][len(g.Decks[deck])-1]
	g.Decks[deck] = g.Decks[deck][:len(g.Decks[deck])]
	g.DeckSizes[deck]--

	return c
}

func (g *Game) GiveTokens(p *Player, col Color, num uint) {
	for g.Tokens[col] > 0 && num > 0 {
		g.Tokens[col]--
		p.Tokens[col]++
		p.NumberTokens++
		num--
	}
}

func (g *Game) DiscardExtra(toDiscard ColorCount, p *Player) {
	if p.NumberTokens <= MAX_TOKENS {
		return
	}

	numDiscard := p.NumberTokens - MAX_TOKENS

DISCARD_LOOP:
	for numDiscard > 0 {
		if p.NumberTokens > 15 {
			os.Exit(1)
		}
		for col, num := range toDiscard {
			if num > 0 {
				if p.Tokens[col] > 0 {
					p.Tokens[col]--
					p.NumberTokens--
					numDiscard--
					g.Tokens[col]++
					continue DISCARD_LOOP
				}
			}
		}
		for col, num := range p.Tokens {
			if num > 0 {
				p.Tokens[col]--
				p.NumberTokens--
				numDiscard--
				g.Tokens[col]++
				continue DISCARD_LOOP
			}
		}
		// Not reachable
		panic("FATAL: Discard loop")
	}
}
