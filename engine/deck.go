package engine

import (
	"errors"
	"math/rand"
	"time"
)

// Deal takes a pointer to deck (slice) of Cards and returns a slice of numDealt
// Cards from it. If the slice of cards is not big enough, it returns an empty
// slice and an error. Deck is reduced in size by numDealt.
func Deal(numDealt uint, deck *[]Card) (hand []Card, err error) {
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
