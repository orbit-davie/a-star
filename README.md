# a-star
A* search algorithm for go

Usage example
```
func main() {
	width := int32(18)
	height := int32(18)
	fp := new(A_star.FindingPath)
	_map := scene_generate.GenerateMap(width, height)
	A_star.NewFindingPath(fp, A_star.Rect{Width: width, Height: height}, _map)
	fp.FindingPath()
	_map.PrintScene()
}
```
