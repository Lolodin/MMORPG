package wmap

import (
	"Test/chunk"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Player struct {
	mut      sync.Mutex
	Name     string `json:"Name"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	password string
	speed    int
	move     bool
	//	AnimKey string
	walkPath chunk.Coordinate
}
type Players struct {
	P []Player `json:"players"`
}

func NewPlayer(n, password string) *Player {
	p := Player{}
	p.X = 16
	p.Y = 16
	p.Name = n
	p.password = password
	p.speed = 15
	p.move = false

	return &p

}

//Проверка, двигается ли персонаж
func (p *Player) isMove() bool {
	return p.move
}

//устанавливаем путь следования персонажа
func (p *Player) SetWalkPath(x, y int, m *WorldMap) {

	if p.X != x && p.Y != y {
		xy := chunk.Coordinate{X: x, Y: y}
		p.walkPath = xy
		// Использовать канал для сигнала?
		if !p.move {
			go p.walk(m)
		} else {
			p.mut.Lock()
			p.move = false
			p.mut.Unlock()
		}

	} else {
		return
	}

}

//Получаем путь куда должен перемещаться персонаж
func (p *Player) GetWalkPath() chunk.Coordinate {
	return p.walkPath
}

func (p *Player) GetPlayerXY() chunk.Coordinate {
	return chunk.Coordinate{X: p.X, Y: p.Y}
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
func (p *Player) moveSwitch() {
	if p.move {
		p.move = false
	} else {
		p.move = true
	}
}
func (p *Player) walk(m *WorldMap) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Path not found", r)
			p.mut.Lock()
			p.move = false
			p.mut.Unlock()
		}
	}()
	a := m.CheckBusyTile(p.walkPath.X, p.walkPath.Y)
	if a {
		log.WithFields(log.Fields{
			"package":  "WorldMap",
			"func":     "walk",
			"BusyTile": a,
			"Person":   chunk.Coordinate{p.X, p.Y},
			"target":   p.walkPath,
		}).Info("Tile Busy")
		return
	}
	p.mut.Lock()
	p.move = true
	p.mut.Unlock()
	graph := createGraph(m, chunk.Coordinate{p.X, p.Y}, p.walkPath)
	path := Astar(graph, chunk.Coordinate{p.X, p.Y}, p.walkPath)
	var s stack = &Node{}
	q := createStackpath(path, s, p.walkPath)
	i := true
	_, err := q.getDataS()
	if err != nil {
		fmt.Println(err.Error())

	}
	log.WithFields(log.Fields{
		"package": "WorldMap",
		"func":    "walk",
		"Person":  chunk.Coordinate{p.X, p.Y},
		"target":  p.walkPath,
		"path":    q,
	}).Info("Walker path")
	for i {

		e, err := q.getDataS()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(e, "Move")
		for true {
			time.Sleep(time.Duration(p.speed) * time.Millisecond)
			fmt.Println("GO,", e)
			if p.Y > e.Y {
				p.Y -= 1
			}
			if p.Y < e.Y {
				p.Y += 1
			}
			if p.X > e.X {
				p.X -= 1
			}
			if p.X < e.X {
				p.X += 1
			}
			if !p.move {
				p.X = e.X
				p.Y = e.Y
				fmt.Println("FALSE")
				return
			}

			if p.X == e.X && p.Y == e.Y {
				break
			}

		}
	}
	p.mut.Lock()
	p.move = false
	fmt.Println("STOP MOVE")
	p.mut.Unlock()
	return

}
