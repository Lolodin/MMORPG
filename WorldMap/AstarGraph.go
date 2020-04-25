package WorldMap

import "Test/Chunk"

type Graphpath map[Chunk.Coordinate][]Chunk.Coordinate

// сделать возврат ошибки
func createGraph(worldMap *WorldMap,person Chunk.Coordinate, target Chunk.Coordinate ) Graphpath {
	chunk:= GetChunkID(person.X, person.Y)
	var stack  stack  = &Node{}
	stack.addInStack(person)
	thisMap := GetCurrentPlayerMap(chunk)
	m:=GetPlayerDrawChunkMap(thisMap, worldMap)
	find := false
	graph := Graphpath{}

	for find {
		currentCoord, e := stack.getDataS()
		if currentCoord == target{
			return graph
		}
		// возврат ошибки
		if e != nil {
			return graph
		}
		coords:= getAllCoordinate(currentCoord, &m)
		graph.addEdge(currentCoord, coords)
		stack = addCoordToStack(stack, coords, &m)
		








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
func addCoordToStack( s stack, coords  []Chunk.Coordinate, m *personMap) stack {
	for _, v := range coords{
		s.addInStack(v)

	}
	return s
}