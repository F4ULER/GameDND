package main

import (
	. "GameOfTrones/Enemies"
	. "GameOfTrones/Interaction"
	. "GameOfTrones/Players"
)

func main() {

	var Player Player
	Player.Possibility_of_player.SetLVL(1)
	Player.SetExperience(0)
	Player.Possibility_of_player.SetHP(5)
	Player.SetName("Tom")
	Player.Possibility_of_player.SetDamage(1)
	Player.Possibility_of_player.SetRace("Человек")
	Enemy1 := CreateEnemy()
	Enemy2 := CreateEnemy()
	Enemy3 := CreateEnemy()
	enemies := make([]Enemy, 0)
	AddEnemy(&enemies, &Enemy1)
	AddEnemy(&enemies, &Enemy2)
	AddEnemy(&enemies, &Enemy3)
	StartGame(&Player, &enemies)
}
