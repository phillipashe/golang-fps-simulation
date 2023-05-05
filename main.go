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

	players := loadPlayers()
	getNames(players)
	wg.Wait()
}

// currently hardcoded
func loadPlayers() []Player {
	var players []Player
	names := [4]string{"Thomas", "Caleb", "Shelby", "Phillip"}
	for _, name := range names {
		players = append(players, Player{Name: name, Health: 100})
	}
	return players
}

func shootLaser(name string) {
	fmt.Printf("Shooting a laser at %s\n", name)
	wg.Done()
}

func getNames(players []Player) {
	for _, player := range players {
		wg.Add(1)
		go shootLaser(player.Name)
	}
}
