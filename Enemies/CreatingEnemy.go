package Enemy

import (
	. "GameOfTrones/Utils"
	"fmt"
)

type IEnemyType interface {
	CreateEnemy()
	chooseEnemyType()
	createGoblin()
	createRat()
	createWitch()
	createOrc()
	setParameters(HP int, LVL int, damagePoints int)
}

func CreateEnemy(enemies *[]Enemy) {
	for {
		var e Enemy
		var key int
		PrintWithDelay("\nСоздание врагов:\n1:Гоблин\n2:Крыса\n3:Ведьма\n4:Орк\n0:Начать игру\n")
		fmt.Scan(&key)
		fmt.Println()
		startGame := chooseEnemyType(&e, key)
		if !startGame {
			break
		}
		addEnemy(enemies, &e)
	}
}

func chooseEnemyType(e *Enemy, enemy int) bool {
	switch enemy {
	case 1:
		createGoblin(e)
	case 2:
		createRat(e)
	case 3:
		createWitch(e)
	case 4:
		createOrc(e)
	case 0:
		return false
	}
	return true
}

func setParameters(e *Enemy, HP int, LVL int, damagePoints int, exp int) {
	e.setPositionEnemy()
	e.Possibility_of_enemy.SetLVL(LVL)
	e.Possibility_of_enemy.SetDamage(damagePoints)
	e.Possibility_of_enemy.SetHP(HP)
	e.Possibility_of_enemy.SetExperience(exp)
}

func createGoblin(e *Enemy) {
	e.Possibility_of_enemy.SetRace("Гоблин")
	setParameters(e, 5, 1, 1, 5)
}

func createRat(e *Enemy) {
	e.Possibility_of_enemy.SetRace("Крыса")
	setParameters(e, 2, 1, 1, 10)
}

func createWitch(e *Enemy) {
	e.Possibility_of_enemy.SetRace("Ведьма")
	setParameters(e, 7, 2, 2, 10)
}

func createOrc(e *Enemy) {
	e.Possibility_of_enemy.SetRace("Орк")
	setParameters(e, 12, 3, 3, 15)
}

func addEnemy(enemies *[]Enemy, e *Enemy) {
	*enemies = append(*enemies, *e)
}
