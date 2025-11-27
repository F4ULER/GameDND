package Enemy

import (
	//. "GameOfTrones/Units"
	"fmt"
)

type IEnemyType interface {
	chooseEnemyType()
	createGoblin()
	createRat()
	createWitch()
	createOrc()
	setParameters(HP int, LVL int, damagePoints int)
}

func chooseEnemyType(e *Enemy) {
	var enemy int
	fmt.Print("\nВыбор врага:\n1:Гоблин\n2:Крыса\n3:Ведьма\n4:Орк\n")
	fmt.Scanln(&enemy)
	fmt.Println()
	switch enemy {
	case 1:
		createGoblin(e)
	case 2:
		createRat(e)
	case 3:
		createWitch(e)
	case 4:
		createOrc(e)
	}
}

func setParameters(e *Enemy, HP int, LVL int, damagePoints int) {
	e.setPositionEnemy()
	e.Possibility_of_enemy.SetLVL(LVL)
	e.Possibility_of_enemy.SetDamage(damagePoints)
	e.Possibility_of_enemy.SetHP(HP)
}

func createGoblin(e *Enemy) {
	e.Possibility_of_enemy.SetRace("Гоблин")
	setParameters(e, 5, 1, 1)
}

func createRat(e *Enemy) {
	e.Possibility_of_enemy.SetRace("Крыса")
	setParameters(e, 2, 1, 1)
}

func createWitch(e *Enemy) {
	e.Possibility_of_enemy.SetRace("Ведьма")
	setParameters(e, 7, 2, 2)
}

func createOrc(e *Enemy) {
	e.Possibility_of_enemy.SetRace("Орк")
	setParameters(e, 12, 3, 3)
}
