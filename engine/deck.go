package engine

import (
	"errors"
	"math/rand"
	"time"
)

func dealHand(numDealt uint, deck *[]Card) (hand []Card, err error) {
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

func stdDeck() (deck []Card) {
	deck = []Card{
		{suit: Spades, value: 0},
		{suit: Spades, value: 1},
		{suit: Spades, value: 2},
		{suit: Spades, value: 3},
		{suit: Spades, value: 4},
		{suit: Spades, value: 5},
		{suit: Spades, value: 6},
		{suit: Spades, value: 7},
		{suit: Spades, value: 8},
		{suit: Spades, value: 9},
		{suit: Spades, value: 10},
		{suit: Spades, value: 11},
		{suit: Spades, value: 12},
		{suit: Diamonds, value: 0},
		{suit: Diamonds, value: 1},
		{suit: Diamonds, value: 2},
		{suit: Diamonds, value: 3},
		{suit: Diamonds, value: 4},
		{suit: Diamonds, value: 5},
		{suit: Diamonds, value: 6},
		{suit: Diamonds, value: 7},
		{suit: Diamonds, value: 8},
		{suit: Diamonds, value: 9},
		{suit: Diamonds, value: 10},
		{suit: Diamonds, value: 11},
		{suit: Diamonds, value: 12},
		{suit: Clubs, value: 0},
		{suit: Clubs, value: 1},
		{suit: Clubs, value: 2},
		{suit: Clubs, value: 3},
		{suit: Clubs, value: 4},
		{suit: Clubs, value: 5},
		{suit: Clubs, value: 6},
		{suit: Clubs, value: 7},
		{suit: Clubs, value: 8},
		{suit: Clubs, value: 9},
		{suit: Clubs, value: 10},
		{suit: Clubs, value: 11},
		{suit: Clubs, value: 12},
		{suit: Hearts, value: 0},
		{suit: Hearts, value: 1},
		{suit: Hearts, value: 2},
		{suit: Hearts, value: 3},
		{suit: Hearts, value: 4},
		{suit: Hearts, value: 5},
		{suit: Hearts, value: 6},
		{suit: Hearts, value: 7},
		{suit: Hearts, value: 8},
		{suit: Hearts, value: 9},
		{suit: Hearts, value: 10},
		{suit: Hearts, value: 11},
		{suit: Hearts, value: 12},
	}

	return deck
}
