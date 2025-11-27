package Menu

import (
	. "GameOfTrones/Enemies"
	. "GameOfTrones/Players"
	. "GameOfTrones/Utils"
	"fmt"
)

type IShowMenu interface {
	Show_atack_menu()
	Show_move_menu()
}

func ShowEnemyInformation(p *Player, e *Enemy) {
	PrintWithDelay("\nВы встретили " + e.Possibility_of_enemy.GetRace())
	Separate_block()
	fmt.Printf("%s - %d уровень\n%s - %d уровень\n\n", p.GetName(), p.Possibility_of_player.GetLVL(), e.Possibility_of_enemy.GetRace(), e.Possibility_of_enemy.GetLVL())
}

func Show_atack_menu(p *Player, e *Enemy) int {
	var choice int = 1
	PrintWithDelay("Атака!\n1:Атаковать\n2:Сбежать\n")
	fmt.Scan(&choice)
	Separate_block()
	return choice
}

func Show_move_menu() int {
	step := 1
	PrintWithDelay("Движение:\n1.Вперед\n2.Вправо\n3.Влево\n4.Назад\n0:Выход")
	fmt.Scan(&step)
	return step
}
