package game

var offeredChips = make(map[string]int)

func bet(name *string) int {
	return offeredChips[*name]
}

func makeBet(name *string, n int) {
	totalChips[*name] -= n
	offeredChips[*name] += n
}

func resetBets() {
	for k := range offeredChips {
		delete(offeredChips, k)
	}
}

func removeBet(name *string) {
	delete(offeredChips, *name)
}
