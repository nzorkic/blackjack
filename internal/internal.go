package internal

import (
	"sort"

	"github.com/nzorkic/deck"
)

// SortedCardsCopy move card of rank Ace to last position
func SortedCardsCopy(c *[]deck.Card) []deck.Card {
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
