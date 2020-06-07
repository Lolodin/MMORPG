package chunk

import (
	"math/rand"
)

var SPECIES = [2]string{"Oak", "Spruce"}

type Tree struct {
	Species string  `json:"tree"` //Ель, береза и т.д
	Age     float32 `json:"age"`  //Возраст дерева = 1.0 максимальный
	X       int     `json:"x"`
	Y       int     `json:"y"`
}

func NewTree(coordinate Coordinate) Tree {
	t := Tree{}
	t.Age = rand.Float32()
	t.Species = SPECIES[rand.Intn(len(SPECIES))]
	t.X = coordinate.X
	t.Y = coordinate.Y
	if t.Age == 0 {
		t.Species = "Oak"
	}
	return t
}
