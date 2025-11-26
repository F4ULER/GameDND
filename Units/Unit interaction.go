package units

import (
	. "GameOfTrones/Utils"
	"fmt"
	"os"
)

type IUnitInteraction interface {
	IPerson
	IAttack
	IShowMenu
	IExperienceHandler
}

type IAttack interface {
	areUnitDead(p *Player, e *Enemy)
	taking_damage(p *Player, e *Enemy)
	areUnitsOnSamePosition(p *Player, e *Enemy) bool
}

func Batle(p *Player, e *Enemy) {
	Separate_block()
	PrintWithDelay("\nВы встретили " + e.Possibility_of_enemy.Race)
	if show_atack_menu(p, e) {
		return
	}
}

func areUnitDead(p *Player, e *Enemy) bool {
	if p.Possibility_of_player.checkUnitHP() {
		fmt.Println("YOU ARE DEAD!")
		os.Exit(0)
	} else if e.Possibility_of_enemy.checkUnitHP() {
		fmt.Println("Вы выйграли!")
		p.SetExperience(5)
		Separate_block()
		return true
	}
	return false
}

func taking_damage(p *Player, e *Enemy) {
	damage := p.Possibility_of_player.Damage()
	e.Possibility_of_enemy.setHP(-damage)
	damage = e.Possibility_of_enemy.Damage()
	p.Possibility_of_player.setHP(-damage)
	fmt.Printf("Вы получили %d урона (%d ОЗ)\n%s получил %d урона (%d ОЗ)\n", e.Possibility_of_enemy.DamagePoints, p.Possibility_of_player.HP, e.Possibility_of_enemy.Race, p.Possibility_of_player.DamagePoints, e.Possibility_of_enemy.HP)
	Separate_block()
}

func areUnitsOnSamePosition(p *Player, e *Enemy) bool {
	return p.Possibility_of_player.Position == e.Possibility_of_enemy.Position
}

func AreUnitsStartInteracting(p *Player) {

}
