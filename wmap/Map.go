package wmap

import (
	"Test/chunk"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
)

type personMap [9]chunk.Chunk

func (m *personMap) getTile(coordinate chunk.Coordinate) (chunk.Tile, error) {
	chunkID := GetChunkID(coordinate.X, coordinate.Y)
	for _, v := range m {
		if v.ChunkID == [2]int{chunkID.X, chunkID.Y} {
			return v.Map[coordinate], nil
		}
	}
	return chunk.Tile{}, fmt.Errorf("Tile not Found")
}

type playerMap struct {
	IDplayer int
	Map      [9]chunk.Chunk `json:"CurrentMap"`
}

/*
Получаем ID чанка из координат(персонажа\объекта и т.д.)
*/
func GetChunkID(x, y int) chunk.Coordinate {
	tileX := float64(float64(x) / float64(chunk.CHUNKIDSIZE))
	tileY := float64(float64(y) / float64(chunk.CHUNKIDSIZE))

	var ChunkID chunk.Coordinate
	if tileX < 0 {
		ChunkID.X = int(math.Floor(tileX / float64(chunk.CHUNKIDSIZE)))
	} else {
		ChunkID.X = int(math.Ceil(tileX / float64(chunk.CHUNKIDSIZE)))
	}
	if tileY < 0 {
		ChunkID.Y = int(math.Floor(tileY / float64(chunk.CHUNKIDSIZE)))
	} else {
		ChunkID.Y = int(math.Ceil(tileY / float64(chunk.CHUNKIDSIZE)))
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
func GetCurrentPlayerMap(currentChunkID chunk.Coordinate) [9]chunk.Coordinate {
	var CurrentMap [9]chunk.Coordinate
	coordinateX := currentChunkID.X * chunk.CHUNK_SIZE
	coordinateY := currentChunkID.Y * chunk.CHUNK_SIZE

	CurrentMap[0] = currentChunkID

	x := coordinateX + chunk.CHUNK_SIZE
	y := coordinateY + chunk.CHUNK_SIZE
	CurrentMap[1] = GetChunkID(x, y)
	x = coordinateX + chunk.CHUNK_SIZE
	y = coordinateY
	CurrentMap[2] = GetChunkID(x, y)
	if coordinateY < 0 {
		x = coordinateX + chunk.CHUNK_SIZE
		y = coordinateY - chunk.CHUNK_SIZE
	} else {
		x = coordinateX + chunk.CHUNK_SIZE
		y = coordinateY - chunk.CHUNK_SIZE - 1
	}
	CurrentMap[3] = GetChunkID(x, y)
	x = coordinateX
	y = coordinateY + chunk.CHUNK_SIZE
	CurrentMap[4] = GetChunkID(x, y)
	if coordinateY < 0 {
		x = coordinateX
		y = coordinateY - chunk.CHUNK_SIZE
	} else {
		x = coordinateX
		y = coordinateY - chunk.CHUNK_SIZE - 1
	}
	CurrentMap[5] = GetChunkID(x, y)
	if coordinateX < 0 {
		x = coordinateX - chunk.CHUNK_SIZE
		y = coordinateY + chunk.CHUNK_SIZE
	} else {
		x = coordinateX - chunk.CHUNK_SIZE - 1
		y = coordinateY + chunk.CHUNK_SIZE
	}
	CurrentMap[6] = GetChunkID(x, y)
	if coordinateX < 0 {
		x = coordinateX - chunk.CHUNK_SIZE
		y = coordinateY
	} else {
		x = coordinateX - chunk.CHUNK_SIZE - 1
		y = coordinateY
	}
	CurrentMap[7] = GetChunkID(x, y)
	if coordinateX < 0 && coordinateY < 0 {
		x = coordinateX - chunk.CHUNK_SIZE
		y = coordinateY - chunk.CHUNK_SIZE
	} else {
		if coordinateX > 0 {
			x = coordinateX - chunk.CHUNK_SIZE - 1
		} else {
			x = coordinateX - chunk.CHUNK_SIZE
		}
		if coordinateY < 0 {
			y = coordinateY - chunk.CHUNK_SIZE
		} else {
			y = coordinateY - chunk.CHUNK_SIZE - 1
		}

	}
	CurrentMap[8] = GetChunkID(x, y)
	return CurrentMap
}

//Получаем готовую карту из 9 чанков для отображения
func GetPlayerDrawChunkMap(currentMap [9]chunk.Coordinate, W *WorldMap) personMap {
	var playerMap personMap
	for i, m := range currentMap {
		if W.isChunkExist(m) {
			playerMap[i], _ = W.GetChunk(m)
		} else {
			chunk := chunk.NewChunk(m)
			playerMap[i] = chunk
			W.AddChunk(m, chunk)
		}

	}
	return playerMap
}

func MapToJSON(m [9]chunk.Chunk, id int) []byte {
	a := playerMap{
		IDplayer: id,
		Map:      m,
	}
	r, e := json.Marshal(a)
	if e != nil {
		log.WithFields(log.Fields{
			"package": "worldMap",
			"func":    "Map to Json",
			"error":   e,
			"map":     a.Map,
			"player":  a.IDplayer,
		}).Warning("Error Marshal player map")
		return nil
	}
	return r
}
