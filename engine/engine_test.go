package engine

import (
	"fmt"
	"testing"
)

var handOne = []Card{
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

var handTwo = []Card{
	{value: 1, suit: Spades},
	{value: 8, suit: Diamonds},
	{value: 12, suit: Diamonds},
	{value: 0, suit: Clubs},
	{value: 3, suit: Clubs},
	{value: 4, suit: Clubs},
	{value: 6, suit: Clubs},
	{value: 7, suit: Clubs},
	{value: 8, suit: Clubs},
	{value: 1, suit: Hearts},
	{value: 3, suit: Hearts},
	{value: 4, suit: Hearts},
	{value: 5, suit: Hearts},
}
var handThree = []Card{
	{value: 2, suit: Spades},
	{value: 5, suit: Spades},
	{value: 4, suit: Spades},
	{value: 6, suit: Hearts},
	{value: 10, suit: Hearts},
	{value: 11, suit: Clubs},
	{value: 5, suit: Clubs},
	{value: 9, suit: Clubs},
	{value: 11, suit: Clubs},
	{value: 6, suit: Clubs},
	{value: 10, suit: Diamonds},
	{value: 11, suit: Diamonds},
	{value: 12, suit: Diamonds},
}
var handFour = []Card{
	{value: 3, suit: Spades},
	{value: 7, suit: Spades},
	{value: 7, suit: Spades},
	{value: 8, suit: Spades},
	{value: 10, suit: Spades},
	{value: 11, suit: Spades},
	{value: 12, suit: Spades},
	{value: 2, suit: Diamonds},
	{value: 7, suit: Diamonds},
	{value: 10, suit: Diamonds},
	{value: 2, suit: Hearts},
	{value: 7, suit: Hearts},
	{value: 9, suit: Hearts},
}

func TestNewGameState(t *testing.T) {
	game := New()
	_, err := game.Player(0)

	for i := range game.players {
		_, err = game.Player(uint8(i))
		if err == nil {
			t.Error("should not be able to get an unset hand")
		}
	}

}

func TestSetAndGet(t *testing.T) {
	game := New()

	err := game.SetPlayer(0, 0, handOne)

	if err != nil {
		t.Error("should be able to set an unset hand")
	}

	firstPlayer, err := game.Player(0)

	if err != nil {
		t.Error("should be able to get a set player")
	}

	err = game.SetPlayer(0, 0, handFour)

	if err == nil {
		t.Error("should not be able to set a player who is already set")
	}

	err = compareHand(handOne, firstPlayer.hand)

	if err != nil {
		t.Error(err)
	}

	for i := 1; i < len(game.players); i++ {
		player := game.players[i]
		if player.points != 0 && player.hand != nil {
			t.Error("players that were not set should have zeroed values")
		} else if _, err := game.Player(uint8(i)); err == nil {
			t.Error("should not be able to get players that are not set")
		}
	}
}

func TestPlayCard(t *testing.T) {
	game := New()
	cardIndex := 3

	game.SetPlayer(0, 0, handOne)
	game.SetPlayer(1, 0, handTwo)
	game.SetPlayer(2, 0, handThree)
	game.SetPlayer(3, 0, handFour)

	player, _ := game.Player(0)

	for _, card := range player.hand {
		if card.exposed || card.played {
			t.Error("Cards should be dealt unplayed and unexposed")
		}
	}

	err := game.PlayUp(0, handTwo[cardIndex+1])

	if err == nil {
		t.Error("should only be able to play a card in the player's hand")
	}

	err = game.PlayUp(0, handOne[cardIndex])

	if err != nil {
		t.Error("should be able to play a card in player's hand")
	}

	for i, card := range player.hand {
		if i == cardIndex && (!card.played || !card.exposed) {
			t.Error("cards should be be playable")
		} else if i != cardIndex && (card.played || card.exposed) {
			t.Error("only the selected card should be played")
		}
	}

	err = game.PlayUp(0, handOne[cardIndex])

	if err == nil {
		t.Error("should not be able to play the same card twice")
	}
}

func TestViewAs(t *testing.T) {
	game := New()

	game.SetPlayer(3, 0, handFour)

}

// PRIVATE HELPER FUNCTIONS

func compareHand(actual []Card, expected []Card) error {
	// If the lengths don't match, they must be different.
	if len(actual) != len(expected) {
		return fmt.Errorf("expected:\n%v\nfound:\n%v", expected, actual)
	}

	// If each card matches, and the length is the same, they must match.
	for i, card := range actual {
		if card.value != handOne[i].value {
			return fmt.Errorf("expected:\n%v\nfound:\n%v", expected, actual)
		}
	}

	return nil
}
