package WorldMap

import (
	"Test/Chunk"
	"fmt"
)

type queue interface {
	addInQueue(n2 Node)
	checkChild() bool
	getData() (Chunk.Coordinate, error)
}
type stack interface {
	addInStack(n2 Node)
	getDataS() (Chunk.Coordinate, error)
}
type Node struct {
	Data Chunk.Coordinate
	NextNode *Node
}
func (n *Node) addInQueue(n2 Node) {
if n.checkChild() {
	n.NextNode.addInQueue(n2)
} else {
	n.NextNode = &n2
}
}
func (n *Node) checkChild() bool {
	return n.NextNode != nil
}
func (n *Node) addInStack(n2 Node) {
	node:= *n
n2.NextNode = &node
n.Data = n2.Data
n.NextNode =n2.NextNode


}
func (n *Node) getData() (Chunk.Coordinate, error) {
		nextNode:= n.NextNode
	if nextNode == nil {
		return Chunk.Coordinate{nil, nil}, fmt.Errorf("queue is empty")
	}
		n.Data = nextNode.Data
		n.NextNode = nextNode.NextNode
		return nextNode.Data, nil
}
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
// Функция поиска пути, возвращает Очередь из координат по которой пойдет персонаж
//func(worldMap *WorldMap) A(person,target Chunk.Coordinate) {
// chunk:= GetChunkID(person.X, person.Y)
// thisMap := GetCurrentPlayerMap(chunk)
// m:=GetPlayerDrawChunkMap(thisMap, worldMap)
// t, e := m.getTile(person)
//if e != nil {
//	fmt.Println(e.Error())
//}
//
//
//}
