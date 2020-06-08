package perlinNoise

import (
	"fmt"
	"testing"
)

type testarg struct {
	value  [2]float32
	result float32
}

var test = testarg{value: [2]float32{16 / 150, 16 / 150}, result: 10}

func TestNoise(t *testing.T) {
	v := Noise(test.value[0], test.value[1])
	fmt.Println(v)
}
