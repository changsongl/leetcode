package main

import (
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	sameCnt := 0
	polarDegrees := []float64{}
	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			sameCnt++
		} else {
			polarDegrees = append(polarDegrees, math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0])))
		}
	}
	sort.Float64s(polarDegrees)

	n := len(polarDegrees)
	for _, deg := range polarDegrees {
		polarDegrees = append(polarDegrees, deg+2*math.Pi)
	}

	maxCnt := 0
	right := 0
	degree := float64(angle) * math.Pi / 180
	for i, deg := range polarDegrees[:n] {
		for right < n*2 && polarDegrees[right] <= deg+degree {
			right++
		}
		if right-i > maxCnt {
			maxCnt = right - i
		}
	}
	return sameCnt + maxCnt
}
