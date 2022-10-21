//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func newPlayer() Player {
	return Player{
		name:      "test",
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
	}
}

func TestHealth(t *testing.T) {
	player := newPlayer()
	updateHealth(&player, 50)
	if player.health != 50 {
		t.Fatalf("health %v", player.health)
	}
}

func TestEnergy(t *testing.T) {
	player := newPlayer()
	updateEnergy(&player, 50)
	if player.energy != 50 {
		t.Fatalf("energy %v", player.energy)
	}
}
