package game

import (
	"fmt"
	"sort"
	"strings"

	"github.com/nzorkic/deck"
)

func calculatePoints(cards *[]deck.Card) int {
	res := 0
	cardArr := sortedCardsCopy(cards)
	for _, card := range cardArr {
		if card.Visible {
			if card.Rank == deck.Ace {
				if res+11 <= 21 {
					card.Point = 11
				}
			}
			res += card.Point
		}
	}
	return res
}

func cardsAsString(cards *[]deck.Card) string {
	strSlice := []string{}
	for _, card := range *cards {
		strSlice = append(strSlice, card.String())
	}
	return strings.Join(strSlice[:], ", ")
}

func sortedCardsCopy(c *[]deck.Card) []deck.Card {
	cardArr := make([]deck.Card, len(*c))
	copy(cardArr, *c)
	sort.Slice(cardArr, less(&cardArr))
	return cardArr
}

func less(cards *[]deck.Card) func(i, j int) bool {
	return func(i, j int) bool {
		return (*cards)[i].Rank > (*cards)[j].Rank
	}
}

func blackjackHand(cards *[]deck.Card) bool {
	if len(*cards) != 2 {
		return false
	}
	if calculatePoints(cards) == MaxScore {
		return true
	}
	return false
}

func resetHand(players *[]player) {
	for idx := range *players {
		(*players)[idx].cards = nil
	}
	dealer.cards = nil
}

func createNewDeckIfNeeded(playerSize int, d *deck.Deck) {
	if ((playerSize * 4) + 4) > len(*d) {
		fmt.Println("Shuffling new deck...")
		fmt.Println()
		*d = createDeck(&playerSize)
	}
}

func setupPoints(d *deck.Deck) {
	(*d).FacePoints(10)
}
