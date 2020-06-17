package main

import (
	"fmt"
	"strings"

	util "github.com/nzorkic/blackjack/internal"
	"github.com/nzorkic/deck/deck"
)

// MaxScore is highest score player can get in Blackjack
const MaxScore = 21

type player struct {
	name   string
	cards  []deck.Card
	points int
}

var results = make(map[string]int)

type action uint8

const (
	hit action = iota
	stand
)

func (p player) String() string {
	return fmt.Sprintf("%s (%d): %v", p.name, calculatePoints(&p.cards), cardsAsString(&p.cards))
}

func cardsAsString(cards *[]deck.Card) string {
	strSlice := []string{}
	for _, card := range *cards {
		strSlice = append(strSlice, card.String())
	}
	return strings.Join(strSlice[:], ", ")
}

func main() {
	numOfPlayers := 1 + 1
	players := createPlayers(numOfPlayers)
	deck := deck.New(deck.Shuffle())
	initPoints(deck)
	deal(&players, &deck)
	start(&players, &deck)
}

func initPoints(deck deck.Deck) {
	deck.FacePoints(10)
}

func calculatePoints(cards *[]deck.Card) int {
	res := 0
	cardArr := util.SortedCardsCopy(cards)
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

func createPlayers(n int) []player {
	players := make([]player, n)
	for i := 0; i < n; i++ {
		players[i] = player{name: fmt.Sprintf("Player #%d", i+1)}
		if isLast(i, n) {
			players[n-1].name = "Dealer"
		}
	}
	return players
}

func deal(p *[]player, deck *deck.Deck) {
	cardIdx := 0
	playerSize := len(*p)
	turns := 2
	for i := 0; i < turns; i++ {
		for j := 0; j < playerSize; j++ {
			cards := (*deck).Draw(1)
			if isLast(i, turns) {
				if isLast(j, playerSize) {
					cards[0].Visible = false
				}
			}
			(*p)[j].cards = append((*p)[j].cards, cards...)
			cardIdx++
		}
	}
	for _, player := range *p {
		results[player.name] += calculatePoints(&player.cards)
	}
	printScores(p)
}

func start(p *[]player, d *deck.Deck) {
	for _, player := range *p {
		if player.name == "Dealer" {
			player.cards[1].Visible = true
			results[player.name] = calculatePoints(&player.cards)
		}
		var playerAction action
		printActionChoices(&player)
		askUserInput(&playerAction)
		switch playerAction {
		case hit:
			hitAction(&player, playerAction, d)
		case stand:
			break
		}
	}
	printFinalScore()
}

func hitAction(player *player, playerAction action, d *deck.Deck) {
	player.cards = append(player.cards, (*d).Draw(1)...)
	for playerAction == hit {
		printActionChoices(player)
		points := calculatePoints(&player.cards)
		results[player.name] = points
		if points > MaxScore {
			fmt.Println("It's a bust!")
			fmt.Println()
			break
		}
		askUserInput(&playerAction)
		if playerAction == hit {
			player.cards = append(player.cards, (*d).Draw(1)...)
		}
	}
}

func printScores(players *[]player) {
	for _, player := range *players {
		fmt.Println(player)
	}
	fmt.Println()
}

func printActionChoices(player *player) {
	fmt.Println(player)
	fmt.Println("Enter 0 to Hit or 1 to Stand")
	fmt.Print("> ")
}

func printFinalScore() {
	winner := "Dealer"
	winnerPts := 0
	runnerUpPts := 0
	for name, score := range results {
		fmt.Printf("%s: %d\n", name, score)
		if score > MaxScore {
			continue
		}
		if score >= winnerPts {
			winner = name
			runnerUpPts = winnerPts
			winnerPts = score
		}
	}
	fmt.Println()
	if runnerUpPts == winnerPts {
		fmt.Println("It's a push!")
	} else {
		fmt.Println("The winner is ", winner)
	}
}

func askUserInput(playerAction *action) {
	fmt.Scanln(playerAction)
	fmt.Println()
}

func isLast(i, j int) bool {
	if i+1 == j {
		return true
	}
	return false
}
