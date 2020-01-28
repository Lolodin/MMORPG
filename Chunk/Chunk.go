package Chunk

import (
	"Test/PerlinNoise"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var TILE_SIZE = 32
var CHANK_SIZE = 32 * 32
var PERLIN_SEED float32 = 1700

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
	X   int `json:"x"`
	Y   int `json:"y"`
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
	fmt.Println("New Chank", idChunk)
	chunk := Chunk{ChunkID: [2]int{idChunk.X, idChunk.Y}}
	var chunkXMax, chunkYMax int
	var chunkMap map[Coordinate]Tile
	var treeMap map[Coordinate]Tree
	chunkMap = make(map[Coordinate]Tile)
	treeMap = make(map[Coordinate]Tree)
	chunkXMax = idChunk.X * CHANK_SIZE
	chunkYMax = idChunk.Y * CHANK_SIZE
	var tree Tree

	switch {
	case chunkXMax < 0 && chunkYMax < 0:
		{
			for x := chunkXMax + CHANK_SIZE; x > chunkXMax; x -= TILE_SIZE {
				for y := chunkYMax + CHANK_SIZE; y > chunkYMax; y -= TILE_SIZE {

					posX := float32(x - 8)
					posY := float32(y + 8)
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
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile
					treeMap[Coordinate{X: tree.X, Y: tree.Y}] = tree

				}
			}
		}
	case chunkXMax < 0:
		{
			for x := chunkXMax + CHANK_SIZE; x > chunkXMax; x -= TILE_SIZE {
				for y := chunkYMax - CHANK_SIZE; y < chunkYMax; y += TILE_SIZE {
					posX := float32(x - 8)
					posY := float32(y + 8)

					tile := Tile{}

					tile.X = int(posX)
					tile.Y = int(posY)

					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.12:
						tile.Key = "Water"
					case perlinValue >= -0.12 && perlinValue <= 0.5:
						tile.Key = "Ground"
						rand.Seed(int64(time.Now().Nanosecond() + x - y))
						randomTree := rand.Float32()
						if randomTree > 0.90 {
							tree = NewTree(Coordinate{X: tile.X, Y: tile.Y})
						}
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}

					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

					treeMap[Coordinate{X: tree.X, Y: tree.Y}] = tree

				}
			}
		}
	case chunkYMax < 0:
		{
			for x := chunkXMax - CHANK_SIZE; x < chunkXMax; x += TILE_SIZE {
				for y := chunkYMax + CHANK_SIZE; y > chunkYMax; y -= TILE_SIZE {

					posX := float32(x + 8)
					posY := float32(y - 8)
					tile := Tile{}

					tile.X = int(posX)
					tile.Y = int(posY)
					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.12:
						tile.Key = "Water"
					case perlinValue >= -0.12 && perlinValue <= 0.5:
						tile.Key = "Ground"
						rand.Seed(int64(time.Now().Nanosecond() + x - y))
						randomTree := rand.Float32()
						if randomTree > 0.90 {
							tree = NewTree(Coordinate{X: tile.X, Y: tile.Y})
						}
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile
					treeMap[Coordinate{X: tree.X, Y: tree.Y}] = tree

				}
			}
		}
	default:
		{
			for x := chunkXMax - CHANK_SIZE; x < chunkXMax; x += TILE_SIZE {
				for y := chunkYMax - CHANK_SIZE; y < chunkYMax; y += TILE_SIZE {
					posX := float32(x + 8)
					posY := float32(y + 8)
					tile := Tile{}

					tile.X = int(posX)
					tile.Y = int(posY)
					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)

					switch {
					case perlinValue < -0.12:
						tile.Key = "Water"
					case perlinValue >= -0.12 && perlinValue <= 0.5:
						tile.Key = "Ground"
						rand.Seed(int64(time.Now().Nanosecond() + x - y))
						randomTree := rand.Float32()
						if randomTree > 0.90 {
							tree = NewTree(Coordinate{X: tile.X, Y: tile.Y})
						}
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile
					treeMap[Coordinate{X: tree.X, Y: tree.Y}] = tree

				}
			}
		}

	}

	chunk.Map = chunkMap
	chunk.Tree = treeMap
	fmt.Println(treeMap)

	return chunk
}
