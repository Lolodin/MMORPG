package WorldMap

import (
	"Test/Chunk"
	"encoding/json"
	"fmt"
	"math"
)

type playerMap struct {
	IDplayer int
	Map      [9]Chunk.Chunk `json:"CurrentMap"`
}

/*
Получаем ID чанка из координат(персонажа\объекта и т.д.)
*/
func GetChunkID(x, y int) Chunk.Coordinate {
	tileX := float64(float64(x) / float64(Chunk.CHUNKIDSIZE))
	tileY := float64(float64(y) / float64(Chunk.CHUNKIDSIZE))

	var ChunkID Chunk.Coordinate
	if tileX < 0 {
		ChunkID.X = int(math.Floor(tileX / float64(Chunk.CHUNKIDSIZE)))
	} else {
		ChunkID.X = int(math.Ceil(tileX / float64(Chunk.CHUNKIDSIZE)))
	}
	if tileY < 0 {
		ChunkID.Y = int(math.Floor(tileY / float64(Chunk.CHUNKIDSIZE)))
	} else {
		ChunkID.Y = int(math.Ceil(tileY / float64(Chunk.CHUNKIDSIZE)))
	}
	if tileX == 0 {
		ChunkID.X = 1
	}
	if tileY == 0 {
		ChunkID.Y = 1
	}
	return ChunkID

}

/*
Получаем текущую карту которую должен видеть персонаж
*/
func GetCurrentPlayerMap(currentChunkID Chunk.Coordinate) [9]Chunk.Coordinate {
	var CurrentMap [9]Chunk.Coordinate
	coordinateX := currentChunkID.X * Chunk.CHUNK_SIZE
	coordinateY := currentChunkID.Y * Chunk.CHUNK_SIZE

	CurrentMap[0] = currentChunkID

	x := coordinateX + Chunk.CHUNK_SIZE
	y := coordinateY + Chunk.CHUNK_SIZE
	CurrentMap[1] = GetChunkID(x, y)
	x = coordinateX + Chunk.CHUNK_SIZE
	y = coordinateY
	CurrentMap[2] = GetChunkID(x, y)
	if coordinateY < 0 {
		x = coordinateX + Chunk.CHUNK_SIZE
		y = coordinateY - Chunk.CHUNK_SIZE
	} else {
		x = coordinateX + Chunk.CHUNK_SIZE
		y = coordinateY - Chunk.CHUNK_SIZE - 1
	}
	CurrentMap[3] = GetChunkID(x, y)
	x = coordinateX
	y = coordinateY + Chunk.CHUNK_SIZE
	CurrentMap[4] = GetChunkID(x, y)
	if coordinateY < 0 {
		x = coordinateX
		y = coordinateY - Chunk.CHUNK_SIZE
	} else {
		x = coordinateX
		y = coordinateY - Chunk.CHUNK_SIZE - 1
	}
	CurrentMap[5] = GetChunkID(x, y)
	if coordinateX < 0 {
		x = coordinateX - Chunk.CHUNK_SIZE
		y = coordinateY + Chunk.CHUNK_SIZE
	} else {
		x = coordinateX - Chunk.CHUNK_SIZE - 1
		y = coordinateY + Chunk.CHUNK_SIZE
	}
	CurrentMap[6] = GetChunkID(x, y)
	if coordinateX < 0 {
		x = coordinateX - Chunk.CHUNK_SIZE
		y = coordinateY
	} else {
		x = coordinateX - Chunk.CHUNK_SIZE - 1
		y = coordinateY
	}
	CurrentMap[7] = GetChunkID(x, y)
	if coordinateX < 0 && coordinateY < 0 {
		x = coordinateX - Chunk.CHUNK_SIZE
		y = coordinateY - Chunk.CHUNK_SIZE
	} else {
		if coordinateX > 0 {
			x = coordinateX - Chunk.CHUNK_SIZE - 1
		} else {
			x = coordinateX - Chunk.CHUNK_SIZE
		}
		if coordinateY < 0 {
			y = coordinateY - Chunk.CHUNK_SIZE
		} else {
			y = coordinateY - Chunk.CHUNK_SIZE - 1
		}

	}
	CurrentMap[8] = GetChunkID(x, y)
	return CurrentMap
}

//Получаем готовую карту из 9 чанков для отображения игроку
func GetPlayerDrawChunkMap(currentMap [9]Chunk.Coordinate, W *WorldMap) [9]Chunk.Chunk {
	var playerMap [9]Chunk.Chunk
	for i, m := range currentMap {
		if W.isChunkExist(m) {
			playerMap[i], _ = W.GetChunk(m)
		} else {
			chunk := Chunk.NewChunk(m)
			playerMap[i] = chunk
			W.AddChunk(m, chunk)
		}

	}
	return playerMap
}

func MapToJSON(m [9]Chunk.Chunk, id int) []byte {
	a := playerMap{
		IDplayer: id,
		Map:      m,
	}
	r, e := json.Marshal(a)
	if e != nil {
		fmt.Println("error", e)
		return nil
	}
	return r
}
