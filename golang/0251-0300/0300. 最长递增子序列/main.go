package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

// 执行用时：68 ms, 在所有 Go 提交中击败了46.52%的用户
// 内存消耗：3.7 MB, 在所有 Go 提交中击败了39.55%的用户
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	ans := 0

	for i := 0; i < len(nums); i++ {
		max := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && max < dp[j]+1 {
				max = dp[j] + 1
			}
		}
		dp[i] = max
		if max > ans {
			ans = max
		}
	}

	return ans
}
