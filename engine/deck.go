package main

import "math/rand"

type CardDeck []*Card

func (d CardDeck) Shuffle() CardDeck {
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
	return d
}
