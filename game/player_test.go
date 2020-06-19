package game

import "testing"

func TestCreatePlayers(t *testing.T) {
	n := 3
	players := createPlayers(&n)
	if len(players) != n {
		t.Errorf("Error while creating players. Expected %d, got %d", n, len(players))
	}
}

func TestRemoveBrokePlayers(t *testing.T) {
	totalChips["p1"] = 100
	totalChips["p2"] = 0
	totalChips["p3"] = 1
	players := []player{{name: "p1"}, {name: "p2"}, {name: "p3"}}
	removeBrokePlayers(&players)
	if len(players) != 2 {
		t.Error("Broke player was not removed from the list.")
	}
	for _, pl := range players {
		if pl.name == "p2" {
			t.Error("Player p1 is still in the list. He is broke tho...")
		}
	}
}
