package WorldMap

import (
	"Test/Chunk"
	"fmt"
)

type Graphpath map[Chunk.Coordinate][]Chunk.Coordinate

// сделать возврат ошибки
func createGraph(worldMap *WorldMap,person Chunk.Coordinate, target Chunk.Coordinate ) Graphpath {
	chunk:= GetChunkID(person.X, person.Y)
	var stack  stack  = &Node{}
	stack.addInStack(person)
	thisMap := GetCurrentPlayerMap(chunk)
	m:=GetPlayerDrawChunkMap(thisMap, worldMap)
	find := true
	graph := Graphpath{}

	for find {

		currentCoord, e := stack.getDataS()

		//if currentCoord == target{
		//	fmt.Println("FIND!", currentCoord)
		//	return graph
		//}
		// возврат ошибки
		if e != nil {
			fmt.Println("Stack empty", currentCoord)
			return graph
		}
		coords:= getAllCoordinate(currentCoord, &m)
		graph.addEdge(currentCoord, coords)
		stack = addCoordToStack(stack, coords, &m, graph)
		








	}
return graph
}

// переделать под переход только на свободные тайлы
func getAllCoordinate(person Chunk.Coordinate,  m *personMap) []Chunk.Coordinate {
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
var noBysyTile []Chunk.Coordinate
	for _, v := range coord {
		 t, e :=m.getTile(v)
		 if e != nil {
			 continue
		 }
		 if !t.Busy {
			noBysyTile = append(noBysyTile, v)
		 }
	}
return noBysyTile
}


func (g Graphpath) checkEdge(coord Chunk.Coordinate) bool  {
_, ok := g[coord]
return ok
}
func (g Graphpath) addEdge(coordParent Chunk.Coordinate, coords []Chunk.Coordinate)   {
for _, v := range coords {
	if g.checkEdge(v) {
		continue
	} else {
		g[coordParent]=append(g[coordParent], v)
	}
}
}
func addCoordToStack( s stack, coords  []Chunk.Coordinate, m *personMap, graphpath Graphpath) stack {
	for _, v := range coords{
		if !graphpath.checkEdge(v) {
			s.addInStack(v)
		}


	}
	return s
}