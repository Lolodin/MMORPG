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
	// Добавляем элемент в стек
	addInStack(coor Chunk.Coordinate)
	getDataS() (Chunk.Coordinate, error)
}
type list interface {
	addNextNode(node *Node)

}


type Node struct {
	Data Chunk.Coordinate
	NextNode *Node
}
func (n *Node) addNextNode(n2 Node) {
	if n.checkChild() {
		n.NextNode.addNextNode(n2)
	} else {
		n.NextNode = &n2
	}
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
func Astar(graphpath Graphpath, person Chunk.Coordinate, target Chunk.Coordinate) Node {
var q stack = &Node{}
q.addInStack(person)
var coord Chunk.Coordinate

visited:= make(map[Chunk.Coordinate]bool)
path:= make(map[Chunk.Coordinate]Node)
for coord!= target {


coord, e := q.getDataS()
	fmt.Println(coord,"newLOOP")
	if visited[coord] {
		fmt.Println("Visited")
		a, e := q.getDataS()
		coord = a
		if e!= nil {
			panic("error get chunk")
		}
		continue
	}
if e!= nil {
	panic("error get chunk")
}
n1:= &Node{}
n1.Data = coord
path[coord] = *n1

visited[coord] = true
g := graphpath[coord]
fmt.Println(g)
for _, v := range g {

	if visited[v] {
		continue
	}
	q.addInStack(v)
n1 = &Node{}
n1.Data = v
n1.addNextNode(path[coord])
path[v] = *n1

}
coord, e = q.getDataS()
	if coord == target {
		return  path[coord]
	}
}
return path[coord]




}
// return nil if stack empty
func createStackpath(node Node, s stack, target Chunk.Coordinate) stack {
	if node.NextNode !=nil {
		s.addInStack(node.Data)
		createStackpath(*node.NextNode, s, target)
	} else {
		s.addInStack(node.Data)
		s.addInStack(target)
		return s
	}
	return s

}