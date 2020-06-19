package game

import (
	"fmt"

	"github.com/nzorkic/deck"
)

type player struct {
	name   string
	cards  []deck.Card
	points int
}

func (p player) String() string {
	return p.name
}

var dealer = player{name: "Mr. Dealer"}

func createPlayers(n *int) []player {
	players := make([]player, *n)
	for i := 0; i < *n; i++ {
		name := fmt.Sprintf("Player #%d", i+1)
		player := player{name: name}
		players[i] = player
		addChips(&player.name, 900)
	}
	return players
}

func removeBrokePlayers(players *[]player) {
	for i := 0; i < len(*players); i++ {
		playerName := (*players)[i].name
		if chips(&playerName) <= 0 {
			*players = append((*players)[:i], (*players)[i+1:]...)
			removeBet(&playerName)
			removeChip(&playerName)
			i--
		}
	}
}
