//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name      string
	health    int
	maxHealth int
	energy    int
	maxEnergy int
}

func updateHealth(player *Player, health int) {
	player.health = health
}

func updateEnergy(player *Player, energy int) {
	player.energy = energy
}

func main() {
	player := Player{
		name:      "Mateus",
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
	}
	fmt.Println(player)
	updateHealth(&player, 50)
	updateEnergy(&player, 50)
	fmt.Println(player)
}
