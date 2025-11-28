package units

import (
	"fmt"
)

type IPerson interface {
	GetDamage() int
	SetDamage()
	move(deltaX, deltaY int)
	CheckUnitHP() bool
	GetPosition()
	GetHP()
	SetHP()
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
	hitPoints    int
	damagePoints int
	position     Coordinates
	race         string
	level        int
	experience   int
}

func (u *Unit) GetDamage() int {
	return u.damagePoints
}

func (u *Unit) SetDamage(damage int) {
	u.damagePoints = damage
}

func (u *Unit) GetRace() string {
	return u.race
}

func (u *Unit) SetRace(race string) {
	u.race = race
}

func (u *Unit) GetExperience() int {
	return u.experience
}

func (u *Unit) SetExperience(exp int) {
	u.experience += exp
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
	u.position.X += deltaX
	u.position.Y += deltaY
}

func (u *Unit) SetHP(HP int) {
	u.hitPoints += HP
}

func (u *Unit) GetHP() int {
	return u.hitPoints
}

func (u *Unit) CheckUnitHP() bool {
	if u.hitPoints <= 0 {
		fmt.Printf("%s УМЕР!\n", u.GetRace())
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
	return u.position
}

func (u *Unit) SetPosition(x int, y int) {
	u.position.X = x
	u.position.Y = y
}
