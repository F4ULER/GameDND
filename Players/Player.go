package players

import . "GameOfTrones/Units"

type Player struct {
	name                  string
	Possibility_of_player Unit
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) SetName(name string) {
	p.name = name
}

func (p *Player) LVLUp() {
	var LVL = map[int]int{
		1: 0,
		2: 30,
		3: 80,
	}
	for level, exp := range LVL {
		if p.Possibility_of_player.GetExperience() >= exp {
			p.Possibility_of_player.SetLVL(level)
		} else {
			return
		}
	}
}
