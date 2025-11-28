package main

import (
	. "GameOfTrones/Enemies"
	. "GameOfTrones/Interaction"
	. "GameOfTrones/Players"
)

func main() {
	enemies := make([]Enemy, 0)
	var Player Player
	Player.Possibility_of_player.SetLVL(1)
	Player.SetExperience(0)
	Player.Possibility_of_player.SetHP(10)
	Player.SetName("Tom")
	Player.Possibility_of_player.SetDamage(1)
	Player.Possibility_of_player.SetRace("Человек")
	CreateEnemy(&enemies)

	StartGame(&Player, &enemies)
}
