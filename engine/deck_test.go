package engine

import "testing"

func TestDealHand(t *testing.T) {

	// test dealing 0
	d := []Card{
		{Value: 0, Suit: Spades},
		{Value: 4, Suit: Spades},
		{Value: 9, Suit: Spades},
		{Value: 0, Suit: Diamonds},
		{Value: 1, Suit: Diamonds},
		{Value: 3, Suit: Diamonds},
		{Value: 5, Suit: Diamonds},
		{Value: 9, Suit: Diamonds},
		{Value: 1, Suit: Clubs},
		{Value: 2, Suit: Clubs},
		{Value: 13, Suit: Clubs},
		{Value: 0, Suit: Hearts},
		{Value: 8, Suit: Hearts},
	}
	deck := &d
	expected := cloneHand(*deck)

	hand, err := dealHand(0, deck)

	if err != nil {
		t.Error("expected no error, but returned one")
	}

	if !handsEq(*deck, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, *deck)

	}

	expected = []Card{}
	if !handsEq(hand, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, hand)
	}

	// test dealing 11
	expected = []Card{
		{Value: 0, Suit: Hearts},
		{Value: 8, Suit: Hearts},
	}

	hand, err = dealHand(11, deck)

	if err != nil {
		t.Error("expected no error, but returned one")
	}

	if !handsEq(*deck, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, *deck)
	}

	expected = []Card{
		{Value: 0, Suit: Spades},
		{Value: 4, Suit: Spades},
		{Value: 9, Suit: Spades},
		{Value: 0, Suit: Diamonds},
		{Value: 1, Suit: Diamonds},
		{Value: 3, Suit: Diamonds},
		{Value: 5, Suit: Diamonds},
		{Value: 9, Suit: Diamonds},
		{Value: 1, Suit: Clubs},
		{Value: 2, Suit: Clubs},
		{Value: 13, Suit: Clubs},
	}

	if !handsEq(hand, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, hand)
	}

	// test dealing to a zero card deck
	expected = []Card{}

	hand, err = dealHand(2, deck)

	if err != nil {
		t.Error("expected no error, but returned one")
	}

	if !handsEq(*deck, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, *deck)
	}

	expected = []Card{
		{Value: 0, Suit: Hearts},
		{Value: 8, Suit: Hearts},
	}

	if !handsEq(hand, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, hand)
	}

	// test error handling
	expected = []Card{}

	hand, err = dealHand(2, deck)

	if err == nil {
		t.Error("expected an error about deleting when not enough cards, but returned none")
	}

	if !handsEq(*deck, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, *deck)
	}

	if !handsEq(hand, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, hand)
	}
}

func TestShuffle(t *testing.T) {
	deck := cloneHand(handOne)

	shuffle(deck)

	if handsEq(deck, handOne) {
		t.Error("Expected cards to be shuffled, but they are unchanged.")
	}

	deck2 := cloneHand(deck)

	shuffle(deck)

	if handsEq(deck2, deck) {
		t.Error("Expected cards to be shuffled, but they are unchanged.")
	}

	if handsEq(deck2, handOne) {
		t.Error("Expected cards to be shuffled, but they are unchanged.")
	}
}
