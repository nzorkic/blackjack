package game

import (
	"fmt"

	"github.com/nzorkic/deck"
)

type action uint8

const (
	stand action = iota
	hit
)

func invokeAction(player *player, d *deck.Deck) {
	var playerAction action
	printActionChoices(player)
	if blackjackHand(&player.cards) {
		results[player.name] = blackjack
		fmt.Println("Blackjack!")
		fmt.Println()
	} else {
		fmt.Scan(&playerAction)
		fmt.Println()
		switch playerAction {
		case hit:
			hitAction(player, d)
		case stand:
			break
		}
	}
}

func hitAction(player *player, d *deck.Deck) {
	player.cards = append(player.cards, (*d).Draw(1)...)
	var playerAction action = hit
	for playerAction == hit {
		printActionChoices(player)
		points := calculatePoints(&player.cards)
		player.points = points
		if points > MaxScore {
			results[player.name] = lost
			fmt.Println("It's a bust!")
			fmt.Println()
			break
		}
		fmt.Scan(&playerAction)
		fmt.Println()
		if playerAction == hit {
			player.cards = append(player.cards, (*d).Draw(1)...)
		}
	}
}
