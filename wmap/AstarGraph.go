package wmap

import (
	"Test/chunk"
)

type Graphpath map[chunk.Coordinate][]chunk.Coordinate

// сделать возврат ошибки
func createGraph(worldMap *WorldMap, person chunk.Coordinate, target chunk.Coordinate) Graphpath {
	Chunk := GetChunkID(person.X, person.Y)

	var stack queue = &Node{}
	stack.addInQueue(person)
	thisMap := GetCurrentPlayerMap(Chunk)
	m := GetPlayerDrawChunkMap(thisMap, worldMap)
	find := true
	graph := Graphpath{}
	visited := make(map[chunk.Coordinate]bool)
	for find {

		currentCoord, e := stack.getData()
		//fmt.Println("Find", currentCoord)
		// возврат ошибки
		if e != nil {
			//fmt.Println("Stack empty", currentCoord)
			return graph
		}
		coords := getAllCoordinate(currentCoord, &m, visited)
		graph.addEdge(currentCoord, coords)
		stack = addCoordToStack(stack, coords, &m, graph, visited)

	}
	return graph
}

// переделать под переход только на свободные тайлы
func getAllCoordinate(person chunk.Coordinate, m *personMap, visited map[chunk.Coordinate]bool) []chunk.Coordinate {
	var coord [8]chunk.Coordinate

	coord[0].X = person.X + 32
	coord[0].Y = person.Y + 32

	coord[1].X = person.X - 32
	coord[1].Y = person.Y - 32

	coord[2].X = person.X - 32
	coord[2].Y = person.Y + 32

	coord[3].X = person.X + 32
	coord[3].Y = person.Y - 32

	coord[4].X = person.X - 32
	coord[4].Y = person.Y

	coord[5].X = person.X + 32
	coord[5].Y = person.Y

	coord[6].X = person.X
	coord[6].Y = person.Y - 32

	coord[7].X = person.X
	coord[7].Y = person.Y + 32
	var noBysyTile []chunk.Coordinate
	for _, v := range coord {
		t, e := m.getTile(v)
		if e != nil || visited[v] {
			continue
		}
		if !t.Busy {
			noBysyTile = append(noBysyTile, v)
		}
	}
	return noBysyTile
}

func (g Graphpath) checkEdge(coord chunk.Coordinate) bool {
	_, ok := g[coord]
	return ok
}
func (g Graphpath) addEdge(coordParent chunk.Coordinate, coords []chunk.Coordinate) {
	for _, v := range coords {
		if g.checkEdge(v) {
			continue
		} else {
			g[coordParent] = append(g[coordParent], v)
		}
	}
}

// Добавить Координаты в стек
func addCoordToStack(s queue, coords []chunk.Coordinate, m *personMap, graphpath Graphpath, visited map[chunk.Coordinate]bool) queue {
	for _, v := range coords {
		if !graphpath.checkEdge(v) && !visited[v] {
			visited[v] = true
			s.addInQueue(v)
		}

	}
	return s
}
