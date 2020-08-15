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
		{Suit: Spades, Value: 0},
		{Suit: Spades, Value: 1},
		{Suit: Spades, Value: 2},
		{Suit: Spades, Value: 3},
		{Suit: Spades, Value: 4},
		{Suit: Spades, Value: 5},
		{Suit: Spades, Value: 6},
		{Suit: Spades, Value: 7},
		{Suit: Spades, Value: 8},
		{Suit: Spades, Value: 9},
		{Suit: Spades, Value: 10},
		{Suit: Spades, Value: 11},
		{Suit: Spades, Value: 12},
		{Suit: Diamonds, Value: 0},
		{Suit: Diamonds, Value: 1},
		{Suit: Diamonds, Value: 2},
		{Suit: Diamonds, Value: 3},
		{Suit: Diamonds, Value: 4},
		{Suit: Diamonds, Value: 5},
		{Suit: Diamonds, Value: 6},
		{Suit: Diamonds, Value: 7},
		{Suit: Diamonds, Value: 8},
		{Suit: Diamonds, Value: 9},
		{Suit: Diamonds, Value: 10},
		{Suit: Diamonds, Value: 11},
		{Suit: Diamonds, Value: 12},
		{Suit: Clubs, Value: 0},
		{Suit: Clubs, Value: 1},
		{Suit: Clubs, Value: 2},
		{Suit: Clubs, Value: 3},
		{Suit: Clubs, Value: 4},
		{Suit: Clubs, Value: 5},
		{Suit: Clubs, Value: 6},
		{Suit: Clubs, Value: 7},
		{Suit: Clubs, Value: 8},
		{Suit: Clubs, Value: 9},
		{Suit: Clubs, Value: 10},
		{Suit: Clubs, Value: 11},
		{Suit: Clubs, Value: 12},
		{Suit: Hearts, Value: 0},
		{Suit: Hearts, Value: 1},
		{Suit: Hearts, Value: 2},
		{Suit: Hearts, Value: 3},
		{Suit: Hearts, Value: 4},
		{Suit: Hearts, Value: 5},
		{Suit: Hearts, Value: 6},
		{Suit: Hearts, Value: 7},
		{Suit: Hearts, Value: 8},
		{Suit: Hearts, Value: 9},
		{Suit: Hearts, Value: 10},
		{Suit: Hearts, Value: 11},
		{Suit: Hearts, Value: 12},
	}

	return deck
}
