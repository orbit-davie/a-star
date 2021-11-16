package A_star

import (
	"path_finding/common"
	"path_finding/lib/utils"
)

const (
	hv_move_cost = int32(10) //each horizontal or vertical square moved cost
	d_move_cost  = int32(14) //diagonal moved cost
)

type direction struct {
	h int32 //horizontal
	v int32 //vertical
}

type checkHandle func(x, y int32) bool

var (
	zero = struct{}{}

	//start at (1, 0), clockwise
	directions = []direction{
		{1, 0},
		//{1, -1},
		{0, -1},
		//{-1, -1},
		{-1, 0},
		//{-1, 1},
		{0, 1},
		//{1, 1},
	}
)

type Rect struct {
	Width  int32
	Height int32
}

func (r *Rect) IsValidPoint(x, y int32) bool {
	if x < 0 || x >= r.Width || y < 0 || y >= r.Height {
		return false
	}
	return true
}

func PosNum(x, y int32) int32 {
	return (x+1)*10000 + y + 1
}

func PosXY(posNum int32) (int32, int32) {
	return posNum/10000 - 1, posNum%10000 - 1
}

func calcG(father *rec, _direction *direction) (g int32) {
	if utils.AbsInt32(_direction.h) == utils.AbsInt32(_direction.v) {
		g = d_move_cost + father.g
	} else {
		g = hv_move_cost + father.g
	}
	return
}

func calcHvG(father *rec, _direction *direction) int32 {
	return hv_move_cost + father.g
}

func equal(n, goal *common.Point) bool {
	return n.X == goal.X && n.Y == goal.Y
}
