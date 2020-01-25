package WorldMap

import "Test/Chunk"

type Player struct {
	Name     string `json:"Name"`
	X        int `json:"X"`
	Y        int `json:"Y"`
	speed    int
	move     bool
//	AnimKey string
	walkPath Chunk.Coordinate
}
type Players struct {
	P []Player `json:"players"`
}

func NewPlayer(n string, id int) Player {
	p := Player{}
	p.X = 0
	p.Y = 0
	p.Name = n

	return p

}

//Проверка, двигается ли персонаж
func (p *Player) isMove() bool {
	return p.move
}
//устанавливаем путь следования персонажа
func (p *Player) setWalkPath(x, y int) {
	if p.X != x && p.Y != y {
		xy := Chunk.Coordinate{X: x, Y: y}
		p.walkPath = xy
	} else {
		return
	}


}
//Получаем путь куда должен перемещаться персонаж
func (p *Player) getWalkPath() Chunk.Coordinate {
	return p.walkPath
}
