package WorldMap

import (
	"Test/Chunk"
)

type Graphpath map[Chunk.Coordinate][]Chunk.Coordinate

// сделать возврат ошибки
func createGraph(worldMap *WorldMap,person Chunk.Coordinate, target Chunk.Coordinate ) Graphpath {
	chunk:= GetChunkID(person.X, person.Y)
	var stack  queue  = &Node{}
	stack.addInQueue(person)
	thisMap := GetCurrentPlayerMap(chunk)
	m:=GetPlayerDrawChunkMap(thisMap, worldMap)
	find := true
	graph := Graphpath{}
	visited:= make(map[Chunk.Coordinate]bool)
	for find {


		currentCoord, e := stack.getData()
		//fmt.Println("Find", currentCoord)
		// возврат ошибки
		if e != nil {
			//fmt.Println("Stack empty", currentCoord)
			return graph
		}
		coords:= getAllCoordinate(currentCoord, &m, visited)
		graph.addEdge(currentCoord, coords)
		stack = addCoordToStack(stack, coords, &m, graph, visited)
		








	}
return graph
}

// переделать под переход только на свободные тайлы
func getAllCoordinate(person Chunk.Coordinate,  m *personMap, visited map[Chunk.Coordinate]bool ) []Chunk.Coordinate {
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
		 if e != nil || visited[v] {
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
func addCoordToStack( s queue, coords  []Chunk.Coordinate, m *personMap, graphpath Graphpath, visited map[Chunk.Coordinate]bool) queue {
	for _, v := range coords{
		if !graphpath.checkEdge(v) && !visited[v]{
			visited[v] = true
			s.addInQueue(v)
		}


	}
	return s
}