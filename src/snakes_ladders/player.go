package snakes_ladders

type Player struct {
	name string
	id   int
}

func NewPlayer(name string, id int) *Player {
	return &Player{
		name: name,
		id:   id,
	}
}

func (p *Player) getPlayerName() string {
	return p.name
}

func (p *Player) setPlayerName(name string) {
	p.name = name
}

func (p *Player) getPlayerID() int {
	return p.id
}

func (p *Player) setPlayerID(id int) {
	p.id = id
}
