package internal

import (
	"sort"
	"strings"

	"github.com/nzorkic/deck"
)

// CalculatePoints calculate points from given cards
func CalculatePoints(cards *[]deck.Card) int {
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

// CardsAsString converts cards to comma separated string
func CardsAsString(cards *[]deck.Card) string {
	strSlice := []string{}
	for _, card := range *cards {
		strSlice = append(strSlice, card.String())
	}
	return strings.Join(strSlice[:], ", ")
}
