package WorldMap

import (
	"Test/Chunk"
	"fmt"
	"testing"
)

type testCoord struct {
	X      int
	Y      int
	Result Chunk.Coordinate
}

var tests = []testCoord{
	{X: 320, Y: 320, Result: Chunk.Coordinate{X: 2, Y: 2}},
	{X: 2560, Y: 2560, Result: Chunk.Coordinate{X: 3, Y: 3}},
}

func TestGetChankID(t *testing.T) {
	for _, testValue := range tests {
		t := GetChunkID(testValue.X, testValue.Y)
		fmt.Println(t == testValue.Result, t)
	}

}
