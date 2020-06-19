package game

type result uint8

const (
	push result = iota
	won
	lost
	blackjack
)

var results = make(map[string]result)

func setResults(players *[]player) {
	for _, player := range *players {
		res, ok := results[player.name]
		if ok {
			if (res == blackjack || res == lost) &&
				results[dealer.name] == res {
				results[player.name] = push
			}
		} else {
			switch results[dealer.name] {
			case blackjack:
				results[player.name] = lost
			case lost:
				results[player.name] = won
			default:
				playerPoints := player.points
				dealerPoints := dealer.points
				switch {
				case playerPoints > dealerPoints:
					results[player.name] = won
				case playerPoints < dealerPoints:
					results[player.name] = lost
				case playerPoints == dealerPoints:
					results[player.name] = push
				}
			}
		}
	}
}

func finalizeResults() {
	for name, res := range results {
		switch res {
		case push:
			addChips(&name, bet(&name))
		case blackjack:
			addChips(&name, bet(&name)+
				int(float64(bet(&name))*BlackjackBonus))
		case won:
			addChips(&name, bet(&name)*2)
		}
	}
}

func resetResults() {
	for k := range results {
		delete(results, k)
	}
}
