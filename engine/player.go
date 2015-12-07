package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Player struct {
	Program       string                              `json:"-"`
	Number        uint                                `json:"number"`
	Tokens        ColorCount                          `json:"tokens"`
	Cards         map[Color][]*Card                   `json:"cards"`
	ReservedCards []*Card                             `json:"reserved_cards"`
	Nobles        []*Noble                            `json:"nobles"`
	NumberCards   uint                                `json:"number_cards"`
	NumberTokens  uint                                `json:"number_tokens"`
	VP            uint                                `json:"vp"`
	Data          map[string]interface{}              `json:"data"`
	executor      func(*Player, *Game) (*Move, error) `json:"-"`
}

func NewForkPlayer(bin string) *Player {
	return &Player{
		Program:       bin,
		Tokens:        ColorCount{},
		Cards:         map[Color][]*Card{},
		ReservedCards: []*Card{},
		Nobles:        []*Noble{},
		executor:      (*Player).ForkMove,
	}
}

func (p *Player) GetMove(g *Game) (*Move, error) {
	if p.executor == nil {
		return nil, errors.New("No executor set")
	}
	return p.executor(p, g)
}

func (p *Player) ForkMove(g *Game) (*Move, error) {

	cmd := exec.Command(p.Program)

	in := &bytes.Buffer{}
	out := &bytes.Buffer{}

	j, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	in.Write(j)
	cmd.Stdin = in
	cmd.Stdout = out

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(PLAYER_TIMEOUT):
		if err := cmd.Process.Kill(); err != nil {
			fmt.Fprintf(os.Stderr, "Player %s (%d): Failed to kill subprocess after timeout: %s", p.Number, p.Program, err)
		}
		<-done
		return nil, errors.New("Timeout")
	case err := <-done:
		if err != nil {
			return nil, err
		}
	}

	move := &Move{}
	err = json.Unmarshal(out.Bytes(), move)
	if err != nil {
		return nil, err
	}

	return move, nil
}

func (p *Player) GetCardCost(c *Card) ColorCount {
	ret := make(ColorCount)

	for col, num := range c.Cost {
		if uint(len(p.Cards[col])) >= num {
			num = 0
		} else {
			num -= uint(len(p.Cards[col]))
		}
		ret[col] = num
	}

	return ret
}
