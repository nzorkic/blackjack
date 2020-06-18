package chips

var totalChips = make(map[string]int)

var offeredChips = make(map[string]int)

// SeeChips returns chip worth of player with provided name
func SeeChips(name *string) int {
	return totalChips[*name]
}

// AddChips adds chips for player with given name
func AddChips(name *string, n int) {
	totalChips[*name] += n
}

// RemoveChips removes chips for player with given name
func RemoveChips(name *string, n int) {
	totalChips[*name] -= n
	if totalChips[*name] <= 0 {
		delete(totalChips, *name)
	}
}

// ResetChips epties the total chips of all players
func ResetChips() {
	for k := range totalChips {
		delete(totalChips, k)
	}
}

// SeeBet returns chip worth of player with provided name
func SeeBet(name *string) int {
	return offeredChips[*name]
}

// MakeBet is amount of money placed by a player
func MakeBet(name *string, n int) {
	totalChips[*name] -= n
	offeredChips[*name] += n
}

// ResetBets epties the total chips of all players
func ResetBets() {
	for k := range offeredChips {
		delete(offeredChips, k)
	}
}

// Remove removes player from chips and bets
func Remove(name *string) {
	delete(totalChips, *name)
	delete(offeredChips, *name)
}
