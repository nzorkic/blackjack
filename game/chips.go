package game

import "fmt"

var totalChips = make(map[string]int)

func chips(name *string) int {
	return totalChips[*name]
}

func addChips(name *string, n int) {
	totalChips[*name] += n
}

func resetChips() {
	for k := range totalChips {
		delete(totalChips, k)
	}
}

func removeChip(name *string) {
	delete(totalChips, *name)
}

func offerChips(players *[]player) {
	for _, player := range *players {
		var offer int
		playerChips := chips(&player.name)
		fmt.Printf("%s place a bet (max. %d$) > ", player.name, playerChips)
		fmt.Scan(&offer)
		for offer > playerChips {
			fmt.Printf("You don't have that much chips, %s, try again (max. %d$) > ", player.name, playerChips)
			fmt.Scan(&offer)
		}
		fmt.Println()
		makeBet(&player.name, offer)
	}
}
