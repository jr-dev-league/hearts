package engine

import (
	"testing"
)

const (
	playerOne   = 0
	playerTwo   = 1
	playerThree = 2
	playerFour  = 3
)

var handOne = []Card{
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

var handTwo = []Card{
	{Value: 1, Suit: Spades},
	{Value: 8, Suit: Diamonds},
	{Value: 12, Suit: Diamonds},
	{Value: 0, Suit: Clubs},
	{Value: 3, Suit: Clubs},
	{Value: 4, Suit: Clubs},
	{Value: 6, Suit: Clubs},
	{Value: 7, Suit: Clubs},
	{Value: 8, Suit: Clubs},
	{Value: 1, Suit: Hearts},
	{Value: 3, Suit: Hearts},
	{Value: 4, Suit: Hearts},
	{Value: 5, Suit: Hearts},
}
var handThree = []Card{
	{Value: 2, Suit: Spades},
	{Value: 5, Suit: Spades},
	{Value: 4, Suit: Spades},
	{Value: 6, Suit: Hearts},
	{Value: 10, Suit: Hearts},
	{Value: 11, Suit: Clubs},
	{Value: 5, Suit: Clubs},
	{Value: 9, Suit: Clubs},
	{Value: 11, Suit: Clubs},
	{Value: 6, Suit: Clubs},
	{Value: 10, Suit: Diamonds},
	{Value: 11, Suit: Diamonds},
	{Value: 12, Suit: Diamonds},
}
var handFour = []Card{
	{Value: 3, Suit: Spades},
	{Value: 7, Suit: Spades},
	{Value: 7, Suit: Spades},
	{Value: 8, Suit: Spades},
	{Value: 10, Suit: Spades},
	{Value: 11, Suit: Spades},
	{Value: 12, Suit: Spades},
	{Value: 2, Suit: Diamonds},
	{Value: 7, Suit: Diamonds},
	{Value: 10, Suit: Diamonds},
	{Value: 2, Suit: Hearts},
	{Value: 7, Suit: Hearts},
	{Value: 9, Suit: Hearts},
}

func TestNewGameState(t *testing.T) {
	game := New()
	_, err := game.Player(0)

	for i := range game.Players {
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

	if !handsEq(handOne, firstPlayer.Hand) {
		t.Errorf("expected:\n\n%v\nfound:%v\n", handOne, firstPlayer.Hand)
	}

	for i := 1; i < len(game.Players); i++ {
		player := game.Players[i]
		if player.Points != 0 && player.Hand != nil {
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

	for _, card := range player.Hand {
		if card.Exposed || card.Played {
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

	for i, card := range player.Hand {
		if i == cardIndex && (!card.Played || !card.Exposed) {
			t.Error("cards should be be playable")
		} else if i != cardIndex && (card.Played || card.Exposed) {
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

	for i, player := range view.Players {
		if player.CardCount != maxHandSize {
			t.Errorf("player %d should have %d cards", i, maxHandSize)
		}

		if i == playerThree {

			if !handsEq(player.Hand, handThree) {
				t.Errorf("expected:\n\n%v\nfound:%v\n", handThree, player.Hand)
			}

		} else if len(player.Hand) != 0 {
			t.Errorf("opponent hands should not be viewable. saw %v", player.Hand)
		}
	}

	playedCard := handFour[5]

	game.PlayUp(playerFour, playedCard)
	view = game.ViewAs(playerThree)

	if len(view.Players[playerFour].Hand) != 1 {
		t.Error("player should only have one visible card")
	}

	actual := view.Players[playerFour].Hand[0]

	if actual.Suit != playedCard.Suit || actual.Value != playedCard.Value {
		t.Error("the wrong card is showing")
	}
}

func TestDeal(t *testing.T) {
	game := New()

	err := game.Deal()

	const compareIdx = 1

	var firstDeal []Card

	if err != nil {
		t.Error("Deal returned an unexpected error")
	}

	for i, player := range game.Players {
		if i == compareIdx {
			firstDeal = player.Hand
		}

		if len(player.Hand) != 13 {
			t.Errorf("expected hand length to be 13, received %d", len(player.Hand))
		}
	}

	err = game.Deal()

	if err == nil {
		t.Error("expected Deal to throw an error, but it did not")
	}

	game = New()
	game.Deal()

	for i, player := range game.Players {
		if i == compareIdx {
			if handsEq(player.Hand, firstDeal) {
				t.Error("shuffle should shuffle hands")
			}
		}

		if len(player.Hand) != 13 {
			t.Errorf("expected hand length to be 13, received %d", len(player.Hand))
		}
	}
}

func TestDiscard(t *testing.T) {
	game := New()

	game.SetPlayer(0, 0, maxHandSize, cloneHand(handOne))

	spadesAce := Card{true, true, Spades, 0}
	err := game.Discard(0, spadesAce)
	if err != nil {
		t.Error("expected no error, but returned one")
	}

	expected := []Card{
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

	actual := game.Players[0].Hand

	if !handsEq(actual, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, actual)
	}

	expectedCardCount := uint8(12)
	actualCardCount := game.Players[0].CardCount

	if expectedCardCount != actualCardCount {
		t.Errorf("expected cardCount of %d, actual: %d", expectedCardCount, actualCardCount)
	}

	heartsEight := Card{true, true, Hearts, 8}

	err = game.Discard(0, heartsEight)
	if err != nil {
		t.Error("expected no error, but returned one")
	}

	expected = []Card{
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
	}

	actual = game.Players[0].Hand

	if !handsEq(actual, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, actual)
	}

	expectedCardCount = uint8(11)
	actualCardCount = game.Players[0].CardCount

	if expectedCardCount != actualCardCount {
		t.Errorf("expected cardCount of %d, actual: %d", expectedCardCount, actualCardCount)
	}

	diamondsFive := Card{true, true, Diamonds, 5}
	game.Discard(0, diamondsFive)

	expected = []Card{
		{Value: 4, Suit: Spades},
		{Value: 9, Suit: Spades},
		{Value: 0, Suit: Diamonds},
		{Value: 1, Suit: Diamonds},
		{Value: 3, Suit: Diamonds},
		{Value: 9, Suit: Diamonds},
		{Value: 1, Suit: Clubs},
		{Value: 2, Suit: Clubs},
		{Value: 13, Suit: Clubs},
		{Value: 0, Suit: Hearts},
	}

	actual = game.Players[0].Hand

	if !handsEq(actual, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, actual)
	}

	expectedCardCount = uint8(10)
	actualCardCount = game.Players[0].CardCount

	if expectedCardCount != actualCardCount {
		t.Errorf("expected cardCount of %d, actual: %d", expectedCardCount, actualCardCount)
	}

	expected = []Card{
		{Value: 4, Suit: Spades},
	}

	game.SetPlayer(1, 0, 1, expected)

	err = game.Discard(1, spadesAce)
	if err == nil {
		t.Error("expected error but did not receive one")
	}

	actual = game.Players[1].Hand

	if !handsEq(actual, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, actual)
	}

	spadesFour := Card{true, true, Spades, 4}
	err = game.Discard(1, spadesFour)
	if err != nil {
		t.Error("expected no error, but returned one")
	}

	actual = game.Players[1].Hand
	expected = []Card{}
	if !handsEq(actual, expected) {
		t.Errorf("expected:\n%v\nfound:\n%v", expected, actual)
	}

	err = game.Discard(1, spadesFour)
	if err == nil {
		t.Error("expected error but did not receive one")
	}
}

func TestDiscardAll(t *testing.T) {
	game := New()

	game.SetPlayer(0, 0, maxHandSize, cloneHand(handOne))
	game.SetPlayer(1, 0, maxHandSize, cloneHand(handTwo))
	game.SetPlayer(2, 0, maxHandSize, cloneHand(handThree))
	game.SetPlayer(3, 0, maxHandSize, cloneHand(handFour))

	game.PlayUp(0, handOne[0])
	game.PlayUp(1, handTwo[0])
	game.PlayUp(2, handThree[0])
	game.PlayUp(3, handFour[0])

	actualDiscarded := game.DiscardPlayed()

	expectedDiscarded := []Card{
		handOne[0],
		handTwo[0],
		handThree[0],
		handFour[0],
	}

	for i := range expectedDiscarded {
		expectedDiscarded[i].Played = true
		expectedDiscarded[i].Exposed = true
	}

	if !handsEq(actualDiscarded, expectedDiscarded) {
		t.Errorf("expected:\n%v\nfound:\n%v", expectedDiscarded, actualDiscarded)

	}

	actualDiscarded = game.DiscardPlayed()
	expectedDiscarded = []Card{}

	if !handsEq(actualDiscarded, expectedDiscarded) {
		t.Errorf("expected:\n%v\nfound:\n%v", expectedDiscarded, actualDiscarded)
	}
}

// PRIVATE HELPER FUNCTIONS

func handsEq(actual []Card, expected []Card) bool {
	// If the lengths don't match, they must be different.
	if len(actual) != len(expected) {
		return false
	}

	// If each card matches, and the length is the same, they must match.
	for i, card := range actual {
		if card != expected[i] {
			return false
		}
	}

	return true
}

// Even though arrays can be copied by value, an array of struct is an array
// of pointers to structs, so copying the array doesn't change the underlying
// structs. We don't want to edit the test structs at the top of this file, so
// we need to clone those arrays to keep them safe!
func cloneHand(hand []Card) []Card {
	clone := make([]Card, 0, maxHandSize)
	for _, card := range hand {
		cardClone := Card{}

		cardClone.Exposed = card.Exposed
		cardClone.Played = card.Played
		cardClone.Suit = card.Suit
		cardClone.Value = card.Value

		clone = append(clone, cardClone)
	}

	return clone
}
