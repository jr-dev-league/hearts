package main

import (
	"github.com/jr-dev-league/hearts/database"
	"github.com/jr-dev-league/hearts/engine"
)

func cvtPoints(g database.GameRecord) (points uint8, cards uint8) {
	points = 26
	cards = 14
	players := g.Players

	for _, palyer := range players {
		for _, card := range palyer.Hand {
			if card.Suit == engine.Hearts {
				cards--
				points--
			} else if card.Suit == engine.Spades && card.Value == 11 {
				cards--
				points -= 13
			}
		}
	}

	return
}

func ctvBroken(points uint8, cards uint8) bool {
	var max uint8

	// Is the queen taken?
	if points/(cards-1) != 2 { // yes
		max = 13

	} else { //no
		max = 26
	}

	return points < max
}

func cvtShootable(players [4]database.Player) bool {
	withPoints := 0

	for _, player := range players {
		if player.Round > 0 {
			withPoints++
		}
	}

	return withPoints > 1
}

func cvtTakenLast(g database.GameRecord) (pIdx uint8) {
	players := g.Players
	played := false
	turn := g.Turn

	// If we count backwards, then the person who took last is the person who played a card
	// after someone who didn't play a card. Otherwise it's the person at the top of the list.
	// Otherwise, no on has played a card yet, and it's whoever's turn it is.
	//
	// For example: [played, didn't, played, played]
	//									^
	// this player went |~~~~~~~~~~~~~~~
	// 			  first |
	for i := uint8(len(players)) - 1; i >= 0; i-- { // count backward
		if player := players[i]; len(player.Active) > 0 { // Have they played a card?
			pIdx = i      // move the marker
			played = true // flag that at least someone played
		} else if played { // they didn't play, but someone else did
			break // so the last person we saw must have been them.
		}
	}

	if !played { // if we never found anyone who played...
		pIdx = turn // then the person whose turn it is will go first
	}

	return
}

func cvtPlayers(g database.GameRecord) (p [4]engine.Player) {
	playing := g.Phase == database.PhasePlay
	pushCard := func(p engine.Player, card database.Card, active bool) {
		eCard := engine.Card{
			Suit:    card.Suit,
			Value:   card.Value,
			Exposed: playing && active,
			Played:  active,
		}

		p.Hand = append(p.Hand, eCard)
	}

	for i, player := range g.Players {
		for _, card := range player.Hand {
			pushCard(p[i], card, false)
		}

		for _, card := range player.Active {
			pushCard(p[i], card, true)
		}
	}

	return
}

func toState(g database.GameRecord) (s engine.State) {
	var broken bool
	var takenLast uint8
	var players [4]engine.Player
	var canShoot bool
	var readonly bool

	// figure out broken
	points, cards := cvtPoints(g)
	broken = ctvBroken(points, cards)

	// figure out takenLast
	if g.Phase == database.PhasePlay {
		takenLast = cvtTakenLast(g)
	}

	// figure out players
	players = cvtPlayers(g)

	// figure out shootable
	canShoot = cvtShootable(g.Players)

	// figure out readonly
	readonly = false

	s = engine.State{
		Broken:    broken,
		TakenLast: takenLast,
		Players:   players,
		Shootable: canShoot,
		Readonly:  readonly,
	}

	return
}
