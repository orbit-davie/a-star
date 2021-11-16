package A_star

import "path_finding/common"

type rec struct {
	priority int32
	g        int32
	value    common.Point
	father   *rec
}

type recHeap []*rec

func (h recHeap) Len() int           { return len(h) }
func (h recHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h recHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *recHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*rec))
}

func (h *recHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (r *rec) equal(t *rec) bool {
	return r.value.X == t.value.X && r.value.Y == t.value.Y
}
