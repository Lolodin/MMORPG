package chunk

import (
	"fmt"
	"testing"
)

var id Coordinate

func Test_NewChunk(t *testing.T) {
	id.X = 32
	id.Y = 32
	v := NewChunk(id)
	fmt.Println(v.Map)
}
