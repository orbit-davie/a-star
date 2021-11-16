package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func AbsInt32(v int32) int32 {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func MsCast(createAt int64, params ...int) {
	var execNum int
	if len(params) > 0 && params[0] > 0 {
		execNum = params[0]
	}

	used := float64(time.Now().UnixNano()-createAt) / 1000000
	fmt.Println("ms cast: ", used)
	if execNum > 0 {
		fmt.Println(" exec num: ", execNum)
		fmt.Println(" op/ms: ", float64(used)/float64(execNum))
	}
}