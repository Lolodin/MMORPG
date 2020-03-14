package WorldMap

import (
	"Test/Chunk"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	var queue queue = &Node{}
	for i := 0; i < 2; i++ {
		queue.addInQueue(Node{Data: Chunk.Coordinate{i, i + 1}})
	}
	fmt.Println(queue.getData())
	fmt.Println(queue.getData())
	fmt.Println(queue.getData())
}
func TestStack(t *testing.T)  {
	var stack stack = &Node{}
	for i := 0; i < 2; i++ {
		stack.addInStack(Node{Data: Chunk.Coordinate{i, i + 1}})
	}
i:=0
for i<3 {
	i++
	s, e:= stack.getDataS()
	if e!=nil {
		fmt.Println(e.Error())
		break
	}
	fmt.Println(s)
}
}
