package PerlinNoise

import (
	"math"
	"math/rand"
)

func Noise(x, y float32) float32 {
	//Coordinate left and top vertex square
	left := float32(math.Floor(float64(x)))
	top := float32(math.Floor(float64(y)))

	//Local coordinate
	localPoinX := x - left
	localPoiny := y - top

	topLeft := getRandomVector(left, top)
	topRight := getRandomVector(left+1, top)
	bottomLeft := getRandomVector(left, top+1)
	bottomRight := getRandomVector(left+1, top+1)
	// Вектора от вершин до точки внутри
	DtopLeft := []float32{localPoinX, localPoiny}
	DtopRight := []float32{localPoinX - 1, localPoiny}
	DbottomLeft := []float32{localPoinX, localPoiny - 1}
	DbottomRight := []float32{localPoinX - 1, localPoiny - 1}

	//Скалярное произведение
	tx1 := dot(DtopLeft, topLeft)
	tx2 := dot(DtopRight, topRight)
	bx1 := dot(DbottomLeft, bottomLeft)
	bx2 := dot(DbottomRight, bottomRight)

	//параметры для нелинейности
	pointX := curve(localPoinX)
	pointY := curve(localPoiny)

	//Интерполяция

	tx := lerp(tx1, tx2, pointX)
	bx := lerp(bx1, bx2, pointX)
	tb := lerp(tx, bx, pointY)
	return tb

}
func getRandomVector(x, y float32) []float32 {
	rand.Seed(int64(x * y))
	v := rand.Intn(3)

	switch v {

	case 0:
		return []float32{-1, 0}
	case 1:
		return []float32{1, 0}
	case 2:
		return []float32{0, 1}
	default:
		return []float32{0, -1}

	}
}
func dot(a []float32, b []float32) float32 {

	return (a[0]*b[0] + b[1]*a[1])
}
func lerp(a, b, c float32) float32 {

	return a*(1-c) + b*c

}
func curve(t float32) float32 {

	return (t * t * t * (t*(t*6-15) + 10))

}
