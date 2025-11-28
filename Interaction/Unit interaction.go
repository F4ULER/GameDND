package interaction

import (
	. "GameOfTrones/Enemies"
	. "GameOfTrones/Menu"
	. "GameOfTrones/Players"
	. "GameOfTrones/Units"
	. "GameOfTrones/Utils"
	"fmt"
	"os"
)

type IUnitInteraction interface {
	IPerson
	IAttack
	IExperienceHandler
}

type IAttack interface {
	areUnitDead(p *Player, e *Enemy)
	taking_damage(p *Player, e *Enemy)
	areUnitsOnSamePosition(p *Player, e *Enemy) bool
}

func StartGame(Player *Player, enemies *[]Enemy) {
	for {
		step := Show_move_menu()
		if step == 0 {
			break
		}
		Player.Possibility_of_player.DirectionMove(step)
		for _, enemy := range *enemies {
			if areUnitsOnSamePosition(Player, &enemy) {
				Batle(Player, &enemy)
				break
			}
		}
	}
}

func Batle(p *Player, e *Enemy) {
	ShowEnemyInformation(p, e)
	for {
		switch Show_atack_menu(p, e) {
		case 1:
			taking_damage(p, e)
			if areUnitDead(p, e) {
				p.LVLUp()
				fmt.Printf("%s - %d уровень!\n", p.GetName(), p.Possibility_of_player.GetLVL())
				return
			}
		case 2:
			return
		}
	}

}

func areUnitDead(p *Player, e *Enemy) bool {
	if p.Possibility_of_player.CheckUnitHP() {
		fmt.Println("YOU ARE DEAD!")
		os.Exit(0)
	} else if e.Possibility_of_enemy.CheckUnitHP() {
		fmt.Println("Вы выйграли!")
		p.SetExperience(5)
		Separate_block()
		return true
	}
	return false
}

func taking_damage(p *Player, e *Enemy) {
	damage := p.Possibility_of_player.GetDamage()
	e.Possibility_of_enemy.SetHP(-damage)
	damage = e.Possibility_of_enemy.GetDamage()
	p.Possibility_of_player.SetHP(-damage)
	fmt.Printf("Вы получили %d урона (%d ОЗ)\n%s получил %d урона (%d ОЗ)\n", e.Possibility_of_enemy.GetDamage(), p.Possibility_of_player.GetHP(), e.Possibility_of_enemy.GetRace(), p.Possibility_of_player.GetDamage(), e.Possibility_of_enemy.GetHP())
	Separate_block()
}

func areUnitsOnSamePosition(p *Player, e *Enemy) bool {
	return p.Possibility_of_player.GetPosition() == e.Possibility_of_enemy.GetPosition()
}
