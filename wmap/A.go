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
	Data     Chunk.Coordinate
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
	n2 := Node{Data: coor}
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
	n2 := Node{Data: coor}
	node := *n
	n2.NextNode = &node
	n.Data = n2.Data
	n.NextNode = n2.NextNode

}

//Возвращаем элемент очереди
func (n *Node) getData() (Chunk.Coordinate, error) {
	nextNode := n.NextNode
	if nextNode == nil {
		return Chunk.Coordinate{}, fmt.Errorf("queue is empty")
	}
	n.Data = nextNode.Data
	n.NextNode = nextNode.NextNode
	return nextNode.Data, nil
}

//Возвращаем элемент стека
func (n *Node) getDataS() (Chunk.Coordinate, error) {
	data := n.Data
	node := n.NextNode
	if node == nil {
		return Chunk.Coordinate{}, fmt.Errorf("Stack is empty")
	}
	n.NextNode = node.NextNode
	n.Data = node.Data
	return data, nil
}
func (n *Node) printChild() {
	fmt.Println(n.Data, "print child")
	if n.NextNode != nil {
		n.NextNode.printChild()
	}
}

//Функция поиска пути, возвращает Очередь из координат
func Astar(graphpath Graphpath, person Chunk.Coordinate, target Chunk.Coordinate) Node {
	var q queue = &Node{}
	q.addInQueue(person)
	var coord Chunk.Coordinate

	visited := make(map[Chunk.Coordinate]bool)
	path := make(map[Chunk.Coordinate]Node)
	n1 := &Node{}
	n1.Data = coord
	path[person] = *n1
	for coord != target {

		coord, e := q.getData()
		if coord == target {
			return path[coord]
		}

		if visited[coord] {
			// fmt.Println("Visited")
			if e != nil {
				panic("error get chunk")
			}
			continue
		}
		if e != nil {
			panic("error get chunk")
		}

		visited[coord] = true
		g := graphpath[coord]
		fmt.Println(g)
		for _, v := range g {
			_, ok := path[v]
			if visited[v] || ok {
				continue
			}
			q.addInQueue(v)
			n2 := &Node{}
			n2.Data = v
			n2.addNextNode(path[coord])
			path[v] = *n2
			if v == target {
				return path[v]
			}
		}

	}
	return path[target]

}

// return nil if stack empty
func createStackpath(node Node, s stack, p Chunk.Coordinate) stack {
	if node.NextNode != nil {
		fmt.Println(s, "+stack")
		s.addInStack(node.Data)
		createStackpath(*node.NextNode, s, p)
	} else {
		s.addInStack(p)

	}
	return s

}
