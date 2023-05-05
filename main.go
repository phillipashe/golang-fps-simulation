package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	todo:
	* implement a mutex
	* read in user values
	* random damage
	* random shooting intervals
*/

var wg sync.WaitGroup

type Player struct {
	Name   string
	Health int
}

func main() {

	startTimer := time.Now()
	// defer runs after the function exits
	// this uses anon func but can pass a named func too
	defer func() {
		fmt.Printf("Match finished in %s\n", time.Since(startTimer))
	}()

	playerList := loadPlayers()
	startBattle(playerList)
	wg.Wait()
}

// currently hardcoded
func loadPlayers() *[]Player {
	var players []Player
	names := [4]string{"Thomas", "Caleb", "Shelby", "Phillip"}
	for _, name := range names {
		players = append(players, Player{Name: name, Health: 100})
	}
	return &players
}

func AttackPlayers(player *Player, playerList *[]Player) {

	for player.Health > 0 {
		fmt.Printf("Shooting a laser at %s\n", player.Name)
		player.Health = player.Health - 50
	}
	wg.Done()
}

func startBattle(playerList *[]Player) {
	for _, player := range *playerList {
		wg.Add(1)
		AttackPlayers(&player, playerList)
	}
}
