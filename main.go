package main

import (
	. "GameOfTrones/Units"
	//. "GameOfTrones/Utils"
)

func main() {

	var Player = Player{Name: "Tom", Possibility_of_player: Unit{HP: 6, DamagePoints: 1, Position: Coordinates{X: 0, Y: 0}, Race: "human"}}
	Player.Possibility_of_player.SetLVL(1)
	Player.SetExperience(0)
	Goblin1 := CreateEnemy()
	Goblin2 := CreateEnemy()
	Goblin3 := CreateEnemy()
	enemies := make([]Enemy, 0)
	AddEnemy(&enemies, &Goblin1)
	AddEnemy(&enemies, &Goblin2)
	AddEnemy(&enemies, &Goblin3)
	Show_move_menu(&Player, &enemies)
}
