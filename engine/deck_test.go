package engine

import "testing"

func TestDeal(t *testing.T) {

	// test dealing 0
	d := []Card{
		{value: 0, suit: Spades},
		{value: 4, suit: Spades},
		{value: 9, suit: Spades},
		{value: 0, suit: Diamonds},
		{value: 1, suit: Diamonds},
		{value: 3, suit: Diamonds},
		{value: 5, suit: Diamonds},
		{value: 9, suit: Diamonds},
		{value: 1, suit: Clubs},
		{value: 2, suit: Clubs},
		{value: 13, suit: Clubs},
		{value: 0, suit: Hearts},
		{value: 8, suit: Hearts},
	}
	deck := &d
	expected := cloneHand(*deck)

	hand, err := deal(0, deck)

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
		{value: 0, suit: Hearts},
		{value: 8, suit: Hearts},
	}

	hand, err = deal(11, deck)

	if err != nil {
		t.Error("expected no error, but returned one")
	}

	if !handsEq(*deck, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, *deck)
	}

	expected = []Card{
		{value: 0, suit: Spades},
		{value: 4, suit: Spades},
		{value: 9, suit: Spades},
		{value: 0, suit: Diamonds},
		{value: 1, suit: Diamonds},
		{value: 3, suit: Diamonds},
		{value: 5, suit: Diamonds},
		{value: 9, suit: Diamonds},
		{value: 1, suit: Clubs},
		{value: 2, suit: Clubs},
		{value: 13, suit: Clubs},
	}

	if !handsEq(hand, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, hand)
	}

	// test dealing to a zero card deck
	expected = []Card{}

	hand, err = deal(2, deck)

	if err != nil {
		t.Error("expected no error, but returned one")
	}

	if !handsEq(*deck, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, *deck)
	}

	expected = []Card{
		{value: 0, suit: Hearts},
		{value: 8, suit: Hearts},
	}

	if !handsEq(hand, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, hand)
	}

	// test error handling
	expected = []Card{}

	hand, err = deal(2, deck)

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

func Test_shuffle(t *testing.T) {
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
