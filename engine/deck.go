package engine

import (
	"errors"
	"math/rand"
	"time"
)

func deal(numDealt uint, deck *[]Card) (hand []Card, err error) {
	d := *deck
	if len(d) < int(numDealt) {
		return hand, errors.New("Deal: not enough cards in deck")
	}
	hand = d[:numDealt]
	*deck = d[numDealt:]
	return
}

func shuffle(cards []Card) {
	src := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(src)
	for i := len(cards) - 1; i >= 0; i-- {
		rIdx := gen.Intn(327) % (i + 1)
		swap(cards[:], rIdx, i)
	}
}

func swap(deck []Card, i int, j int) {
	tmp := deck[j]
	deck[j] = deck[i]
	deck[i] = tmp
}
