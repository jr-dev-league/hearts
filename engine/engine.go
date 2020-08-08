package engine

import "errors"

// New creates a new game status
func New() (game State) {
	game = State{
		broken:    false,
		players:   [4]Player{},
		takenLast: 4,
		shootable: true,
		readonly:  false,
	}

	return
}

// SetPlayer sets a player at the given index
func (game *State) SetPlayer(i uint8, points int8, cardCount uint8, hand []Card) error {
	if game.readonly {
		return errors.New("cannot edit a readonly game")
	}

	if game.players[i].hand != nil {
		return errors.New("this player has already been set")
	}

	game.players[i] = Player{cardCount, hand, points}

	return nil
}

// Player returns a player by the given index
func (game *State) Player(i uint8) (*Player, error) {
	hand := game.players[i].hand

	if hand == nil {
		return nil, errors.New("this player had not been set")
	}

	return &game.players[i], nil
}

// ViewAs returns a game state as known by a given (by index) player
func (game *State) ViewAs(p uint8) (view State) {
	view = New()

	for i := uint8(0); i < uint8(len(game.players)); i++ {
		player := &game.players[i]
		if i == p {
			view.SetPlayer(i, player.points, maxHandSize, player.hand)
		} else {
			hand := make([]Card, 0, maxHandSize)

			for c := range game.players[i].hand {
				card := game.players[i].hand[c]

				if card.played {
					hand = append(hand, card)
				}
			}

			view.SetPlayer(i, player.points, maxHandSize, hand)
		}
	}

	view.readonly = true

	return
}

// PlayUp plays one card face up by player index, p, and card index, c.
// Because played cards are still owned by the player that played them,
// there is no need to remove it from their hand. It should be up to the
// view to display the card in the middle of the table.
func (game *State) PlayUp(p uint8, c Card) error {
	if game.readonly {
		return errors.New("cannot edit a readonly game")
	}

	player := game.players[p]
	i, err := findCard(player.hand, c)

	if err != nil {
		return err
	}

	if player.hand[i].played {
		return errors.New("selected card was already played")
	}

	player.hand[i].played = true
	player.hand[i].exposed = true
	player.cardCount--

	return nil
}

// Deal shuffles a deck of cards and deals it out to each player.
func (game *State) Deal() error {
	deck := stdDeck()

	shuffle(deck)

	for i := 0; i < len(game.players); i++ {
		player := &game.players[i]
		if len(player.hand) > 0 {
			return errors.New("one or more players already has cards")
		}

		hand, err := dealHand(13, &deck)

		if err != nil {
			return err
		}

		player.hand = hand
	}

	return nil
}

// Discard deletes a card from the hand of the player p
func (game *State) Discard(p uint8, card Card) error {
	pl := &game.players[p]
	i, err := findCard(pl.hand, card)

	if err != nil {
		return errors.New("engine.Game.Discard: card not found")
	}

	if i != -1 {
		pl.discard(i, card)
	}

	return err
}

// DiscardPlayed played deletes all cards from the game that have been played,
// and returns a slice of the cards deleted in this way
func (game *State) DiscardPlayed() (stack []Card) {
	for p := range game.players {
		pl := &game.players[p]
		for i, card := range pl.hand {
			if card.played == true {
				pl.discard(i, card)
				stack = append(stack, card)
			}
		}
	}

	return
}

// PRIVATE HELPER FUNCTIONS

// discard (lowercase) is a helper method to call on a player to remove a
// card at index i in their hand
func (player *Player) discard(i int, card Card) {
	player.hand = append(player.hand[:i], player.hand[i+1:]...)
	player.cardCount--
}

// findCard searches a given hand for a card. Returns the card index if found
// and an error if not.
func findCard(stack []Card, target Card) (int, error) {
	for i, card := range stack {
		if target.suit == card.suit &&
			target.value == card.value {
			return i, nil
		}
	}

	return -1, errors.New("could not find target card")
}
