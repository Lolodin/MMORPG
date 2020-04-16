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

