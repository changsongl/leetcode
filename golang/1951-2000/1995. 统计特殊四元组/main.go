package main

import "fmt"

func main() {
	fmt.Println(countQuadruplets([]int{1, 1, 1, 3, 5}))
}

func countQuadruplets(nums []int) int {
	ans := 0
	n := len(nums)
	dict := make(map[int]int, n)
	for b := n - 3; b >= 1; b-- {
		for d := b + 2; d < n; d++ {
			dict[nums[d]-nums[b+1]]++
		}
		for a := 0; a < b; a++ {
			if num := dict[nums[a]+nums[b]]; num > 0 {
				ans += num
			}
		}
	}
	return ans
}
