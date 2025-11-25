package main

import (
	. "GameOfTrones/Units"
	. "GameOfTrones/Utils"
	"fmt"
)

func main() {

	var Player = Player{Name: "Tom", Possibility_of_player: Unit{HP: 6, DamagePoints: 1, Position: Coordinates{X: 0, Y: 0}, Race: "human"}}
	var Goblin1 = Enemy{Possibility_of_enemy: Unit{HP: 2, DamagePoints: 1, Position: Coordinates{X: 1, Y: 1}, Race: "goblin"}}
	Goblin1.Possibility_of_enemy.SetLVL(1)
	Player.Possibility_of_player.SetLVL(1)
	Player.SetExperience(0)
	var Goblin2 = Enemy{Possibility_of_enemy: Unit{HP: 2, DamagePoints: 1, Position: Coordinates{X: 1, Y: 1}, Race: "goblin"}}
	Goblin2.Possibility_of_enemy.SetLVL(1)

	step := 1
	for {

		PrintWithDelay("Движение:\n1.Вперед\n2.Вправо\n3.Влево\n4.Назад\n0:Выход")
		fmt.Scan(&step)
		if step == 0 {
			break
		}
		Player.Possibility_of_player.Direction(step)
		Separate_block()
		Batle(&Player, &Goblin1)

	}
}
