package main

import (
	"fmt"
	"sort"
)

func main() {
	//a := []int{1, 2, 3}
	//fmt.Println(a[:1])
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

// 枚举出所有可能
func subsetsWithDupMask(nums []int) (ans [][]int) {
	sort.Ints(nums)
	// 所有可能性
	comb := 1 << len(nums)
outer:
	for i := 0; i < comb; i++ {
		var t []int
		for j, num := range nums {
			// 如果当前位是1代表这个数字被选中，假如上个数字和这个数字一样，但是上个数字在这个
			// 组合里没有被选中，则代表此时已经是重复的组合了，因此直接跳到下一个组合。
			if i>>j&1 == 1 {
				if j > 0 && i>>(j-1)&1 == 0 && nums[j] == nums[j-1] {
					continue outer
				}
				t = append(t, num)
			}
		}
		ans = append(ans, append([]int(nil), t...))
	}
	return
}

// 回溯算法
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	l := len(nums)
	var t []int
	var dfs func(prevSelect bool, i int)
	dfs = func(prevSelect bool, i int) {
		if i == l {
			ans = append(ans, append([]int(nil), t...))
			return
		}

		dfs(false, i+1)

		if !prevSelect && i > 0 && nums[i] == nums[i-1] {
			return
		}

		t = append(t, nums[i])
		dfs(true, i+1)

		t = t[:len(t)-1]
	}

	dfs(false, 0)
	return ans
}
