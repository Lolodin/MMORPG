package Chunk

import (
	"Test/PerlinNoise"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
)

const CHUNKIDSIZE = 32
const TILE_SIZE = 32
const CHUNK_SIZE = 32 * 32
const PERLIN_SEED float32 = 2300
// Чанк который хранит тайтлы и другие игровые объекты
type Chunk struct {
	ChunkID [2]int
	Map     map[Coordinate]Tile
	Tree    map[Coordinate]Tree
}

/*
Тайтл игрового мира
*/
type Tile struct {
	Key string `json:"key"`
	X   int    `json:"x"`
	Y   int    `json:"y"`
	Busy bool
}

/*
Универсальная структура для хранения координат
*/
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (t Coordinate) MarshalText() ([]byte, error) {

	return []byte("[" + strconv.Itoa(t.X) + "," + strconv.Itoa(t.Y) + "]"), nil
}

/*
Создает карту чанка из тайтлов, генерирует карту на основе координаты чанка
Например [1,1]
*/
func NewChunk(idChunk Coordinate) Chunk {

	log.WithFields(log.Fields{
		"package": "Chunk",
		"func" : "NewChunk",
		"idChunk": idChunk,

	}).Info("Create new Chunck")
	chunk := Chunk{ChunkID: [2]int{idChunk.X, idChunk.Y}}
	var chunkXMax, chunkYMax int
	var chunkMap map[Coordinate]Tile
	var treeMap map[Coordinate]Tree
	chunkMap = make(map[Coordinate]Tile)
	treeMap = make(map[Coordinate]Tree)
	chunkXMax = idChunk.X * CHUNK_SIZE
	chunkYMax = idChunk.Y * CHUNK_SIZE
	var tree Tree

	switch {
	case chunkXMax < 0 && chunkYMax < 0:
		{
			for x := chunkXMax + CHUNK_SIZE; x > chunkXMax; x -= TILE_SIZE {
				for y := chunkYMax + CHUNK_SIZE; y > chunkYMax; y -= TILE_SIZE {

					posX := float32(x - (TILE_SIZE / 2))
					posY := float32(y + (TILE_SIZE / 2))
					tile := Tile{}

					tile.X = int(posX)
					tile.Y = int(posY)

					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.01:
						tile.Key = "Water"
					case perlinValue >= -0.01 && perlinValue < 0:
						tile.Key = "Sand"
					case perlinValue >= 0 && perlinValue <= 0.5:
						tile.Key = "Ground"
						rand.Seed(int64(time.Now().Nanosecond() + x - y))
						randomTree := rand.Float32()

						if randomTree > 0.99 {
							tree = NewTree(Coordinate{X: tile.X, Y: tile.Y})
						}
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}

					if tree.Species != "" {
						treeMap[Coordinate{X: tree.X, Y: tree.Y}] = tree
						tile.Busy = true
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

				}
			}
		}
	case chunkXMax < 0:
		{
			for x := chunkXMax + CHUNK_SIZE; x > chunkXMax; x -= TILE_SIZE {
				for y := chunkYMax - CHUNK_SIZE; y < chunkYMax; y += TILE_SIZE {
					posX := float32(x - (TILE_SIZE / 2))
					posY := float32(y + (TILE_SIZE / 2))

					tile := Tile{}

					tile.X = int(posX)
					tile.Y = int(posY)

					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.012:
						tile.Key = "Water"
					case perlinValue >= -0.012 && perlinValue < 0:
						tile.Key = "Sand"
					case perlinValue >= 0 && perlinValue <= 0.5:
						tile.Key = "Ground"
						rand.Seed(int64(time.Now().Nanosecond() + x - y))
						randomTree := rand.Float32()
						if randomTree > 0.95 {
							tree = NewTree(Coordinate{X: tile.X, Y: tile.Y})
						}
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}

					if tree.Species != "" {
						treeMap[Coordinate{X: tree.X, Y: tree.Y}] = tree
						tile.Busy = true
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

				}
			}
		}
	case chunkYMax < 0:
		{
			for x := chunkXMax - CHUNK_SIZE; x < chunkXMax; x += TILE_SIZE {
				for y := chunkYMax + CHUNK_SIZE; y > chunkYMax; y -= TILE_SIZE {

					posX := float32(x + (TILE_SIZE / 2))
					posY := float32(y - (TILE_SIZE / 2))
					tile := Tile{}

					tile.X = int(posX)
					tile.Y = int(posY)
					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.012:
						tile.Key = "Water"
					case perlinValue >= -0.012 && perlinValue < 0:
						tile.Key = "Sand"
					case perlinValue >= 0 && perlinValue <= 0.5:
						tile.Key = "Ground"
						rand.Seed(int64(time.Now().Nanosecond() + x - y))
						randomTree := rand.Float32()
						if randomTree > 0.95 {
							tree = NewTree(Coordinate{X: tile.X, Y: tile.Y})
						}
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile
					if tree.Species != "" {
						treeMap[Coordinate{X: tree.X, Y: tree.Y}] = tree
						tile.Busy = true
					}

				}
			}
		}
	default:
		{
			for x := chunkXMax - CHUNK_SIZE; x < chunkXMax; x += TILE_SIZE {
				for y := chunkYMax - CHUNK_SIZE; y < chunkYMax; y += TILE_SIZE {
					posX := float32(x + (TILE_SIZE / 2))
					posY := float32(y + (TILE_SIZE / 2))
					tile := Tile{}

					tile.X = int(posX)
					tile.Y = int(posY)
					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)

					switch {
					case perlinValue < -0.012:
						tile.Key = "Water"
					case perlinValue >= -0.012 && perlinValue < 0:
						tile.Key = "Sand"
					case perlinValue >= 0 && perlinValue <= 0.5:
						tile.Key = "Ground"
						rand.Seed(int64(time.Now().Nanosecond() + x - y))
						randomTree := rand.Float32()
						if randomTree > 0.95 {
							tree = NewTree(Coordinate{X: tile.X, Y: tile.Y})
						}
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}
					if tree.Species != "" {
						treeMap[Coordinate{X: tree.X, Y: tree.Y}] = tree
						tile.Busy = true
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

				}
			}
		}

	}

	chunk.Map = chunkMap
	chunk.Tree = treeMap


	return chunk
}
