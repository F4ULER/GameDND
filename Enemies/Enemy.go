package Enemy

import (
	Units "GameOfTrones/Units"
	. "GameOfTrones/Utils"
	"fmt"
)

type IEnemy interface {
	setPositionEnemy()
	CreateEnemy()
	addEnemy()
	getExperienceForEnemyKill()
}

type Enemy struct {
	Possibility_of_enemy Units.Unit
}

func (e *Enemy) setPositionEnemy() {
	var x, y int
	PrintWithDelay("Введите координаты врага\nX:")
	fmt.Scan(&x)
	PrintWithDelay("Y:")
	fmt.Scan(&y)
	fmt.Println()

	e.Possibility_of_enemy.SetPosition(x, y)
}

func (e *Enemy) getExperienceForEnemyKill() {

}
