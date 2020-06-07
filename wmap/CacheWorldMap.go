package wmap

import (
	"Test/chunk"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
)

type WorldMap struct {
	sync.Mutex
	Chunks map[chunk.Coordinate]chunk.Chunk
	Player map[string]*Player
	Tree   *chunk.Tree
}

//Возвращает чанк соответствующий координатам, возвращает ошибку если такого чанка не существует
func (w *WorldMap) GetChunk(coordinate chunk.Coordinate) (chunk.Chunk, error) {
	c, ok := w.Chunks[coordinate]
	if ok == true {
		return c, nil
	} else {
		return c, fmt.Errorf("Chunck is not Exist")
	}

}

// Добавлявет в мир чанк
func (w *WorldMap) AddChunk(coordinate chunk.Coordinate, chunk chunk.Chunk) {

	isExist := w.isChunkExist(coordinate)
	if isExist {
		return
	} else {
		w.Lock()
		w.Chunks[coordinate] = chunk
		log.WithFields(log.Fields{
			"package":  "WorldMap",
			"func":     "AddChunk",
			"Chunk":    chunk,
			"map Tree": chunk.Tree,
		}).Info("Create new Chunk")
		w.Unlock()
	}

}

//Проверяет, существует чанк в мире или нет
func (w *WorldMap) isChunkExist(coordinate chunk.Coordinate) bool {
	_, ok := w.Chunks[coordinate]

	return ok
}

func NewCacheWorldMap() WorldMap {
	world := WorldMap{}
	world.Chunks = make(map[chunk.Coordinate]chunk.Chunk)
	world.Player = make(map[string]*Player)
	return world
}

//Добавляем нового игрока в карту
func (w *WorldMap) AddPlayer(player *Player) {

	_, ok := w.Player[player.Name]
	if !ok {
		fmt.Println(player.Name)
		w.Lock()
		w.Player[player.Name] = player
		w.Unlock()
	} else {
		fmt.Println("Relogin: " + player.Name)
	}
}

// Обновляем данные персонажа в мире
func (w *WorldMap) UpdatePlayer(player Player) {
	w.Lock()
	p, ok := w.Player[player.Name]
	w.Unlock()
	if ok {

		p.X = player.X
		p.Y = player.Y
		w.Lock()
		w.Player[player.Name] = p
		w.Unlock()

	} else {
		fmt.Println("Player is not Exile")
	}

}

//map players
func (w *WorldMap) GetPlayers() Players {
	pls := Players{}
	w.Lock()
	for _, P := range w.Player {
		pls.P = append(pls.P, *P)
	}
	w.Unlock()
	return pls
}

// return true if Tree busy tile
func (w *WorldMap) CheckBusyTile(PX, PY int) bool {
	chunkId := GetChunkID(PX, PY)
	w.Lock()
	defer w.Unlock()
	b := w.Chunks[chunkId].Map[chunk.Coordinate{X: PX, Y: PY}].Busy
	c := w.Chunks[chunkId].Map[chunk.Coordinate{X: PX, Y: PY}].Key
	return b && c == "Water"

}

//Получить player
func (w *WorldMap) GetPlayer(name string) (*Player, bool) {
	pl, ok := w.Player[name]
	if ok {
		return pl, ok
	} else {
		return &Player{}, ok
	}

}
