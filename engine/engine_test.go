package engine

import (
	"fmt"
	"testing"
)

const (
	playerOne   = 0
	playerTwo   = 1
	playerThree = 2
	playerFour  = 3
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

	err := game.SetPlayer(0, 0, maxHandSize, cloneHand(handOne))

	if err != nil {
		t.Error("should be able to set an unset hand")
	}

	firstPlayer, err := game.Player(0)

	if err != nil {
		t.Error("should be able to get a set player")
	}

	err = game.SetPlayer(0, 0, maxHandSize, cloneHand(handFour))

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

	game.SetPlayer(0, 0, maxHandSize, cloneHand(handOne))
	game.SetPlayer(1, 0, maxHandSize, cloneHand(handTwo))
	game.SetPlayer(2, 0, maxHandSize, cloneHand(handThree))
	game.SetPlayer(3, 0, maxHandSize, cloneHand(handFour))

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

	game.SetPlayer(0, 0, maxHandSize, cloneHand(handOne))
	game.SetPlayer(1, 0, maxHandSize, cloneHand(handTwo))
	game.SetPlayer(2, 0, maxHandSize, cloneHand(handThree))
	game.SetPlayer(3, 0, maxHandSize, cloneHand(handFour))

	view := game.ViewAs(playerThree)

	err := view.PlayUp(2, handThree[3])

	if err == nil {
		t.Error("should not be able to play cards in a readonly state")
	}

	err = view.SetPlayer(0, 10, 12, handOne)

	if err == nil {
		t.Error("should not be able to set players in a readonly state")
	}

	for i, player := range view.players {
		if player.cardCount != maxHandSize {
			t.Errorf("player %d should have %d cards", i, maxHandSize)
		}

		if i == playerThree {
			err := compareHand(player.hand, handThree)

			if err != nil {
				t.Errorf("expected:\n\n%v\nfound:%v\n", handThree, player.hand)
			}

		} else if len(player.hand) != 0 {
			t.Errorf("oppenent hands should not be viewable. saw %v", player.hand)
		}
	}

	playedCard := handFour[5]

	game.PlayUp(playerFour, playedCard)
	view = game.ViewAs(playerThree)

	if len(view.players[playerFour].hand) != 1 {
		t.Error("player should only have one visible card")
	}

	actual := view.players[playerFour].hand[0]

	if err := compareCard(actual, playedCard); err != nil {
		t.Error("the wrong card is showing")
	}
}

// PRIVATE HELPER FUNCTIONS

func compareHand(actual []Card, expected []Card) error {
	// If the lengths don't match, they must be different.
	if len(actual) != len(expected) {
		return fmt.Errorf("expected:\n%v\nfound:\n%v", expected, actual)
	}

	// If each card matches, and the length is the same, they must match.
	for i, card := range actual {
		return compareCard(card, expected[i])
	}

	return nil
}

func compareCard(actual Card, expected Card) error {
	if actual.value != expected.value {
		return fmt.Errorf("expected:\n%v\nfound:\n%v", expected, actual)
	}

	return nil
}

// Even though arrays can be copied by value, an array of struct is an array
// of pointers to structs, so copying the array doesn't change the underlying
// structs. We don't want to edit the test structs at the top of this file, so
// we need to clone those arrays to keep them safe!
func cloneHand(hand []Card) []Card {
	clone := make([]Card, 0, maxHandSize)
	for _, card := range hand {
		cardClone := Card{}

		cardClone.exposed = card.exposed
		cardClone.played = card.played
		cardClone.suit = card.suit
		cardClone.value = card.value

		clone = append(clone, cardClone)
	}

	return clone
}
