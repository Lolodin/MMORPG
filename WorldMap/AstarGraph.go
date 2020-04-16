package WorldMap

import "Test/Chunk"

type Grathpath map[Chunk.Coordinate][]Chunk.Coordinate
func createGrath(worldMap *WorldMap,person Chunk.Coordinate, target Chunk.Coordinate ) {
	chunk:= GetChunkID(person.X, person.Y)
	thisMap := GetCurrentPlayerMap(chunk)
	m:=GetPlayerDrawChunkMap(thisMap, worldMap)
	find := false

	for find {








	}

}
func getAllCoordinate(person Chunk.Coordinate) [8]Chunk.Coordinate {
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
	return coord
}
func addEdge(person Chunk.Coordinate, target Chunk.Coordinate, grathpath *Grathpath) {

}

func (g Grathpath) checkEdge(coord Chunk.Coordinate) bool  {
_, ok := g[coord]
return ok
}