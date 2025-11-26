package units

import (
	. "GameOfTrones/Utils"
	"fmt"
)

type IShowMenu interface {
	show_atack_menu()
	Show_move_menu()
}

func show_atack_menu(p *Player, e *Enemy) bool {
	fmt.Printf("%s - %d уровень\n%s - %d уровень\n\n", p.Name, p.Possibility_of_player.GetLVL(), e.Possibility_of_enemy.Race, e.Possibility_of_enemy.GetLVL())
	var choice int = 1
	for {
		PrintWithDelay("Атака!\n1:Атаковать\n2:Сбежать\n")
		fmt.Scan(&choice)
		Separate_block()
		switch choice {
		case 1:
			taking_damage(p, e)
			if areUnitDead(p, e) {
				p.LVLUp()
				fmt.Printf("%s - %d уровень!\n", p.Name, p.Possibility_of_player.GetLVL())
				return true
			}
		case 2:
			return true
		}
	}
}

func Show_move_menu(Player *Player, enemies *[]Enemy) {
	step := 1
	for {

		PrintWithDelay("Движение:\n1.Вперед\n2.Вправо\n3.Влево\n4.Назад\n0:Выход")
		fmt.Scan(&step)
		if step == 0 {
			break
		}
		if step == 5 {
			fmt.Print(Player.Possibility_of_player.GetPosition())
		}
		Player.Possibility_of_player.DirectionMove(step)
		for _, enemy := range *enemies {
			if areUnitsOnSamePosition(Player, &enemy) {
				Batle(Player, &enemy)
			}
		}
		Separate_block()

	}
}
