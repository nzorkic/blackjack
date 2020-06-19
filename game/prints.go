package game

import "fmt"

func printGameIntro() {
	fmt.Println("#############")
	fmt.Println("# BLACKJACK #")
	fmt.Println("#############")
	fmt.Println()
}

func printScores(players *[]player) {
	for _, player := range *players {
		printPlayerInfoLine(&player)
	}
	printDealerInfoLine(&dealer)
	fmt.Println()
}

func printActionChoices(player *player) {
	printPlayerStats(player)
	fmt.Println("Enter 0 to Stand or 1 to Hit")
	fmt.Print("> ")
}

func printPlayerStats(player *player) {
	fmt.Println(player)
	fmt.Println("Hand: ", cardsAsString(&player.cards))
	fmt.Printf("Bank: $%d\n", chips(&player.name))
	fmt.Printf("Bet: $%d\n", bet(&player.name))
	fmt.Printf("Points: %d\n", calculatePoints(&player.cards))
}

func printPlayerInfoLine(player *player) {
	fmt.Printf("%s [Bet: $%d (Bank: $%d)] (%d pts): %v\n",
		player.name,
		bet(&player.name),
		chips(&player.name),
		calculatePoints(&player.cards),
		cardsAsString(&player.cards))
}

func printDealerInfoLine(dealer *player) {
	fmt.Printf("%s (%d): %v\n",
		dealer.name,
		calculatePoints(&dealer.cards),
		cardsAsString(&dealer.cards))
}

func printFinalScore() {
	delete(results, dealer.name)
	for name, score := range results {
		fmt.Printf("%s: ", name)
		switch score {
		case blackjack:
			fmt.Printf("Won with Blackjack! (+$%d)",
				int(float64(bet(&name))*BlackjackBonus))
		case won:
			fmt.Printf("Won! (+$%d)", bet(&name))
		case lost:
			fmt.Printf("Lost (-$%d)", bet(&name))
		case push:
			fmt.Print("Push")
		}
		fmt.Println()
	}
	fmt.Println()
}

func printEndgameOptions() {
	fmt.Println("Want to play another game?")
	fmt.Println("Enter 0 for another round")
	fmt.Println("Enter 1 for new game")
	fmt.Println("Enter 2 to quit")
	fmt.Print("> ")
}
