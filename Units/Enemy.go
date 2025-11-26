package units

import "fmt"

type IEnemy interface {
	setPosition(x int, y int)
	createEnemy()
	AddEnemy()
}

type Enemy struct {
	Possibility_of_enemy Unit
}

func CreateEnemy() Enemy {
	var Goblin = Enemy{Possibility_of_enemy: Unit{HP: 5, DamagePoints: 1, Race: "Гоблин"}}
	var x, y int
	fmt.Print("Введите координаты врага\nX:")
	fmt.Scan(&x)
	fmt.Print("Y:")
	fmt.Scan(&y)
	fmt.Println()
	Goblin.setPosition(x, y)
	Goblin.Possibility_of_enemy.SetLVL(1)
	return Goblin
}

func (e *Enemy) setPosition(x int, y int) {
	e.Possibility_of_enemy.Position.X = x
	e.Possibility_of_enemy.Position.Y = y
}

func AddEnemy(enemies *[]Enemy, e *Enemy) {
	*enemies = append(*enemies, *e)
}
