package Enemy

import (
	Units "GameOfTrones/Units"
	"fmt"
)

type IEnemy interface {
	setPositionEnemy()
	createEnemy()
	AddEnemy()
	getExperienceForEnemyKill()
}

type Enemy struct {
	Possibility_of_enemy Units.Unit
}

func CreateEnemy() Enemy {
	var e Enemy
	chooseEnemyType(&e)
	return e
}

func (e *Enemy) setPositionEnemy() {
	var x, y int
	fmt.Print("Введите координаты врага\nX:")
	fmt.Scan(&x)
	fmt.Print("Y:")
	fmt.Scan(&y)
	fmt.Println()
	e.Possibility_of_enemy.SetPosition(x, y)
}

func AddEnemy(enemies *[]Enemy, e *Enemy) {
	*enemies = append(*enemies, *e)
}

func (e *Enemy) getExperienceForEnemyKill() {

}
