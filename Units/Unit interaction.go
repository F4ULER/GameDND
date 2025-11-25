package units

import (
	. "GameOfTrones/Utils"
	"fmt"
	"os"
)

type UnitInteraction interface {
	Person
	Attack
	ShowMenu
	ExperienceHandler
}

type Attack interface {
	areUnitDead(p *Player, e *Enemy)
	taking_damage(p *Player, e *Enemy)
	areUnitsOnSamePosition(p *Player, e *Enemy) bool
}

func Batle(p *Player, e *Enemy) {
	if areUnitsOnSamePosition(p, e) {
		PrintWithDelay("Вы встретили " + e.Possibility_of_enemy.Race)
		if show_atack_menu(p, e) {
			return
		}
	}
}

func show_atack_menu(p *Player, e *Enemy) bool {
	fmt.Printf("%s - %d уровень\n%s - %d уровень\n", p.Name, p.Possibility_of_player.GetLVL(), e.Possibility_of_enemy.Race, e.Possibility_of_enemy.GetLVL())
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
