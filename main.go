package main

import (
	"path_finding/A_star"
	"path_finding/scene_generate"
)

func main() {
	width := int32(18)
	height := int32(18)
	fp := new(A_star.FindingPath)
	_map := scene_generate.GenerateMap(width, height)
	A_star.NewFindingPath(fp, A_star.Rect{Width: width, Height: height}, _map)
	fp.FindingPath()
	_map.PrintScene()
}