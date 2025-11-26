package units

import (
	"fmt"
)

type IPerson interface {
	Damage() int
	move(deltaX, deltaY int)
	setHP(HP int)
	checkUnitHP() bool
	GetPosition()
}

type IExperienceHandler interface {
	GetExperience()
	SetExperience(exp int)
	LVLUp()
	GetLVL() int
	SetLVL(level int)
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

func (u *Unit) Damage() int {
	return u.DamagePoints
}

func (u *Unit) DirectionMove(step int) {
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

func (u *Unit) GetLVL() int {
	return u.level
}

func (u *Unit) SetLVL(level int) {
	u.level = level
}

func (u *Unit) GetPosition() Coordinates {
	return u.Position
}
