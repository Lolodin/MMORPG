package WorldMap

import (
	"Test/Chunk"
	"fmt"
	"sync"
)

type WorldMap struct {
	sync.Mutex
	Chunks map[Chunk.Coordinate]Chunk.Chunk
	Player map[string]*Player
	Tree   *Chunk.Tree
}

//Возвращает чанк соответствующий координатам, возвращает ошибку если такого чанка не существует
func (w *WorldMap) GetChunk(coordinate Chunk.Coordinate) (Chunk.Chunk, error) {
	c, ok := w.Chunks[coordinate]
	if ok == true {
		return c, nil
	} else {
		return c, fmt.Errorf("Chunck is not Exist")
	}

}

// Добавлявет в мир чанк
func (w *WorldMap) AddChunk(coordinate Chunk.Coordinate, chunk Chunk.Chunk) {

	isExist := w.isChunkExist(coordinate)
	if isExist {
		return
	} else {
		w.Chunks[coordinate] = chunk
	}

}

//Проверяет, существует чанк в мире или нет
func (w *WorldMap) isChunkExist(coordinate Chunk.Coordinate) bool {
	_, ok := w.Chunks[coordinate]

	return ok
}

func NewCacheWorldMap() WorldMap {
	world := WorldMap{}
	world.Chunks = make(map[Chunk.Coordinate]Chunk.Chunk)
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
	PX = PX
	PY = PY
	chunkId := GetChankID(PX, PY)
	w.Lock()
	objChunk := w.Chunks[chunkId]
	_, ok := objChunk.Tree[Chunk.Coordinate{X: PX, Y: PY}]
	w.Unlock()
	return ok
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
