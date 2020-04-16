package WorldMap

import (
	"Test/Chunk"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	var queue queue = &Node{}
	for i := 0; i < 5; i++ {
		queue.addInQueue(Chunk.Coordinate{i, i + 1})
	}
	i:=0
	for i<10 {
		i++
		s, e:= queue.getData()
		if e!=nil {
			fmt.Println(e.Error())
			break
		}
		fmt.Println(s)
	}
	m:= make(map[string]bool)
	m["lol1"] = true
	fmt.Println(m["lol2"])
}
func TestStack(t *testing.T)  {
	var stack stack = &Node{Chunk.Coordinate{1,1}, nil}
	for i := 0; i < 5; i++ {
		stack.addInStack(Chunk.Coordinate{i, i + 1})
	}
i:=0
for i<10 {
	i++
	s, e:= stack.getDataS()
	if e!=nil {
		fmt.Println(e.Error())
		break
	}
	fmt.Println(s)
}
}
