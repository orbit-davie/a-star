package A_star

import (
	"path_finding/lib/utils"
	"path_finding/scene_generate"
	"testing"
	"time"
)

func TestAccuracyVerityAStar(t *testing.T) {
	width := int32(18)
	height := int32(18)
	fp := new(FindingPath)
	_map := scene_generate.GenerateMap(width, height)
	NewFindingPath(fp, Rect{Width: width, Height: height}, _map)
	fp.FindingPath()
	_map.PrintScene()
}

func TestAStar(t *testing.T) {
	width := int32(18)
	height := int32(18)
	fp := new(FindingPath)
	_map := scene_generate.GenerateMap(width, height)
	NewFindingPath(fp, Rect{Width: width, Height: height}, _map)
	startAt := time.Now().UnixNano()
	fp.FindingPath()
	utils.MsCast(startAt)
	_map.PrintScene()
}
