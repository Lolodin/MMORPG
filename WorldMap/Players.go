package WorldMap

import (
	"Test/Chunk"
	"time"
)

type Player struct {
	Name     string `json:"Name"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	password string
	speed    int
	move     bool
	//	AnimKey string
	walkPath Chunk.Coordinate
}
type Players struct {
	P []Player `json:"players"`
}

func NewPlayer(n, password string) *Player {
	p := Player{}
	p.X = 0
	p.Y = 0
	p.Name = n
	p.password = password
	p.speed = 5
	p.walkPath.X = p.X
	p.walkPath.Y = p.Y
	go p.walk()

	return &p

}

//Проверка, двигается ли персонаж
func (p *Player) isMove() bool {
	return p.move
}

//устанавливаем путь следования персонажа
func (p *Player) SetWalkPath(x, y int) {
	if p.X != x && p.Y != y {
		xy := Chunk.Coordinate{X: x, Y: y}
		p.walkPath = xy
	} else {
		return
	}

}

//Получаем путь куда должен перемещаться персонаж
func (p *Player) GetWalkPath() Chunk.Coordinate {
	return p.walkPath
}

func (p *Player) SetPassword(pass string) {
	p.password = pass
}
func (p *Player) GetPassword() string {
	return p.password
}

// bool true if pass == player.password
func (p *Player) ComparePassword(pass string) bool {
	if pass == p.password {
		return true
	} else {
		return false
	}
}
func (p *Player) walk() {
	for  {
		time.Sleep(25 * time.Millisecond)
		if p.Y >p.walkPath.Y {
 				p.Y -= p.speed
		}
		if p.Y <p.walkPath.Y {
			p.Y += p.speed
		}
		if p.X > p.walkPath.X {
			p.X -= p.speed
		}
		if p.X < p.walkPath.X {
			p.X += p.speed
		}
	}

}
