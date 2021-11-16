package scene_generate

import (
	"fmt"
	"math/rand"
	"path_finding/common"
	"path_finding/lib/utils"
	"time"
)

type Scene struct {
	Data  [][]int
	Start common.Point
	Goal  common.Point
}

func GenerateMap(width, height int32) *Scene {
	_map := new(Scene)
	_map.Data = make([][]int, height)
	for i := height - 1; i >= 0; i-- {
		list := make([]int, 0, width)
		for j := int32(0); j < width; j++ {
			list = append(list, 0)
		}
		_map.Data[i] = list
	}

	rand.Seed(time.Now().UTC().UnixNano()) // Seed Random properly
	num := (width * height) / 10
	for i := int32(0); i < num; i++ {
		x := int32(float64(width) * rand.Float64())
		y := int32(float64(height) * rand.Float64())
		_map.Data[y][x] = 1
	}

	// generate start
	_map.genStart(width, height)
	// generate goal
	_map.genGoal(width, height)
	return _map
}

func (ins *Scene) genStart(width, height int32) {
	x := int32(float64(width) * utils.RandFloat64(0.8, 1))
	y := int32(float64(height) * utils.RandFloat64(0.8, 1))
	ins.Data[y][x] = 5
	ins.Start = common.NewPoint(x, y)
}

func (ins *Scene) GetStart() common.Point {
	return ins.Start
}

func (ins *Scene) GetGoal() common.Point {
	return ins.Goal
}

func (ins *Scene) genGoal(width, height int32) {
	x := int32(float64(width) * utils.RandFloat64(0, 0.3))
	y := int32(float64(height) * utils.RandFloat64(0, 0.2))
	ins.Data[y][x] = 3
	ins.Goal = common.NewPoint(x, y)
}

func (ins *Scene) IsBlock(x, y int32) bool {
	pattern := ins.Data[y][x]
	return pattern == 1
}

func (ins *Scene) IsGoal(x, y int32) bool {
	pattern := ins.Data[y][x]
	return pattern == 3
}

func (ins *Scene) IsStart(x, y int32) bool {
	pattern := ins.Data[y][x]
	return pattern == 5
}

func (ins *Scene) MarkPath(x, y int32) {
	if ins.IsStart(x, y) {
		return
	}
	ins.Data[y][x] = 2
}

func (ins *Scene) MarkClosed(x, y int32) {
	if ins.IsStart(x, y) {
		return
	}
	//if pattern := ins.Data[y][x]; pattern != TYPE_NONE {
	//	return
	//}
	ins.Data[y][x] = TypeClosed
}

func (ins *Scene) PrintScene() {
	format := "\033[1;40;32m%v\033[0m"
	formatGreen := "\033[1;32m%v "
	formatWhite := "\033[1;37m%v "
	formatRed := "\033[1;31m%v "
	formatPurple := "\033[1;35m%v "
	formatBlue := "\033[1;34m%v "

	for i := 0; i < len(ins.Data); i++ {
		var str string
		for _, v := range ins.Data[i] {
			switch v {
			case TypeNone:
				str += fmt.Sprintf(formatWhite, " ")
			case TypeBlock:
				str += fmt.Sprintf(formatRed, "#")
			case TypePath: //path
				str += fmt.Sprintf(formatGreen, "P")
			case TypeGoal: //goal
				str += fmt.Sprintf(formatPurple, "E")
			case TypeClosed:
				str += fmt.Sprintf(formatBlue, "C")
			case TypeStart:
				str += fmt.Sprintf(formatPurple, "S")
			}
		}
		str += "\033[0m"
		fmt.Println(fmt.Sprintf(format, str))
	}
}
