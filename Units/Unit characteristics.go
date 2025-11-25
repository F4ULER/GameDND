package units

import (
	"fmt"
)

type Person interface {
	Damage() int
	move(deltaX, deltaY int)
	setHP(HP int)
	checkUnitHP() bool
}

type ExperienceHandler interface {
	GetExperience()
	SetExperience()
	LVLUp()
}

type Enemies interface {
	getEnemyLVL()
}

type Coordinates struct {
	X int
	Y int
}

type Unit struct {
	HP           int
	DamagePoints int
	Position     Coordinates
	Race         string
	level        int
}
type Player struct {
	Name                  string
	Possibility_of_player Unit
	experience            int
}

type Enemy struct {
	Possibility_of_enemy Unit
}

func (u *Unit) Damage() int {
	return u.DamagePoints
}

func (u *Unit) Direction(step int) {
	switch step {
	case 1:
		u.move(1, 0)
	case 2:
		u.move(0, -1)
	case 3:
		u.move(0, 1)
	case 4:
		u.move(-1, 0)
	}

}
func (u *Unit) move(deltaX, deltaY int) {
	u.Position.X += deltaX
	u.Position.Y += deltaY
}

func (u *Unit) setHP(HP int) {
	u.HP += HP
}

func (u *Unit) checkUnitHP() bool {
	if u.HP <= 0 {
		fmt.Printf("%s УМЕР!\n", u.Race)
		return true
	} else {
		return false
	}
}

func (p *Player) GetExperience() int {
	return p.experience
}

func (p *Player) SetExperience(exp int) {
	p.experience += exp
}

func (p *Player) LVLUp() {
	var LVL = map[int]int{
		1: 0,
		2: 30,
		3: 80,
	}
	for level, exp := range LVL {
		if p.experience >= exp {
			p.Possibility_of_player.SetLVL(level)
		} else {
			return
		}
	}
}

func (u *Unit) GetLVL() int {
	return u.level
}

func (u *Unit) SetLVL(level int) {
	u.level = level
}
