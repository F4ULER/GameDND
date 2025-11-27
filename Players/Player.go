package players

import . "GameOfTrones/Units"

type Player struct {
	name                  string
	Possibility_of_player Unit
	experience            int
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) SetName(name string) {
	p.name = name
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
