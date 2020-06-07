package wmap

import (
	"Test/chunk"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	var queue queue = &Node{}
	for i := 0; i < 5; i++ {
		queue.addInQueue(chunk.Coordinate{i, i + 1})
	}
	i := 0
	for i < 10 {
		i++
		s, e := queue.getData()
		if e != nil {
			fmt.Println(e.Error())
			break
		}
		fmt.Println(s)
	}
	m := make(map[string]bool)
	m["lol1"] = true
	fmt.Println(m["lol2"])
}
func TestStack(t *testing.T) {
	var stack stack = &Node{chunk.Coordinate{1, 1}, nil}
	for i := 0; i < 5; i++ {
		stack.addInStack(chunk.Coordinate{i, i + 1})
	}
	i := 0
	for i < 10 {
		i++
		s, e := stack.getDataS()
		if e != nil {
			fmt.Println(e.Error())
			break
		}
		fmt.Println(s)
	}
}
func TestAstar(t *testing.T) {
	World := NewCacheWorldMap()
	person := chunk.Coordinate{16, 16}
	target := chunk.Coordinate{304, 240}
	c := GetChunkID(person.X, person.Y)
	d := GetCurrentPlayerMap(c)
	_ = GetPlayerDrawChunkMap(d, &World)
	a := World.CheckBusyTile(target.X, target.Y)
	fmt.Println(target)
	if a {
		panic("Chunk is Busy")
	}
	g := createGraph(&World, person, target)
	fmt.Println("exist?", g[target])
	fmt.Println(g)
	q := Astar(g, person, target)
	q.printChild()
	var s stack = &Node{Data: person}

	s = createStackpath(q, s, person)
	fmt.Println(q)
	fmt.Println("start stack")
	fmt.Println(s.getDataS())
	fmt.Println(s.getDataS())
	fmt.Println(s.getDataS())
	fmt.Println(s.getDataS())
	fmt.Println(s.getDataS())
	fmt.Println(s.getDataS())
	fmt.Println(s.getDataS())
	fmt.Println(s.getDataS())
	fmt.Println(s.getDataS())
	fmt.Println(s.getDataS())

}
