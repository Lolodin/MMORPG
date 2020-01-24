package WorldMap

import "Test/Chunk"

type Player struct {
	Name     string
	X        int
	Y        int
	speed    int
	move     bool
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
func (p Player) isMove() bool {
	return p.move
}
func (p Player) setWalkPath(x, y int) {
	xy := Chunk.Coordinate{X: x, Y: y}
	p.walkPath = xy
}
func (p Player) getWalkPath() Chunk.Coordinate {
	return p.walkPath
}
