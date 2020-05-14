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
	
	hand, err := Deal(0, deck)

	if err != nil {
		t.Error("expected no error, but returned one")
	}

	if err := compareHand(*deck, expected); err != nil {
		t.Error(err)
	}

	expected = []Card{}
	if err := compareHand(hand, expected); err != nil {
		t.Error(err)
	}

	// test dealing 11
	expected = []Card{
		{value: 0, suit: Hearts},
		{value: 8, suit: Hearts},
	}

	hand, err = Deal(11, deck)

	if err != nil {
		t.Error("expected no error, but returned one")
	}

	if err := compareHand(*deck, expected); err != nil {
		t.Error(err)
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

	if err := compareHand(hand, expected); err != nil {
		t.Error(err)
	}

	// test dealing to a zero card deck
	expected = []Card{}

	hand, err = Deal(2, deck)

	if err != nil {
		t.Error("expected no error, but returned one")
	}

	if err := compareHand(*deck, expected); err != nil {
		t.Error(err)
	}

	expected = []Card{
		{value: 0, suit: Hearts},
		{value: 8, suit: Hearts},
	}

	if err := compareHand(hand, expected); err != nil {
		t.Error(err)
	}

	// test error handling
	expected = []Card{}

	hand, err = Deal(2, deck)

	if err == nil {
		t.Error("expected an error about deleting when not enough cards, but returned none")
	}

	if err := compareHand(*deck, expected); err != nil {
		t.Error(err)
	}

	if err := compareHand(hand, expected); err != nil {
		t.Error(err)
	}
}
