package WorldMap

import (
	"Test/Chunk"
	"fmt"
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
	walkPath Chunk.Coordinate
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
	p.speed = 5
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
		xy := Chunk.Coordinate{X: x, Y: y}
		p.walkPath = xy
		if !p.move {
			go p.walk(m)
		}

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
func (p *Player) moveSwitch() {
	if p.move {
		p.move = false
	} else {
		p.move = true
	}
}
func (p *Player) walk(m *WorldMap) {
	//p.mut.Lock()
	//p.move = true
	//p.mut.Unlock()
	//for p.move {
	//	time.Sleep(25 * time.Millisecond)
	//	if p.Y > p.walkPath.Y {
	//		p.Y -= p.speed
	//	}
	//	if p.Y < p.walkPath.Y {
	//		p.Y += p.speed
	//	}
	//	if p.X > p.walkPath.X {
	//		p.X -= p.speed
	//	}
	//	if p.X < p.walkPath.X {
	//		p.X += p.speed
	//	}
	//	if (p.X == p.walkPath.X && p.Y == p.walkPath.Y) || p.move == false {
	//		p.mut.Lock()
	//		p.move = false
	//		p.mut.Unlock()
	//		return
	//	}

	//__________________________________________________________
	p.mut.Lock()
	p.move = true
	p.mut.Unlock()
	graph := createGraph(m, Chunk.Coordinate{p.X,p.Y}, p.walkPath )
	path:=Astar(graph,Chunk.Coordinate{p.X,p.Y}, p.walkPath)
	var s stack = &Node{}
	q:= createStackpath(path, s, p.walkPath)
	i := true
	for i {
		fmt.Println("Move")
		time.Sleep(1 * time.Second)
		e, err := q.getDataS()
		if err!=nil{
			fmt.Println(err.Error())
			break
		}
		p.X = e.X
		p.Y = e.Y
		fmt.Println(e)
	}
	p.mut.Lock()
	p.move = false
	p.mut.Unlock()
	return

}

