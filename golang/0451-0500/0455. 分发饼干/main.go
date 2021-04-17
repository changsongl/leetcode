package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gLen := len(g)
	sLen := len(s)
	ans, gIndex, sIndex := 0, 0, 0

	for gIndex < gLen && sIndex < sLen {
		if s[sIndex] >= g[gIndex] {
			gIndex++
			ans++
		}
		sIndex++
	}
	return ans
}
