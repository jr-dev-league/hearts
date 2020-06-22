package main

import (
	"fmt"

	"github.com/nolwn/hearts/engine"
)

func newGame() {
	// create a game, returning the game with everything but the id
	// gamedata

	// TODO needs to call deal on new game

	newGame := engine.New()
	fmt.Println(newGame)
}
