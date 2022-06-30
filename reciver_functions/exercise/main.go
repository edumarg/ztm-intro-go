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
	Name              string
	Health, maxHealth int16
	Energy, maxEnergy int16
}

func (p *Player) changeHealth(health int16) {
	p.Health += health
	if p.Health < 0 {
		p.Health = 0
	}
	if p.Health > p.maxHealth {
		p.Health = p.maxHealth
	}
}

func (p *Player) changeEnergy(energy int16) {
	p.Energy += energy
	if p.Energy < 0 {
		p.Energy = 0
	}
	if p.Energy > p.maxEnergy {
		p.Energy = p.maxEnergy
	}
}

func main() {
	player1 := Player{Name: "Moolok", Energy: 0, Health: 0, maxHealth: 1000, maxEnergy: 1000}
	player2 := Player{Name: "IronWidow", Energy: 0, Health: 0, maxHealth: 1000, maxEnergy: 1000}
	player3 := Player{Name: "DarthLuck", Energy: 0, Health: 0, maxHealth: 1000, maxEnergy: 1000}

	player1.changeEnergy(30)
	player1.changeHealth(500)
	fmt.Println(player1)
	player1.changeHealth(-80)
	fmt.Println(player1)

	player2.changeEnergy(100)
	player2.changeHealth(1500)
	fmt.Println(player2)
	player1.changeEnergy(-90)
	player2.changeHealth(-400)
	fmt.Println(player2)

	player3.changeEnergy(1000)
	player3.changeHealth(500)
	fmt.Println(player3)
	player1.changeEnergy(10)
	player3.changeHealth(-600)
	fmt.Println(player3)
}
