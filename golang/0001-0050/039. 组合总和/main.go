package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)

	n := len(candidates)
	var dfs func(comb []int, target int, start int)
	dfs = func(comb []int, target int, start int) {
		if target == 0 {
			ans = append(ans, append([]int{}, comb...))
		}

		for i := start; i < n && candidates[i] <= target; i++ {
			comb = append(comb, candidates[i])
			dfs(comb, target-candidates[i], i)
			comb = comb[:len(comb)-1]
		}
	}

	dfs([]int{}, target, 0)
	return
}
