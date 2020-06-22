package engine

import (
	"errors"
	"math/rand"
	"time"
)

// TODO:
// correct values on stdDeck, deal needs to use stdDeck

var stdDeck = []Card{
	{value: 0, suit: Spades},
	{value: 0, suit: Spades},
	{value: 4, suit: Spades},
	{value: 9, suit: Spades},
	{value: 0, suit: Spades},
	{value: 4, suit: Spades},
	{value: 9, suit: Spades},
	{value: 0, suit: Spades},
	{value: 4, suit: Spades},
	{value: 9, suit: Spades},
	{value: 4, suit: Spades},
	{value: 9, suit: Spades},
	{value: 0, suit: Spades},
	{value: 0, suit: Diamonds},
	{value: 1, suit: Diamonds},
	{value: 3, suit: Diamonds},
	{value: 5, suit: Diamonds},
	{value: 9, suit: Diamonds},
	{value: 0, suit: Diamonds},
	{value: 1, suit: Diamonds},
	{value: 3, suit: Diamonds},
	{value: 5, suit: Diamonds},
	{value: 9, suit: Diamonds},
	{value: 3, suit: Diamonds},
	{value: 5, suit: Diamonds},
	{value: 9, suit: Diamonds},
	{value: 1, suit: Clubs},
	{value: 2, suit: Clubs},
	{value: 13, suit: Clubs},
	{value: 13, suit: Clubs},
	{value: 13, suit: Clubs},
	{value: 13, suit: Clubs},
	{value: 13, suit: Clubs},
	{value: 13, suit: Clubs},
	{value: 13, suit: Clubs},
	{value: 13, suit: Clubs},
	{value: 13, suit: Clubs},
	{value: 13, suit: Clubs},
	{value: 13, suit: Hearts},
	{value: 1, suit: Hearts},
	{value: 2, suit: Hearts},
	{value: 13, suit: Hearts},
	{value: 13, suit: Hearts},
	{value: 13, suit: Hearts},
	{value: 13, suit: Hearts},
	{value: 13, suit: Hearts},
	{value: 13, suit: Hearts},
	{value: 13, suit: Hearts},
	{value: 13, suit: Hearts},
	{value: 13, suit: Hearts},
	{value: 13, suit: Hearts},
	{value: 13, suit: Hearts},
}

// TODO: change name of func to dealHand, will have to change in unit tests as well
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
