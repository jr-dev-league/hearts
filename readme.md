# Hearts Game, v1

## Outline

- [ ] User can view a list of available seats from the homepage
- [ ] User can select a seat and join a game
- [ ] Game starts when all users are seated
- [ ] User can play a full game of hearts with other players
- [ ] The user's view of the game updates automatically

## Technical Overview

- [ ] Games are played in memory, no need for a db in v1
- [ ] Users cannot log in, but web app session should remember their seat
- [ ] Hearts games should be complete and contain all game rules
- [ ] Automatic game updates can be done via polling

## Backend

The game needs to be divided into resources in the REST style. It will be
written in Go. It will avoid external dependecies as much as is practical. It
will use the net/http package.

### Games

The game resource manages the state of the game. Requests to this resource
do things like:

- [x] Get the state of the game
- [ ] Get the state of the game for a particular player
- [ ] Play a card from a players hand

#### /api/games

**GET** lists all games

**POST** creates a new game

#### /api/games/{gameid}

**GET** gets the state of a given game by its id. Takes `seatno` as a query param, which
shows only data that should be viewable to players in that seat. This
functionality should be replaced by proper authentication/authorization later.

**PATCH** plays a given card.

#### /api/games/{gameid}/join

**PATCH** marks one open game seat as closed.

#### /api/games/{gameid}/leave

**PATCH** marks one closed game seat as open.

## Frontend

The frontend will be written in TypeScript React. It will use Semantic UI for
its UX. It will need to keep track of what the player is supposed to be doing—
it should cache the games a player is in, as well as their seat numbers. Later,
these features will be moved to the backend (where they belong), but v1 isn't
concerned with that kind of correctness.

### Homepage

The homepage needs to list games with open seats. It is possible that v1 will
only support a single game. In that case, it should still show a list of 0–1
games, depending on whether that game has open seats. Players should be able to
click to join the game.

### Lobby page

When a player joins a game that has not yet started, they should be taken to a
lobby page. That page should display the number of seats that are not filled.
In v1, this page does not have much utility, but it will serve to make it clear
when a game has started vs. when it is pending. In later versions, it will
serve a more important function. Players will be able to click a leave game
button and be returned to the homepage, marking a closed seat as open.

### Game page

Once the game is started it should not be joinable or leaveable in v1. Players
should be able to interact with the game state from the page by playing cards.
The game page should display a user's cards, the cards that have been put on
the table, all the players' scores (by their seat, since there is no user
login). Each edge of the screen should represent a seat, with the bottom edge
representing the user. Animations should clearly show who took a trick.
