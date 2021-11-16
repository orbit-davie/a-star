package A_star

import (
	"container/heap"
	"fmt"
	"path_finding/common"
	"path_finding/lib/utils"
	"path_finding/scene_generate"
)

type FindingPath struct {
	openList  *recHeap
	openMap   map[int32]*rec
	closeList map[int32]struct{}
	rect      Rect
	check     checkHandle

	scene *scene_generate.Scene
}

func NewFindingPath(fp *FindingPath, rect Rect, _map *scene_generate.Scene, option ...PathFindingOption) {
	total := rect.Width * rect.Height
	opList := make(recHeap, 0, total)
	fp.openList = &opList
	heap.Init(fp.openList)
	fp.openMap = make(map[int32]*rec, total)
	fp.closeList = make(map[int32]struct{}, total)
	fp.rect = rect
	fp.scene = _map

	if len(option) > 0 {
		op := option[0]
		if op.Check != nil {
			fp.check = op.Check
		}
	}
}

func (ins *FindingPath) FindingPath() {
	sPoint := ins.scene.GetStart()
	gPoint := ins.scene.GetGoal()
	ins.push(&rec{value: sPoint, priority: 0}, PosNum(sPoint.X, sPoint.Y))
	for ins.openList.Len() > 0 {
		curRec := ins.pop()
		hPoint := curRec.value
		ins.closeList[PosNum(hPoint.X, hPoint.Y)] = zero

		for _, direction := range directions {
			nPos := common.Point{X: hPoint.X + direction.h, Y: hPoint.Y + direction.v}
			if !ins.rect.IsValidPoint(nPos.X, nPos.Y) {
				continue
			}

			nPosNum := PosNum(nPos.X, nPos.Y)
			if _, ok := ins.closeList[nPosNum]; ok {
				continue
			}

			if ins.scene.IsGoal(nPos.X, nPos.Y) {
				ins.markScene(curRec)
				return
			}

			//block
			if ins.scene.IsBlock(nPos.X, nPos.Y) {
				continue
			}

			h := heuristic(&nPos, &gPoint)
			g := calcHvG(curRec, &direction)

			if remain, ok := ins.openMap[nPosNum]; ok {
				if remain.g > g {
					ins.fix(g, remain, curRec)
				}
			} else {
				rec := &rec{
					value:    nPos,
					father:   curRec,
					g:        g,
					priority: h + g,
				}
				ins.push(rec, nPosNum)
			}
		}
	}

	fmt.Println("finding path failed...")
}

//启发函数
func heuristic(n, goal *common.Point) int32 {
	//Manhattan Distance
	D := hv_move_cost
	return D * (utils.AbsInt32(n.X-goal.X) + utils.AbsInt32(n.Y-goal.Y))
}

func (ins *FindingPath) fix(g int32, remain, curRec *rec) {
	remain.g = g
	remain.father = curRec
	for index, item := range *ins.openList {
		if item.equal(remain) {
			heap.Fix(ins.openList, index)
			return
		}
	}
}

func (ins *FindingPath) push(rec *rec, rPosNum int32) {
	heap.Push(ins.openList, rec)
	ins.openMap[rPosNum] = rec
}

func (ins *FindingPath) pop() *rec {
	cur := heap.Pop(ins.openList).(*rec)
	hPosNum := PosNum(cur.value.X, cur.value.Y)
	delete(ins.openMap, hPosNum)
	return cur
}

func (ins *FindingPath) markScene(n *rec) {
	_map := ins.scene
	point := n.value
	_map.MarkPath(point.X, point.Y)
	for n.father != nil {
		n = n.father
		point = n.value
		_map.MarkPath(point.X, point.Y)
	}
}
