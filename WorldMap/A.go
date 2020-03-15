package WorldMap

import (
	"Test/Chunk"
	"fmt"
)

type queue interface {
	addInQueue(coor Chunk.Coordinate)
	checkChild() bool
	getData() (Chunk.Coordinate, error)
}
type stack interface {
	addInStack(coor Chunk.Coordinate)
	getDataS() (Chunk.Coordinate, error)
}

type Node struct {
	Data Chunk.Coordinate
	NextNode *Node
}
// Добавляем элемент в очередь
func (n *Node) addInQueue(coor Chunk.Coordinate) {
	n2:= Node{Data:coor}
if n.checkChild() {
	n.NextNode.addInQueue(coor)
} else {
	n.NextNode = &n2
}
}
// Првоеряем содержит очередь потомков или нет
func (n *Node) checkChild() bool {
	return n.NextNode != nil
}
// Добавляем элемент в стек
func (n *Node) addInStack(coor Chunk.Coordinate) {
	n2:= Node{Data:coor}
	node:= *n
n2.NextNode = &node
n.Data = n2.Data
n.NextNode =n2.NextNode


}
//Возвращаем элемент очереди
func (n *Node) getData() (Chunk.Coordinate, error) {
		nextNode:= n.NextNode
	if nextNode == nil {
		return Chunk.Coordinate{}, fmt.Errorf("queue is empty")
	}
		n.Data = nextNode.Data
		n.NextNode = nextNode.NextNode
		return nextNode.Data, nil
}
//Возвращаем элемент стека
func (n *Node) getDataS() (Chunk.Coordinate, error) {
 			data:= n.Data
 			node := n.NextNode
	if node == nil {
		return Chunk.Coordinate{}, fmt.Errorf("Stack is empty")
	}
 			n.NextNode = node.NextNode
 			n.Data= node.Data
 			return data, nil
}
//Функция поиска пути, возвращает Очередь из координат по которой пойдет персонаж
func(worldMap *WorldMap) A(person, target Chunk.Coordinate) queue {
	find:= true
	visited:= make(map[Chunk.Coordinate]bool)
	var path queue = &Node{}
chunk:= GetChunkID(person.X, person.Y)
thisMap := GetCurrentPlayerMap(chunk)
m:=GetPlayerDrawChunkMap(thisMap, worldMap)
var stack stack = &Node{}

	stack.addInStack(person)
	for find {
		person, e := stack.getDataS()
		el, _:=m.getTile(person)
		fmt.Println(person, "dddddddddddddd", el.Busy)
		path.addInQueue(person)
		if e!= nil {
fmt.Println(e.Error())
		}
		var optCoord Chunk.Coordinate
		optCoord = getOptCoordinate(person, target)
		optTile, e:= m.getTile(optCoord)
		if (optTile.Busy || e!= nil) || visited[optCoord]{
			coords:= getAllCoordinate(person)
			for _, v:= range coords{
				tile, e:= m.getTile(v)
				if e!= nil || visited[v] || tile.Busy{
					continue
				}
				if !tile.Busy && !visited[v]{
					stack.addInStack(v)
				}
			}
			visited[person] = true


		} else {
			visited[person] = true
			stack.addInStack(optCoord)
		}
		if person == target {
			find = false
			fmt.Println("exit")
		}


	}
	return path

}

func getOptCoordinate(person, target Chunk.Coordinate) Chunk.Coordinate {
	var optCoord Chunk.Coordinate
	switch {
	case person.X<target.X && person.Y<target.Y:
		optCoord.X = person.X+32
		optCoord.Y = person.Y+32
	case person.X>target.X && person.Y>target.Y :
		optCoord.X = person.X-32
		optCoord.Y = person.Y-32
	case person.X>target.X && person.Y<target.Y :
		optCoord.X = person.X-32
		optCoord.Y = person.Y+32
	case person.X<target.X && person.Y>target.Y :
		optCoord.X = person.X+32
		optCoord.Y = person.Y-32
	case person.X>target.X && person.Y==target.Y :
		optCoord.X = person.X-32
		optCoord.Y = person.Y
	case person.X<target.X && person.Y==target.Y :
		optCoord.X = person.X+32
		optCoord.Y = person.Y
	case person.X==target.X && person.Y>target.Y :
		optCoord.X = person.X
		optCoord.Y = person.Y-32
	case person.X==target.X && person.Y<target.Y :
		optCoord.X = person.X
		optCoord.Y = person.Y+32
	}
	return optCoord
}
func getAllCoordinate(person Chunk.Coordinate) [8]Chunk.Coordinate {
var coord [8]Chunk.Coordinate

	coord[0].X = person.X+32
	coord[0].Y = person.Y+32

	coord[1].X = person.X-32
	coord[1].Y = person.Y-32

	coord[2].X = person.X-32
	coord[2].Y = person.Y+32

	coord[3].X = person.X+32
	coord[3].Y = person.Y-32

	coord[4].X = person.X-32
	coord[4].Y = person.Y

	coord[5].X = person.X+32
	coord[5].Y = person.Y

	coord[6].X = person.X
	coord[6].Y = person.Y-32

	coord[7].X = person.X
	coord[7].Y = person.Y+32
return coord
}