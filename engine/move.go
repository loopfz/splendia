package main

import (
	"errors"
	"fmt"

	"github.com/satori/go.uuid"
)

type Move struct {
	Action  string     `json:"action,omitempty"`
	Card    string     `json:"card,omitempty"`
	Tokens  ColorCount `json:"tokens,omitempty"`
	Discard ColorCount `json:"discard,omitempty"`
}

func (m *Move) Apply(p *Player, g *Game) error {
	switch m.Action {
	case BUY_CARD_ACT:
		return m.BuyCard(p, g)
	case RESERVE_CARD_ACT:
		return m.ReserveCard(p, g)
	case TOKENS_ACT:
		return m.TakeTokens(p, g)
	}
	return fmt.Errorf("Invalid action '%s'", m.Action)
}

func doBuyCard(c *Card, p *Player, g *Game) error {

	cost := p.GetCardCost(c)

	availGold := p.Tokens[GOLD]
	goldToPay := 0

	for col, num := range cost {
		availToks := p.Tokens[col]
		if num > availToks && col != GOLD {
			missing := num - availToks
			if missing > availGold {
				return fmt.Errorf("not enough %s tokens", col)
			}
			availGold -= missing
			goldToPay += int(missing)
		}
	}

	totalToPay := 0
	toPay := make(ColorCount)

	// TODO honor token preferred token use
	for col, num := range cost {
		availToks := p.Tokens[col]
		if availToks < num {
			num = availToks
		}
		toPay[col] = num
		totalToPay += int(num)
	}
	toPay[GOLD] = uint(goldToPay)

PAY_LOOP:
	for totalToPay > 0 {
		for col, num := range toPay {
			if num > 0 {
				p.Tokens[col]--
				p.NumberTokens--
				toPay[col]--
				totalToPay--
				g.Tokens[col]++
				continue PAY_LOOP
			}
		}
		// Not reachable
		panic("FATAL: Pay loop")
	}

	p.Cards[c.Color] = append(p.Cards[c.Color], c)
	p.NumberCards++
	p.VP += c.VP

	return nil
}

func CheckNobles(p *Player, g *Game) {

	// TODO honor noble choice
	for i, n := range g.Nobles {
		if p.CanGetNoble(n) {
			g.Nobles[i] = g.Nobles[len(g.Nobles)-1]
			g.Nobles[len(g.Nobles)-1] = nil
			g.Nobles = g.Nobles[:len(g.Nobles)-1]
			p.Nobles = append(p.Nobles, n)
			p.VP += n.VP
			fmt.Printf("Player %d (%s) wins noble %s (%v)!\n", p.Number, p.Program, n.Name, n)
			break
		}
	}
}

func (m *Move) BuyCard(p *Player, g *Game) error {

	c, ok := g.Cards[m.Card]
	if !ok {
		return fmt.Errorf("Cannot buy card %s: no such card", m.Card)
	}

	err := doBuyCard(c, p, g)
	if err != nil {
		return fmt.Errorf("Cannot buy card %s: %s", m.Card, err)
	}

	g.RevealCard(m.Card)

	fmt.Printf("Player %d (%s) buys card %s (%v)\n", p.Number, p.Program, m.Card, c)

	CheckNobles(p, g)

	return nil
}

func (m *Move) BuyReservedCard(p *Player, g *Game) error {

	c, ok := p.ReservedCards[m.Card]
	if !ok {
		return fmt.Errorf("Cannot buy card %s: no such card", m.Card)
	}

	err := doBuyCard(c, p, g)
	if err != nil {
		return fmt.Errorf("Cannot buy reserved card %s: %s", m.Card, err)
	}

	delete(p.ReservedCards, m.Card)

	fmt.Printf("Player %d (%s) buys reserved card %s (%v)\n", p.Number, p.Program, m.Card, c)

	CheckNobles(p, g)

	return nil
}

func (m *Move) ReserveCard(p *Player, g *Game) error {
	if len(p.ReservedCards) == MAX_RESERVED {
		return fmt.Errorf("Cannot reserve card %s: max reserved reached (%d)", m.Card, MAX_RESERVED)
	}
	c, ok := g.Cards[m.Card]
	if !ok {
		if m.Card == TIER1 || m.Card == TIER2 || m.Card == TIER3 {
			deck := m.Card
			c = g.DrawCard(deck)
			if c == nil {
				return fmt.Errorf("Cannot reserve random %s card: empty deck", deck)
			}
		} else {
			return fmt.Errorf("Cannot reserve card %s: no such card", m.Card)
		}
	} else {
		g.RevealCard(m.Card)
	}
	p.ReservedCards[uuid.NewV4().String()] = c
	//p.ReservedCards = append(p.ReservedCards, c)
	g.GiveTokens(p, GOLD, 1)
	g.DiscardExtra(m.Discard, p)

	fmt.Printf("Player %d (%s) reserves card %s (%v)\n", p.Number, p.Program, m.Card, c)
	return nil
}

func (m *Move) TakeTokens(p *Player, g *Game) error {

	numColors := 0
	numDouble := 0
	numTotal := 0
	for col, num := range m.Tokens {
		if col == GOLD && num > 0 {
			return fmt.Errorf("Cannot take tokens: tried to take %d %s tokens", num, col)
		}
		if g.Tokens[col] < num {
			return fmt.Errorf("Cannot take tokens: tried to take %d %s tokens, %d available.", num, col, g.Tokens[col])
		}
		if num > 2 {
			return fmt.Errorf("Cannot take tokens: tried to take %d %s tokens", num, col)
		}
		if num == 2 {
			if g.Tokens[col] < 4 {
				return fmt.Errorf("Cannot take tokens: tried to take double %s tokens, only %d remaining", col, g.Tokens[col])
			}
			numDouble++
		}
		if num > 0 {
			numColors++
		}
		numTotal += int(num)
	}

	if numColors > 1 && numDouble > 0 {
		return errors.New("Cannot take tokens: tried to take double tokens of several colors")
	}
	if numTotal > 3 {
		return fmt.Errorf("Cannot take tokens: attempted to take too many (%d)", numTotal)
	}

	for col, num := range m.Tokens {
		g.GiveTokens(p, col, num)
	}

	fmt.Printf("Player %d (%s) takes tokens (%v)\n", p.Number, p.Program, m.Tokens)

	g.DiscardExtra(m.Discard, p)

	return nil
}

func DefaultMove(p *Player, g *Game) error {

	taken := 0
	toTake := make(ColorCount)
	for col, num := range g.Tokens {
		if col != GOLD && num > 0 {
			if taken < 3 {
				taken++
				toTake[col]++
			}
		}
	}
	if taken > 0 {
		return (&Move{Tokens: toTake}).TakeTokens(p, g)
	}

	decks := []string{TIER1, TIER2, TIER3}

	for _, d := range decks {
		err := (&Move{Card: d}).ReserveCard(p, g)
		if err == nil {
			return nil
		}
	}

	for cName, _ := range p.ReservedCards {
		err := (&Move{Card: cName}).BuyReservedCard(p, g)
		if err == nil {
			return nil
		}
	}

	for cName, _ := range g.Cards {
		err := (&Move{Card: cName}).BuyCard(p, g)
		if err == nil {
			return nil
		}
	}

	fmt.Printf("WARNING! Player %d (%s) had no valid default move to make\n", p.Number, p.Program)
	return nil
}
