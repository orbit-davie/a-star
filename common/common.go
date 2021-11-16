package common

type Point struct {
	X int32
	Y int32
}

func NewPoint(x, y int32) Point {
	return Point{
		X: x,
		Y: y,
	}
}